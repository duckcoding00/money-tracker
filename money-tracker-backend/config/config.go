package config

type Config struct {
	PortAddress  string
	RedisAddress string
	DbConfig     DBConfig
	JwtConfig    JwtConfig
}

type DBConfig struct {
	DbAddr      string
	MaxOpenCons int
	MaxIdleCons int
	MaxIdleTime string
}

type JwtConfig struct {
	Secret string
}
