// Code generated by entc, DO NOT EDIT.

package k8slabel

import (
	"time"
)

const (
	// Label holds the string label denoting the k8slabel type in the database.
	Label = "k8s_label"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldLabel holds the string denoting the label field in the database.
	FieldLabel = "label"
	// FieldK8sObjectId holds the string denoting the k8sobjectid field in the database.
	FieldK8sObjectId = "k8s_object_id"
	// Table holds the table name of the k8slabel in the database.
	Table = "k8sLabels"
)

// Columns holds all SQL columns for k8slabel fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldLabel,
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
	// LabelValidator is a validator for the "label" field. It is called by the builders before save.
	LabelValidator func(string) error
)
