package nextserver

import (
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent"
	"time"
)

type IncidentItem struct {
	ClusterId   uint
	NodeId      uint
	ProcessId   uint
	ContainerId uint
	PodId       uint
	TargetType  string
	Target      string
	Value       float64
	Condition   float64
	EventName   string
	ReportedTs  time.Time
	DetectedTs  time.Time
}

type Incident struct {
	incidentMap map[string][]*IncidentItem
}

func (s *NextServer) InitBasicRuleChecker() {
	s.CheckNodeBasicIncident(s.metricChannel)
}

func (s *NextServer) CheckNodeBasicIncident(nodeMetricChan chan ent.Metric) {
	//gaugeType := Controller.FindMetricType()
}
