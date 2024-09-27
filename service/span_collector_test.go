package service

import (
	"fmt"
	"github.com/latifrons/distributed-event-collector/pbgo/dec"
	"testing"
	"time"
)

func TestSpanCollector_GetEventStatistics(t *testing.T) {
	s := Prepare(t)
	statistics, err := s.GetEventStatistics(10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range statistics.EventStatisticList {
		t.Logf("event: %s, count: %d, nano: %d", v.EventType, v.Count, v.SumTimeNano)
	}
}

func TestSpanCollector_GetSpan(t *testing.T) {
	s := Prepare(t)
	first, last, err := s.GetSpan("owner-1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("first: %d, last: %d", first, last)

	first, last, err = s.GetSpan("owner-xx")
	if err != nil {
		return
	}
	t.Logf("first: %d, last: %d", first, last)
}

func Prepare(t *testing.T) (s SpanCollector) {
	s = SpanCollector{
		CacheSize: 100,
	}
	s.InitDefault()

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			err := s.Report(&dec.ReportRequest{
				Owner:     fmt.Sprintf("owner-%d", a),
				EventType: fmt.Sprintf("event-type-%d", b),
			})
			time.Sleep(time.Millisecond)
			if err != nil {
				t.Error(err)
				return
			}
		}
	}

	return s
}
