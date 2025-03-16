package domain

import "time"

type Base struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func (b *Base) SetID(id string) {
	b.id = id
}

func (b *Base) GetID() string {
	return b.id
}

func (b *Base) SetCreatedAt(createdAt time.Time) {
	b.createdAt = createdAt
}

func (b *Base) GetCreatedAt() time.Time {
	return b.createdAt

}

func (b *Base) SetUpdatedAt(updatedAt time.Time) {
	b.updatedAt = updatedAt
}

func (b *Base) GetUpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Base) SetDeletedAt(deletedAt *time.Time) {
	b.deletedAt = deletedAt
}

func (b *Base) GetDeletedAt() *time.Time {
	return b.deletedAt
}
