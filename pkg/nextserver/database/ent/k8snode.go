// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8snode"
)

// K8sNode is the model entity for the K8sNode schema.
type K8sNode struct {
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
func (*K8sNode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case k8snode.FieldID, k8snode.FieldK8sClusterId, k8snode.FieldK8sObjectId:
			values[i] = new(sql.NullInt64)
		case k8snode.FieldName:
			values[i] = new(sql.NullString)
		case k8snode.FieldCreatedAt, k8snode.FieldUpdatedAt, k8snode.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type K8sNode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sNode fields.
func (kn *K8sNode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case k8snode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kn.ID = uint(value.Int64)
		case k8snode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kn.CreatedAt = value.Time
			}
		case k8snode.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kn.UpdatedAt = value.Time
			}
		case k8snode.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kn.DeletedAt = value.Time
			}
		case k8snode.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				kn.Name = value.String
			}
		case k8snode.FieldK8sClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sClusterId", values[i])
			} else if value.Valid {
				kn.K8sClusterId = uint(value.Int64)
			}
		case k8snode.FieldK8sObjectId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sObjectId", values[i])
			} else if value.Valid {
				kn.K8sObjectId = uint(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this K8sNode.
// Note that you need to call K8sNode.Unwrap() before calling this method if this K8sNode
// was returned from a transaction, and the transaction was committed or rolled back.
func (kn *K8sNode) Update() *K8sNodeUpdateOne {
	return (&K8sNodeClient{config: kn.config}).UpdateOne(kn)
}

// Unwrap unwraps the K8sNode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kn *K8sNode) Unwrap() *K8sNode {
	tx, ok := kn.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sNode is not a transactional entity")
	}
	kn.config.driver = tx.drv
	return kn
}

// String implements the fmt.Stringer.
func (kn *K8sNode) String() string {
	var builder strings.Builder
	builder.WriteString("K8sNode(")
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

// K8sNodes is a parsable slice of K8sNode.
type K8sNodes []*K8sNode

func (kn K8sNodes) config(cfg config) {
	for _i := range kn {
		kn[_i].config = cfg
	}
}
