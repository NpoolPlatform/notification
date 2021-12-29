package schema

import "entgo.io/ent"

// MailBox holds the schema definition for the MailBox entity.
type MailBox struct {
	ent.Schema
}

// Fields of the MailBox.
func (MailBox) Fields() []ent.Field {
	return nil
}

// Edges of the MailBox.
func (MailBox) Edges() []ent.Edge {
	return nil
}
