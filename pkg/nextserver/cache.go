package nextserver

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
)

func (s *NextServer) initCache() (*ristretto.Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cache: %v\n", err)
	}
	s.cache = cache
	return cache, nil
}

//func (s *NextServer) getMetricEndpoint(endpoint string) {
//	var metricEndpoint = new(*ent.MetricEndpoint)
//	value, found := s.cache.Get(fmt.Sprintf("ME_%s", endpoint))
//	if !found {
//		metricEndpoint := s.FindMetricEndpoint(endpoint)
//		s.cache.Set(fmt.Sprintf("ME_%s", endpoint), metricEndpoint, 1)
//
//		return metricEndpoint
//	}
//
//	metricEndpoint := value.(ent.MetricEndpoint)
//
//}
