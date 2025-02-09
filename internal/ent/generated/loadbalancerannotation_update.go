// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerannotation"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
)

// LoadBalancerAnnotationUpdate is the builder for updating LoadBalancerAnnotation entities.
type LoadBalancerAnnotationUpdate struct {
	config
	hooks    []Hook
	mutation *LoadBalancerAnnotationMutation
}

// Where appends a list predicates to the LoadBalancerAnnotationUpdate builder.
func (lbau *LoadBalancerAnnotationUpdate) Where(ps ...predicate.LoadBalancerAnnotation) *LoadBalancerAnnotationUpdate {
	lbau.mutation.Where(ps...)
	return lbau
}

// Mutation returns the LoadBalancerAnnotationMutation object of the builder.
func (lbau *LoadBalancerAnnotationUpdate) Mutation() *LoadBalancerAnnotationMutation {
	return lbau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lbau *LoadBalancerAnnotationUpdate) Save(ctx context.Context) (int, error) {
	lbau.defaults()
	return withHooks(ctx, lbau.sqlSave, lbau.mutation, lbau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lbau *LoadBalancerAnnotationUpdate) SaveX(ctx context.Context) int {
	affected, err := lbau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lbau *LoadBalancerAnnotationUpdate) Exec(ctx context.Context) error {
	_, err := lbau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbau *LoadBalancerAnnotationUpdate) ExecX(ctx context.Context) {
	if err := lbau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lbau *LoadBalancerAnnotationUpdate) defaults() {
	if _, ok := lbau.mutation.UpdatedAt(); !ok {
		v := loadbalancerannotation.UpdateDefaultUpdatedAt()
		lbau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lbau *LoadBalancerAnnotationUpdate) check() error {
	if _, ok := lbau.mutation.LoadBalancerID(); lbau.mutation.LoadBalancerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "LoadBalancerAnnotation.load_balancer"`)
	}
	return nil
}

func (lbau *LoadBalancerAnnotationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lbau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(loadbalancerannotation.Table, loadbalancerannotation.Columns, sqlgraph.NewFieldSpec(loadbalancerannotation.FieldID, field.TypeString))
	if ps := lbau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lbau.mutation.UpdatedAt(); ok {
		_spec.SetField(loadbalancerannotation.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lbau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loadbalancerannotation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lbau.mutation.done = true
	return n, nil
}

// LoadBalancerAnnotationUpdateOne is the builder for updating a single LoadBalancerAnnotation entity.
type LoadBalancerAnnotationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoadBalancerAnnotationMutation
}

// Mutation returns the LoadBalancerAnnotationMutation object of the builder.
func (lbauo *LoadBalancerAnnotationUpdateOne) Mutation() *LoadBalancerAnnotationMutation {
	return lbauo.mutation
}

// Where appends a list predicates to the LoadBalancerAnnotationUpdate builder.
func (lbauo *LoadBalancerAnnotationUpdateOne) Where(ps ...predicate.LoadBalancerAnnotation) *LoadBalancerAnnotationUpdateOne {
	lbauo.mutation.Where(ps...)
	return lbauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lbauo *LoadBalancerAnnotationUpdateOne) Select(field string, fields ...string) *LoadBalancerAnnotationUpdateOne {
	lbauo.fields = append([]string{field}, fields...)
	return lbauo
}

// Save executes the query and returns the updated LoadBalancerAnnotation entity.
func (lbauo *LoadBalancerAnnotationUpdateOne) Save(ctx context.Context) (*LoadBalancerAnnotation, error) {
	lbauo.defaults()
	return withHooks(ctx, lbauo.sqlSave, lbauo.mutation, lbauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lbauo *LoadBalancerAnnotationUpdateOne) SaveX(ctx context.Context) *LoadBalancerAnnotation {
	node, err := lbauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lbauo *LoadBalancerAnnotationUpdateOne) Exec(ctx context.Context) error {
	_, err := lbauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbauo *LoadBalancerAnnotationUpdateOne) ExecX(ctx context.Context) {
	if err := lbauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lbauo *LoadBalancerAnnotationUpdateOne) defaults() {
	if _, ok := lbauo.mutation.UpdatedAt(); !ok {
		v := loadbalancerannotation.UpdateDefaultUpdatedAt()
		lbauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lbauo *LoadBalancerAnnotationUpdateOne) check() error {
	if _, ok := lbauo.mutation.LoadBalancerID(); lbauo.mutation.LoadBalancerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "LoadBalancerAnnotation.load_balancer"`)
	}
	return nil
}

func (lbauo *LoadBalancerAnnotationUpdateOne) sqlSave(ctx context.Context) (_node *LoadBalancerAnnotation, err error) {
	if err := lbauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(loadbalancerannotation.Table, loadbalancerannotation.Columns, sqlgraph.NewFieldSpec(loadbalancerannotation.FieldID, field.TypeString))
	id, ok := lbauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "LoadBalancerAnnotation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lbauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loadbalancerannotation.FieldID)
		for _, f := range fields {
			if !loadbalancerannotation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != loadbalancerannotation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lbauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lbauo.mutation.UpdatedAt(); ok {
		_spec.SetField(loadbalancerannotation.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &LoadBalancerAnnotation{config: lbauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lbauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{loadbalancerannotation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lbauo.mutation.done = true
	return _node, nil
}
