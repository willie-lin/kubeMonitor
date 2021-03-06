// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sstatefulset"
)

// K8sStatefulSet is the model entity for the K8sStatefulSet schema.
type K8sStatefulSet struct {
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
	// K8sNamespaceId holds the value of the "k8sNamespaceId" field.
	K8sNamespaceId uint `json:"k8sNamespaceId,omitempty"`
	// K8sObjectId holds the value of the "k8sObjectId" field.
	K8sObjectId uint `json:"k8sObjectId,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*K8sStatefulSet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case k8sstatefulset.FieldID, k8sstatefulset.FieldK8sClusterId, k8sstatefulset.FieldK8sNamespaceId, k8sstatefulset.FieldK8sObjectId:
			values[i] = new(sql.NullInt64)
		case k8sstatefulset.FieldName:
			values[i] = new(sql.NullString)
		case k8sstatefulset.FieldCreatedAt, k8sstatefulset.FieldUpdatedAt, k8sstatefulset.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type K8sStatefulSet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sStatefulSet fields.
func (kss *K8sStatefulSet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case k8sstatefulset.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			kss.ID = uint(value.Int64)
		case k8sstatefulset.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				kss.CreatedAt = value.Time
			}
		case k8sstatefulset.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				kss.UpdatedAt = value.Time
			}
		case k8sstatefulset.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				kss.DeletedAt = value.Time
			}
		case k8sstatefulset.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				kss.Name = value.String
			}
		case k8sstatefulset.FieldK8sClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sClusterId", values[i])
			} else if value.Valid {
				kss.K8sClusterId = uint(value.Int64)
			}
		case k8sstatefulset.FieldK8sNamespaceId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sNamespaceId", values[i])
			} else if value.Valid {
				kss.K8sNamespaceId = uint(value.Int64)
			}
		case k8sstatefulset.FieldK8sObjectId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sObjectId", values[i])
			} else if value.Valid {
				kss.K8sObjectId = uint(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this K8sStatefulSet.
// Note that you need to call K8sStatefulSet.Unwrap() before calling this method if this K8sStatefulSet
// was returned from a transaction, and the transaction was committed or rolled back.
func (kss *K8sStatefulSet) Update() *K8sStatefulSetUpdateOne {
	return (&K8sStatefulSetClient{config: kss.config}).UpdateOne(kss)
}

// Unwrap unwraps the K8sStatefulSet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (kss *K8sStatefulSet) Unwrap() *K8sStatefulSet {
	tx, ok := kss.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sStatefulSet is not a transactional entity")
	}
	kss.config.driver = tx.drv
	return kss
}

// String implements the fmt.Stringer.
func (kss *K8sStatefulSet) String() string {
	var builder strings.Builder
	builder.WriteString("K8sStatefulSet(")
	builder.WriteString(fmt.Sprintf("id=%v", kss.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(kss.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(kss.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(kss.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(kss.Name)
	builder.WriteString(", k8sClusterId=")
	builder.WriteString(fmt.Sprintf("%v", kss.K8sClusterId))
	builder.WriteString(", k8sNamespaceId=")
	builder.WriteString(fmt.Sprintf("%v", kss.K8sNamespaceId))
	builder.WriteString(", k8sObjectId=")
	builder.WriteString(fmt.Sprintf("%v", kss.K8sObjectId))
	builder.WriteByte(')')
	return builder.String()
}

// K8sStatefulSets is a parsable slice of K8sStatefulSet.
type K8sStatefulSets []*K8sStatefulSet

func (kss K8sStatefulSets) config(cfg config) {
	for _i := range kss {
		kss[_i].config = cfg
	}
}
