package modelstat

import "time"

func (i *Info) GetID() string {
	return i.ID
}

func (i *Info) GetCreatedAt() time.Time {
	return i.CreatedAt
}

func (i *Info) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}
