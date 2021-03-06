// Code generated by entc, DO NOT EDIT.

package metricname

import (
	"time"
)

const (
	// Label holds the string label denoting the metricname type in the database.
	Label = "metric_name"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHelp holds the string denoting the help field in the database.
	FieldHelp = "help"
	// FieldTypeId holds the string denoting the typeid field in the database.
	FieldTypeId = "type_id"
	// EdgeMetrics holds the string denoting the metrics edge name in mutations.
	EdgeMetrics = "metrics"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// EdgeOwners holds the string denoting the owners edge name in mutations.
	EdgeOwners = "owners"
	// Table holds the table name of the metricname in the database.
	Table = "metricNames"
	// MetricsTable is the table that holds the metrics relation/edge.
	MetricsTable = "metrics"
	// MetricsInverseTable is the table name for the Metric entity.
	// It exists in this package in order to avoid circular dependency with the "metric" package.
	MetricsInverseTable = "metrics"
	// MetricsColumn is the table column denoting the metrics relation/edge.
	MetricsColumn = "metric_name_metrics"
	// EventsTable is the table that holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "metric_name_events"
	// OwnersTable is the table that holds the owners relation/edge.
	OwnersTable = "metricNames"
	// OwnersInverseTable is the table name for the MetricType entity.
	// It exists in this package in order to avoid circular dependency with the "metrictype" package.
	OwnersInverseTable = "metricTypes"
	// OwnersColumn is the table column denoting the owners relation/edge.
	OwnersColumn = "metric_type_metric_names"
)

// Columns holds all SQL columns for metricname fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldHelp,
	FieldTypeId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "metricNames"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"metric_type_metric_names",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
