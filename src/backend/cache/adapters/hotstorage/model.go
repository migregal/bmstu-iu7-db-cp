package hotstorage

func (c *Cache) UpdateModelInfo(id string, info interface{}) error {
	return c.i.Upsert("model", []interface{}{id, info}, []interface{}{[]interface{}{"=", 1, info}})
}

func (c *Cache) GetModelInfo(id string) ([]interface{}, error) {
	return c.i.Get("model", []interface{}{id})
}

func (c *Cache) DeleteModelInfo(id string) error {
	return c.i.Delete("model", []interface{}{id})
}
