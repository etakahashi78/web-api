package entity

import "time"

type Todo struct {
	ID        int64
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewName(name, content string) (*Todo, error) {

	now := time.Now()
	return &Todo{
		Name:      name,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
