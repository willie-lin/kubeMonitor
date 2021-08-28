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
)

// Container is the model entity for the Container schema.
type Container struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// ContainerId holds the value of the "containerId" field.
	ContainerId string `json:"containerId,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Info holds the value of the "info" field.
	Info []string `json:"info,omitempty"`
	// ClusterId holds the value of the "clusterId" field.
	ClusterId uint `json:"clusterId,omitempty"`
	// NodeId holds the value of the "nodeId" field.
	NodeId uint `json:"nodeId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContainerQuery when eager-loading is set.
	Edges           ContainerEdges `json:"edges"`
	node_containers *uint
}

// ContainerEdges holds the relations/edges for other nodes in the graph.
type ContainerEdges struct {
	// Process holds the value of the process edge.
	Process []*Proces `json:"process,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Node `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProcessOrErr returns the Process value or an error if the edge
// was not loaded in eager-loading.
func (e ContainerEdges) ProcessOrErr() ([]*Proces, error) {
	if e.loadedTypes[0] {
		return e.Process, nil
	}
	return nil, &NotLoadedError{edge: "process"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContainerEdges) OwnerOrErr() (*Node, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Container) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case container.FieldInfo:
			values[i] = new([]byte)
		case container.FieldID, container.FieldClusterId, container.FieldNodeId:
			values[i] = new(sql.NullInt64)
		case container.FieldType, container.FieldContainerId, container.FieldName, container.FieldImage:
			values[i] = new(sql.NullString)
		case container.FieldCreatedAt, container.FieldUpdatedAt, container.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case container.ForeignKeys[0]: // node_containers
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Container", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Container fields.
func (c *Container) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case container.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint(value.Int64)
		case container.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case container.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case container.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case container.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case container.FieldContainerId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field containerId", values[i])
			} else if value.Valid {
				c.ContainerId = value.String
			}
		case container.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case container.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				c.Image = value.String
			}
		case container.FieldInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Info); err != nil {
					return fmt.Errorf("unmarshal field info: %w", err)
				}
			}
		case container.FieldClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field clusterId", values[i])
			} else if value.Valid {
				c.ClusterId = uint(value.Int64)
			}
		case container.FieldNodeId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nodeId", values[i])
			} else if value.Valid {
				c.NodeId = uint(value.Int64)
			}
		case container.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field node_containers", value)
			} else if value.Valid {
				c.node_containers = new(uint)
				*c.node_containers = uint(value.Int64)
			}
		}
	}
	return nil
}

// QueryProcess queries the "process" edge of the Container entity.
func (c *Container) QueryProcess() *ProcesQuery {
	return (&ContainerClient{config: c.config}).QueryProcess(c)
}

// QueryOwner queries the "owner" edge of the Container entity.
func (c *Container) QueryOwner() *NodeQuery {
	return (&ContainerClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Container.
// Note that you need to call Container.Unwrap() before calling this method if this Container
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Container) Update() *ContainerUpdateOne {
	return (&ContainerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Container entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Container) Unwrap() *Container {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Container is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Container) String() string {
	var builder strings.Builder
	builder.WriteString("Container(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", type=")
	builder.WriteString(c.Type)
	builder.WriteString(", containerId=")
	builder.WriteString(c.ContainerId)
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", image=")
	builder.WriteString(c.Image)
	builder.WriteString(", info=")
	builder.WriteString(fmt.Sprintf("%v", c.Info))
	builder.WriteString(", clusterId=")
	builder.WriteString(fmt.Sprintf("%v", c.ClusterId))
	builder.WriteString(", nodeId=")
	builder.WriteString(fmt.Sprintf("%v", c.NodeId))
	builder.WriteByte(')')
	return builder.String()
}

// Containers is a parsable slice of Container.
type Containers []*Container

func (c Containers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
