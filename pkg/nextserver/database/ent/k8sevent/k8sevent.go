// Code generated by entc, DO NOT EDIT.

package k8sevent

import (
	"time"
)

const (
	// Label holds the string label denoting the k8sevent type in the database.
	Label = "k8s_event"
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
	// FieldNamespaceId holds the string denoting the namespaceid field in the database.
	FieldNamespaceId = "namespace_id"
	// FieldNodeId holds the string denoting the nodeid field in the database.
	FieldNodeId = "node_id"
	// FieldContainerId holds the string denoting the containerid field in the database.
	FieldContainerId = "container_id"
	// FieldPodId holds the string denoting the podid field in the database.
	FieldPodId = "pod_id"
	// Table holds the table name of the k8sevent in the database.
	Table = "k8sEvents"
)

// Columns holds all SQL columns for k8sevent fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTs,
	FieldValue,
	FieldEndpointId,
	FieldTypeId,
	FieldNameId,
	FieldLabelId,
	FieldClusterId,
	FieldNamespaceId,
	FieldNodeId,
	FieldContainerId,
	FieldPodId,
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
