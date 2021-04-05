package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EntityID = uuid.UUID

type Entity interface {
	String() string
	ID() EntityID
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type entity struct {
	id        EntityID
	createdAt time.Time
	updatedAt time.Time
}

func (e *entity) String() string {
	return fmt.Sprintf("entity<id=%v, created_at=%v, updated_at=%v>", e.id, e.createdAt, e.updatedAt)
}

func (ent *entity) ID() uuid.UUID {
	return ent.id
}

func (ent *entity) CreatedAt() time.Time {
	return ent.createdAt
}

func (ent *entity) UpdatedAt() time.Time {
	return ent.updatedAt
}
