package magent

import "a-module/errors"

var DBName = "mtool_db"
var DBAddress = "http://127.0.0.1:8086"
var AggTime = []string{"7d", "30d"}
var AggRP = "agg_rp"
var DefaultRP = "default_rp"
var TimeGroupsDefault = map[string]string{
	"1m":  "1s",
	"5m":  "1s",
	"15m": "1m",
	"1h":  "1m",
	"6h":  "1m",
	"12h": "5m",
	"24h": "10m",
	"7d":  "1h",
	"30d": "6h",
	"60d": "60d",
}

var (
	errConnInfluxDB = errors.New("could not connect to Influxdb")
	errQuery        = errors.New("could not query to database")
	errData         = errors.New("no data available")
	errEndPoint     = errors.New("use time from 1m,5m,15m,1h,6h,12h,24h,7d,30d")
)

var netAggRPQ = "SELECT mean(bytes_recv) ,mean(bytes_sent), mean(drop_in), mean(drop_out), mean(err_in), mean(err_out), mean(packets_recv), mean(packets_sent)  FROM %s.%s.mean_net WHERE time > now() - %s"
var netDefaultRPQ = "SELECT mean(bytes_recv) ,mean(bytes_sent), mean(drop_in), mean(drop_out), mean(err_in), mean(err_out), mean(packets_recv), mean(packets_sent)  FROM %s.%s.net WHERE time > now() - %s GROUP BY time(%s)"
var netLastRecordQ = "SELECT mean(bytes_recv), mean(bytes_sent), mean(drop_in), mean(drop_out), mean(err_in), mean(err_out), mean(packets_recv), mean(packets_sent)  FROM %s.%s.net LIMIT 1"
var netAddQ = "SELECT time, \"name\", address FROM %s.autogen.ethernet"
var netDriverQ = "SELECT time, \"name\", driver FROM %s.autogen.ethernet"
var cpuAggRPQ = "SELECT usage_user AS mean_usage_user FROM %s.%s.mean_cpu WHERE time > now() - %s"
var cpuDefaultRPQ = "SELECT mean(usage_user) AS mean_usage_user FROM %s.%s.cpu WHERE time > now() - %s GROUP BY time(%s)"
var cpuLastRecordQ = "SELECT last(usage_user) AS mean_usage_user FROM %s.%s.cpu LIMIT 1"
var diskAggRPQ = "SELECT used AS mean_used_bytes FROM %s.%s.mean_disk WHERE time > now() - %s"
var diskDefaultRPQ = "SELECT mean(used) AS mean_used_bytes FROM %s.%s.disk WHERE time > now() - %s GROUP BY time(%s)"
var diskLastRecordQ = "SELECT last(used) AS mean_used_bytes FROM %s.%s.disk LIMIT 1"
var memoryAggRPQ = "SELECT used_percent AS mean_used_percent FROM %s.%s.mean_mem WHERE time > now() - %s"
var memoryDefaultRPQ = "SELECT mean(used_percent) AS mean_used_percent FROM %s.%s.mem WHERE time > now() - %s GROUP BY time(%s)"
var memoryLastRecordQ = "SELECT last(used_percent) AS mean_used_percent FROM %s.%s.mem LIMIT 1"

var ReadBandwidthAggRPQ = `SELECT /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_bw_read$/, /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "bw" FROM "mtool_db"."%s"."mean_air" WHERE time > now() - %s FILL(null)`
var ReadBandwidthDefaultRPQ = `SELECT mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_bw_read$/), mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/) as "bw" FROM "mtool_db"."%s"."air" WHERE time > now() - %s GROUP BY time(%s) FILL(null)`
var ReadBandwidthLastRecordQ = `SELECT /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_bw_read$/, /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "bw" FROM "mtool_db"."%s"."air" order by time desc limit 1`

var WriteBandwidthAggRPQ = `SELECT /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_bw_write$/, /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "bw" FROM "mtool_db"."%s"."mean_air" WHERE time > now() - %s FILL(null)`
var WriteBandwidthDefaultRPQ = `SELECT mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_bw_write$/), mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/) as "bw" FROM "mtool_db"."%s"."air" WHERE time > now() - %s GROUP BY time(%s) FILL(null)`
var WriteBandwidthLastRecordQ = `SELECT /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_bw_write$/, /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "bw" FROM "mtool_db"."%s"."air" order by time desc limit 1`

var ReadIOPSAggRPQ = `SELECT /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_iops_read$/, /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "iops" FROM "mtool_db"."%s"."mean_air" WHERE time > now() - %s FILL(null)`
var ReadIOPSDefaultRPQ = `SELECT mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_iops_read$/), mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/) as "iops" FROM "mtool_db"."%s"."air" WHERE time > now() - %s GROUP BY time(%s) FILL(null)`
var ReadIOPSLastRecordQ = `SELECT /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_iops_read$/, /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "iops" FROM "mtool_db"."%s"."air" order by time desc limit 1`

var WriteIOPSAggRPQ = `SELECT /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_iops_write$/, /^mean_perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "iops" FROM "mtool_db"."%s"."mean_air" WHERE time > now() - %s FILL(null)`
var WriteIOPSDefaultRPQ = `SELECT mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_iops_write$/), mean(/^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/) as "iops" FROM "mtool_db"."%s"."air" WHERE time > now() - %s GROUP BY time(%s) FILL(null)`
var WriteIOPSLastRecordQ = `SELECT /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_iops_write$/, /^perf_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "iops" FROM "mtool_db"."%s"."air" order by time desc limit 1`

var LatencyAggRPQ = `SELECT /^mean_lat_data_0_tid_arr_[\S]_aid_arr_[\S]_timelag_arr_0_mean$/, /^mean_lat_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "latency" FROM "mtool_db"."%s"."mean_air" WHERE time > now() - %s FILL(null)`
var LatencyDefaultRPQ = `SELECT mean(/^lat_data_0_tid_arr_[\S]_aid_arr_[\S]_timelag_arr_0_mean$/), mean(/^lat_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/) as "latency" FROM "mtool_db"."%s"."air" WHERE time > now() - %s GROUP BY time(%s) FILL(null)`
var LatencyLastRecordQ = `SELECT /^lat_data_0_tid_arr_[\S]_aid_arr_[\S]_timelag_arr_0_mean$/, /^lat_data_0_tid_arr_[\S]_aid_arr_[\S]_aid$/ as "latency" FROM "mtool_db"."%s"."air" order by time desc limit 1`

var RebuildingLogQ = `SELECT "value" FROM "mtool_db"."autogen"."rebuilding_status" WHERE time > now() - %s`
