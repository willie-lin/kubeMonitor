// Code generated by entc, DO NOT EDIT.

package k8spod

import (
	"time"
)

const (
	// Label holds the string label denoting the k8spod type in the database.
	Label = "k8s_pod"
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
	// FieldQos holds the string denoting the qos field in the database.
	FieldQos = "qos"
	// FieldK8sClusterId holds the string denoting the k8sclusterid field in the database.
	FieldK8sClusterId = "k8s_cluster_id"
	// FieldK8sNamespaceId holds the string denoting the k8snamespaceid field in the database.
	FieldK8sNamespaceId = "k8s_namespace_id"
	// FieldK8sObjectId holds the string denoting the k8sobjectid field in the database.
	FieldK8sObjectId = "k8s_object_id"
	// Table holds the table name of the k8spod in the database.
	Table = "k8sPods"
)

// Columns holds all SQL columns for k8spod fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldQos,
	FieldK8sClusterId,
	FieldK8sNamespaceId,
	FieldK8sObjectId,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// QosValidator is a validator for the "qos" field. It is called by the builders before save.
	QosValidator func(string) error
)