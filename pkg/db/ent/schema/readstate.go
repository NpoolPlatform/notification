package schema

import "entgo.io/ent"

// ReadState holds the schema definition for the ReadState entity.
type ReadState struct {
	ent.Schema
}

// Fields of the ReadState.
func (ReadState) Fields() []ent.Field {
	return nil
}

// Edges of the ReadState.
func (ReadState) Edges() []ent.Edge {
	return nil
}
