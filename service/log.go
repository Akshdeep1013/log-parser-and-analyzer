package service

import (
	"log-parser-and-analyzer/model"
	"strconv"
	"strings"
)

func GetLogInsight(logs *[]model.Log) (model.LogInsight, error) {
	result := model.LogInsight{
		Overall: getOverallInsight(logs),
		Success: getSuccessInsight(logs),
		Failed:  getFailureInsight(logs),
		Traffic: getTrafficInsight(logs),
	}

	return result, nil
}

func getOverallInsight(logs *[]model.Log) model.OverallInsight {
	ipLookUP, httpLookUp := make(map[string]bool), make(map[string]int)
	for _, log := range *logs {
		ipLookUP[log.IP] = true
		httpLookUp[log.HttpVersion]++
	}
	http := make([]model.HttpVersion, 0)
	for k, v := range httpLookUp {
		http = append(http, model.HttpVersion{
			Version: k,
			Total:   v,
		})
	}

	return model.OverallInsight{
		TotalRequest: len(*logs),
		UniqueIPs:    len(ipLookUP),
		Http:         http,
	}
}

func getSuccessInsight(logs *[]model.Log) model.SuccessInsight {
	ipLookUP, pageLookUp := make(map[string]int), make(map[string]int)
	var count, max int
	var maxResponseSize, sum int64
	for _, log := range *logs {
		if log.StatusCode >= 200 && log.StatusCode < 400 {
			ipLookUP[log.IP]++
			pageLookUp[log.Page]++
			count++
			if log.ResponseSize > maxResponseSize {
				maxResponseSize = log.ResponseSize
			}
			sum += log.ResponseSize
		}
	}

	var mostPageVisited string
	max = 0
	for k, v := range pageLookUp {
		if v > max {
			max = v
			mostPageVisited = k
		}
	}

	var mostActiveIP string
	max = 0
	for k, v := range ipLookUP {
		if v > max {
			max = v
			mostActiveIP = k
		}
	}

	return model.SuccessInsight{
		TotalSucceedRequest: count,
		UniqueIPs:           len(ipLookUP),
		MostActiveIP:        mostActiveIP,
		MostPageVisited:     mostPageVisited,
		MaxResponseSize:     maxResponseSize,
		TotalResponseSize:   sum,
	}
}

func getFailureInsight(logs *[]model.Log) model.FailureInsight {
	ipLookUP, pageLookUp, errorLookUp := make(map[string]int), make(map[string]int), make(map[int]int)
	var count, max int
	for _, log := range *logs {
		if log.StatusCode >= 400 {
			ipLookUP[log.IP]++
			pageLookUp[log.Page]++
			errorLookUp[log.StatusCode]++
			count++
		}
	}

	var mostPageVisited string
	max = 0
	for k, v := range pageLookUp {
		if v > max {
			max = v
			mostPageVisited = k
		}
	}

	var mostActiveIP string
	max = 0
	for k, v := range ipLookUP {
		if v > max {
			max = v
			mostActiveIP = k
		}
	}

	return model.FailureInsight{
		TotalFailureRequest: count,
		UniqueIPs:           len(ipLookUP),
		MostPageVisited:     mostPageVisited,
		MostActiveIP:        mostActiveIP,
		ErrorCode:           errorLookUp,
	}
}

func getTrafficInsight(logs *[]model.Log) map[string]model.TrafficInsight {
	lookUp := make(map[string]model.TrafficInsight)

	for _, log := range *logs {
		hour := getHour(log.Time)

		if ti, exist := lookUp[log.Timezone]; exist {
			//if date exist
			if _, dateExist := ti.HourlyTraffic[log.Date]; dateExist {
				ti.HourlyTraffic[log.Date][hour]++
			} else {
				//need to create date
				ti.HourlyTraffic[log.Date] = make(map[string]int)
				ti.HourlyTraffic[log.Date][hour] = 1
			}
			lookUp[log.Timezone] = ti

		} else {
			hourLookUp := make(map[string]int)
			hourLookUp[hour] = 1

			dateLookUp := make(map[string]map[string]int)
			dateLookUp[log.Date] = hourLookUp

			lookUp[log.Timezone] = model.TrafficInsight{
				HourlyTraffic: dateLookUp,
			}
		}
	}

	return lookUp
}

func getHour(time string) string {
	timeSplits := strings.Split(time, ":")
	hour, _ := strconv.Atoi(timeSplits[0])
	return strconv.Itoa(hour)
}
