// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/process"
)

// Process is the model entity for the Process schema.
type Process struct {
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
	// PId holds the value of the "pId" field.
	PId int32 `json:"pId,omitempty"`
	// Cmd holds the value of the "cmd" field.
	Cmd string `json:"cmd,omitempty"`
	// Info holds the value of the "info" field.
	Info []string `json:"info,omitempty"`
	// ClusterId holds the value of the "clusterId" field.
	ClusterId string `json:"clusterId,omitempty"`
	// NodeId holds the value of the "nodeId" field.
	NodeId string `json:"nodeId,omitempty"`
	// ContainerId holds the value of the "containerId" field.
	ContainerId         string `json:"containerId,omitempty"`
	container_processes *uint
	node_processes      *uint
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Process) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case process.FieldInfo:
			values[i] = new([]byte)
		case process.FieldID, process.FieldPId:
			values[i] = new(sql.NullInt64)
		case process.FieldName, process.FieldCmd, process.FieldClusterId, process.FieldNodeId, process.FieldContainerId:
			values[i] = new(sql.NullString)
		case process.FieldCreatedAt, process.FieldUpdatedAt, process.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case process.ForeignKeys[0]: // container_processes
			values[i] = new(sql.NullInt64)
		case process.ForeignKeys[1]: // node_processes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Process", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Process fields.
func (pr *Process) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case process.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint(value.Int64)
		case process.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case process.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case process.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = value.Time
			}
		case process.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case process.FieldPId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pId", values[i])
			} else if value.Valid {
				pr.PId = int32(value.Int64)
			}
		case process.FieldCmd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cmd", values[i])
			} else if value.Valid {
				pr.Cmd = value.String
			}
		case process.FieldInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Info); err != nil {
					return fmt.Errorf("unmarshal field info: %w", err)
				}
			}
		case process.FieldClusterId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clusterId", values[i])
			} else if value.Valid {
				pr.ClusterId = value.String
			}
		case process.FieldNodeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nodeId", values[i])
			} else if value.Valid {
				pr.NodeId = value.String
			}
		case process.FieldContainerId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field containerId", values[i])
			} else if value.Valid {
				pr.ContainerId = value.String
			}
		case process.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field container_processes", value)
			} else if value.Valid {
				pr.container_processes = new(uint)
				*pr.container_processes = uint(value.Int64)
			}
		case process.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field node_processes", value)
			} else if value.Valid {
				pr.node_processes = new(uint)
				*pr.node_processes = uint(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Process.
// Note that you need to call Process.Unwrap() before calling this method if this Process
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Process) Update() *ProcessUpdateOne {
	return (&ProcessClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Process entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Process) Unwrap() *Process {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Process is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Process) String() string {
	var builder strings.Builder
	builder.WriteString("Process(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(pr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", pId=")
	builder.WriteString(fmt.Sprintf("%v", pr.PId))
	builder.WriteString(", cmd=")
	builder.WriteString(pr.Cmd)
	builder.WriteString(", info=")
	builder.WriteString(fmt.Sprintf("%v", pr.Info))
	builder.WriteString(", clusterId=")
	builder.WriteString(pr.ClusterId)
	builder.WriteString(", nodeId=")
	builder.WriteString(pr.NodeId)
	builder.WriteString(", containerId=")
	builder.WriteString(pr.ContainerId)
	builder.WriteByte(')')
	return builder.String()
}

// Processes is a parsable slice of Process.
type Processes []*Process

func (pr Processes) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
