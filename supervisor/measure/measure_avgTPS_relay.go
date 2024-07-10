package measure

import (
	"blockEmulator/message"
	"time"
)

// to test average TPS in this system
type TestModule_avgTPS_Relay struct {
	epochID      int
	excutedTxNum []float64   // record how many excuted txs in a epoch, maybe the cross shard tx will be calculated as a 0.5 tx
	startTime    []time.Time // record when the epoch starts
	endTime      []time.Time // record when the epoch ends
}

func NewTestModule_avgTPS_Relay() *TestModule_avgTPS_Relay {
	return &TestModule_avgTPS_Relay{
		epochID:      -1,
		excutedTxNum: make([]float64, 0),
		startTime:    make([]time.Time, 0),
		endTime:      make([]time.Time, 0),
	}
}

func (tat *TestModule_avgTPS_Relay) OutputMetricName() string {
	return "Average_TPS"
}

// add the number of excuted txs, and change the time records
func (tat *TestModule_avgTPS_Relay) UpdateMeasureRecord(b *message.BlockInfoMsg) {
	if b.BlockBodyLength == 0 { // empty block
		return
	}

	epochid := b.Epoch
	earliestTime := b.ProposeTime
	latestTime := b.CommitTime
	r1TxNum := len(b.Relay1Txs)
	r2TxNum := len(b.Relay2Txs)

	// extend
	for tat.epochID < epochid {
		tat.excutedTxNum = append(tat.excutedTxNum, 0)
		tat.startTime = append(tat.startTime, time.Time{})
		tat.endTime = append(tat.endTime, time.Time{})
		tat.epochID++
	}

	// modify the local epoch data
	tat.excutedTxNum[epochid] += float64(r1TxNum+r2TxNum) / 2
	tat.excutedTxNum[epochid] += float64(len(b.InterShardTxs))

	if tat.startTime[epochid].IsZero() || tat.startTime[epochid].After(earliestTime) {
		tat.startTime[epochid] = earliestTime
	}
	if tat.endTime[epochid].IsZero() || latestTime.After(tat.endTime[epochid]) {
		tat.endTime[epochid] = latestTime
	}
}

func (tat *TestModule_avgTPS_Relay) HandleExtraMessage([]byte) {}

// output the average TPS
func (tat *TestModule_avgTPS_Relay) OutputRecord() (perEpochTPS []float64, totalTPS float64) {
	perEpochTPS = make([]float64, tat.epochID+1)
	totalTxNum := 0.0
	eTime := time.Now()
	lTime := time.Time{}
	for eid, exTxNum := range tat.excutedTxNum {
		timeGap := tat.endTime[eid].Sub(tat.startTime[eid]).Seconds()
		perEpochTPS[eid] = exTxNum / timeGap
		totalTxNum += exTxNum
		if eTime.After(tat.startTime[eid]) {
			eTime = tat.startTime[eid]
		}
		if tat.endTime[eid].After(lTime) {
			lTime = tat.endTime[eid]
		}
	}
	totalTPS = totalTxNum / (lTime.Sub(eTime).Seconds())
	return
}
