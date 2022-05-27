package modelstat

import "time"

type Info struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
}

func New(id string, createdAt, updatedAt time.Time) *Info {
	return &Info{id: id, createdAt: createdAt, updatedAt: updatedAt}
}
