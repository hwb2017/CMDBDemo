package config

type DatabaseConfig struct {
	Username string
	Password string
	Host string
    DBName string
}

type ServerConfig struct {
	Host string
}

type LogConfig struct {
	Level string
	Path string
}

func (c *Configuration) ReadSection(k string, v interface{}) error {
	err := c.vp.UnmarshalKey(k, v)
	if err != nil {
	    return err
	}
	return nil
}