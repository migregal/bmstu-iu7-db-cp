package layer

type Layer struct {
	ID             int    `gorm:"-"`
	InnerID        string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	StructureID    string `gorm:"type:uuid;column:structure_id;"`
	LimitFunc      string `gorm:"column:limit_func"`
	ActivationFunc string `gorm:"column:activation_func"`
}
