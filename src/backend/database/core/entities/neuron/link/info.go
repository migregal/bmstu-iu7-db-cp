package link

type Link struct {
	ID        int    `gorm:"column:link_id;"`
	From      int    `gorm:"-"`
	To        int    `gorm:"-"`
	InnerID   string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	InnerFrom string `gorm:"type:uuid;column:from_id;"`
	InnerTo   string `gorm:"type:uuid;column:to_id;"`
}

func (Link) TableName() string {
	return "neuron_links"
}
