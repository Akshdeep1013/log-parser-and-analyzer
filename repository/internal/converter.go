package internal

import (
	"fmt"
	"log-parser-and-analyzer/model"
	"regexp"
	"strconv"
	"strings"
)

func ParseLogFile(data string) (*[]model.Log, error) {
	rows := strings.Split(data, "\n")

	result := make([]model.Log, 0)
	//192.168.1.1 - - [25/Dec/2023:10:00:01 +0000] "GET /index.html HTTP/1.1" 200 1234

	//define pattern
	pattern := `(\d+\.\d+\.\d+\.\d+) - - \[([^:]+):([^\s]+) ([^]]+)\] "(\w+) ([^\s]+) ([^"]+)" (\d+) (\d+)`

	//compile pattern
	regex := regexp.MustCompile(pattern)

	for i, row := range rows {

		//extract data
		matches := regex.FindStringSubmatch(row)

		if len(matches) != 10 {
			return nil, fmt.Errorf("invalid log format at line i:%d", i)
		}
		statusCode, _ := strconv.Atoi(matches[8])
		responseSize, _ := strconv.Atoi(matches[9])
		log := model.Log{
			//marches[0] stores the entire matched string
			IP:           matches[1],
			Date:         matches[2],
			Time:         matches[3],
			Timezone:     matches[4],
			HttpMethod:   matches[5],
			Page:         matches[6],
			HttpVersion:  matches[7],
			StatusCode:   statusCode,
			ResponseSize: int64(responseSize),
		}
		result = append(result, log)
	}
	return &result, nil
}
