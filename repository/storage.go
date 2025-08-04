package repository

import (
	"errors"
	"fmt"
	"log-parser-and-analyzer/model"
	"log-parser-and-analyzer/repository/internal"
	"os"
)

const (
	FilePath    = "/home/deep/personal-projects/log-parser-and-analyzer/log.txt"
	MaxFileSize = 100 * 1024 * 1024 //100MB
)

func Load() (*[]model.Log, error) {
	info, err := os.Stat(FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist :%v", err)
		}
		return nil, errors.New("error in reading file info")
	}

	if info.Size() < MaxFileSize {
		data, err := os.ReadFile(FilePath)
		if err != nil {
			return nil, errors.New("error in reading file")
		}

		rows, err := internal.ParseLogFile(string(data))
		if err != nil {
			return nil, errors.New("error in parsing log file")
		}
		return rows, nil
	} else {
		// TODO: need to implement Streaming
		return nil, fmt.Errorf("file size is %d bytes, thus exceed the max limit allowed %d bytes. Streaming will be implemented soon", info.Size(), MaxFileSize)
	}

}
