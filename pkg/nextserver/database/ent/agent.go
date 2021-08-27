// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/agent"
)

// Agent is the model entity for the Agent schema.
type Agent struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Online holds the value of the "online" field.
	Online bool `json:"online,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Ipv4 holds the value of the "ipv4" field.
	Ipv4 string `json:"ipv4,omitempty"`
	// Ipv6 holds the value of the "ipv6" field.
	Ipv6 string `json:"ipv6,omitempty"`
	// PublicIpv4 holds the value of the "public_ipv4" field.
	PublicIpv4 string `json:"public_ipv4,omitempty"`
	// PublicIpv6 holds the value of the "public_ipv6" field.
	PublicIpv6 string `json:"public_ipv6,omitempty"`
	// LastContact holds the value of the "last_contact" field.
	LastContact time.Time `json:"last_contact,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// MachineId holds the value of the "machineId" field.
	MachineId string `json:"machineId,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ClusterId holds the value of the "clusterId" field.
	ClusterId uint `json:"clusterId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AgentQuery when eager-loading is set.
	Edges          AgentEdges `json:"edges"`
	cluster_agents *uint
}

// AgentEdges holds the relations/edges for other nodes in the graph.
type AgentEdges struct {
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"nodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[0] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Agent) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case agent.FieldOnline, agent.FieldDisabled:
			values[i] = new(sql.NullBool)
		case agent.FieldID, agent.FieldClusterId:
			values[i] = new(sql.NullInt64)
		case agent.FieldVersion, agent.FieldIpv4, agent.FieldIpv6, agent.FieldPublicIpv4, agent.FieldPublicIpv6, agent.FieldUUID, agent.FieldMachineId, agent.FieldDescription:
			values[i] = new(sql.NullString)
		case agent.FieldCreatedAt, agent.FieldUpdatedAt, agent.FieldDeletedAt, agent.FieldLastContact:
			values[i] = new(sql.NullTime)
		case agent.ForeignKeys[0]: // cluster_agents
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Agent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Agent fields.
func (a *Agent) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case agent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint(value.Int64)
		case agent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case agent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case agent.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case agent.FieldOnline:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field online", values[i])
			} else if value.Valid {
				a.Online = value.Bool
			}
		case agent.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				a.Version = value.String
			}
		case agent.FieldIpv4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ipv4", values[i])
			} else if value.Valid {
				a.Ipv4 = value.String
			}
		case agent.FieldIpv6:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ipv6", values[i])
			} else if value.Valid {
				a.Ipv6 = value.String
			}
		case agent.FieldPublicIpv4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_ipv4", values[i])
			} else if value.Valid {
				a.PublicIpv4 = value.String
			}
		case agent.FieldPublicIpv6:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_ipv6", values[i])
			} else if value.Valid {
				a.PublicIpv6 = value.String
			}
		case agent.FieldLastContact:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_contact", values[i])
			} else if value.Valid {
				a.LastContact = value.Time
			}
		case agent.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				a.Disabled = value.Bool
			}
		case agent.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				a.UUID = value.String
			}
		case agent.FieldMachineId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field machineId", values[i])
			} else if value.Valid {
				a.MachineId = value.String
			}
		case agent.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case agent.FieldClusterId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field clusterId", values[i])
			} else if value.Valid {
				a.ClusterId = uint(value.Int64)
			}
		case agent.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field cluster_agents", value)
			} else if value.Valid {
				a.cluster_agents = new(uint)
				*a.cluster_agents = uint(value.Int64)
			}
		}
	}
	return nil
}

// QueryNodes queries the "nodes" edge of the Agent entity.
func (a *Agent) QueryNodes() *NodeQuery {
	return (&AgentClient{config: a.config}).QueryNodes(a)
}

// Update returns a builder for updating this Agent.
// Note that you need to call Agent.Unwrap() before calling this method if this Agent
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Agent) Update() *AgentUpdateOne {
	return (&AgentClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Agent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Agent) Unwrap() *Agent {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Agent is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Agent) String() string {
	var builder strings.Builder
	builder.WriteString("Agent(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", online=")
	builder.WriteString(fmt.Sprintf("%v", a.Online))
	builder.WriteString(", version=")
	builder.WriteString(a.Version)
	builder.WriteString(", ipv4=")
	builder.WriteString(a.Ipv4)
	builder.WriteString(", ipv6=")
	builder.WriteString(a.Ipv6)
	builder.WriteString(", public_ipv4=")
	builder.WriteString(a.PublicIpv4)
	builder.WriteString(", public_ipv6=")
	builder.WriteString(a.PublicIpv6)
	builder.WriteString(", last_contact=")
	builder.WriteString(a.LastContact.Format(time.ANSIC))
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", a.Disabled))
	builder.WriteString(", uuid=")
	builder.WriteString(a.UUID)
	builder.WriteString(", machineId=")
	builder.WriteString(a.MachineId)
	builder.WriteString(", description=")
	builder.WriteString(a.Description)
	builder.WriteString(", clusterId=")
	builder.WriteString(fmt.Sprintf("%v", a.ClusterId))
	builder.WriteByte(')')
	return builder.String()
}

// Agents is a parsable slice of Agent.
type Agents []*Agent

func (a Agents) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}