package weightsstat

import "time"

func (w *Info) GetID() string {
	return w.ID
}

func (w *Info) GetCreatedAt() time.Time {
	return w.CreatedAt
}

func (w *Info) GetUpdatedAt() time.Time {
	return w.UpdatedAt
}
