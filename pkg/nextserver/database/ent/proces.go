// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/container"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/node"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/proces"
)

// Proces is the model entity for the Proces schema.
type Proces struct {
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
	ContainerId string `json:"containerId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProcesQuery when eager-loading is set.
	Edges             ProcesEdges `json:"edges"`
	container_process *uint
	node_process      *uint
}

// ProcesEdges holds the relations/edges for other nodes in the graph.
type ProcesEdges struct {
	// NodeProcess holds the value of the node_process edge.
	NodeProcess *Node `json:"node_process,omitempty"`
	// ContainerProcess holds the value of the container_process edge.
	ContainerProcess *Container `json:"container_process,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// NodeProcessOrErr returns the NodeProcess value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProcesEdges) NodeProcessOrErr() (*Node, error) {
	if e.loadedTypes[0] {
		if e.NodeProcess == nil {
			// The edge node_process was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.NodeProcess, nil
	}
	return nil, &NotLoadedError{edge: "node_process"}
}

// ContainerProcessOrErr returns the ContainerProcess value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProcesEdges) ContainerProcessOrErr() (*Container, error) {
	if e.loadedTypes[1] {
		if e.ContainerProcess == nil {
			// The edge container_process was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: container.Label}
		}
		return e.ContainerProcess, nil
	}
	return nil, &NotLoadedError{edge: "container_process"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Proces) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case proces.FieldInfo:
			values[i] = new([]byte)
		case proces.FieldID, proces.FieldPId:
			values[i] = new(sql.NullInt64)
		case proces.FieldName, proces.FieldCmd, proces.FieldClusterId, proces.FieldNodeId, proces.FieldContainerId:
			values[i] = new(sql.NullString)
		case proces.FieldCreatedAt, proces.FieldUpdatedAt, proces.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case proces.ForeignKeys[0]: // container_process
			values[i] = new(sql.NullInt64)
		case proces.ForeignKeys[1]: // node_process
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Proces", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Proces fields.
func (pr *Proces) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case proces.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint(value.Int64)
		case proces.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case proces.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case proces.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = value.Time
			}
		case proces.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case proces.FieldPId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pId", values[i])
			} else if value.Valid {
				pr.PId = int32(value.Int64)
			}
		case proces.FieldCmd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cmd", values[i])
			} else if value.Valid {
				pr.Cmd = value.String
			}
		case proces.FieldInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Info); err != nil {
					return fmt.Errorf("unmarshal field info: %w", err)
				}
			}
		case proces.FieldClusterId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clusterId", values[i])
			} else if value.Valid {
				pr.ClusterId = value.String
			}
		case proces.FieldNodeId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nodeId", values[i])
			} else if value.Valid {
				pr.NodeId = value.String
			}
		case proces.FieldContainerId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field containerId", values[i])
			} else if value.Valid {
				pr.ContainerId = value.String
			}
		case proces.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field container_process", value)
			} else if value.Valid {
				pr.container_process = new(uint)
				*pr.container_process = uint(value.Int64)
			}
		case proces.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field node_process", value)
			} else if value.Valid {
				pr.node_process = new(uint)
				*pr.node_process = uint(value.Int64)
			}
		}
	}
	return nil
}

// QueryNodeProcess queries the "node_process" edge of the Proces entity.
func (pr *Proces) QueryNodeProcess() *NodeQuery {
	return (&ProcesClient{config: pr.config}).QueryNodeProcess(pr)
}

// QueryContainerProcess queries the "container_process" edge of the Proces entity.
func (pr *Proces) QueryContainerProcess() *ContainerQuery {
	return (&ProcesClient{config: pr.config}).QueryContainerProcess(pr)
}

// Update returns a builder for updating this Proces.
// Note that you need to call Proces.Unwrap() before calling this method if this Proces
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Proces) Update() *ProcesUpdateOne {
	return (&ProcesClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Proces entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Proces) Unwrap() *Proces {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Proces is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Proces) String() string {
	var builder strings.Builder
	builder.WriteString("Proces(")
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

// ProcesSlice is a parsable slice of Proces.
type ProcesSlice []*Proces

func (pr ProcesSlice) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
