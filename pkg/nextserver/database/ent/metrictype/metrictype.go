// Code generated by entc, DO NOT EDIT.

package metrictype

import (
	"time"
)

const (
	// Label holds the string label denoting the metrictype type in the database.
	Label = "metric_type"
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
	// EdgeMetricNames holds the string denoting the metricnames edge name in mutations.
	EdgeMetricNames = "metricNames"
	// Table holds the table name of the metrictype in the database.
	Table = "metricTypes"
	// MetricNamesTable is the table that holds the metricNames relation/edge.
	MetricNamesTable = "metricNames"
	// MetricNamesInverseTable is the table name for the MetricName entity.
	// It exists in this package in order to avoid circular dependency with the "metricname" package.
	MetricNamesInverseTable = "metricNames"
	// MetricNamesColumn is the table column denoting the metricNames relation/edge.
	MetricNamesColumn = "metric_type_metric_names"
)

// Columns holds all SQL columns for metrictype fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
