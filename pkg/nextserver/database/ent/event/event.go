// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTs holds the string denoting the ts field in the database.
	FieldTs = "ts"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldAcked holds the string denoting the acked field in the database.
	FieldAcked = "acked"
	// FieldAckedTs holds the string denoting the ackedts field in the database.
	FieldAckedTs = "acked_ts"
	// FieldEndpointId holds the string denoting the endpointid field in the database.
	FieldEndpointId = "endpoint_id"
	// FieldTypeId holds the string denoting the typeid field in the database.
	FieldTypeId = "type_id"
	// FieldNameId holds the string denoting the nameid field in the database.
	FieldNameId = "name_id"
	// FieldLabelId holds the string denoting the labelid field in the database.
	FieldLabelId = "label_id"
	// FieldClusterId holds the string denoting the clusterid field in the database.
	FieldClusterId = "cluster_id"
	// FieldAgentId holds the string denoting the agentid field in the database.
	FieldAgentId = "agent_id"
	// FieldNodeId holds the string denoting the nodeid field in the database.
	FieldNodeId = "node_id"
	// FieldProcesId holds the string denoting the procesid field in the database.
	FieldProcesId = "proces_id"
	// FieldContainerId holds the string denoting the containerid field in the database.
	FieldContainerId = "container_id"
	// FieldPodId holds the string denoting the podid field in the database.
	FieldPodId = "pod_id"
	// EdgeMetricNameEvents holds the string denoting the metricname_events edge name in mutations.
	EdgeMetricNameEvents = "MetricName_events"
	// EdgeMetricLabelEvents holds the string denoting the metriclabel_events edge name in mutations.
	EdgeMetricLabelEvents = "MetricLabel_events"
	// EdgeMetricEndpointEvents holds the string denoting the metricendpoint_events edge name in mutations.
	EdgeMetricEndpointEvents = "MetricEndpoint_events"
	// Table holds the table name of the event in the database.
	Table = "events"
	// MetricNameEventsTable is the table that holds the MetricName_events relation/edge.
	MetricNameEventsTable = "events"
	// MetricNameEventsInverseTable is the table name for the MetricName entity.
	// It exists in this package in order to avoid circular dependency with the "metricname" package.
	MetricNameEventsInverseTable = "metricNames"
	// MetricNameEventsColumn is the table column denoting the MetricName_events relation/edge.
	MetricNameEventsColumn = "metric_name_events"
	// MetricLabelEventsTable is the table that holds the MetricLabel_events relation/edge.
	MetricLabelEventsTable = "events"
	// MetricLabelEventsInverseTable is the table name for the MetricLabel entity.
	// It exists in this package in order to avoid circular dependency with the "metriclabel" package.
	MetricLabelEventsInverseTable = "metricLabels"
	// MetricLabelEventsColumn is the table column denoting the MetricLabel_events relation/edge.
	MetricLabelEventsColumn = "metric_label_events"
	// MetricEndpointEventsTable is the table that holds the MetricEndpoint_events relation/edge.
	MetricEndpointEventsTable = "events"
	// MetricEndpointEventsInverseTable is the table name for the MetricEndpoint entity.
	// It exists in this package in order to avoid circular dependency with the "metricendpoint" package.
	MetricEndpointEventsInverseTable = "metricEndpoints"
	// MetricEndpointEventsColumn is the table column denoting the MetricEndpoint_events relation/edge.
	MetricEndpointEventsColumn = "metric_endpoint_events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTs,
	FieldValue,
	FieldAcked,
	FieldAckedTs,
	FieldEndpointId,
	FieldTypeId,
	FieldNameId,
	FieldLabelId,
	FieldClusterId,
	FieldAgentId,
	FieldNodeId,
	FieldProcesId,
	FieldContainerId,
	FieldPodId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"metric_endpoint_events",
	"metric_label_events",
	"metric_name_events",
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
)
