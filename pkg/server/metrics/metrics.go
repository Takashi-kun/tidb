// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"strconv"

	"github.com/pingcap/tidb/pkg/metrics"
	"github.com/pingcap/tidb/pkg/parser/mysql"
	"github.com/pingcap/tidb/pkg/resourcegroup"
	"github.com/prometheus/client_golang/prometheus"
)

// server metrics vars
var (
	QueryTotalCountOk  []prometheus.Counter
	QueryTotalCountErr []prometheus.Counter

	DisconnectNormal            prometheus.Counter
	DisconnectByClientWithError prometheus.Counter
	DisconnectErrorUndetermined prometheus.Counter

	ConnIdleDurationHistogramNotInTxn prometheus.Observer
	ConnIdleDurationHistogramInTxn    prometheus.Observer

	InPacketBytes  prometheus.Counter
	OutPacketBytes prometheus.Counter
)

func init() {
	InitMetricsVars()
}

// CmdToString convert command type to string.
func CmdToString(cmd byte) string {
	switch cmd {
	case mysql.ComSleep:
		return "Sleep"
	case mysql.ComQuit:
		return "Quit"
	case mysql.ComInitDB:
		return "InitDB"
	case mysql.ComQuery:
		return "Query"
	case mysql.ComPing:
		return "Ping"
	case mysql.ComFieldList:
		return "FieldList"
	case mysql.ComStmtPrepare:
		return "StmtPrepare"
	case mysql.ComStmtExecute:
		return "StmtExecute"
	case mysql.ComStmtFetch:
		return "StmtFetch"
	case mysql.ComStmtClose:
		return "StmtClose"
	case mysql.ComStmtSendLongData:
		return "StmtSendLongData"
	case mysql.ComStmtReset:
		return "StmtReset"
	case mysql.ComSetOption:
		return "SetOption"
	case mysql.ComCreateDB:
		return "CreateDB"
	case mysql.ComDropDB:
		return "DropDB"
	case mysql.ComRefresh:
		return "Refresh"
	case mysql.ComShutdown:
		return "Shutdown"
	case mysql.ComStatistics:
		return "Statistics"
	case mysql.ComProcessInfo:
		return "ProcessInfo"
	case mysql.ComConnect:
		return "Connect"
	case mysql.ComProcessKill:
		return "ProcessKill"
	case mysql.ComDebug:
		return "Debug"
	case mysql.ComTime:
		return "Time"
	case mysql.ComDelayedInsert:
		return "DelayedInsert"
	case mysql.ComChangeUser:
		return "ChangeUser"
	case mysql.ComBinlogDump:
		return "BinlogDump"
	case mysql.ComTableDump:
		return "TableDump"
	case mysql.ComConnectOut:
		return "ConnectOut"
	case mysql.ComRegisterSlave:
		return "RegisterSlave"
	case mysql.ComDaemon:
		return "Daemon"
	case mysql.ComBinlogDumpGtid:
		return "BinlogDumpGtid"
	case mysql.ComResetConnection:
		return "ResetConnection"
	case mysql.ComEnd:
		return "End"
	}
	return strconv.Itoa(int(cmd))
}

// InitMetricsVars init server metrics vars.
func InitMetricsVars() {
	QueryTotalCountOk = []prometheus.Counter{
		mysql.ComSleep:            metrics.QueryTotalCounter.WithLabelValues("Sleep", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComQuit:             metrics.QueryTotalCounter.WithLabelValues("Quit", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComInitDB:           metrics.QueryTotalCounter.WithLabelValues("InitDB", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComQuery:            metrics.QueryTotalCounter.WithLabelValues("Query", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComPing:             metrics.QueryTotalCounter.WithLabelValues("Ping", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComFieldList:        metrics.QueryTotalCounter.WithLabelValues("FieldList", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtPrepare:      metrics.QueryTotalCounter.WithLabelValues("StmtPrepare", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtExecute:      metrics.QueryTotalCounter.WithLabelValues("StmtExecute", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtFetch:        metrics.QueryTotalCounter.WithLabelValues("StmtFetch", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtClose:        metrics.QueryTotalCounter.WithLabelValues("StmtClose", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtSendLongData: metrics.QueryTotalCounter.WithLabelValues("StmtSendLongData", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtReset:        metrics.QueryTotalCounter.WithLabelValues("StmtReset", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComSetOption:        metrics.QueryTotalCounter.WithLabelValues("SetOption", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComCreateDB:         metrics.QueryTotalCounter.WithLabelValues("CreateDB", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComDropDB:           metrics.QueryTotalCounter.WithLabelValues("DropDB", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComRefresh:          metrics.QueryTotalCounter.WithLabelValues("Refresh", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComShutdown:         metrics.QueryTotalCounter.WithLabelValues("Shutdown", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComStatistics:       metrics.QueryTotalCounter.WithLabelValues("Statistics", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComProcessInfo:      metrics.QueryTotalCounter.WithLabelValues("ProcessInfo", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComConnect:          metrics.QueryTotalCounter.WithLabelValues("Connect", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComProcessKill:      metrics.QueryTotalCounter.WithLabelValues("ProcessKill", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComDebug:            metrics.QueryTotalCounter.WithLabelValues("Debug", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComTime:             metrics.QueryTotalCounter.WithLabelValues("Time", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComDelayedInsert:    metrics.QueryTotalCounter.WithLabelValues("DelayedInsert", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComChangeUser:       metrics.QueryTotalCounter.WithLabelValues("ChangeUser", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComBinlogDump:       metrics.QueryTotalCounter.WithLabelValues("BinlogDump", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComTableDump:        metrics.QueryTotalCounter.WithLabelValues("TableDump", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComConnectOut:       metrics.QueryTotalCounter.WithLabelValues("ConnectOut", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComRegisterSlave:    metrics.QueryTotalCounter.WithLabelValues("RegisterSlave", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComDaemon:           metrics.QueryTotalCounter.WithLabelValues("Daemon", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComBinlogDumpGtid:   metrics.QueryTotalCounter.WithLabelValues("BinlogDumpGtid", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComResetConnection:  metrics.QueryTotalCounter.WithLabelValues("ResetConnection", "OK", resourcegroup.DefaultResourceGroupName),
		mysql.ComEnd:              metrics.QueryTotalCounter.WithLabelValues("End", "OK", resourcegroup.DefaultResourceGroupName),
	}
	QueryTotalCountErr = []prometheus.Counter{
		mysql.ComSleep:            metrics.QueryTotalCounter.WithLabelValues("Sleep", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComQuit:             metrics.QueryTotalCounter.WithLabelValues("Quit", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComInitDB:           metrics.QueryTotalCounter.WithLabelValues("InitDB", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComQuery:            metrics.QueryTotalCounter.WithLabelValues("Query", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComPing:             metrics.QueryTotalCounter.WithLabelValues("Ping", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComFieldList:        metrics.QueryTotalCounter.WithLabelValues("FieldList", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtPrepare:      metrics.QueryTotalCounter.WithLabelValues("StmtPrepare", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtExecute:      metrics.QueryTotalCounter.WithLabelValues("StmtExecute", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtFetch:        metrics.QueryTotalCounter.WithLabelValues("StmtFetch", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtClose:        metrics.QueryTotalCounter.WithLabelValues("StmtClose", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtSendLongData: metrics.QueryTotalCounter.WithLabelValues("StmtSendLongData", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStmtReset:        metrics.QueryTotalCounter.WithLabelValues("StmtReset", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComSetOption:        metrics.QueryTotalCounter.WithLabelValues("SetOption", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComCreateDB:         metrics.QueryTotalCounter.WithLabelValues("CreateDB", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComDropDB:           metrics.QueryTotalCounter.WithLabelValues("DropDB", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComRefresh:          metrics.QueryTotalCounter.WithLabelValues("Refresh", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComShutdown:         metrics.QueryTotalCounter.WithLabelValues("Shutdown", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComStatistics:       metrics.QueryTotalCounter.WithLabelValues("Statistics", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComProcessInfo:      metrics.QueryTotalCounter.WithLabelValues("ProcessInfo", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComConnect:          metrics.QueryTotalCounter.WithLabelValues("Connect", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComProcessKill:      metrics.QueryTotalCounter.WithLabelValues("ProcessKill", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComDebug:            metrics.QueryTotalCounter.WithLabelValues("Debug", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComTime:             metrics.QueryTotalCounter.WithLabelValues("Time", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComDelayedInsert:    metrics.QueryTotalCounter.WithLabelValues("DelayedInsert", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComChangeUser:       metrics.QueryTotalCounter.WithLabelValues("ChangeUser", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComBinlogDump:       metrics.QueryTotalCounter.WithLabelValues("BinlogDump", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComTableDump:        metrics.QueryTotalCounter.WithLabelValues("TableDump", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComConnectOut:       metrics.QueryTotalCounter.WithLabelValues("ConnectOut", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComRegisterSlave:    metrics.QueryTotalCounter.WithLabelValues("RegisterSlave", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComDaemon:           metrics.QueryTotalCounter.WithLabelValues("Daemon", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComBinlogDumpGtid:   metrics.QueryTotalCounter.WithLabelValues("BinlogDumpGtid", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComResetConnection:  metrics.QueryTotalCounter.WithLabelValues("ResetConnection", "Error", resourcegroup.DefaultResourceGroupName),
		mysql.ComEnd:              metrics.QueryTotalCounter.WithLabelValues("End", "Error", resourcegroup.DefaultResourceGroupName),
	}

	DisconnectNormal = metrics.DisconnectionCounter.WithLabelValues(metrics.LblOK)
	DisconnectByClientWithError = metrics.DisconnectionCounter.WithLabelValues(metrics.LblError)
	DisconnectErrorUndetermined = metrics.DisconnectionCounter.WithLabelValues("undetermined")

	ConnIdleDurationHistogramNotInTxn = metrics.ConnIdleDurationHistogram.WithLabelValues("0")
	ConnIdleDurationHistogramInTxn = metrics.ConnIdleDurationHistogram.WithLabelValues("1")

	InPacketBytes = metrics.PacketIOCounter.WithLabelValues("In")
	OutPacketBytes = metrics.PacketIOCounter.WithLabelValues("Out")
}
