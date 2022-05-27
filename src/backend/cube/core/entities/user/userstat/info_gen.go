package userstat

import "time"

func (i *Info) ID() string {
	return i.id
}

func (i *Info) CreatedAt() time.Time {
	return i.registeredAt
}

func (i *Info) UpdatedAt() time.Time {
	return i.updatedAt
}
