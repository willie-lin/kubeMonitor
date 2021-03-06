// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8smetric"
)

// K8sMetric is the model entity for the K8sMetric schema.
type K8sMetric struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Ts holds the value of the "ts" field.
	Ts time.Time `json:"ts,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// EndpointId holds the value of the "endpointId" field.
	EndpointId uint `json:"endpointId,omitempty"`
	// TypeId holds the value of the "typeId" field.
	TypeId uint `json:"typeId,omitempty"`
	// NameId holds the value of the "nameId" field.
	NameId uint `json:"nameId,omitempty"`
	// LabelId holds the value of the "labelId" field.
	LabelId uint `json:"labelId,omitempty"`
	// K8sClusterId holds the value of the "k8sClusterId" field.
	K8sClusterId uint `json:"k8sClusterId,omitempty"`
	// K8sNodeId holds the value of the "k8sNodeId" field.
	K8sNodeId uint `json:"k8sNodeId,omitempty"`
	// K8sNamespaceId holds the value of the "k8sNamespaceId" field.
	K8sNamespaceId uint `json:"k8sNamespaceId,omitempty"`
	// K8sContainerId holds the value of the "k8sContainerId" field.
	K8sContainerId uint `json:"k8sContainerId,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*K8sMetric) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case k8smetric.FieldValue:
			values[i] = new(sql.NullFloat64)
		case k8smetric.FieldID, k8smetric.FieldEndpointId, k8smetric.FieldTypeId, k8smetric.FieldNameId, k8smetric.FieldLabelId, k8smetric.FieldK8sClusterId, k8smetric.FieldK8sNodeId, k8smetric.FieldK8sNamespaceId, k8smetric.FieldK8sContainerId:
			values[i] = new(sql.NullInt64)
		case k8smetric.FieldCreatedAt, k8smetric.FieldUpdatedAt, k8smetric.FieldDeletedAt, k8smetric.FieldTs:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type K8sMetric", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sMetric fields.
func (km *K8sMetric) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case k8smetric.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			km.ID = int(value.Int64)
		case k8smetric.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				km.CreatedAt = value.Time
			}
		case k8smetric.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				km.UpdatedAt = value.Time
			}
		case k8smetric.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				km.DeletedAt = value.Time
			}
		case k8smetric.FieldTs:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ts", values[i])
			} else if value.Valid {
				km.Ts = value.Time
			}
		case k8smetric.FieldValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				km.Value = value.Float64
			}
		case k8smetric.FieldEndpointId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field endpointId", values[i])
			} else if value.Valid {
				km.EndpointId = uint(value.Int64)
			}
		case k8smetric.FieldTypeId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field typeId", values[i])
			} else if value.Valid {
				km.TypeId = uint(value.Int64)
			}
		case k8smetric.FieldNameId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nameId", values[i])
			} else if value.Valid {
				km.NameId = uint(value.Int64)
			}
		case k8smetric.FieldLabelId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field labelId", values[i])
			} else if value.Valid {
				km.LabelId = uint(value.Int64)
			}
		case k8smetric.FieldK8sClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sClusterId", values[i])
			} else if value.Valid {
				km.K8sClusterId = uint(value.Int64)
			}
		case k8smetric.FieldK8sNodeId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sNodeId", values[i])
			} else if value.Valid {
				km.K8sNodeId = uint(value.Int64)
			}
		case k8smetric.FieldK8sNamespaceId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sNamespaceId", values[i])
			} else if value.Valid {
				km.K8sNamespaceId = uint(value.Int64)
			}
		case k8smetric.FieldK8sContainerId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field k8sContainerId", values[i])
			} else if value.Valid {
				km.K8sContainerId = uint(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this K8sMetric.
// Note that you need to call K8sMetric.Unwrap() before calling this method if this K8sMetric
// was returned from a transaction, and the transaction was committed or rolled back.
func (km *K8sMetric) Update() *K8sMetricUpdateOne {
	return (&K8sMetricClient{config: km.config}).UpdateOne(km)
}

// Unwrap unwraps the K8sMetric entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (km *K8sMetric) Unwrap() *K8sMetric {
	tx, ok := km.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sMetric is not a transactional entity")
	}
	km.config.driver = tx.drv
	return km
}

// String implements the fmt.Stringer.
func (km *K8sMetric) String() string {
	var builder strings.Builder
	builder.WriteString("K8sMetric(")
	builder.WriteString(fmt.Sprintf("id=%v", km.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(km.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(km.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(km.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ts=")
	builder.WriteString(km.Ts.Format(time.ANSIC))
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", km.Value))
	builder.WriteString(", endpointId=")
	builder.WriteString(fmt.Sprintf("%v", km.EndpointId))
	builder.WriteString(", typeId=")
	builder.WriteString(fmt.Sprintf("%v", km.TypeId))
	builder.WriteString(", nameId=")
	builder.WriteString(fmt.Sprintf("%v", km.NameId))
	builder.WriteString(", labelId=")
	builder.WriteString(fmt.Sprintf("%v", km.LabelId))
	builder.WriteString(", k8sClusterId=")
	builder.WriteString(fmt.Sprintf("%v", km.K8sClusterId))
	builder.WriteString(", k8sNodeId=")
	builder.WriteString(fmt.Sprintf("%v", km.K8sNodeId))
	builder.WriteString(", k8sNamespaceId=")
	builder.WriteString(fmt.Sprintf("%v", km.K8sNamespaceId))
	builder.WriteString(", k8sContainerId=")
	builder.WriteString(fmt.Sprintf("%v", km.K8sContainerId))
	builder.WriteByte(')')
	return builder.String()
}

// K8sMetrics is a parsable slice of K8sMetric.
type K8sMetrics []*K8sMetric

func (km K8sMetrics) config(cfg config) {
	for _i := range km {
		km[_i].config = cfg
	}
}
