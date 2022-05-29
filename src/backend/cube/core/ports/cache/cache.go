package cache

type CacheInteractor interface {
	UpdateModelInfo(id string, info interface{}) error
	GetModelInfo(id string) ([]interface{}, error)
	DeleteModelInfo(id string) error
}
