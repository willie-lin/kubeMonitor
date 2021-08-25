// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/k8slabel"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/predicate"
)

// K8sLabelDelete is the builder for deleting a K8sLabel entity.
type K8sLabelDelete struct {
	config
	hooks    []Hook
	mutation *K8sLabelMutation
}

// Where appends a list predicates to the K8sLabelDelete builder.
func (kld *K8sLabelDelete) Where(ps ...predicate.K8sLabel) *K8sLabelDelete {
	kld.mutation.Where(ps...)
	return kld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (kld *K8sLabelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kld.hooks) == 0 {
		affected, err = kld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*K8sLabelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kld.mutation = mutation
			affected, err = kld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kld.hooks) - 1; i >= 0; i-- {
			if kld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (kld *K8sLabelDelete) ExecX(ctx context.Context) int {
	n, err := kld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (kld *K8sLabelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: k8slabel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: k8slabel.FieldID,
			},
		},
	}
	if ps := kld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, kld.driver, _spec)
}

// K8sLabelDeleteOne is the builder for deleting a single K8sLabel entity.
type K8sLabelDeleteOne struct {
	kld *K8sLabelDelete
}

// Exec executes the deletion query.
func (kldo *K8sLabelDeleteOne) Exec(ctx context.Context) error {
	n, err := kldo.kld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{k8slabel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (kldo *K8sLabelDeleteOne) ExecX(ctx context.Context) {
	kldo.kld.ExecX(ctx)
}
