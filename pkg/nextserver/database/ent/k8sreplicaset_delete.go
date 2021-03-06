// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8sreplicaset"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sReplicaSetDelete is the builder for deleting a K8sReplicaSet entity.
type K8sReplicaSetDelete struct {
	config
	hooks    []Hook
	mutation *K8sReplicaSetMutation
}

// Where appends a list predicates to the K8sReplicaSetDelete builder.
func (krsd *K8sReplicaSetDelete) Where(ps ...predicate.K8sReplicaSet) *K8sReplicaSetDelete {
	krsd.mutation.Where(ps...)
	return krsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (krsd *K8sReplicaSetDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(krsd.hooks) == 0 {
		affected, err = krsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sReplicaSetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			krsd.mutation = mutation
			affected, err = krsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(krsd.hooks) - 1; i >= 0; i-- {
			if krsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = krsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, krsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (krsd *K8sReplicaSetDelete) ExecX(ctx context.Context) int {
	n, err := krsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (krsd *K8sReplicaSetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: k8sreplicaset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8sreplicaset.FieldID,
			},
		},
	}
	if ps := krsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, krsd.driver, _spec)
}

// K8sReplicaSetDeleteOne is the builder for deleting a single K8sReplicaSet entity.
type K8sReplicaSetDeleteOne struct {
	krsd *K8sReplicaSetDelete
}

// Exec executes the deletion query.
func (krsdo *K8sReplicaSetDeleteOne) Exec(ctx context.Context) error {
	n, err := krsdo.krsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{k8sreplicaset.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (krsdo *K8sReplicaSetDeleteOne) ExecX(ctx context.Context) {
	krsdo.krsd.ExecX(ctx)
}
