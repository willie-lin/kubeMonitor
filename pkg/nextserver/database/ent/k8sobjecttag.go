// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sobjecttag"
)

// K8sObjectTag is the model entity for the K8sObjectTag schema.
type K8sObjectTag struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// K8sObjectId holds the value of the "k8sObjectId" field.
	K8sObjectId uint `json:"k8sObjectId,omitempty"`
	// K8sLabelId holds the value of the "k8sLabelId" field.
	K8sLabelId uint `json:"k8sLabelId,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*K8sObjectTag) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case k8sobjecttag.FieldID, k8sobjecttag.FieldK8sObjectId, k8sobjecttag.FieldK8sLabelId:
			values[i] = new(sql.NullInt64)
		case k8sobjecttag.FieldCreatedAt, k8sobjecttag.FieldUpdatedAt, k8sobjecttag.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type K8sObjectTag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sObjectTag fields.
func (kot *K8sObjectTag) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case k8sobjecttag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kot.ID = uint(value.Int64)
		case k8sobjecttag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kot.CreatedAt = value.Time
			}
		case k8sobjecttag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kot.UpdatedAt = value.Time
			}
		case k8sobjecttag.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kot.DeletedAt = value.Time
			}
		case k8sobjecttag.FieldK8sObjectId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sObjectId", values[i])
			} else if value.Valid {
				kot.K8sObjectId = uint(value.Int64)
			}
		case k8sobjecttag.FieldK8sLabelId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sLabelId", values[i])
			} else if value.Valid {
				kot.K8sLabelId = uint(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this K8sObjectTag.
// Note that you need to call K8sObjectTag.Unwrap() before calling this method if this K8sObjectTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (kot *K8sObjectTag) Update() *K8sObjectTagUpdateOne {
	return (&K8sObjectTagClient{config: kot.config}).UpdateOne(kot)
}

// Unwrap unwraps the K8sObjectTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kot *K8sObjectTag) Unwrap() *K8sObjectTag {
	tx, ok := kot.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sObjectTag is not a transactional entity")
	}
	kot.config.driver = tx.drv
	return kot
}

// String implements the fmt.Stringer.
func (kot *K8sObjectTag) String() string {
	var builder strings.Builder
	builder.WriteString("K8sObjectTag(")
	builder.WriteString(fmt.Sprintf("id=%v", kot.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(kot.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(kot.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(kot.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", k8sObjectId=")
	builder.WriteString(fmt.Sprintf("%v", kot.K8sObjectId))
	builder.WriteString(", k8sLabelId=")
	builder.WriteString(fmt.Sprintf("%v", kot.K8sLabelId))
	builder.WriteByte(')')
	return builder.String()
}

// K8sObjectTags is a parsable slice of K8sObjectTag.
type K8sObjectTags []*K8sObjectTag

func (kot K8sObjectTags) config(cfg config) {
	for _i := range kot {
		kot[_i].config = cfg
	}
}