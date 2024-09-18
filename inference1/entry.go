package inference1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

const (
	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"
	accessLogInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst) VALUES"

	EntryIdName   = "entry_id"
	AgentIdName   = "agent_id"
	CreatedTSName = "created_ts"
	RegionName    = "region"
	ZoneName      = "zone"
	SubZoneName   = "sub_zone"
	HostName      = "host"
	RouteName     = "route"

	ActualPercentName = "actual_percent"
	ActualValueName   = "actual_value"
	ActualMinimumName = "actual_minimum"

	LimitPercentName = "limit_percent"
	LimitValueName   = "limit_value"
	LimitMinimumName = "limit_minimum"
)

// Entry - host
type Entry struct {
	AgentId string      `json:"agent-id"`
	Origin  core.Origin `json:"origin"`
	Actual  Threshold   `json:"actual"`
	Limit   Threshold   `json:"limit"`
}

func NewEntry() *Entry {
	e := new(Entry)
	return e //NewPercentileThreshold(&e.Actual)
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case AgentIdName:
			e.AgentId = values[i].(string)
		//case CreatedTSName:
		//	e.CreatedTS = values[i].(time.Time)
		case RegionName:
			e.Origin.Region = values[i].(string)
		case ZoneName:
			e.Origin.Zone = values[i].(string)
		case SubZoneName:
			e.Origin.SubZone = values[i].(string)
		case HostName:
			e.Origin.Host = values[i].(string)
		case RouteName:
			e.Origin.Route = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.AgentId,
		//e.CreatedTS,

		e.Origin.Region,
		e.Origin.Zone,
		e.Origin.SubZone,
		e.Origin.Host,
		e.Origin.Route,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
