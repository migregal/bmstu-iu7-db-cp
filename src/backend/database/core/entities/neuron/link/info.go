package link

type Link struct {
	ID        string `gorm:"-"`
	InnerID   string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	From      string `gorm:"type:uuid;column:from_id;"`
	To        string `gorm:"type:uuid;column:to_id;"`
}

func (Link) TableName() string {
	return "neuron_links"
}
