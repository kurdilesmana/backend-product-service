package config

type Config struct {
	AppName            string `json:"app_name,omitempty"`
	HTTPPort           int    `json:"http_port,omitempty"`
	HTTPMaxRequestTime int    `json:"max_request_time,omitempty"`
}

type PgsqlConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Schema          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

type Database struct {
	Engine          string `json:"engine,omitempty"`
	Host            string `json:"host,omitempty"`
	Port            int    `json:"port,omitempty"`
	Username        string `json:"username,omitempty"`
	Password        string `json:"-"`
	DBName          string `json:"database_name,omitempty"`
	Schema          string `json:"schema,omitempty"`
	MaxIdle         int    `json:"max_idle,omitempty"`
	MaxConn         int    `json:"max_conn,omitempty"`
	ConnMaxLifetime int    `json:"conn_max_lifetime,omitempty"`
}

func (d *Database) KbDBConfig() *PgsqlConfig {
	return &PgsqlConfig{
		Host:            d.Host,
		Port:            d.Port,
		User:            d.Username,
		Password:        d.Password,
		Schema:          d.Schema,
		MaxIdleConns:    d.MaxIdle,
		MaxOpenConns:    d.MaxConn,
		ConnMaxLifetime: d.ConnMaxLifetime,
	}
}

type Redis struct {
	Host         string `json:"host,omitempty"`
	Port         int    `json:"port,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"-"`
	DB           int    `json:"db,omitempty"`
	UseTLS       bool   `json:"use_tls,omitempty"`
	MaxRetries   int    `json:"max_retries"`
	MinIdleConns int    `json:"min_idle_conns"`
	PoolSize     int    `json:"pool_size"`
	PoolTimeout  int    `json:"pool_timeout"`
	MaxConnAge   int    `json:"max_conn_age"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}
