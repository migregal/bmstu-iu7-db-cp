package weights

type Weights struct {
	ID   string `gorm:"type:uuid;column:id;default:generated();"`
	Name string `gorm:"column:name"`
}

func (Weights) TableName() string {
	return "weights_info"
}
