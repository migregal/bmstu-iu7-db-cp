package link

type Link struct {
	StructureID string `gorm:"type:uuid;column:structure_id;"`
	LinkID      string `gorm:"type:uuid;column:link_id;"`
	FromID      string `gorm:"type:uuid;column:from_id;"`
	ToID        string `gorm:"type:uuid;column:to_id;"`
}
