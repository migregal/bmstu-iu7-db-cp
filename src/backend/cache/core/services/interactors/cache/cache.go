package cache

import (
	"fmt"

	"github.com/tarantool/go-tarantool"
)

type Params struct {
	Host     string
	Port     string
	User     string
	Password string
}

type CacheInteractor struct {
	params Params
}

func New(p Params) CacheInteractor {
	return CacheInteractor{params: p}
}

func (i *CacheInteractor) newConnection() (*tarantool.Connection, error) {
	opts := tarantool.Opts{User: i.params.User, Pass: i.params.Password}
	conn, err := tarantool.Connect(fmt.Sprintf("%s:%s", i.params.Host, i.params.Port), opts)
	if err != nil {
		return nil, fmt.Errorf("cache error: %w", err)
	}

	return conn, nil
}

func (i *CacheInteractor) Upsert(space string, value []interface{}, ops []interface{}) error {
	conn, err := i.newConnection()
	if err != nil {
		return fmt.Errorf("cache error: %w", err)
	}
	defer conn.Close()

	_, err = conn.Upsert(space, value, ops)
	if err != nil {
		return fmt.Errorf("cache error: %w", err)
	}

	return nil
}

func (i *CacheInteractor) Get(space string, key []interface{}) ([]interface{}, error) {
	conn, err := i.newConnection()
	if err != nil {
		return nil, fmt.Errorf("cache error: %w", err)
	}
	defer conn.Close()

	resp, err := conn.Select(space, "primary", 0, 1, tarantool.IterEq, key)
	if err != nil {
		return nil, fmt.Errorf("cache error: %w", err)
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("no record found")
	}

		if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	tuple, ok := resp.Data[0].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}

	return tuple, nil
}

func (i *CacheInteractor) Delete(space string, key []interface{}) error {
	conn, err := i.newConnection()
	if err != nil {
		return fmt.Errorf("cache error: %w", err)
	}
	defer conn.Close()

	_, err = conn.Delete(space, "primary", key)
	if err != nil {
		return fmt.Errorf("cache error: %w", err)
	}

	return nil
}
