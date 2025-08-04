package repository

import (
	"errors"
	"fmt"
	"log-parser-and-analyzer/model"
	"log-parser-and-analyzer/repository/internal"
	"os"
)

const (
	FilePath = "/home/deep/personal-projects/log-parser-and-analyzer/log.txt"
)

func Load() (*[]model.Log, error) {
	data, err := os.ReadFile(FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("File does not exist :%v", err)
		}
		return nil, errors.New("Error in reading file")
	}
	rows, err := internal.ParseLogFile(string(data))
	if err != nil {
		return nil, errors.New("Error in parsing log file")
	}
	return rows, nil
}
