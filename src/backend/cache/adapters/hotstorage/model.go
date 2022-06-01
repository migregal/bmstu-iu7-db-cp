package hotstorage

func (c *Cache) Update(store string, id string, info interface{}) error {
	return c.i.Upsert(store, []interface{}{id, info}, []interface{}{[]interface{}{"=", 1, info}})
}

func (c *Cache) Get(store string, id string) ([]interface{}, error) {
	return c.i.Get(store, []interface{}{id})
}

func (c *Cache) Delete(store string, id string) error {
	return c.i.Delete(store, []interface{}{id})
}
