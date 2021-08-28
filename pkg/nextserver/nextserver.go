package main

import (
	"github.com/dgraph-io/ristretto"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent"
	"google.golang.org/grpc/keepalive"
	"sync"
	"time"
)

const (
	AppName           = "NextServer"
	AppDescription    = "NextServer for KubeMonitor Monitoring System"
	NextServerVersion = "1.0"
	ConfigFilename    = "nextserver.yaml"
)

var kaep = keepalive.EnforcementPolicy{

	MinTime:             5 * time.Second,
	PermitWithoutStream: true,
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle: 15 * time.Second,
	Time:              5 * time.Second,
	Timeout:           1 * time.Second,
}

// ServerConfig server config
type ServerConfig struct {
	BindAddress     string `json:"bind_address"`
	AgentListenPort int    `json:"agent_listen_port"`
	ApiPort         int    `json:"api_port"`
}

// DatabaseConfig database config
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"prot"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	Type     string `json:"type"`
	SslMode  string `json:"ssl_mode"`
}

// TLSConfig tls config
type TLSConfig struct {
	Use      bool   `json:"use"`
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
}

// BasicRuleConfig basic rule config
type BasicRuleConfig struct {
	NodeCpuLoad    float64
	NodeMemoryFree float64
	NodeDiskFree   float64
}

// Config config
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	TLS       TLSConfig
	BasicRule BasicRuleConfig
}

type IncidentItem struct {
}

type Metric struct {
}

type NextServer struct {
	sync.RWMutex

	config *Config
	db     *ent.Client
	dbLock map[string]*sync.RWMutex

	agentMap map[string]*ent.Agent
	nodeMap  map[string]*ent.Node

	cache *ristretto.Cache

	serverStartTs         time.Time
	metricSaveCounter     uint64
	metricSaveCounterLock sync.RWMutex

	incidentMap map[string]*IncidentItem

	metricChannel chan Metric
}
