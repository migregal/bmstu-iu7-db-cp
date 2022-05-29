package userstat

import "time"

type Info struct {
	id           string
	registeredAt time.Time
	updatedAt    time.Time
}

func New(id string, registeredAt, updatedAt time.Time) *Info {
	return &Info{id: id, registeredAt: registeredAt, updatedAt: updatedAt}
}
