package action1

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
	DetailsName   = "details"
	ActionName    = "action"
)
