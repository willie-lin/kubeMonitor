// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8snamespace"
)

// K8sNamespace is the model entity for the K8sNamespace schema.
type K8sNamespace struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// K8sClusterId holds the value of the "k8sClusterId" field.
	K8sClusterId uint `json:"k8sClusterId,omitempty"`
	// K8sObjectId holds the value of the "k8sObjectId" field.
	K8sObjectId uint `json:"k8sObjectId,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*K8sNamespace) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case k8snamespace.FieldID, k8snamespace.FieldK8sClusterId, k8snamespace.FieldK8sObjectId:
			values[i] = new(sql.NullInt64)
		case k8snamespace.FieldName:
			values[i] = new(sql.NullString)
		case k8snamespace.FieldCreatedAt, k8snamespace.FieldUpdatedAt, k8snamespace.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type K8sNamespace", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sNamespace fields.
func (kn *K8sNamespace) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case k8snamespace.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kn.ID = uint(value.Int64)
		case k8snamespace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kn.CreatedAt = value.Time
			}
		case k8snamespace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kn.UpdatedAt = value.Time
			}
		case k8snamespace.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kn.DeletedAt = value.Time
			}
		case k8snamespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				kn.Name = value.String
			}
		case k8snamespace.FieldK8sClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sClusterId", values[i])
			} else if value.Valid {
				kn.K8sClusterId = uint(value.Int64)
			}
		case k8snamespace.FieldK8sObjectId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sObjectId", values[i])
			} else if value.Valid {
				kn.K8sObjectId = uint(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this K8sNamespace.
// Note that you need to call K8sNamespace.Unwrap() before calling this method if this K8sNamespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (kn *K8sNamespace) Update() *K8sNamespaceUpdateOne {
	return (&K8sNamespaceClient{config: kn.config}).UpdateOne(kn)
}

// Unwrap unwraps the K8sNamespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kn *K8sNamespace) Unwrap() *K8sNamespace {
	tx, ok := kn.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sNamespace is not a transactional entity")
	}
	kn.config.driver = tx.drv
	return kn
}

// String implements the fmt.Stringer.
func (kn *K8sNamespace) String() string {
	var builder strings.Builder
	builder.WriteString("K8sNamespace(")
	builder.WriteString(fmt.Sprintf("id=%v", kn.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(kn.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(kn.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(kn.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(kn.Name)
	builder.WriteString(", k8sClusterId=")
	builder.WriteString(fmt.Sprintf("%v", kn.K8sClusterId))
	builder.WriteString(", k8sObjectId=")
	builder.WriteString(fmt.Sprintf("%v", kn.K8sObjectId))
	builder.WriteByte(')')
	return builder.String()
}

// K8sNamespaces is a parsable slice of K8sNamespace.
type K8sNamespaces []*K8sNamespace

func (kn K8sNamespaces) config(cfg config) {
	for _i := range kn {
		kn[_i].config = cfg
	}
}