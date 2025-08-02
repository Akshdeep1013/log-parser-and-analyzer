package model

//192.168.1.1 - - [25/Dec/2023:10:00:01 +0000] "GET /index.html HTTP/1.1" 200 1234

// just a structure to store all rows one by one
type Log struct {
	IP           string
	Timestamp    string
	HttpMethod   string
	Page         string
	HttpVersion  string
	StatusCode   int
	ResponseSize int64
}

// prepare insight from above logs
type LogInsight struct {
	Overall OverallInsight

	Success SuccessInsight
	Failed  FailureInsight

	Traffic map[string]TrafficInsight //traffic by timezone
}

type OverallInsight struct {
	TotalRequest int
	UniqueIPs    int
	Http         []HttpVersion
}

type TrafficInsight struct {
	HourlyTraffic map[string]map[string]int //each hour of each day traffic
}

type HttpVersion struct {
	Version string
	Total   int
}

type SuccessInsight struct {
	TotalSucceedRequest int
	UniqueIPs           int
	MostPageVisited     string
	MostActiveIP        string
	MaxResponseSize     int64
	TotalResponseSize   int64
}

type FailureInsight struct {
	TotalFailureRequest int
	UniqueIPs           int
	MostPageVisited     string
	MostActiveIP        string
	ErrorCode           map[int]int
}
