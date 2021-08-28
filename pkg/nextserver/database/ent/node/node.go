// Code generated by entc, DO NOT EDIT.

package node

import (
	"time"
)

const (
	// Label holds the string label denoting the node type in the database.
	Label = "node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldIpv4 holds the string denoting the ipv4 field in the database.
	FieldIpv4 = "ipv4"
	// FieldIpv6 holds the string denoting the ipv6 field in the database.
	FieldIpv6 = "ipv6"
	// FieldPublicIpv4 holds the string denoting the public_ipv4 field in the database.
	FieldPublicIpv4 = "public_ipv4"
	// FieldPublicIpv6 holds the string denoting the public_ipv6 field in the database.
	FieldPublicIpv6 = "public_ipv6"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldPlatformFamily holds the string denoting the platformfamily field in the database.
	FieldPlatformFamily = "platform_family"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldAgentId holds the string denoting the agentid field in the database.
	FieldAgentId = "agent_id"
	// FieldClusterId holds the string denoting the clusterid field in the database.
	FieldClusterId = "cluster_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeContainers holds the string denoting the containers edge name in mutations.
	EdgeContainers = "containers"
	// EdgeProcess holds the string denoting the process edge name in mutations.
	EdgeProcess = "process"
	// Table holds the table name of the node in the database.
	Table = "nodes"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "nodes"
	// OwnerInverseTable is the table name for the Cluster entity.
	// It exists in this package in order to avoid circular dependency with the "cluster" package.
	OwnerInverseTable = "clusters"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "cluster_nodes"
	// ContainersTable is the table that holds the containers relation/edge.
	ContainersTable = "containers"
	// ContainersInverseTable is the table name for the Container entity.
	// It exists in this package in order to avoid circular dependency with the "container" package.
	ContainersInverseTable = "containers"
	// ContainersColumn is the table column denoting the containers relation/edge.
	ContainersColumn = "node_containers"
	// ProcessTable is the table that holds the process relation/edge.
	ProcessTable = "process"
	// ProcessInverseTable is the table name for the Proces entity.
	// It exists in this package in order to avoid circular dependency with the "proces" package.
	ProcessInverseTable = "process"
	// ProcessColumn is the table column denoting the process relation/edge.
	ProcessColumn = "node_process"
)

// Columns holds all SQL columns for node fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldHost,
	FieldIpv4,
	FieldIpv6,
	FieldPublicIpv4,
	FieldPublicIpv6,
	FieldOs,
	FieldPlatform,
	FieldPlatformFamily,
	FieldInfo,
	FieldUUID,
	FieldDescription,
	FieldDisabled,
	FieldAgentId,
	FieldClusterId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "nodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"agent_node",
	"cluster_nodes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() time.Time
	// UpdateDefaultDeletedAt holds the default value on update for the "deleted_at" field.
	UpdateDefaultDeletedAt func() time.Time
	// HostValidator is a validator for the "host" field. It is called by the builders before save.
	HostValidator func(string) error
	// Ipv4Validator is a validator for the "ipv4" field. It is called by the builders before save.
	Ipv4Validator func(string) error
	// Ipv6Validator is a validator for the "ipv6" field. It is called by the builders before save.
	Ipv6Validator func(string) error
	// PublicIpv4Validator is a validator for the "public_ipv4" field. It is called by the builders before save.
	PublicIpv4Validator func(string) error
	// PublicIpv6Validator is a validator for the "public_ipv6" field. It is called by the builders before save.
	PublicIpv6Validator func(string) error
	// OsValidator is a validator for the "os" field. It is called by the builders before save.
	OsValidator func(string) error
	// PlatformValidator is a validator for the "platform" field. It is called by the builders before save.
	PlatformValidator func(string) error
	// PlatformFamilyValidator is a validator for the "platformFamily" field. It is called by the builders before save.
	PlatformFamilyValidator func(string) error
)
