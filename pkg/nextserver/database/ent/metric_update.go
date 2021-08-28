// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metric"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricendpoint"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metriclabel"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricname"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// MetricUpdate is the builder for updating Metric entities.
type MetricUpdate struct {
	config
	hooks    []Hook
	mutation *MetricMutation
}

// Where appends a list predicates to the MetricUpdate builder.
func (mu *MetricUpdate) Where(ps ...predicate.Metric) *MetricUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MetricUpdate) SetUpdatedAt(t time.Time) *MetricUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MetricUpdate) SetDeletedAt(t time.Time) *MetricUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetTs sets the "ts" field.
func (mu *MetricUpdate) SetTs(t time.Time) *MetricUpdate {
	mu.mutation.SetTs(t)
	return mu
}

// SetValue sets the "value" field.
func (mu *MetricUpdate) SetValue(f float64) *MetricUpdate {
	mu.mutation.ResetValue()
	mu.mutation.SetValue(f)
	return mu
}

// AddValue adds f to the "value" field.
func (mu *MetricUpdate) AddValue(f float64) *MetricUpdate {
	mu.mutation.AddValue(f)
	return mu
}

// SetEndpointId sets the "endpointId" field.
func (mu *MetricUpdate) SetEndpointId(u uint) *MetricUpdate {
	mu.mutation.ResetEndpointId()
	mu.mutation.SetEndpointId(u)
	return mu
}

// AddEndpointId adds u to the "endpointId" field.
func (mu *MetricUpdate) AddEndpointId(u uint) *MetricUpdate {
	mu.mutation.AddEndpointId(u)
	return mu
}

// SetTypeId sets the "typeId" field.
func (mu *MetricUpdate) SetTypeId(u uint) *MetricUpdate {
	mu.mutation.ResetTypeId()
	mu.mutation.SetTypeId(u)
	return mu
}

// AddTypeId adds u to the "typeId" field.
func (mu *MetricUpdate) AddTypeId(u uint) *MetricUpdate {
	mu.mutation.AddTypeId(u)
	return mu
}

// SetNameId sets the "nameId" field.
func (mu *MetricUpdate) SetNameId(u uint) *MetricUpdate {
	mu.mutation.ResetNameId()
	mu.mutation.SetNameId(u)
	return mu
}

// AddNameId adds u to the "nameId" field.
func (mu *MetricUpdate) AddNameId(u uint) *MetricUpdate {
	mu.mutation.AddNameId(u)
	return mu
}

// SetLabelId sets the "labelId" field.
func (mu *MetricUpdate) SetLabelId(u uint) *MetricUpdate {
	mu.mutation.ResetLabelId()
	mu.mutation.SetLabelId(u)
	return mu
}

// AddLabelId adds u to the "labelId" field.
func (mu *MetricUpdate) AddLabelId(u uint) *MetricUpdate {
	mu.mutation.AddLabelId(u)
	return mu
}

// SetClusterId sets the "clusterId" field.
func (mu *MetricUpdate) SetClusterId(u uint) *MetricUpdate {
	mu.mutation.ResetClusterId()
	mu.mutation.SetClusterId(u)
	return mu
}

// AddClusterId adds u to the "clusterId" field.
func (mu *MetricUpdate) AddClusterId(u uint) *MetricUpdate {
	mu.mutation.AddClusterId(u)
	return mu
}

// SetNodeId sets the "nodeId" field.
func (mu *MetricUpdate) SetNodeId(u uint) *MetricUpdate {
	mu.mutation.ResetNodeId()
	mu.mutation.SetNodeId(u)
	return mu
}

// AddNodeId adds u to the "nodeId" field.
func (mu *MetricUpdate) AddNodeId(u uint) *MetricUpdate {
	mu.mutation.AddNodeId(u)
	return mu
}

// SetProcesId sets the "procesId" field.
func (mu *MetricUpdate) SetProcesId(u uint) *MetricUpdate {
	mu.mutation.ResetProcesId()
	mu.mutation.SetProcesId(u)
	return mu
}

// AddProcesId adds u to the "procesId" field.
func (mu *MetricUpdate) AddProcesId(u uint) *MetricUpdate {
	mu.mutation.AddProcesId(u)
	return mu
}

// SetContainerId sets the "containerId" field.
func (mu *MetricUpdate) SetContainerId(u uint) *MetricUpdate {
	mu.mutation.ResetContainerId()
	mu.mutation.SetContainerId(u)
	return mu
}

// AddContainerId adds u to the "containerId" field.
func (mu *MetricUpdate) AddContainerId(u uint) *MetricUpdate {
	mu.mutation.AddContainerId(u)
	return mu
}

// SetMetricNameMetricsID sets the "MetricName_Metrics" edge to the MetricName entity by ID.
func (mu *MetricUpdate) SetMetricNameMetricsID(id uint) *MetricUpdate {
	mu.mutation.SetMetricNameMetricsID(id)
	return mu
}

// SetNillableMetricNameMetricsID sets the "MetricName_Metrics" edge to the MetricName entity by ID if the given value is not nil.
func (mu *MetricUpdate) SetNillableMetricNameMetricsID(id *uint) *MetricUpdate {
	if id != nil {
		mu = mu.SetMetricNameMetricsID(*id)
	}
	return mu
}

// SetMetricNameMetrics sets the "MetricName_Metrics" edge to the MetricName entity.
func (mu *MetricUpdate) SetMetricNameMetrics(m *MetricName) *MetricUpdate {
	return mu.SetMetricNameMetricsID(m.ID)
}

// SetMetricEndpointMetricsID sets the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity by ID.
func (mu *MetricUpdate) SetMetricEndpointMetricsID(id uint) *MetricUpdate {
	mu.mutation.SetMetricEndpointMetricsID(id)
	return mu
}

// SetNillableMetricEndpointMetricsID sets the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity by ID if the given value is not nil.
func (mu *MetricUpdate) SetNillableMetricEndpointMetricsID(id *uint) *MetricUpdate {
	if id != nil {
		mu = mu.SetMetricEndpointMetricsID(*id)
	}
	return mu
}

// SetMetricEndpointMetrics sets the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity.
func (mu *MetricUpdate) SetMetricEndpointMetrics(m *MetricEndpoint) *MetricUpdate {
	return mu.SetMetricEndpointMetricsID(m.ID)
}

// SetMetricLabelMetricsID sets the "MetricLabel_Metrics" edge to the MetricLabel entity by ID.
func (mu *MetricUpdate) SetMetricLabelMetricsID(id uint) *MetricUpdate {
	mu.mutation.SetMetricLabelMetricsID(id)
	return mu
}

// SetNillableMetricLabelMetricsID sets the "MetricLabel_Metrics" edge to the MetricLabel entity by ID if the given value is not nil.
func (mu *MetricUpdate) SetNillableMetricLabelMetricsID(id *uint) *MetricUpdate {
	if id != nil {
		mu = mu.SetMetricLabelMetricsID(*id)
	}
	return mu
}

// SetMetricLabelMetrics sets the "MetricLabel_Metrics" edge to the MetricLabel entity.
func (mu *MetricUpdate) SetMetricLabelMetrics(m *MetricLabel) *MetricUpdate {
	return mu.SetMetricLabelMetricsID(m.ID)
}

// Mutation returns the MetricMutation object of the builder.
func (mu *MetricUpdate) Mutation() *MetricMutation {
	return mu.mutation
}

// ClearMetricNameMetrics clears the "MetricName_Metrics" edge to the MetricName entity.
func (mu *MetricUpdate) ClearMetricNameMetrics() *MetricUpdate {
	mu.mutation.ClearMetricNameMetrics()
	return mu
}

// ClearMetricEndpointMetrics clears the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity.
func (mu *MetricUpdate) ClearMetricEndpointMetrics() *MetricUpdate {
	mu.mutation.ClearMetricEndpointMetrics()
	return mu
}

// ClearMetricLabelMetrics clears the "MetricLabel_Metrics" edge to the MetricLabel entity.
func (mu *MetricUpdate) ClearMetricLabelMetrics() *MetricUpdate {
	mu.mutation.ClearMetricLabelMetrics()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetricUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetricUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetricUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetricUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MetricUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := metric.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	if _, ok := mu.mutation.DeletedAt(); !ok {
		v := metric.UpdateDefaultDeletedAt()
		mu.mutation.SetDeletedAt(v)
	}
}

func (mu *MetricUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metric.Table,
			Columns: metric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldDeletedAt,
		})
	}
	if value, ok := mu.mutation.Ts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldTs,
		})
	}
	if value, ok := mu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: metric.FieldValue,
		})
	}
	if value, ok := mu.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: metric.FieldValue,
		})
	}
	if value, ok := mu.mutation.EndpointId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldEndpointId,
		})
	}
	if value, ok := mu.mutation.AddedEndpointId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldEndpointId,
		})
	}
	if value, ok := mu.mutation.TypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldTypeId,
		})
	}
	if value, ok := mu.mutation.AddedTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldTypeId,
		})
	}
	if value, ok := mu.mutation.NameId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNameId,
		})
	}
	if value, ok := mu.mutation.AddedNameId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNameId,
		})
	}
	if value, ok := mu.mutation.LabelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldLabelId,
		})
	}
	if value, ok := mu.mutation.AddedLabelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldLabelId,
		})
	}
	if value, ok := mu.mutation.ClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldClusterId,
		})
	}
	if value, ok := mu.mutation.AddedClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldClusterId,
		})
	}
	if value, ok := mu.mutation.NodeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNodeId,
		})
	}
	if value, ok := mu.mutation.AddedNodeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNodeId,
		})
	}
	if value, ok := mu.mutation.ProcesId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldProcesId,
		})
	}
	if value, ok := mu.mutation.AddedProcesId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldProcesId,
		})
	}
	if value, ok := mu.mutation.ContainerId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldContainerId,
		})
	}
	if value, ok := mu.mutation.AddedContainerId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldContainerId,
		})
	}
	if mu.mutation.MetricNameMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricNameMetricsTable,
			Columns: []string{metric.MetricNameMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricname.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MetricNameMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricNameMetricsTable,
			Columns: []string{metric.MetricNameMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricname.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MetricEndpointMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricEndpointMetricsTable,
			Columns: []string{metric.MetricEndpointMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricendpoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MetricEndpointMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricEndpointMetricsTable,
			Columns: []string{metric.MetricEndpointMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricendpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MetricLabelMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricLabelMetricsTable,
			Columns: []string{metric.MetricLabelMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metriclabel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MetricLabelMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricLabelMetricsTable,
			Columns: []string{metric.MetricLabelMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metriclabel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metric.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MetricUpdateOne is the builder for updating a single Metric entity.
type MetricUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetricMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MetricUpdateOne) SetUpdatedAt(t time.Time) *MetricUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MetricUpdateOne) SetDeletedAt(t time.Time) *MetricUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetTs sets the "ts" field.
func (muo *MetricUpdateOne) SetTs(t time.Time) *MetricUpdateOne {
	muo.mutation.SetTs(t)
	return muo
}

// SetValue sets the "value" field.
func (muo *MetricUpdateOne) SetValue(f float64) *MetricUpdateOne {
	muo.mutation.ResetValue()
	muo.mutation.SetValue(f)
	return muo
}

// AddValue adds f to the "value" field.
func (muo *MetricUpdateOne) AddValue(f float64) *MetricUpdateOne {
	muo.mutation.AddValue(f)
	return muo
}

// SetEndpointId sets the "endpointId" field.
func (muo *MetricUpdateOne) SetEndpointId(u uint) *MetricUpdateOne {
	muo.mutation.ResetEndpointId()
	muo.mutation.SetEndpointId(u)
	return muo
}

// AddEndpointId adds u to the "endpointId" field.
func (muo *MetricUpdateOne) AddEndpointId(u uint) *MetricUpdateOne {
	muo.mutation.AddEndpointId(u)
	return muo
}

// SetTypeId sets the "typeId" field.
func (muo *MetricUpdateOne) SetTypeId(u uint) *MetricUpdateOne {
	muo.mutation.ResetTypeId()
	muo.mutation.SetTypeId(u)
	return muo
}

// AddTypeId adds u to the "typeId" field.
func (muo *MetricUpdateOne) AddTypeId(u uint) *MetricUpdateOne {
	muo.mutation.AddTypeId(u)
	return muo
}

// SetNameId sets the "nameId" field.
func (muo *MetricUpdateOne) SetNameId(u uint) *MetricUpdateOne {
	muo.mutation.ResetNameId()
	muo.mutation.SetNameId(u)
	return muo
}

// AddNameId adds u to the "nameId" field.
func (muo *MetricUpdateOne) AddNameId(u uint) *MetricUpdateOne {
	muo.mutation.AddNameId(u)
	return muo
}

// SetLabelId sets the "labelId" field.
func (muo *MetricUpdateOne) SetLabelId(u uint) *MetricUpdateOne {
	muo.mutation.ResetLabelId()
	muo.mutation.SetLabelId(u)
	return muo
}

// AddLabelId adds u to the "labelId" field.
func (muo *MetricUpdateOne) AddLabelId(u uint) *MetricUpdateOne {
	muo.mutation.AddLabelId(u)
	return muo
}

// SetClusterId sets the "clusterId" field.
func (muo *MetricUpdateOne) SetClusterId(u uint) *MetricUpdateOne {
	muo.mutation.ResetClusterId()
	muo.mutation.SetClusterId(u)
	return muo
}

// AddClusterId adds u to the "clusterId" field.
func (muo *MetricUpdateOne) AddClusterId(u uint) *MetricUpdateOne {
	muo.mutation.AddClusterId(u)
	return muo
}

// SetNodeId sets the "nodeId" field.
func (muo *MetricUpdateOne) SetNodeId(u uint) *MetricUpdateOne {
	muo.mutation.ResetNodeId()
	muo.mutation.SetNodeId(u)
	return muo
}

// AddNodeId adds u to the "nodeId" field.
func (muo *MetricUpdateOne) AddNodeId(u uint) *MetricUpdateOne {
	muo.mutation.AddNodeId(u)
	return muo
}

// SetProcesId sets the "procesId" field.
func (muo *MetricUpdateOne) SetProcesId(u uint) *MetricUpdateOne {
	muo.mutation.ResetProcesId()
	muo.mutation.SetProcesId(u)
	return muo
}

// AddProcesId adds u to the "procesId" field.
func (muo *MetricUpdateOne) AddProcesId(u uint) *MetricUpdateOne {
	muo.mutation.AddProcesId(u)
	return muo
}

// SetContainerId sets the "containerId" field.
func (muo *MetricUpdateOne) SetContainerId(u uint) *MetricUpdateOne {
	muo.mutation.ResetContainerId()
	muo.mutation.SetContainerId(u)
	return muo
}

// AddContainerId adds u to the "containerId" field.
func (muo *MetricUpdateOne) AddContainerId(u uint) *MetricUpdateOne {
	muo.mutation.AddContainerId(u)
	return muo
}

// SetMetricNameMetricsID sets the "MetricName_Metrics" edge to the MetricName entity by ID.
func (muo *MetricUpdateOne) SetMetricNameMetricsID(id uint) *MetricUpdateOne {
	muo.mutation.SetMetricNameMetricsID(id)
	return muo
}

// SetNillableMetricNameMetricsID sets the "MetricName_Metrics" edge to the MetricName entity by ID if the given value is not nil.
func (muo *MetricUpdateOne) SetNillableMetricNameMetricsID(id *uint) *MetricUpdateOne {
	if id != nil {
		muo = muo.SetMetricNameMetricsID(*id)
	}
	return muo
}

// SetMetricNameMetrics sets the "MetricName_Metrics" edge to the MetricName entity.
func (muo *MetricUpdateOne) SetMetricNameMetrics(m *MetricName) *MetricUpdateOne {
	return muo.SetMetricNameMetricsID(m.ID)
}

// SetMetricEndpointMetricsID sets the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity by ID.
func (muo *MetricUpdateOne) SetMetricEndpointMetricsID(id uint) *MetricUpdateOne {
	muo.mutation.SetMetricEndpointMetricsID(id)
	return muo
}

// SetNillableMetricEndpointMetricsID sets the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity by ID if the given value is not nil.
func (muo *MetricUpdateOne) SetNillableMetricEndpointMetricsID(id *uint) *MetricUpdateOne {
	if id != nil {
		muo = muo.SetMetricEndpointMetricsID(*id)
	}
	return muo
}

// SetMetricEndpointMetrics sets the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity.
func (muo *MetricUpdateOne) SetMetricEndpointMetrics(m *MetricEndpoint) *MetricUpdateOne {
	return muo.SetMetricEndpointMetricsID(m.ID)
}

// SetMetricLabelMetricsID sets the "MetricLabel_Metrics" edge to the MetricLabel entity by ID.
func (muo *MetricUpdateOne) SetMetricLabelMetricsID(id uint) *MetricUpdateOne {
	muo.mutation.SetMetricLabelMetricsID(id)
	return muo
}

// SetNillableMetricLabelMetricsID sets the "MetricLabel_Metrics" edge to the MetricLabel entity by ID if the given value is not nil.
func (muo *MetricUpdateOne) SetNillableMetricLabelMetricsID(id *uint) *MetricUpdateOne {
	if id != nil {
		muo = muo.SetMetricLabelMetricsID(*id)
	}
	return muo
}

// SetMetricLabelMetrics sets the "MetricLabel_Metrics" edge to the MetricLabel entity.
func (muo *MetricUpdateOne) SetMetricLabelMetrics(m *MetricLabel) *MetricUpdateOne {
	return muo.SetMetricLabelMetricsID(m.ID)
}

// Mutation returns the MetricMutation object of the builder.
func (muo *MetricUpdateOne) Mutation() *MetricMutation {
	return muo.mutation
}

// ClearMetricNameMetrics clears the "MetricName_Metrics" edge to the MetricName entity.
func (muo *MetricUpdateOne) ClearMetricNameMetrics() *MetricUpdateOne {
	muo.mutation.ClearMetricNameMetrics()
	return muo
}

// ClearMetricEndpointMetrics clears the "MetricEndpoint_Metrics" edge to the MetricEndpoint entity.
func (muo *MetricUpdateOne) ClearMetricEndpointMetrics() *MetricUpdateOne {
	muo.mutation.ClearMetricEndpointMetrics()
	return muo
}

// ClearMetricLabelMetrics clears the "MetricLabel_Metrics" edge to the MetricLabel entity.
func (muo *MetricUpdateOne) ClearMetricLabelMetrics() *MetricUpdateOne {
	muo.mutation.ClearMetricLabelMetrics()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetricUpdateOne) Select(field string, fields ...string) *MetricUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metric entity.
func (muo *MetricUpdateOne) Save(ctx context.Context) (*Metric, error) {
	var (
		err  error
		node *Metric
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetricUpdateOne) SaveX(ctx context.Context) *Metric {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetricUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetricUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MetricUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := metric.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	if _, ok := muo.mutation.DeletedAt(); !ok {
		v := metric.UpdateDefaultDeletedAt()
		muo.mutation.SetDeletedAt(v)
	}
}

func (muo *MetricUpdateOne) sqlSave(ctx context.Context) (_node *Metric, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metric.Table,
			Columns: metric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Metric.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metric.FieldID)
		for _, f := range fields {
			if !metric.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != metric.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldDeletedAt,
		})
	}
	if value, ok := muo.mutation.Ts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldTs,
		})
	}
	if value, ok := muo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: metric.FieldValue,
		})
	}
	if value, ok := muo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: metric.FieldValue,
		})
	}
	if value, ok := muo.mutation.EndpointId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldEndpointId,
		})
	}
	if value, ok := muo.mutation.AddedEndpointId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldEndpointId,
		})
	}
	if value, ok := muo.mutation.TypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldTypeId,
		})
	}
	if value, ok := muo.mutation.AddedTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldTypeId,
		})
	}
	if value, ok := muo.mutation.NameId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNameId,
		})
	}
	if value, ok := muo.mutation.AddedNameId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNameId,
		})
	}
	if value, ok := muo.mutation.LabelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldLabelId,
		})
	}
	if value, ok := muo.mutation.AddedLabelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldLabelId,
		})
	}
	if value, ok := muo.mutation.ClusterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldClusterId,
		})
	}
	if value, ok := muo.mutation.AddedClusterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldClusterId,
		})
	}
	if value, ok := muo.mutation.NodeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNodeId,
		})
	}
	if value, ok := muo.mutation.AddedNodeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldNodeId,
		})
	}
	if value, ok := muo.mutation.ProcesId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldProcesId,
		})
	}
	if value, ok := muo.mutation.AddedProcesId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldProcesId,
		})
	}
	if value, ok := muo.mutation.ContainerId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldContainerId,
		})
	}
	if value, ok := muo.mutation.AddedContainerId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: metric.FieldContainerId,
		})
	}
	if muo.mutation.MetricNameMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricNameMetricsTable,
			Columns: []string{metric.MetricNameMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricname.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MetricNameMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricNameMetricsTable,
			Columns: []string{metric.MetricNameMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricname.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MetricEndpointMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricEndpointMetricsTable,
			Columns: []string{metric.MetricEndpointMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricendpoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MetricEndpointMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricEndpointMetricsTable,
			Columns: []string{metric.MetricEndpointMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metricendpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MetricLabelMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricLabelMetricsTable,
			Columns: []string{metric.MetricLabelMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metriclabel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MetricLabelMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.MetricLabelMetricsTable,
			Columns: []string{metric.MetricLabelMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: metriclabel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Metric{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metric.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
