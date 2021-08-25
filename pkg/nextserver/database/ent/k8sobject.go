// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sobject"
)

// K8sObject is the model entity for the K8sObject schema.
type K8sObject struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// K8sClusterId holds the value of the "k8sClusterId" field.
	K8sClusterId uint `json:"k8sClusterId,omitempty"`
	// ApiVersion holds the value of the "apiVersion" field.
	ApiVersion string `json:"apiVersion,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata []string `json:"metadata,omitempty"`
	// Spec holds the value of the "spec" field.
	Spec []string `json:"spec,omitempty"`
	// Status holds the value of the "status" field.
	Status []string `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*K8sObject) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case k8sobject.FieldMetadata, k8sobject.FieldSpec, k8sobject.FieldStatus:
			values[i] = new([]byte)
		case k8sobject.FieldID, k8sobject.FieldK8sClusterId:
			values[i] = new(sql.NullInt64)
		case k8sobject.FieldApiVersion, k8sobject.FieldKind, k8sobject.FieldName:
			values[i] = new(sql.NullString)
		case k8sobject.FieldCreatedAt, k8sobject.FieldUpdatedAt, k8sobject.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type K8sObject", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sObject fields.
func (ko *K8sObject) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case k8sobject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ko.ID = uint(value.Int64)
		case k8sobject.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ko.CreatedAt = value.Time
			}
		case k8sobject.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ko.UpdatedAt = value.Time
			}
		case k8sobject.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ko.DeletedAt = value.Time
			}
		case k8sobject.FieldK8sClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sClusterId", values[i])
			} else if value.Valid {
				ko.K8sClusterId = uint(value.Int64)
			}
		case k8sobject.FieldApiVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field apiVersion", values[i])
			} else if value.Valid {
				ko.ApiVersion = value.String
			}
		case k8sobject.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				ko.Kind = value.String
			}
		case k8sobject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ko.Name = value.String
			}
		case k8sobject.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ko.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case k8sobject.FieldSpec:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field spec", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ko.Spec); err != nil {
					return fmt.Errorf("unmarshal field spec: %w", err)
				}
			}
		case k8sobject.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ko.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this K8sObject.
// Note that you need to call K8sObject.Unwrap() before calling this method if this K8sObject
// was returned from a transaction, and the transaction was committed or rolled back.
func (ko *K8sObject) Update() *K8sObjectUpdateOne {
	return (&K8sObjectClient{config: ko.config}).UpdateOne(ko)
}

// Unwrap unwraps the K8sObject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ko *K8sObject) Unwrap() *K8sObject {
	tx, ok := ko.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sObject is not a transactional entity")
	}
	ko.config.driver = tx.drv
	return ko
}

// String implements the fmt.Stringer.
func (ko *K8sObject) String() string {
	var builder strings.Builder
	builder.WriteString("K8sObject(")
	builder.WriteString(fmt.Sprintf("id=%v", ko.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(ko.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ko.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(ko.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", k8sClusterId=")
	builder.WriteString(fmt.Sprintf("%v", ko.K8sClusterId))
	builder.WriteString(", apiVersion=")
	builder.WriteString(ko.ApiVersion)
	builder.WriteString(", kind=")
	builder.WriteString(ko.Kind)
	builder.WriteString(", name=")
	builder.WriteString(ko.Name)
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", ko.Metadata))
	builder.WriteString(", spec=")
	builder.WriteString(fmt.Sprintf("%v", ko.Spec))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ko.Status))
	builder.WriteByte(')')
	return builder.String()
}

// K8sObjects is a parsable slice of K8sObject.
type K8sObjects []*K8sObject

func (ko K8sObjects) config(cfg config) {
	for _i := range ko {
		ko[_i].config = cfg
	}
}
