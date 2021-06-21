package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Timestamp holds the schema definition for the Timestamp entity.
type Timestamp struct {
	ent.Schema
}

// Fields of the Timestamp.
func (Timestamp) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("start_time").
			Default(time.Now),
		field.Time("end_time"),
		field.Bool("active").
			Default(true).
			StructTag(`json:"active"`),
		field.String("comment").
			Optional(),
		field.String("category").
			Optional(),
		field.String("project"),
	}
}

// Edges of the Timestamp.
func (Timestamp) Edges() []ent.Edge {
	return nil
}
