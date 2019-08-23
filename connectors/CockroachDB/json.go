package cockroachdb

import "math/big"

/**
**
** Json objects : TODO : Is there a better way of doing this ?
**
 */

/*
type ClusterHealth struct {
	Status            uint   // Status of cluster (ok, error, etc.)
	Ranges            uint   // Number of ranges
	Replicas          uint   // Number of Replicas
	Nodes             uint   // Number of nodes
	Init              uint   // Number of nodes that are initialising
	Out               uint   // Number of nodes that are decommissioned or shutting down
	Dead              uint   // Number of nodes that are unavailable
	CapacityUsage     uint64 // Number of bytes of used Capacity
	CapacityReserved  uint64 //
	MemoryUsage       uint64 //
	MemoryReserved    uint64
	UnavailableRanges uint
	QueriesPerSecond  float32
	Latency99         float32
	MeanClockOffset   uint32
}*/

// NodeStats concentrates information and statistics about a node
type NodeStats struct {
	ID                int  // Id in the cluster
	Up                bool // Is it Up ?
	Init              bool // Is it in init state ?
	Out               bool // Is it decommissioned or shutting down ?
	Dead              bool // Is it completely unavailable ?
	Ranges            uint // Number of ranges this node holds
	Replicas          uint // Number of Replicas
	CapacityUsage     uint64
	CapacityReserved  uint64
	CapacityAvailable uint64
	Capacity          uint64
	TotalSystemMemory int64
	MemoryUsage       uint64
	MemoryReserved    uint64
	MemoryAvailable   uint64
	QueriesPerSecond  float64
	Latency99         uint64
	ClockOffset       uint32
}

func newNodeStats() *NodeStats {
	return &NodeStats{
		ID:                0,
		Up:                false,
		Init:              false,
		Out:               false,
		Dead:              false,
		Ranges:            0,
		Replicas:          0,
		CapacityUsage:     0,
		CapacityReserved:  0,
		CapacityAvailable: 0,
		Capacity:          0,
		TotalSystemMemory: 0,
		MemoryUsage:       0,
		MemoryReserved:    0,
		MemoryAvailable:   0,
		QueriesPerSecond:  0,
		Latency99:         0,
		ClockOffset:       0,
	}
}

/*
Top structures specific to requests
*/

// Health holds general information about a node
// /health
// /healthready
type Health struct {
	NodeID  int `json:"nodeId"`
	Address struct {
		NetworkField string `json:"networkField"`
		AddressField string `json:"addressField"`
	} `json:"address"`
	BuildInfo struct {
		GoVersion       string      `json:"goVersion"`
		Tag             string      `json:"tag"`
		Time            string      `json:"time"`
		Revision        string      `json:"revision"`
		CgoCompiler     string      `json:"cgoCompiler"`
		CgoTargetTriple string      `json:"cgoTargetTriple"`
		Platform        string      `json:"platform"`
		Distribution    string      `json:"distribution"`
		Type            string      `json:"type"`
		Channel         string      `json:"channel"`
		EnvChannel      string      `json:"envChannel"`
		Dependencies    interface{} `json:"dependencies"`
	} `json:"buildInfo"`
	SystemInfo struct {
		SystemInfo string `json:"systemInfo"`
		KernelInfo string `json:"kernelInfo"`
	} `json:"systemInfo"`
}

type metrics struct {
	BuildTimestamp                           int     `json:"build.timestamp"`
	ChangefeedBufferEntriesIn                int     `json:"changefeed.buffer_entries.in"`
	ChangefeedBufferEntriesOut               int     `json:"changefeed.buffer_entries.out"`
	ChangefeedEmitNanos                      int     `json:"changefeed.emit_nanos"`
	ChangefeedEmittedBytes                   int     `json:"changefeed.emitted_bytes"`
	ChangefeedEmittedMessages                int     `json:"changefeed.emitted_messages"`
	ChangefeedErrorRetries                   int     `json:"changefeed.error_retries"`
	ChangefeedFlushNanos                     int     `json:"changefeed.flush_nanos"`
	ChangefeedFlushes                        int     `json:"changefeed.flushes"`
	ChangefeedMaxBehindNanos                 int     `json:"changefeed.max_behind_nanos"`
	ChangefeedMinHighWater                   big.Int `json:"changefeed.min_high_water"`
	ChangefeedPollRequestNanosMax            int     `json:"changefeed.poll_request_nanos-max"`
	ChangefeedPollRequestNanosP50            int     `json:"changefeed.poll_request_nanos-p50"`
	ChangefeedPollRequestNanosP75            int     `json:"changefeed.poll_request_nanos-p75"`
	ChangefeedPollRequestNanosP90            int     `json:"changefeed.poll_request_nanos-p90"`
	ChangefeedPollRequestNanosP99            int     `json:"changefeed.poll_request_nanos-p99"`
	ChangefeedPollRequestNanosP999           int     `json:"changefeed.poll_request_nanos-p99.9"`
	ChangefeedPollRequestNanosP9999          int     `json:"changefeed.poll_request_nanos-p99.99"`
	ChangefeedPollRequestNanosP99999         int     `json:"changefeed.poll_request_nanos-p99.999"`
	ChangefeedProcessingNanos                int     `json:"changefeed.processing_nanos"`
	ChangefeedTableMetadataNanos             int     `json:"changefeed.table_metadata_nanos"`
	ClockOffsetMeannanos                     int     `json:"clock-offset.meannanos"`
	ClockOffsetStddevnanos                   int     `json:"clock-offset.stddevnanos"`
	DistsenderBatches                        int     `json:"distsender.batches"`
	DistsenderBatchesAsyncSent               int     `json:"distsender.batches.async.sent"`
	DistsenderBatchesAsyncThrottled          int     `json:"distsender.batches.async.throttled"`
	DistsenderBatchesPartial                 int     `json:"distsender.batches.partial"`
	DistsenderErrorsInleasetransferbackoffs  int     `json:"distsender.errors.inleasetransferbackoffs"`
	DistsenderErrorsNotleaseholder           int     `json:"distsender.errors.notleaseholder"`
	DistsenderRPCSent                        int     `json:"distsender.rpc.sent"`
	DistsenderRPCSentLocal                   int     `json:"distsender.rpc.sent.local"`
	DistsenderRPCSentNextreplicaerror        int     `json:"distsender.rpc.sent.nextreplicaerror"`
	ExecError                                int     `json:"exec.error"`
	ExecLatencyMax                           int     `json:"exec.latency-max"`
	ExecLatencyP50                           int     `json:"exec.latency-p50"`
	ExecLatencyP75                           int     `json:"exec.latency-p75"`
	ExecLatencyP90                           int     `json:"exec.latency-p90"`
	ExecLatencyP99                           int     `json:"exec.latency-p99"`
	ExecLatencyP999                          int     `json:"exec.latency-p99.9"`
	ExecLatencyP9999                         int     `json:"exec.latency-p99.99"`
	ExecLatencyP99999                        int     `json:"exec.latency-p99.999"`
	ExecSuccess                              int     `json:"exec.success"`
	GossipBytesReceived                      int     `json:"gossip.bytes.received"`
	GossipBytesSent                          int     `json:"gossip.bytes.sent"`
	GossipConnectionsIncoming                int     `json:"gossip.connections.incoming"`
	GossipConnectionsOutgoing                int     `json:"gossip.connections.outgoing"`
	GossipConnectionsRefused                 int     `json:"gossip.connections.refused"`
	GossipInfosReceived                      int     `json:"gossip.infos.received"`
	GossipInfosSent                          int     `json:"gossip.infos.sent"`
	LivenessEpochincrements                  int     `json:"liveness.epochincrements"`
	LivenessHeartbeatfailures                int     `json:"liveness.heartbeatfailures"`
	LivenessHeartbeatlatencyMax              int     `json:"liveness.heartbeatlatency-max"`
	LivenessHeartbeatlatencyP50              int     `json:"liveness.heartbeatlatency-p50"`
	LivenessHeartbeatlatencyP75              int     `json:"liveness.heartbeatlatency-p75"`
	LivenessHeartbeatlatencyP90              int     `json:"liveness.heartbeatlatency-p90"`
	LivenessHeartbeatlatencyP99              int     `json:"liveness.heartbeatlatency-p99"`
	LivenessHeartbeatlatencyP999             int     `json:"liveness.heartbeatlatency-p99.9"`
	LivenessHeartbeatlatencyP9999            int     `json:"liveness.heartbeatlatency-p99.99"`
	LivenessHeartbeatlatencyP99999           int     `json:"liveness.heartbeatlatency-p99.999"`
	LivenessHeartbeatsuccesses               int     `json:"liveness.heartbeatsuccesses"`
	LivenessLivenodes                        int     `json:"liveness.livenodes"`
	NodeID                                   int     `json:"node-id"`
	RoundTripLatencyMax                      int     `json:"round-trip-latency-max"`
	RoundTripLatencyP50                      int     `json:"round-trip-latency-p50"`
	RoundTripLatencyP75                      int     `json:"round-trip-latency-p75"`
	RoundTripLatencyP90                      int     `json:"round-trip-latency-p90"`
	RoundTripLatencyP99                      int     `json:"round-trip-latency-p99"`
	RoundTripLatencyP999                     int     `json:"round-trip-latency-p99.9"`
	RoundTripLatencyP9999                    int     `json:"round-trip-latency-p99.99"`
	RoundTripLatencyP99999                   int     `json:"round-trip-latency-p99.999"`
	SQLBytesin                               int     `json:"sql.bytesin"`
	SQLBytesout                              int     `json:"sql.bytesout"`
	SQLConns                                 int     `json:"sql.conns"`
	SQLDdlCount                              int     `json:"sql.ddl.count"`
	SQLDdlCountInternal                      int     `json:"sql.ddl.count.internal"`
	SQLDeleteCount                           int     `json:"sql.delete.count"`
	SQLDeleteCountInternal                   int     `json:"sql.delete.count.internal"`
	SQLDistsqlExecLatencyMax                 int     `json:"sql.distsql.exec.latency-max"`
	SQLDistsqlExecLatencyP50                 int     `json:"sql.distsql.exec.latency-p50"`
	SQLDistsqlExecLatencyP75                 int     `json:"sql.distsql.exec.latency-p75"`
	SQLDistsqlExecLatencyP90                 int     `json:"sql.distsql.exec.latency-p90"`
	SQLDistsqlExecLatencyP99                 int     `json:"sql.distsql.exec.latency-p99"`
	SQLDistsqlExecLatencyP999                int     `json:"sql.distsql.exec.latency-p99.9"`
	SQLDistsqlExecLatencyP9999               int     `json:"sql.distsql.exec.latency-p99.99"`
	SQLDistsqlExecLatencyP99999              int     `json:"sql.distsql.exec.latency-p99.999"`
	SQLDistsqlExecLatencyInternalMax         int     `json:"sql.distsql.exec.latency.internal-max"`
	SQLDistsqlExecLatencyInternalP50         int     `json:"sql.distsql.exec.latency.internal-p50"`
	SQLDistsqlExecLatencyInternalP75         int     `json:"sql.distsql.exec.latency.internal-p75"`
	SQLDistsqlExecLatencyInternalP90         int     `json:"sql.distsql.exec.latency.internal-p90"`
	SQLDistsqlExecLatencyInternalP99         int     `json:"sql.distsql.exec.latency.internal-p99"`
	SQLDistsqlExecLatencyInternalP999        int     `json:"sql.distsql.exec.latency.internal-p99.9"`
	SQLDistsqlExecLatencyInternalP9999       int     `json:"sql.distsql.exec.latency.internal-p99.99"`
	SQLDistsqlExecLatencyInternalP99999      int     `json:"sql.distsql.exec.latency.internal-p99.999"`
	SQLDistsqlFlowsActive                    int     `json:"sql.distsql.flows.active"`
	SQLDistsqlFlowsQueueWaitMax              int     `json:"sql.distsql.flows.queue_wait-max"`
	SQLDistsqlFlowsQueueWaitP50              int     `json:"sql.distsql.flows.queue_wait-p50"`
	SQLDistsqlFlowsQueueWaitP75              int     `json:"sql.distsql.flows.queue_wait-p75"`
	SQLDistsqlFlowsQueueWaitP90              int     `json:"sql.distsql.flows.queue_wait-p90"`
	SQLDistsqlFlowsQueueWaitP99              int     `json:"sql.distsql.flows.queue_wait-p99"`
	SQLDistsqlFlowsQueueWaitP999             int     `json:"sql.distsql.flows.queue_wait-p99.9"`
	SQLDistsqlFlowsQueueWaitP9999            int     `json:"sql.distsql.flows.queue_wait-p99.99"`
	SQLDistsqlFlowsQueueWaitP99999           int     `json:"sql.distsql.flows.queue_wait-p99.999"`
	SQLDistsqlFlowsQueued                    int     `json:"sql.distsql.flows.queued"`
	SQLDistsqlFlowsTotal                     int     `json:"sql.distsql.flows.total"`
	SQLDistsqlQueriesActive                  int     `json:"sql.distsql.queries.active"`
	SQLDistsqlQueriesTotal                   int     `json:"sql.distsql.queries.total"`
	SQLDistsqlSelectCount                    int     `json:"sql.distsql.select.count"`
	SQLDistsqlSelectCountInternal            int     `json:"sql.distsql.select.count.internal"`
	SQLDistsqlServiceLatencyMax              int     `json:"sql.distsql.service.latency-max"`
	SQLDistsqlServiceLatencyP50              int     `json:"sql.distsql.service.latency-p50"`
	SQLDistsqlServiceLatencyP75              int     `json:"sql.distsql.service.latency-p75"`
	SQLDistsqlServiceLatencyP90              int     `json:"sql.distsql.service.latency-p90"`
	SQLDistsqlServiceLatencyP99              int     `json:"sql.distsql.service.latency-p99"`
	SQLDistsqlServiceLatencyP999             int     `json:"sql.distsql.service.latency-p99.9"`
	SQLDistsqlServiceLatencyP9999            int     `json:"sql.distsql.service.latency-p99.99"`
	SQLDistsqlServiceLatencyP99999           int     `json:"sql.distsql.service.latency-p99.999"`
	SQLDistsqlServiceLatencyInternalMax      int     `json:"sql.distsql.service.latency.internal-max"`
	SQLDistsqlServiceLatencyInternalP50      int     `json:"sql.distsql.service.latency.internal-p50"`
	SQLDistsqlServiceLatencyInternalP75      int     `json:"sql.distsql.service.latency.internal-p75"`
	SQLDistsqlServiceLatencyInternalP90      int     `json:"sql.distsql.service.latency.internal-p90"`
	SQLDistsqlServiceLatencyInternalP99      int     `json:"sql.distsql.service.latency.internal-p99"`
	SQLDistsqlServiceLatencyInternalP999     int     `json:"sql.distsql.service.latency.internal-p99.9"`
	SQLDistsqlServiceLatencyInternalP9999    int     `json:"sql.distsql.service.latency.internal-p99.99"`
	SQLDistsqlServiceLatencyInternalP99999   int     `json:"sql.distsql.service.latency.internal-p99.999"`
	SQLExecLatencyMax                        int     `json:"sql.exec.latency-max"`
	SQLExecLatencyP50                        int     `json:"sql.exec.latency-p50"`
	SQLExecLatencyP75                        int     `json:"sql.exec.latency-p75"`
	SQLExecLatencyP90                        int     `json:"sql.exec.latency-p90"`
	SQLExecLatencyP99                        int     `json:"sql.exec.latency-p99"`
	SQLExecLatencyP999                       int     `json:"sql.exec.latency-p99.9"`
	SQLExecLatencyP9999                      int     `json:"sql.exec.latency-p99.99"`
	SQLExecLatencyP99999                     int     `json:"sql.exec.latency-p99.999"`
	SQLExecLatencyInternalMax                int     `json:"sql.exec.latency.internal-max"`
	SQLExecLatencyInternalP50                int     `json:"sql.exec.latency.internal-p50"`
	SQLExecLatencyInternalP75                int     `json:"sql.exec.latency.internal-p75"`
	SQLExecLatencyInternalP90                int     `json:"sql.exec.latency.internal-p90"`
	SQLExecLatencyInternalP99                int     `json:"sql.exec.latency.internal-p99"`
	SQLExecLatencyInternalP999               int     `json:"sql.exec.latency.internal-p99.9"`
	SQLExecLatencyInternalP9999              int     `json:"sql.exec.latency.internal-p99.99"`
	SQLExecLatencyInternalP99999             int     `json:"sql.exec.latency.internal-p99.999"`
	SQLFailureCount                          int     `json:"sql.failure.count"`
	SQLFailureCountInternal                  int     `json:"sql.failure.count.internal"`
	SQLInsertCount                           int     `json:"sql.insert.count"`
	SQLInsertCountInternal                   int     `json:"sql.insert.count.internal"`
	SQLMemAdminCurrent                       int     `json:"sql.mem.admin.current"`
	SQLMemAdminMaxMax                        int     `json:"sql.mem.admin.max-max"`
	SQLMemAdminMaxP50                        int     `json:"sql.mem.admin.max-p50"`
	SQLMemAdminMaxP75                        int     `json:"sql.mem.admin.max-p75"`
	SQLMemAdminMaxP90                        int     `json:"sql.mem.admin.max-p90"`
	SQLMemAdminMaxP99                        int     `json:"sql.mem.admin.max-p99"`
	SQLMemAdminMaxP999                       int     `json:"sql.mem.admin.max-p99.9"`
	SQLMemAdminMaxP9999                      int     `json:"sql.mem.admin.max-p99.99"`
	SQLMemAdminMaxP99999                     int     `json:"sql.mem.admin.max-p99.999"`
	SQLMemAdminSessionCurrent                int     `json:"sql.mem.admin.session.current"`
	SQLMemAdminSessionMaxMax                 int     `json:"sql.mem.admin.session.max-max"`
	SQLMemAdminSessionMaxP50                 int     `json:"sql.mem.admin.session.max-p50"`
	SQLMemAdminSessionMaxP75                 int     `json:"sql.mem.admin.session.max-p75"`
	SQLMemAdminSessionMaxP90                 int     `json:"sql.mem.admin.session.max-p90"`
	SQLMemAdminSessionMaxP99                 int     `json:"sql.mem.admin.session.max-p99"`
	SQLMemAdminSessionMaxP999                int     `json:"sql.mem.admin.session.max-p99.9"`
	SQLMemAdminSessionMaxP9999               int     `json:"sql.mem.admin.session.max-p99.99"`
	SQLMemAdminSessionMaxP99999              int     `json:"sql.mem.admin.session.max-p99.999"`
	SQLMemAdminTxnCurrent                    int     `json:"sql.mem.admin.txn.current"`
	SQLMemAdminTxnMaxMax                     int     `json:"sql.mem.admin.txn.max-max"`
	SQLMemAdminTxnMaxP50                     int     `json:"sql.mem.admin.txn.max-p50"`
	SQLMemAdminTxnMaxP75                     int     `json:"sql.mem.admin.txn.max-p75"`
	SQLMemAdminTxnMaxP90                     int     `json:"sql.mem.admin.txn.max-p90"`
	SQLMemAdminTxnMaxP99                     int     `json:"sql.mem.admin.txn.max-p99"`
	SQLMemAdminTxnMaxP999                    int     `json:"sql.mem.admin.txn.max-p99.9"`
	SQLMemAdminTxnMaxP9999                   int     `json:"sql.mem.admin.txn.max-p99.99"`
	SQLMemAdminTxnMaxP99999                  int     `json:"sql.mem.admin.txn.max-p99.999"`
	SQLMemConnsCurrent                       int     `json:"sql.mem.conns.current"`
	SQLMemConnsMaxMax                        int     `json:"sql.mem.conns.max-max"`
	SQLMemConnsMaxP50                        int     `json:"sql.mem.conns.max-p50"`
	SQLMemConnsMaxP75                        int     `json:"sql.mem.conns.max-p75"`
	SQLMemConnsMaxP90                        int     `json:"sql.mem.conns.max-p90"`
	SQLMemConnsMaxP99                        int     `json:"sql.mem.conns.max-p99"`
	SQLMemConnsMaxP999                       int     `json:"sql.mem.conns.max-p99.9"`
	SQLMemConnsMaxP9999                      int     `json:"sql.mem.conns.max-p99.99"`
	SQLMemConnsMaxP99999                     int     `json:"sql.mem.conns.max-p99.999"`
	SQLMemConnsSessionCurrent                int     `json:"sql.mem.conns.session.current"`
	SQLMemConnsSessionMaxMax                 int     `json:"sql.mem.conns.session.max-max"`
	SQLMemConnsSessionMaxP50                 int     `json:"sql.mem.conns.session.max-p50"`
	SQLMemConnsSessionMaxP75                 int     `json:"sql.mem.conns.session.max-p75"`
	SQLMemConnsSessionMaxP90                 int     `json:"sql.mem.conns.session.max-p90"`
	SQLMemConnsSessionMaxP99                 int     `json:"sql.mem.conns.session.max-p99"`
	SQLMemConnsSessionMaxP999                int     `json:"sql.mem.conns.session.max-p99.9"`
	SQLMemConnsSessionMaxP9999               int     `json:"sql.mem.conns.session.max-p99.99"`
	SQLMemConnsSessionMaxP99999              int     `json:"sql.mem.conns.session.max-p99.999"`
	SQLMemConnsTxnCurrent                    int     `json:"sql.mem.conns.txn.current"`
	SQLMemConnsTxnMaxMax                     int     `json:"sql.mem.conns.txn.max-max"`
	SQLMemConnsTxnMaxP50                     int     `json:"sql.mem.conns.txn.max-p50"`
	SQLMemConnsTxnMaxP75                     int     `json:"sql.mem.conns.txn.max-p75"`
	SQLMemConnsTxnMaxP90                     int     `json:"sql.mem.conns.txn.max-p90"`
	SQLMemConnsTxnMaxP99                     int     `json:"sql.mem.conns.txn.max-p99"`
	SQLMemConnsTxnMaxP999                    int     `json:"sql.mem.conns.txn.max-p99.9"`
	SQLMemConnsTxnMaxP9999                   int     `json:"sql.mem.conns.txn.max-p99.99"`
	SQLMemConnsTxnMaxP99999                  int     `json:"sql.mem.conns.txn.max-p99.999"`
	SQLMemDistsqlCurrent                     int     `json:"sql.mem.distsql.current"`
	SQLMemDistsqlMaxMax                      int     `json:"sql.mem.distsql.max-max"`
	SQLMemDistsqlMaxP50                      int     `json:"sql.mem.distsql.max-p50"`
	SQLMemDistsqlMaxP75                      int     `json:"sql.mem.distsql.max-p75"`
	SQLMemDistsqlMaxP90                      int     `json:"sql.mem.distsql.max-p90"`
	SQLMemDistsqlMaxP99                      int     `json:"sql.mem.distsql.max-p99"`
	SQLMemDistsqlMaxP999                     int     `json:"sql.mem.distsql.max-p99.9"`
	SQLMemDistsqlMaxP9999                    int     `json:"sql.mem.distsql.max-p99.99"`
	SQLMemDistsqlMaxP99999                   int     `json:"sql.mem.distsql.max-p99.999"`
	SQLMemInternalCurrent                    int     `json:"sql.mem.internal.current"`
	SQLMemInternalMaxMax                     int     `json:"sql.mem.internal.max-max"`
	SQLMemInternalMaxP50                     int     `json:"sql.mem.internal.max-p50"`
	SQLMemInternalMaxP75                     int     `json:"sql.mem.internal.max-p75"`
	SQLMemInternalMaxP90                     int     `json:"sql.mem.internal.max-p90"`
	SQLMemInternalMaxP99                     int     `json:"sql.mem.internal.max-p99"`
	SQLMemInternalMaxP999                    int     `json:"sql.mem.internal.max-p99.9"`
	SQLMemInternalMaxP9999                   int     `json:"sql.mem.internal.max-p99.99"`
	SQLMemInternalMaxP99999                  int     `json:"sql.mem.internal.max-p99.999"`
	SQLMemInternalSessionCurrent             int     `json:"sql.mem.internal.session.current"`
	SQLMemInternalSessionMaxMax              int     `json:"sql.mem.internal.session.max-max"`
	SQLMemInternalSessionMaxP50              int     `json:"sql.mem.internal.session.max-p50"`
	SQLMemInternalSessionMaxP75              int     `json:"sql.mem.internal.session.max-p75"`
	SQLMemInternalSessionMaxP90              int     `json:"sql.mem.internal.session.max-p90"`
	SQLMemInternalSessionMaxP99              int     `json:"sql.mem.internal.session.max-p99"`
	SQLMemInternalSessionMaxP999             int     `json:"sql.mem.internal.session.max-p99.9"`
	SQLMemInternalSessionMaxP9999            int     `json:"sql.mem.internal.session.max-p99.99"`
	SQLMemInternalSessionMaxP99999           int     `json:"sql.mem.internal.session.max-p99.999"`
	SQLMemInternalTxnCurrent                 int     `json:"sql.mem.internal.txn.current"`
	SQLMemInternalTxnMaxMax                  int     `json:"sql.mem.internal.txn.max-max"`
	SQLMemInternalTxnMaxP50                  int     `json:"sql.mem.internal.txn.max-p50"`
	SQLMemInternalTxnMaxP75                  int     `json:"sql.mem.internal.txn.max-p75"`
	SQLMemInternalTxnMaxP90                  int     `json:"sql.mem.internal.txn.max-p90"`
	SQLMemInternalTxnMaxP99                  int     `json:"sql.mem.internal.txn.max-p99"`
	SQLMemInternalTxnMaxP999                 int     `json:"sql.mem.internal.txn.max-p99.9"`
	SQLMemInternalTxnMaxP9999                int     `json:"sql.mem.internal.txn.max-p99.99"`
	SQLMemInternalTxnMaxP99999               int     `json:"sql.mem.internal.txn.max-p99.999"`
	SQLMemSQLCurrent                         int     `json:"sql.mem.sql.current"`
	SQLMemSQLMaxMax                          int     `json:"sql.mem.sql.max-max"`
	SQLMemSQLMaxP50                          int     `json:"sql.mem.sql.max-p50"`
	SQLMemSQLMaxP75                          int     `json:"sql.mem.sql.max-p75"`
	SQLMemSQLMaxP90                          int     `json:"sql.mem.sql.max-p90"`
	SQLMemSQLMaxP99                          int     `json:"sql.mem.sql.max-p99"`
	SQLMemSQLMaxP999                         int     `json:"sql.mem.sql.max-p99.9"`
	SQLMemSQLMaxP9999                        int     `json:"sql.mem.sql.max-p99.99"`
	SQLMemSQLMaxP99999                       int     `json:"sql.mem.sql.max-p99.999"`
	SQLMemSQLSessionCurrent                  int     `json:"sql.mem.sql.session.current"`
	SQLMemSQLSessionMaxMax                   int     `json:"sql.mem.sql.session.max-max"`
	SQLMemSQLSessionMaxP50                   int     `json:"sql.mem.sql.session.max-p50"`
	SQLMemSQLSessionMaxP75                   int     `json:"sql.mem.sql.session.max-p75"`
	SQLMemSQLSessionMaxP90                   int     `json:"sql.mem.sql.session.max-p90"`
	SQLMemSQLSessionMaxP99                   int     `json:"sql.mem.sql.session.max-p99"`
	SQLMemSQLSessionMaxP999                  int     `json:"sql.mem.sql.session.max-p99.9"`
	SQLMemSQLSessionMaxP9999                 int     `json:"sql.mem.sql.session.max-p99.99"`
	SQLMemSQLSessionMaxP99999                int     `json:"sql.mem.sql.session.max-p99.999"`
	SQLMemSQLTxnCurrent                      int     `json:"sql.mem.sql.txn.current"`
	SQLMemSQLTxnMaxMax                       int     `json:"sql.mem.sql.txn.max-max"`
	SQLMemSQLTxnMaxP50                       int     `json:"sql.mem.sql.txn.max-p50"`
	SQLMemSQLTxnMaxP75                       int     `json:"sql.mem.sql.txn.max-p75"`
	SQLMemSQLTxnMaxP90                       int     `json:"sql.mem.sql.txn.max-p90"`
	SQLMemSQLTxnMaxP99                       int     `json:"sql.mem.sql.txn.max-p99"`
	SQLMemSQLTxnMaxP999                      int     `json:"sql.mem.sql.txn.max-p99.9"`
	SQLMemSQLTxnMaxP9999                     int     `json:"sql.mem.sql.txn.max-p99.99"`
	SQLMemSQLTxnMaxP99999                    int     `json:"sql.mem.sql.txn.max-p99.999"`
	SQLMiscCount                             int     `json:"sql.misc.count"`
	SQLMiscCountInternal                     int     `json:"sql.misc.count.internal"`
	SQLOptimizerCount                        int     `json:"sql.optimizer.count"`
	SQLOptimizerCountInternal                int     `json:"sql.optimizer.count.internal"`
	SQLOptimizerFallbackCount                int     `json:"sql.optimizer.fallback.count"`
	SQLOptimizerFallbackCountInternal        int     `json:"sql.optimizer.fallback.count.internal"`
	SQLOptimizerPlanCacheHits                int     `json:"sql.optimizer.plan_cache.hits"`
	SQLOptimizerPlanCacheHitsInternal        int     `json:"sql.optimizer.plan_cache.hits.internal"`
	SQLOptimizerPlanCacheMisses              int     `json:"sql.optimizer.plan_cache.misses"`
	SQLOptimizerPlanCacheMissesInternal      int     `json:"sql.optimizer.plan_cache.misses.internal"`
	SQLQueryCount                            int     `json:"sql.query.count"`
	SQLQueryCountInternal                    int     `json:"sql.query.count.internal"`
	SQLRestartSavepointCount                 int     `json:"sql.restart_savepoint.count"`
	SQLRestartSavepointCountInternal         int     `json:"sql.restart_savepoint.count.internal"`
	SQLRestartSavepointReleaseCount          int     `json:"sql.restart_savepoint.release.count"`
	SQLRestartSavepointReleaseCountInternal  int     `json:"sql.restart_savepoint.release.count.internal"`
	SQLRestartSavepointRollbackCount         int     `json:"sql.restart_savepoint.rollback.count"`
	SQLRestartSavepointRollbackCountInternal int     `json:"sql.restart_savepoint.rollback.count.internal"`
	SQLSavepointCount                        int     `json:"sql.savepoint.count"`
	SQLSavepointCountInternal                int     `json:"sql.savepoint.count.internal"`
	SQLSelectCount                           int     `json:"sql.select.count"`
	SQLSelectCountInternal                   int     `json:"sql.select.count.internal"`
	SQLServiceLatencyMax                     int     `json:"sql.service.latency-max"`
	SQLServiceLatencyP50                     int     `json:"sql.service.latency-p50"`
	SQLServiceLatencyP75                     int     `json:"sql.service.latency-p75"`
	SQLServiceLatencyP90                     int     `json:"sql.service.latency-p90"`
	SQLServiceLatencyP99                     int     `json:"sql.service.latency-p99"`
	SQLServiceLatencyP999                    int     `json:"sql.service.latency-p99.9"`
	SQLServiceLatencyP9999                   int     `json:"sql.service.latency-p99.99"`
	SQLServiceLatencyP99999                  int     `json:"sql.service.latency-p99.999"`
	SQLServiceLatencyInternalMax             int     `json:"sql.service.latency.internal-max"`
	SQLServiceLatencyInternalP50             int     `json:"sql.service.latency.internal-p50"`
	SQLServiceLatencyInternalP75             int     `json:"sql.service.latency.internal-p75"`
	SQLServiceLatencyInternalP90             int     `json:"sql.service.latency.internal-p90"`
	SQLServiceLatencyInternalP99             int     `json:"sql.service.latency.internal-p99"`
	SQLServiceLatencyInternalP999            int     `json:"sql.service.latency.internal-p99.9"`
	SQLServiceLatencyInternalP9999           int     `json:"sql.service.latency.internal-p99.99"`
	SQLServiceLatencyInternalP99999          int     `json:"sql.service.latency.internal-p99.999"`
	SQLTxnAbortCount                         int     `json:"sql.txn.abort.count"`
	SQLTxnAbortCountInternal                 int     `json:"sql.txn.abort.count.internal"`
	SQLTxnBeginCount                         int     `json:"sql.txn.begin.count"`
	SQLTxnBeginCountInternal                 int     `json:"sql.txn.begin.count.internal"`
	SQLTxnCommitCount                        int     `json:"sql.txn.commit.count"`
	SQLTxnCommitCountInternal                int     `json:"sql.txn.commit.count.internal"`
	SQLTxnRollbackCount                      int     `json:"sql.txn.rollback.count"`
	SQLTxnRollbackCountInternal              int     `json:"sql.txn.rollback.count.internal"`
	SQLUpdateCount                           int     `json:"sql.update.count"`
	SQLUpdateCountInternal                   int     `json:"sql.update.count.internal"`
	SysCgoAllocbytes                         int     `json:"sys.cgo.allocbytes"`
	SysCgoTotalbytes                         int     `json:"sys.cgo.totalbytes"`
	SysCgocalls                              int     `json:"sys.cgocalls"`
	SysCPUCombinedPercentNormalized          float64 `json:"sys.cpu.combined.percent-normalized"`
	SysCPUSysNs                              int64   `json:"sys.cpu.sys.ns"`
	SysCPUSysPercent                         float64 `json:"sys.cpu.sys.percent"`
	SysCPUUserNs                             int64   `json:"sys.cpu.user.ns"`
	SysCPUUserPercent                        float64 `json:"sys.cpu.user.percent"`
	SysFdOpen                                int     `json:"sys.fd.open"`
	SysFdSoftlimit                           int     `json:"sys.fd.softlimit"`
	SysGcCount                               int     `json:"sys.gc.count"`
	SysGcPauseNs                             int     `json:"sys.gc.pause.ns"`
	SysGcPausePercent                        float64 `json:"sys.gc.pause.percent"`
	SysGoAllocbytes                          int     `json:"sys.go.allocbytes"`
	SysGoTotalbytes                          int     `json:"sys.go.totalbytes"`
	SysGoroutines                            int     `json:"sys.goroutines"`
	SysHostDiskIoTime                        int64   `json:"sys.host.disk.io.time"`
	SysHostDiskIopsinprogress                int     `json:"sys.host.disk.iopsinprogress"`
	SysHostDiskReadBytes                     int     `json:"sys.host.disk.read.bytes"`
	SysHostDiskReadCount                     int     `json:"sys.host.disk.read.count"`
	SysHostDiskReadTime                      int64   `json:"sys.host.disk.read.time"`
	SysHostDiskWeightedioTime                int64   `json:"sys.host.disk.weightedio.time"`
	SysHostDiskWriteBytes                    int     `json:"sys.host.disk.write.bytes"`
	SysHostDiskWriteCount                    int     `json:"sys.host.disk.write.count"`
	SysHostDiskWriteTime                     int64   `json:"sys.host.disk.write.time"`
	SysHostNetRecvBytes                      int     `json:"sys.host.net.recv.bytes"`
	SysHostNetRecvPackets                    int     `json:"sys.host.net.recv.packets"`
	SysHostNetSendBytes                      int     `json:"sys.host.net.send.bytes"`
	SysHostNetSendPackets                    int     `json:"sys.host.net.send.packets"`
	SysRss                                   int     `json:"sys.rss"`
	SysUptime                                int     `json:"sys.uptime"`
	TimeseriesWriteBytes                     int     `json:"timeseries.write.bytes"`
	TimeseriesWriteErrors                    int     `json:"timeseries.write.errors"`
	TimeseriesWriteSamples                   int     `json:"timeseries.write.samples"`
	TxnAborts                                int     `json:"txn.aborts"`
	TxnAutoretries                           int     `json:"txn.autoretries"`
	TxnCommits                               int     `json:"txn.commits"`
	TxnCommits1PC                            int     `json:"txn.commits1PC"`
	TxnDurationsMax                          int     `json:"txn.durations-max"`
	TxnDurationsP50                          int     `json:"txn.durations-p50"`
	TxnDurationsP75                          int     `json:"txn.durations-p75"`
	TxnDurationsP90                          int     `json:"txn.durations-p90"`
	TxnDurationsP99                          int     `json:"txn.durations-p99"`
	TxnDurationsP999                         int     `json:"txn.durations-p99.9"`
	TxnDurationsP9999                        int     `json:"txn.durations-p99.99"`
	TxnDurationsP99999                       int     `json:"txn.durations-p99.999"`
	TxnRestartsMax                           int     `json:"txn.restarts-max"`
	TxnRestartsP50                           int     `json:"txn.restarts-p50"`
	TxnRestartsP75                           int     `json:"txn.restarts-p75"`
	TxnRestartsP90                           int     `json:"txn.restarts-p90"`
	TxnRestartsP99                           int     `json:"txn.restarts-p99"`
	TxnRestartsP999                          int     `json:"txn.restarts-p99.9"`
	TxnRestartsP9999                         int     `json:"txn.restarts-p99.99"`
	TxnRestartsP99999                        int     `json:"txn.restarts-p99.999"`
	TxnRestartsAsyncwritefailure             int     `json:"txn.restarts.asyncwritefailure"`
	TxnRestartsPossiblereplay                int     `json:"txn.restarts.possiblereplay"`
	TxnRestartsReadwithinuncertainty         int     `json:"txn.restarts.readwithinuncertainty"`
	TxnRestartsSerializable                  int     `json:"txn.restarts.serializable"`
	TxnRestartsTxnaborted                    int     `json:"txn.restarts.txnaborted"`
	TxnRestartsTxnpush                       int     `json:"txn.restarts.txnpush"`
	TxnRestartsUnknown                       int     `json:"txn.restarts.unknown"`
	TxnRestartsWritetooold                   int     `json:"txn.restarts.writetooold"`
	TxnRestartsWritetoooldmulti              int     `json:"txn.restarts.writetoooldmulti"`
}

// AllNodes holds the json response from /_status/nodes
type AllNodes struct {
	Nodes []struct {
		Desc struct {
			NodeID  int `json:"nodeId"`
			Address struct {
				NetworkField string `json:"networkField"`
				AddressField string `json:"addressField"`
			} `json:"address"`
			Attrs struct {
				Attrs []interface{} `json:"attrs"`
			} `json:"attrs"`
			Locality struct {
				Tiers []interface{} `json:"tiers"`
			} `json:"locality"`
			ServerVersion struct {
				MajorVal int `json:"majorVal"`
				MinorVal int `json:"minorVal"`
				Patch    int `json:"patch"`
				Unstable int `json:"unstable"`
			} `json:"ServerVersion"`
			BuildTag        string        `json:"buildTag"`
			StartedAt       string        `json:"startedAt"`
			LocalityAddress []interface{} `json:"localityAddress"`
		} `json:"desc"`
		BuildInfo struct {
			GoVersion       string      `json:"goVersion"`
			Tag             string      `json:"tag"`
			Time            string      `json:"time"`
			Revision        string      `json:"revision"`
			CgoCompiler     string      `json:"cgoCompiler"`
			CgoTargetTriple string      `json:"cgoTargetTriple"`
			Platform        string      `json:"platform"`
			Distribution    string      `json:"distribution"`
			Type            string      `json:"type"`
			Channel         string      `json:"channel"`
			EnvChannel      string      `json:"envChannel"`
			Dependencies    interface{} `json:"dependencies"`
		} `json:"buildInfo"`
		StartedAt     string  `json:"startedAt"`
		UpdatedAt     string  `json:"updatedAt"`
		Metrics       metrics `json:"metrics"`
		StoreStatuses []struct {
			Desc struct {
				StoreID int `json:"storeId"`
				Attrs   struct {
					Attrs []interface{} `json:"attrs"`
				} `json:"attrs"`
				Node struct {
					NodeID  int `json:"nodeId"`
					Address struct {
						NetworkField string `json:"networkField"`
						AddressField string `json:"addressField"`
					} `json:"address"`
					Attrs struct {
						Attrs []interface{} `json:"attrs"`
					} `json:"attrs"`
					Locality struct {
						Tiers []interface{} `json:"tiers"`
					} `json:"locality"`
					ServerVersion struct {
						MajorVal int `json:"majorVal"`
						MinorVal int `json:"minorVal"`
						Patch    int `json:"patch"`
						Unstable int `json:"unstable"`
					} `json:"ServerVersion"`
					BuildTag        string        `json:"buildTag"`
					StartedAt       string        `json:"startedAt"`
					LocalityAddress []interface{} `json:"localityAddress"`
				} `json:"node"`
				Capacity struct {
					Capacity         string  `json:"capacity"`
					Available        string  `json:"available"`
					Used             string  `json:"used"`
					LogicalBytes     string  `json:"logicalBytes"`
					RangeCount       int     `json:"rangeCount"`
					LeaseCount       int     `json:"leaseCount"`
					QueriesPerSecond float64 `json:"queriesPerSecond"`
					WritesPerSecond  float64 `json:"writesPerSecond"`
					BytesPerReplica  struct {
						P10  int `json:"p10"`
						P25  int `json:"p25"`
						P50  int `json:"p50"`
						P75  int `json:"p75"`
						P90  int `json:"p90"`
						PMax int `json:"pMax"`
					} `json:"bytesPerReplica"`
					WritesPerReplica struct {
						P10  int     `json:"p10"`
						P25  int     `json:"p25"`
						P50  int     `json:"p50"`
						P75  float64 `json:"p75"`
						P90  float64 `json:"p90"`
						PMax float64 `json:"pMax"`
					} `json:"writesPerReplica"`
				} `json:"capacity"`
			} `json:"desc"`
			Metrics struct {
				AddsstableApplications                 int     `json:"addsstable.applications"`
				AddsstableCopies                       int     `json:"addsstable.copies"`
				AddsstableProposals                    int     `json:"addsstable.proposals"`
				Capacity                               int     `json:"capacity"`
				CapacityAvailable                      int     `json:"capacity.available"`
				CapacityReserved                       int     `json:"capacity.reserved"`
				CapacityUsed                           int     `json:"capacity.used"`
				CompactorCompactingnanos               int     `json:"compactor.compactingnanos"`
				CompactorCompactionsFailure            int     `json:"compactor.compactions.failure"`
				CompactorCompactionsSuccess            int     `json:"compactor.compactions.success"`
				CompactorSuggestionbytesCompacted      int     `json:"compactor.suggestionbytes.compacted"`
				CompactorSuggestionbytesQueued         int     `json:"compactor.suggestionbytes.queued"`
				CompactorSuggestionbytesSkipped        int     `json:"compactor.suggestionbytes.skipped"`
				FollowerReadsSuccessCount              int     `json:"follower_reads.success_count"`
				Gcbytesage                             int64   `json:"gcbytesage"`
				Intentage                              int     `json:"intentage"`
				Intentbytes                            int     `json:"intentbytes"`
				Intentcount                            int     `json:"intentcount"`
				IntentresolverAsyncThrottled           int     `json:"intentresolver.async.throttled"`
				IntentsAbortAttempts                   int     `json:"intents.abort-attempts"`
				IntentsPoisonAttempts                  int     `json:"intents.poison-attempts"`
				IntentsResolveAttempts                 int     `json:"intents.resolve-attempts"`
				Keybytes                               int     `json:"keybytes"`
				Keycount                               int     `json:"keycount"`
				KvClosedTimestampMaxBehindNanos        int64   `json:"kv.closed_timestamp.max_behind_nanos"`
				KvRangefeedCatchupScanNanos            int     `json:"kv.rangefeed.catchup_scan_nanos"`
				Lastupdatenanos                        int64   `json:"lastupdatenanos"`
				LeasesEpoch                            int     `json:"leases.epoch"`
				LeasesError                            int     `json:"leases.error"`
				LeasesExpiration                       int     `json:"leases.expiration"`
				LeasesSuccess                          int     `json:"leases.success"`
				LeasesTransfersError                   int     `json:"leases.transfers.error"`
				LeasesTransfersSuccess                 int     `json:"leases.transfers.success"`
				Livebytes                              int     `json:"livebytes"`
				Livecount                              int     `json:"livecount"`
				QueueConsistencyPending                int     `json:"queue.consistency.pending"`
				QueueConsistencyProcessFailure         int     `json:"queue.consistency.process.failure"`
				QueueConsistencyProcessSuccess         int     `json:"queue.consistency.process.success"`
				QueueConsistencyProcessingnanos        int64   `json:"queue.consistency.processingnanos"`
				QueueGcInfoAbortspanconsidered         int     `json:"queue.gc.info.abortspanconsidered"`
				QueueGcInfoAbortspangcnum              int     `json:"queue.gc.info.abortspangcnum"`
				QueueGcInfoAbortspanscanned            int     `json:"queue.gc.info.abortspanscanned"`
				QueueGcInfoIntentsconsidered           int     `json:"queue.gc.info.intentsconsidered"`
				QueueGcInfoIntenttxns                  int     `json:"queue.gc.info.intenttxns"`
				QueueGcInfoNumkeysaffected             int     `json:"queue.gc.info.numkeysaffected"`
				QueueGcInfoPushtxn                     int     `json:"queue.gc.info.pushtxn"`
				QueueGcInfoResolvesuccess              int     `json:"queue.gc.info.resolvesuccess"`
				QueueGcInfoResolvetotal                int     `json:"queue.gc.info.resolvetotal"`
				QueueGcInfoTransactionspangcaborted    int     `json:"queue.gc.info.transactionspangcaborted"`
				QueueGcInfoTransactionspangccommitted  int     `json:"queue.gc.info.transactionspangccommitted"`
				QueueGcInfoTransactionspangcpending    int     `json:"queue.gc.info.transactionspangcpending"`
				QueueGcInfoTransactionspanscanned      int     `json:"queue.gc.info.transactionspanscanned"`
				QueueGcPending                         int     `json:"queue.gc.pending"`
				QueueGcProcessFailure                  int     `json:"queue.gc.process.failure"`
				QueueGcProcessSuccess                  int     `json:"queue.gc.process.success"`
				QueueGcProcessingnanos                 int     `json:"queue.gc.processingnanos"`
				QueueMergePending                      int     `json:"queue.merge.pending"`
				QueueMergeProcessFailure               int     `json:"queue.merge.process.failure"`
				QueueMergeProcessSuccess               int     `json:"queue.merge.process.success"`
				QueueMergeProcessingnanos              int64   `json:"queue.merge.processingnanos"`
				QueueMergePurgatory                    int     `json:"queue.merge.purgatory"`
				QueueRaftlogPending                    int     `json:"queue.raftlog.pending"`
				QueueRaftlogProcessFailure             int     `json:"queue.raftlog.process.failure"`
				QueueRaftlogProcessSuccess             int     `json:"queue.raftlog.process.success"`
				QueueRaftlogProcessingnanos            int64   `json:"queue.raftlog.processingnanos"`
				QueueRaftsnapshotPending               int     `json:"queue.raftsnapshot.pending"`
				QueueRaftsnapshotProcessFailure        int     `json:"queue.raftsnapshot.process.failure"`
				QueueRaftsnapshotProcessSuccess        int     `json:"queue.raftsnapshot.process.success"`
				QueueRaftsnapshotProcessingnanos       int     `json:"queue.raftsnapshot.processingnanos"`
				QueueReplicagcPending                  int     `json:"queue.replicagc.pending"`
				QueueReplicagcProcessFailure           int     `json:"queue.replicagc.process.failure"`
				QueueReplicagcProcessSuccess           int     `json:"queue.replicagc.process.success"`
				QueueReplicagcProcessingnanos          int     `json:"queue.replicagc.processingnanos"`
				QueueReplicagcRemovereplica            int     `json:"queue.replicagc.removereplica"`
				QueueReplicateAddreplica               int     `json:"queue.replicate.addreplica"`
				QueueReplicatePending                  int     `json:"queue.replicate.pending"`
				QueueReplicateProcessFailure           int     `json:"queue.replicate.process.failure"`
				QueueReplicateProcessSuccess           int     `json:"queue.replicate.process.success"`
				QueueReplicateProcessingnanos          int     `json:"queue.replicate.processingnanos"`
				QueueReplicatePurgatory                int     `json:"queue.replicate.purgatory"`
				QueueReplicateRebalancereplica         int     `json:"queue.replicate.rebalancereplica"`
				QueueReplicateRemovedeadreplica        int     `json:"queue.replicate.removedeadreplica"`
				QueueReplicateRemovereplica            int     `json:"queue.replicate.removereplica"`
				QueueReplicateTransferlease            int     `json:"queue.replicate.transferlease"`
				QueueSplitPending                      int     `json:"queue.split.pending"`
				QueueSplitProcessFailure               int     `json:"queue.split.process.failure"`
				QueueSplitProcessSuccess               int     `json:"queue.split.process.success"`
				QueueSplitProcessingnanos              int64   `json:"queue.split.processingnanos"`
				QueueSplitPurgatory                    int     `json:"queue.split.purgatory"`
				QueueTsmaintenancePending              int     `json:"queue.tsmaintenance.pending"`
				QueueTsmaintenanceProcessFailure       int     `json:"queue.tsmaintenance.process.failure"`
				QueueTsmaintenanceProcessSuccess       int     `json:"queue.tsmaintenance.process.success"`
				QueueTsmaintenanceProcessingnanos      int     `json:"queue.tsmaintenance.processingnanos"`
				RaftCommandsapplied                    int     `json:"raft.commandsapplied"`
				RaftEnqueuedPending                    int     `json:"raft.enqueued.pending"`
				RaftEntrycacheAccesses                 int     `json:"raft.entrycache.accesses"`
				RaftEntrycacheBytes                    int     `json:"raft.entrycache.bytes"`
				RaftEntrycacheHits                     int     `json:"raft.entrycache.hits"`
				RaftEntrycacheSize                     int     `json:"raft.entrycache.size"`
				RaftHeartbeatsPending                  int     `json:"raft.heartbeats.pending"`
				RaftProcessApplycommittedLatencyMax    int     `json:"raft.process.applycommitted.latency-max"`
				RaftProcessApplycommittedLatencyP50    int     `json:"raft.process.applycommitted.latency-p50"`
				RaftProcessApplycommittedLatencyP75    int     `json:"raft.process.applycommitted.latency-p75"`
				RaftProcessApplycommittedLatencyP90    int     `json:"raft.process.applycommitted.latency-p90"`
				RaftProcessApplycommittedLatencyP99    int     `json:"raft.process.applycommitted.latency-p99"`
				RaftProcessApplycommittedLatencyP999   int     `json:"raft.process.applycommitted.latency-p99.9"`
				RaftProcessApplycommittedLatencyP9999  int     `json:"raft.process.applycommitted.latency-p99.99"`
				RaftProcessApplycommittedLatencyP99999 int     `json:"raft.process.applycommitted.latency-p99.999"`
				RaftProcessCommandcommitLatencyMax     int     `json:"raft.process.commandcommit.latency-max"`
				RaftProcessCommandcommitLatencyP50     int     `json:"raft.process.commandcommit.latency-p50"`
				RaftProcessCommandcommitLatencyP75     int     `json:"raft.process.commandcommit.latency-p75"`
				RaftProcessCommandcommitLatencyP90     int     `json:"raft.process.commandcommit.latency-p90"`
				RaftProcessCommandcommitLatencyP99     int     `json:"raft.process.commandcommit.latency-p99"`
				RaftProcessCommandcommitLatencyP999    int     `json:"raft.process.commandcommit.latency-p99.9"`
				RaftProcessCommandcommitLatencyP9999   int     `json:"raft.process.commandcommit.latency-p99.99"`
				RaftProcessCommandcommitLatencyP99999  int     `json:"raft.process.commandcommit.latency-p99.999"`
				RaftProcessHandlereadyLatencyMax       int     `json:"raft.process.handleready.latency-max"`
				RaftProcessHandlereadyLatencyP50       int     `json:"raft.process.handleready.latency-p50"`
				RaftProcessHandlereadyLatencyP75       int     `json:"raft.process.handleready.latency-p75"`
				RaftProcessHandlereadyLatencyP90       int     `json:"raft.process.handleready.latency-p90"`
				RaftProcessHandlereadyLatencyP99       int     `json:"raft.process.handleready.latency-p99"`
				RaftProcessHandlereadyLatencyP999      int     `json:"raft.process.handleready.latency-p99.9"`
				RaftProcessHandlereadyLatencyP9999     int     `json:"raft.process.handleready.latency-p99.99"`
				RaftProcessHandlereadyLatencyP99999    int     `json:"raft.process.handleready.latency-p99.999"`
				RaftProcessLogcommitLatencyMax         int     `json:"raft.process.logcommit.latency-max"`
				RaftProcessLogcommitLatencyP50         int     `json:"raft.process.logcommit.latency-p50"`
				RaftProcessLogcommitLatencyP75         int     `json:"raft.process.logcommit.latency-p75"`
				RaftProcessLogcommitLatencyP90         int     `json:"raft.process.logcommit.latency-p90"`
				RaftProcessLogcommitLatencyP99         int     `json:"raft.process.logcommit.latency-p99"`
				RaftProcessLogcommitLatencyP999        int     `json:"raft.process.logcommit.latency-p99.9"`
				RaftProcessLogcommitLatencyP9999       int     `json:"raft.process.logcommit.latency-p99.99"`
				RaftProcessLogcommitLatencyP99999      int     `json:"raft.process.logcommit.latency-p99.999"`
				RaftProcessTickingnanos                int64   `json:"raft.process.tickingnanos"`
				RaftProcessWorkingnanos                int64   `json:"raft.process.workingnanos"`
				RaftRcvdApp                            int     `json:"raft.rcvd.app"`
				RaftRcvdAppresp                        int     `json:"raft.rcvd.appresp"`
				RaftRcvdDropped                        int     `json:"raft.rcvd.dropped"`
				RaftRcvdHeartbeat                      int     `json:"raft.rcvd.heartbeat"`
				RaftRcvdHeartbeatresp                  int     `json:"raft.rcvd.heartbeatresp"`
				RaftRcvdPrevote                        int     `json:"raft.rcvd.prevote"`
				RaftRcvdPrevoteresp                    int     `json:"raft.rcvd.prevoteresp"`
				RaftRcvdProp                           int     `json:"raft.rcvd.prop"`
				RaftRcvdSnap                           int     `json:"raft.rcvd.snap"`
				RaftRcvdTimeoutnow                     int     `json:"raft.rcvd.timeoutnow"`
				RaftRcvdTransferleader                 int     `json:"raft.rcvd.transferleader"`
				RaftRcvdVote                           int     `json:"raft.rcvd.vote"`
				RaftRcvdVoteresp                       int     `json:"raft.rcvd.voteresp"`
				RaftTicks                              int     `json:"raft.ticks"`
				RaftlogBehind                          int     `json:"raftlog.behind"`
				RaftlogTruncated                       int     `json:"raftlog.truncated"`
				RangeAdds                              int     `json:"range.adds"`
				RangeMerges                            int     `json:"range.merges"`
				RangeRaftleadertransfers               int     `json:"range.raftleadertransfers"`
				RangeRemoves                           int     `json:"range.removes"`
				RangeSnapshotsGenerated                int     `json:"range.snapshots.generated"`
				RangeSnapshotsNormalApplied            int     `json:"range.snapshots.normal-applied"`
				RangeSnapshotsPreemptiveApplied        int     `json:"range.snapshots.preemptive-applied"`
				RangeSplits                            int     `json:"range.splits"`
				Ranges                                 int     `json:"ranges"`
				RangesOverreplicated                   int     `json:"ranges.overreplicated"`
				RangesUnavailable                      int     `json:"ranges.unavailable"`
				RangesUnderreplicated                  int     `json:"ranges.underreplicated"`
				RebalancingLeaseTransfers              int     `json:"rebalancing.lease.transfers"`
				RebalancingQueriespersecond            float64 `json:"rebalancing.queriespersecond"`
				RebalancingRangeRebalances             int     `json:"rebalancing.range.rebalances"`
				RebalancingWritespersecond             float64 `json:"rebalancing.writespersecond"`
				Replicas                               int     `json:"replicas"`
				ReplicasLeaders                        int     `json:"replicas.leaders"`
				ReplicasLeadersNotLeaseholders         int     `json:"replicas.leaders_not_leaseholders"`
				ReplicasLeaseholders                   int     `json:"replicas.leaseholders"`
				ReplicasQuiescent                      int     `json:"replicas.quiescent"`
				ReplicasReserved                       int     `json:"replicas.reserved"`
				RequestsBackpressureSplit              int     `json:"requests.backpressure.split"`
				RequestsSlowLatch                      int     `json:"requests.slow.latch"`
				RequestsSlowLease                      int     `json:"requests.slow.lease"`
				RequestsSlowRaft                       int     `json:"requests.slow.raft"`
				RocksdbBlockCacheHits                  int     `json:"rocksdb.block.cache.hits"`
				RocksdbBlockCacheMisses                int     `json:"rocksdb.block.cache.misses"`
				RocksdbBlockCachePinnedUsage           int     `json:"rocksdb.block.cache.pinned-usage"`
				RocksdbBlockCacheUsage                 int     `json:"rocksdb.block.cache.usage"`
				RocksdbBloomFilterPrefixChecked        int     `json:"rocksdb.bloom.filter.prefix.checked"`
				RocksdbBloomFilterPrefixUseful         int     `json:"rocksdb.bloom.filter.prefix.useful"`
				RocksdbCompactions                     int     `json:"rocksdb.compactions"`
				RocksdbEncryptionAlgorithm             int     `json:"rocksdb.encryption.algorithm"`
				RocksdbFlushes                         int     `json:"rocksdb.flushes"`
				RocksdbMemtableTotalSize               int     `json:"rocksdb.memtable.total-size"`
				RocksdbNumSstables                     int     `json:"rocksdb.num-sstables"`
				RocksdbReadAmplification               int     `json:"rocksdb.read-amplification"`
				RocksdbTableReadersMemEstimate         int     `json:"rocksdb.table-readers-mem-estimate"`
				Sysbytes                               int     `json:"sysbytes"`
				Syscount                               int     `json:"syscount"`
				Totalbytes                             int     `json:"totalbytes"`
				TscacheSklReadPages                    int     `json:"tscache.skl.read.pages"`
				TscacheSklReadRotations                int     `json:"tscache.skl.read.rotations"`
				TscacheSklWritePages                   int     `json:"tscache.skl.write.pages"`
				TscacheSklWriteRotations               int     `json:"tscache.skl.write.rotations"`
				TxnwaitqueueDeadlocksTotal             int     `json:"txnwaitqueue.deadlocks_total"`
				TxnwaitqueuePusheeWaiting              int     `json:"txnwaitqueue.pushee.waiting"`
				TxnwaitqueuePusherSlow                 int     `json:"txnwaitqueue.pusher.slow"`
				TxnwaitqueuePusherWaitTimeMax          int     `json:"txnwaitqueue.pusher.wait_time-max"`
				TxnwaitqueuePusherWaitTimeP50          int     `json:"txnwaitqueue.pusher.wait_time-p50"`
				TxnwaitqueuePusherWaitTimeP75          int     `json:"txnwaitqueue.pusher.wait_time-p75"`
				TxnwaitqueuePusherWaitTimeP90          int     `json:"txnwaitqueue.pusher.wait_time-p90"`
				TxnwaitqueuePusherWaitTimeP99          int     `json:"txnwaitqueue.pusher.wait_time-p99"`
				TxnwaitqueuePusherWaitTimeP999         int     `json:"txnwaitqueue.pusher.wait_time-p99.9"`
				TxnwaitqueuePusherWaitTimeP9999        int     `json:"txnwaitqueue.pusher.wait_time-p99.99"`
				TxnwaitqueuePusherWaitTimeP99999       int     `json:"txnwaitqueue.pusher.wait_time-p99.999"`
				TxnwaitqueuePusherWaiting              int     `json:"txnwaitqueue.pusher.waiting"`
				TxnwaitqueueQueryWaitTimeMax           int     `json:"txnwaitqueue.query.wait_time-max"`
				TxnwaitqueueQueryWaitTimeP50           int     `json:"txnwaitqueue.query.wait_time-p50"`
				TxnwaitqueueQueryWaitTimeP75           int     `json:"txnwaitqueue.query.wait_time-p75"`
				TxnwaitqueueQueryWaitTimeP90           int     `json:"txnwaitqueue.query.wait_time-p90"`
				TxnwaitqueueQueryWaitTimeP99           int     `json:"txnwaitqueue.query.wait_time-p99"`
				TxnwaitqueueQueryWaitTimeP999          int     `json:"txnwaitqueue.query.wait_time-p99.9"`
				TxnwaitqueueQueryWaitTimeP9999         int     `json:"txnwaitqueue.query.wait_time-p99.99"`
				TxnwaitqueueQueryWaitTimeP99999        int     `json:"txnwaitqueue.query.wait_time-p99.999"`
				TxnwaitqueueQueryWaiting               int     `json:"txnwaitqueue.query.waiting"`
				Valbytes                               int     `json:"valbytes"`
				Valcount                               int     `json:"valcount"`
			} `json:"metrics"`
		} `json:"storeStatuses"`
		Args      []string `json:"args"`
		Env       []string `json:"env"`
		Latencies struct {
		} `json:"latencies"`
		Activity struct {
			Num1 struct {
				Incoming string `json:"incoming"`
				Outgoing string `json:"outgoing"`
				Latency  string `json:"latency"`
			} `json:"1"`
			Num2 struct {
				Incoming string `json:"incoming"`
				Outgoing string `json:"outgoing"`
				Latency  string `json:"latency"`
			} `json:"2"`
			Num3 struct {
				Incoming string `json:"incoming"`
				Outgoing string `json:"outgoing"`
				Latency  string `json:"latency"`
			} `json:"3"`
		} `json:"activity"`
		TotalSystemMemory string `json:"totalSystemMemory"`
		NumCpus           int    `json:"numCpus"`
	} `json:"nodes"`
}

// ClusterRaft holds the json response from /_status/raft
type ClusterRaft struct {
	/*Ranges struct {
		Num1 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"expiration"`
								Replica struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"deprecatedStartStasis"`
								ProposedTs struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"expiration"`
						Replica struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"deprecatedStartStasis"`
						ProposedTs struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int     `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Replica struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"deprecatedStartStasis"`
							ProposedTs struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int         `json:"state"`
						Liveness interface{} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"1"`
		Num2 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"expiration"`
								Replica struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"deprecatedStartStasis"`
								ProposedTs struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"expiration"`
						Replica struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"deprecatedStartStasis"`
						ProposedTs struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int     `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Replica struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"deprecatedStartStasis"`
							ProposedTs struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int         `json:"state"`
						Liveness interface{} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"2"`
		Num3 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"3"`
		Num4 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int    `json:"nextReplicaId"`
								Generation    string `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"4"`
		Num5 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"5"`
		Num6 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  int     `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"6"`
		Num7 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"7"`
		Num8 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"8"`
		Num9 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int     `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"9"`
		Num10 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"10"`
		Num11 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  int     `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"11"`
		Num12 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"12"`
		Num13 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"13"`
		Num14 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"14"`
		Num15 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"15"`
		Num16 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"16"`
		Num17 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"17"`
		Num18 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"18"`
		Num19 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"19"`
		Num20 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int    `json:"nextReplicaId"`
								Generation    string `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"20"`
		Num21 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"21"`
		Num35 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int    `json:"nextReplicaId"`
								Generation    string `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"35"`
		Num37 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int    `json:"nextReplicaId"`
								Generation    string `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"37"`
		Num40 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int    `json:"nextReplicaId"`
								Generation    string `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int     `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"40"`
		Num42 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int `json:"queriesPerSecond"`
						WritesPerSecond  int `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"42"`
		Num43 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"43"`
		Num44 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
							Num1 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"1"`
							Num2 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"2"`
							Num3 struct {
								Match           string `json:"match"`
								Next            string `json:"next"`
								State           string `json:"state"`
								Paused          bool   `json:"paused"`
								PendingSnapshot string `json:"pendingSnapshot"`
							} `json:"3"`
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int    `json:"nextReplicaId"`
								Generation    string `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond float64 `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"44"`
		Num61 struct {
			RangeID string        `json:"rangeId"`
			Errors  []interface{} `json:"errors"`
			Nodes   []struct {
				NodeID int `json:"nodeId"`
				Range  struct {
					Span struct {
						StartKey string `json:"startKey"`
						EndKey   string `json:"endKey"`
					} `json:"span"`
					RaftState struct {
						ReplicaID string `json:"replicaId"`
						HardState struct {
							Term   string `json:"term"`
							Vote   string `json:"vote"`
							Commit string `json:"commit"`
						} `json:"hardState"`
						Lead     string `json:"lead"`
						State    string `json:"state"`
						Applied  string `json:"applied"`
						Progress struct {
						} `json:"progress"`
						LeadTransferee string `json:"leadTransferee"`
					} `json:"raftState"`
					State struct {
						State struct {
							RaftAppliedIndex  string `json:"raftAppliedIndex"`
							LeaseAppliedIndex string `json:"leaseAppliedIndex"`
							Desc              struct {
								RangeID  string      `json:"rangeId"`
								StartKey interface{} `json:"startKey"`
								EndKey   interface{} `json:"endKey"`
								Replicas []struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replicas"`
								NextReplicaID int         `json:"nextReplicaId"`
								Generation    interface{} `json:"generation"`
							} `json:"desc"`
							Lease struct {
								Start struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"start"`
								Expiration interface{} `json:"expiration"`
								Replica    struct {
									NodeID    int `json:"nodeId"`
									StoreID   int `json:"storeId"`
									ReplicaID int `json:"replicaId"`
								} `json:"replica"`
								DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
								ProposedTs            struct {
									WallTime string `json:"wallTime"`
									Logical  int    `json:"logical"`
								} `json:"proposedTs"`
								Epoch    string `json:"epoch"`
								Sequence string `json:"sequence"`
							} `json:"lease"`
							TruncatedState struct {
								Index string `json:"index"`
								Term  string `json:"term"`
							} `json:"truncatedState"`
							GcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"gcThreshold"`
							Stats struct {
								ContainsEstimates bool   `json:"containsEstimates"`
								LastUpdateNanos   string `json:"lastUpdateNanos"`
								IntentAge         string `json:"intentAge"`
								GcBytesAge        string `json:"gcBytesAge"`
								LiveBytes         string `json:"liveBytes"`
								LiveCount         string `json:"liveCount"`
								KeyBytes          string `json:"keyBytes"`
								KeyCount          string `json:"keyCount"`
								ValBytes          string `json:"valBytes"`
								ValCount          string `json:"valCount"`
								IntentBytes       string `json:"intentBytes"`
								IntentCount       string `json:"intentCount"`
								SysBytes          string `json:"sysBytes"`
								SysCount          string `json:"sysCount"`
							} `json:"stats"`
							TxnSpanGcThreshold struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"txnSpanGcThreshold"`
							UsingAppliedStateKey bool `json:"usingAppliedStateKey"`
						} `json:"state"`
						LastIndex                string `json:"lastIndex"`
						NumPending               string `json:"numPending"`
						NumDropped               string `json:"numDropped"`
						RaftLogSize              string `json:"raftLogSize"`
						RaftLogSizeTrusted       bool   `json:"raftLogSizeTrusted"`
						ApproximateProposalQuota string `json:"approximateProposalQuota"`
						RangeMaxBytes            string `json:"rangeMaxBytes"`
						NewestClosedTimestamp    struct {
							NodeID          int `json:"nodeId"`
							ClosedTimestamp struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"closedTimestamp"`
							Mlai string `json:"mlai"`
						} `json:"newestClosedTimestamp"`
						ActiveClosedTimestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"activeClosedTimestamp"`
					} `json:"state"`
					SourceNodeID  int    `json:"sourceNodeId"`
					SourceStoreID int    `json:"sourceStoreId"`
					ErrorMessage  string `json:"errorMessage"`
					LeaseHistory  []struct {
						Start struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"start"`
						Expiration interface{} `json:"expiration"`
						Replica    struct {
							NodeID    int `json:"nodeId"`
							StoreID   int `json:"storeId"`
							ReplicaID int `json:"replicaId"`
						} `json:"replica"`
						DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
						ProposedTs            struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"proposedTs"`
						Epoch    string `json:"epoch"`
						Sequence string `json:"sequence"`
					} `json:"leaseHistory"`
					Problems struct {
						Unavailable            bool `json:"unavailable"`
						LeaderNotLeaseHolder   bool `json:"leaderNotLeaseHolder"`
						NoRaftLeader           bool `json:"noRaftLeader"`
						Underreplicated        bool `json:"underreplicated"`
						Overreplicated         bool `json:"overreplicated"`
						NoLease                bool `json:"noLease"`
						QuiescentEqualsTicking bool `json:"quiescentEqualsTicking"`
						RaftLogTooLarge        bool `json:"raftLogTooLarge"`
					} `json:"problems"`
					Stats struct {
						QueriesPerSecond int     `json:"queriesPerSecond"`
						WritesPerSecond  float64 `json:"writesPerSecond"`
					} `json:"stats"`
					LatchesLocal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesLocal"`
					LatchesGlobal struct {
						ReadCount  string `json:"readCount"`
						WriteCount string `json:"writeCount"`
					} `json:"latchesGlobal"`
					LeaseStatus struct {
						Lease struct {
							Start struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"start"`
							Expiration interface{} `json:"expiration"`
							Replica    struct {
								NodeID    int `json:"nodeId"`
								StoreID   int `json:"storeId"`
								ReplicaID int `json:"replicaId"`
							} `json:"replica"`
							DeprecatedStartStasis interface{} `json:"deprecatedStartStasis"`
							ProposedTs            struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"proposedTs"`
							Epoch    string `json:"epoch"`
							Sequence string `json:"sequence"`
						} `json:"lease"`
						Timestamp struct {
							WallTime string `json:"wallTime"`
							Logical  int    `json:"logical"`
						} `json:"timestamp"`
						State    int `json:"state"`
						Liveness struct {
							NodeID     int    `json:"nodeId"`
							Epoch      string `json:"epoch"`
							Expiration struct {
								WallTime string `json:"wallTime"`
								Logical  int    `json:"logical"`
							} `json:"expiration"`
							Draining        bool `json:"draining"`
							Decommissioning bool `json:"decommissioning"`
						} `json:"liveness"`
					} `json:"leaseStatus"`
					Quiescent bool `json:"quiescent"`
					Ticking   bool `json:"ticking"`
				} `json:"range"`
			} `json:"nodes"`
		} `json:"61"`
	} `json:"ranges"`*/
	Ranges []interface{} `json:"ranges"`
	Errors []interface{} `json:"errors"`
}
