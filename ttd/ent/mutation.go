// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/zackartz/ttd/ent/predicate"
	"github.com/zackartz/ttd/ent/timestamp"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTimestamp = "Timestamp"
)

// TimestampMutation represents an operation that mutates the Timestamp nodes in the graph.
type TimestampMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	start_time    *time.Time
	end_time      *time.Time
	active        *bool
	comment       *string
	category      *string
	project       *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Timestamp, error)
	predicates    []predicate.Timestamp
}

var _ ent.Mutation = (*TimestampMutation)(nil)

// timestampOption allows management of the mutation configuration using functional options.
type timestampOption func(*TimestampMutation)

// newTimestampMutation creates new mutation for the Timestamp entity.
func newTimestampMutation(c config, op Op, opts ...timestampOption) *TimestampMutation {
	m := &TimestampMutation{
		config:        c,
		op:            op,
		typ:           TypeTimestamp,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTimestampID sets the ID field of the mutation.
func withTimestampID(id uuid.UUID) timestampOption {
	return func(m *TimestampMutation) {
		var (
			err   error
			once  sync.Once
			value *Timestamp
		)
		m.oldValue = func(ctx context.Context) (*Timestamp, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Timestamp.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTimestamp sets the old Timestamp of the mutation.
func withTimestamp(node *Timestamp) timestampOption {
	return func(m *TimestampMutation) {
		m.oldValue = func(context.Context) (*Timestamp, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TimestampMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TimestampMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Timestamp entities.
func (m *TimestampMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *TimestampMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetStartTime sets the "start_time" field.
func (m *TimestampMutation) SetStartTime(t time.Time) {
	m.start_time = &t
}

// StartTime returns the value of the "start_time" field in the mutation.
func (m *TimestampMutation) StartTime() (r time.Time, exists bool) {
	v := m.start_time
	if v == nil {
		return
	}
	return *v, true
}

// OldStartTime returns the old "start_time" field's value of the Timestamp entity.
// If the Timestamp object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimestampMutation) OldStartTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStartTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStartTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartTime: %w", err)
	}
	return oldValue.StartTime, nil
}

// ResetStartTime resets all changes to the "start_time" field.
func (m *TimestampMutation) ResetStartTime() {
	m.start_time = nil
}

// SetEndTime sets the "end_time" field.
func (m *TimestampMutation) SetEndTime(t time.Time) {
	m.end_time = &t
}

// EndTime returns the value of the "end_time" field in the mutation.
func (m *TimestampMutation) EndTime() (r time.Time, exists bool) {
	v := m.end_time
	if v == nil {
		return
	}
	return *v, true
}

// OldEndTime returns the old "end_time" field's value of the Timestamp entity.
// If the Timestamp object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimestampMutation) OldEndTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEndTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEndTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEndTime: %w", err)
	}
	return oldValue.EndTime, nil
}

// ResetEndTime resets all changes to the "end_time" field.
func (m *TimestampMutation) ResetEndTime() {
	m.end_time = nil
}

// SetActive sets the "active" field.
func (m *TimestampMutation) SetActive(b bool) {
	m.active = &b
}

// Active returns the value of the "active" field in the mutation.
func (m *TimestampMutation) Active() (r bool, exists bool) {
	v := m.active
	if v == nil {
		return
	}
	return *v, true
}

// OldActive returns the old "active" field's value of the Timestamp entity.
// If the Timestamp object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimestampMutation) OldActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActive: %w", err)
	}
	return oldValue.Active, nil
}

// ResetActive resets all changes to the "active" field.
func (m *TimestampMutation) ResetActive() {
	m.active = nil
}

// SetComment sets the "comment" field.
func (m *TimestampMutation) SetComment(s string) {
	m.comment = &s
}

// Comment returns the value of the "comment" field in the mutation.
func (m *TimestampMutation) Comment() (r string, exists bool) {
	v := m.comment
	if v == nil {
		return
	}
	return *v, true
}

// OldComment returns the old "comment" field's value of the Timestamp entity.
// If the Timestamp object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimestampMutation) OldComment(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldComment is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldComment requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldComment: %w", err)
	}
	return oldValue.Comment, nil
}

// ClearComment clears the value of the "comment" field.
func (m *TimestampMutation) ClearComment() {
	m.comment = nil
	m.clearedFields[timestamp.FieldComment] = struct{}{}
}

// CommentCleared returns if the "comment" field was cleared in this mutation.
func (m *TimestampMutation) CommentCleared() bool {
	_, ok := m.clearedFields[timestamp.FieldComment]
	return ok
}

// ResetComment resets all changes to the "comment" field.
func (m *TimestampMutation) ResetComment() {
	m.comment = nil
	delete(m.clearedFields, timestamp.FieldComment)
}

// SetCategory sets the "category" field.
func (m *TimestampMutation) SetCategory(s string) {
	m.category = &s
}

// Category returns the value of the "category" field in the mutation.
func (m *TimestampMutation) Category() (r string, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategory returns the old "category" field's value of the Timestamp entity.
// If the Timestamp object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimestampMutation) OldCategory(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCategory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCategory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategory: %w", err)
	}
	return oldValue.Category, nil
}

// ClearCategory clears the value of the "category" field.
func (m *TimestampMutation) ClearCategory() {
	m.category = nil
	m.clearedFields[timestamp.FieldCategory] = struct{}{}
}

// CategoryCleared returns if the "category" field was cleared in this mutation.
func (m *TimestampMutation) CategoryCleared() bool {
	_, ok := m.clearedFields[timestamp.FieldCategory]
	return ok
}

// ResetCategory resets all changes to the "category" field.
func (m *TimestampMutation) ResetCategory() {
	m.category = nil
	delete(m.clearedFields, timestamp.FieldCategory)
}

// SetProject sets the "project" field.
func (m *TimestampMutation) SetProject(s string) {
	m.project = &s
}

// Project returns the value of the "project" field in the mutation.
func (m *TimestampMutation) Project() (r string, exists bool) {
	v := m.project
	if v == nil {
		return
	}
	return *v, true
}

// OldProject returns the old "project" field's value of the Timestamp entity.
// If the Timestamp object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimestampMutation) OldProject(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldProject is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldProject requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProject: %w", err)
	}
	return oldValue.Project, nil
}

// ResetProject resets all changes to the "project" field.
func (m *TimestampMutation) ResetProject() {
	m.project = nil
}

// Op returns the operation name.
func (m *TimestampMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Timestamp).
func (m *TimestampMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TimestampMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.start_time != nil {
		fields = append(fields, timestamp.FieldStartTime)
	}
	if m.end_time != nil {
		fields = append(fields, timestamp.FieldEndTime)
	}
	if m.active != nil {
		fields = append(fields, timestamp.FieldActive)
	}
	if m.comment != nil {
		fields = append(fields, timestamp.FieldComment)
	}
	if m.category != nil {
		fields = append(fields, timestamp.FieldCategory)
	}
	if m.project != nil {
		fields = append(fields, timestamp.FieldProject)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TimestampMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case timestamp.FieldStartTime:
		return m.StartTime()
	case timestamp.FieldEndTime:
		return m.EndTime()
	case timestamp.FieldActive:
		return m.Active()
	case timestamp.FieldComment:
		return m.Comment()
	case timestamp.FieldCategory:
		return m.Category()
	case timestamp.FieldProject:
		return m.Project()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TimestampMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case timestamp.FieldStartTime:
		return m.OldStartTime(ctx)
	case timestamp.FieldEndTime:
		return m.OldEndTime(ctx)
	case timestamp.FieldActive:
		return m.OldActive(ctx)
	case timestamp.FieldComment:
		return m.OldComment(ctx)
	case timestamp.FieldCategory:
		return m.OldCategory(ctx)
	case timestamp.FieldProject:
		return m.OldProject(ctx)
	}
	return nil, fmt.Errorf("unknown Timestamp field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TimestampMutation) SetField(name string, value ent.Value) error {
	switch name {
	case timestamp.FieldStartTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartTime(v)
		return nil
	case timestamp.FieldEndTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEndTime(v)
		return nil
	case timestamp.FieldActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActive(v)
		return nil
	case timestamp.FieldComment:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetComment(v)
		return nil
	case timestamp.FieldCategory:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategory(v)
		return nil
	case timestamp.FieldProject:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProject(v)
		return nil
	}
	return fmt.Errorf("unknown Timestamp field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TimestampMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TimestampMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TimestampMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Timestamp numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TimestampMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(timestamp.FieldComment) {
		fields = append(fields, timestamp.FieldComment)
	}
	if m.FieldCleared(timestamp.FieldCategory) {
		fields = append(fields, timestamp.FieldCategory)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TimestampMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TimestampMutation) ClearField(name string) error {
	switch name {
	case timestamp.FieldComment:
		m.ClearComment()
		return nil
	case timestamp.FieldCategory:
		m.ClearCategory()
		return nil
	}
	return fmt.Errorf("unknown Timestamp nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TimestampMutation) ResetField(name string) error {
	switch name {
	case timestamp.FieldStartTime:
		m.ResetStartTime()
		return nil
	case timestamp.FieldEndTime:
		m.ResetEndTime()
		return nil
	case timestamp.FieldActive:
		m.ResetActive()
		return nil
	case timestamp.FieldComment:
		m.ResetComment()
		return nil
	case timestamp.FieldCategory:
		m.ResetCategory()
		return nil
	case timestamp.FieldProject:
		m.ResetProject()
		return nil
	}
	return fmt.Errorf("unknown Timestamp field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TimestampMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TimestampMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TimestampMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TimestampMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TimestampMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TimestampMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TimestampMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Timestamp unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TimestampMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Timestamp edge %s", name)
}
