package link

type Link struct {
	ID        string `gorm:"type:uuid;column:link_id;"`
	Structure string `gorm:"type:uuid;column:structure_id;"`
	From      string `gorm:"type:uuid;column:from_id;"`
	To        string `gorm:"type:uuid;column:to_id;"`
}

func (Link) TableName() string {
	return "neuron_links"
}
