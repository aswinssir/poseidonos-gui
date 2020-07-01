/*
In this code we are using gin framework and influxdb golang client libraries
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster than other frameworks


DESCRIPTION: Collects the logs from InfluxDB
NAME : logs.go
@AUTHORS: Aswin K K
@Version : 1.0 *
@REVISION HISTORY
[6/25/2020] [aswin.kk] : Code for collecting Log

*/

package magent

import (
	"a-module/routers/m9k/model"
	"fmt"
)

type LogsField struct {
	Time  string
	Value string
}

type LogsFields []LogsField

// GetRebuildLogs gets the logs from influxdb and returns a JSON response
func GetRebuildLogs(param interface{}) (model.Response, error) {
	var res model.Response
	fieldsList := make(LogsFields, 0)
	timeInterval := param.(model.MAgentParam).Time
	query := fmt.Sprintf(RebuildingLogQ, timeInterval)
	result, err := ExecuteQuery(query)
	if err != nil {
		res.Result.Status.Description = err.Error()
		return res, nil
	}
	if len(result) == 0 || len(result[0].Series) == 0 {
		res.Result.Status.Description = errData.Error()
		return res, nil

	}
	for _, Values := range result[0].Series[0].Values {
		if Values[1] != nil {
			fieldsList = append(fieldsList, LogsField{Values[0].(string), Values[1].(string)})
		}
	}
	res.Result.Status.Code = 0
	res.Result.Status.Description = "DONE"
	res.Result.Data = fieldsList

	return res, nil
}