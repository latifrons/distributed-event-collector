package service

import (
	"errors"
	"github.com/bluele/gcache"
	"github.com/latifrons/distributed-event-collector/pbgo/dec"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sort"
	"time"
)

type EventRecord struct {
	EventType string
	Timestamp int64
}

type SpanCollector struct {
	cache     gcache.Cache
	CacheSize int
}

func (s *SpanCollector) InitDefault() {
	s.cache = gcache.New(s.CacheSize).LRU().
		LoaderFunc(func(key interface{}) (interface{}, error) {
			m := cmap.New[int64]()
			return &m, nil
		}).Build()
}

func (s *SpanCollector) Report(rr *dec.ReportRequest) (err error) {
	item, err := s.cache.Get(rr.Owner)
	if err != nil {
		return
	}
	m, ok := item.(*cmap.ConcurrentMap[string, int64])
	if !ok {
		err = errors.New("invalid cache item type")
		return
	}
	m.Set(rr.EventType, time.Now().UnixNano())
	return
}

// for given key, get the first timestamp and last tinmstamp in the cache.
func (s *SpanCollector) GetSpan(owner string) (first int64, last int64, err error) {
	item, err := s.cache.GetIFPresent(owner)
	if err != nil {
		return
	}
	m, ok := item.(*cmap.ConcurrentMap[string, int64])
	if !ok {
		err = errors.New("invalid cache item type")
		return
	}
	for _, key := range m.Keys() {
		value, ok := m.Get(key)
		if !ok {
			continue
		}
		if first == 0 || value < first {
			first = value
		}
		if last == 0 || value > last {
			last = value
		}
	}
	return
}

func (s *SpanCollector) GetEventFlow(owner string) (events *dec.GetEventFlowResponse, err error) {
	item, err := s.cache.GetIFPresent(owner)
	if err != nil {
		return
	}
	m, ok := item.(*cmap.ConcurrentMap[string, int64])
	if !ok {
		err = errors.New("invalid cache item type")
		return
	}
	eventFlowList := []dec.EventFlow{}
	for key, value := range m.Items() {
		eventFlowList = append(eventFlowList, dec.EventFlow{
			EventType:     key,
			TimestampNano: value,
		})
	}
	return
}

func (s *SpanCollector) GetEventStatistics(samples int64) (resp *dec.GetEventStatisticsResponse, err error) {
	owners := s.cache.Keys(false)
	// randomly select samples owners
	if int64(len(owners)) > samples {
		owners = owners[:samples]
	}

	eventStatistics := map[string]*dec.EventStatistic{}
	for _, eventOwner := range owners {
		item, err := s.cache.GetIFPresent(eventOwner)
		if err != nil {
			return nil, err
		}
		m, ok := item.(*cmap.ConcurrentMap[string, int64])
		if !ok {
			return nil, errors.New("invalid cache item type")
		}
		arr := []EventRecord{}
		for eventType, timestamp := range m.Items() {
			arr = append(arr, EventRecord{
				EventType: eventType,
				Timestamp: timestamp,
			})
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Timestamp < arr[j].Timestamp
		})

		t := int64(0)
		for _, record := range arr {
			if t == 0 {
				t = record.Timestamp
				continue
			}
			eventType := record.EventType
			duration := record.Timestamp - t
			if _, ok := eventStatistics[eventType]; !ok {
				eventStatistics[eventType] = &dec.EventStatistic{
					EventType: eventType,
					Count:     0,
				}
			}
			v := eventStatistics[eventType]
			v.Count += 1
			v.SumTimeNano += duration
			t = record.Timestamp
		}
	}

	resp = &dec.GetEventStatisticsResponse{}
	resp.EventStatisticList = []*dec.EventStatistic{}
	for _, v := range eventStatistics {
		resp.EventStatisticList = append(resp.EventStatisticList, v)
	}
	return
}
