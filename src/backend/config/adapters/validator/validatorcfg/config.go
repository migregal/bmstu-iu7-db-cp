package validatorcfg

type Config struct {
	MinUserNameLen int
	MaxUserNameLen int
	MinPasswordLen int
	MaxPasswordLen int
}

func (c *Config) IsMocked() bool {
	return false
}

func (c *Config) MinUnameLen() int {
	return c.MinUserNameLen
}

func (c *Config) MaxUnameLen() int {
	return c.MaxUserNameLen
}

func (c *Config) MinPwdLen() int {
	return c.MinPasswordLen
}
func (c *Config) MaxPwdLen() int {
	return c.MaxPasswordLen
}
