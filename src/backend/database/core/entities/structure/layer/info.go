package layer

type Layer struct {
	ID             string `gorm:"primaryKey;type:uuid;column:id;default:generated();"`
	StructID       string `gorm:"type:uuid;column:struct_id;"`
	LimitFunc      string `gorm:"column:limit_func"`
	ActivationFunc string `gorm:"column:activetion_func"`
}
