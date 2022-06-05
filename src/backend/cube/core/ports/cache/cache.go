package cache

type CacheInteractor interface {
	Update(storage string, id string, info interface{}) error
	Get(storage string, id string) ([]interface{}, error)
	Delete(storage string, id string) error
}
