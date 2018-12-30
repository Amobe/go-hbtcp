package admin

import (
	"encoding/json"
	"sync"
	"sync/atomic"
)

var pOnce sync.Once
var gProcStats *ProcStats

type ProcStats struct {
	ClinetOnlineConn    uint64 `json:"clientOnlineConn"`
	ServerInPktAcc      uint64 `json:"serverInPktAcc"`
	ServerOutPktAcc     uint64 `json:"serverOutPktAcc"`
	ReguestQueueSize    uint32 `json:"reguestQueueSize"`
	ReguestQueuePadding uint32 `json:"reguestQueuePadding"`
}

func (s ProcStats) ToJson() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		return []byte(err.Error())
	}
	return data
}

func (s *ProcStats) IncClientConn() {
	atomic.AddUint64(&s.ClinetOnlineConn, 1)
}

func (s *ProcStats) DecClientConn() {
	atomic.AddUint64(&s.ClinetOnlineConn, ^uint64(0))
}

func (s *ProcStats) IncServerInPktAcc() {
	atomic.AddUint64(&s.ServerInPktAcc, 1)
}

func (s *ProcStats) IncServerOutPktAcc() {
	atomic.AddUint64(&s.ServerOutPktAcc, 1)
}

func (s *ProcStats) SetReguestQueueSize(size int) {
	atomic.StoreUint32(&s.ReguestQueueSize, uint32(size))
}

func (s *ProcStats) IncReguestQueuePadding() {
	atomic.AddUint32(&s.ReguestQueuePadding, 1)
}

func (s *ProcStats) DecReguestQueuePadding() {
	atomic.AddUint32(&s.ReguestQueuePadding, ^uint32(0))
}

func createProcStatsInstance() {
	if gProcStats == nil {
		gProcStats = new(ProcStats)
	}
}

// GetProcStatsInstance return a global instance of the process statistic object
func GetProcStatsInstance() *ProcStats {
	if gProcStats == nil {
		pOnce.Do(createProcStatsInstance)
	}
	return gProcStats
}
