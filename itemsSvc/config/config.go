package config

type Config struct {
	Database DbConfig `yaml:"database" env:"DATABASE"`
}

type DbConfig struct {
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"postgres"`
	DbName   string `yaml:"dbName" env:"DB_NAME" env-default:"postgres"`
	Host     string `yaml:"host" env:"DB_HOST" env-default:"127.0.0.1"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	SslMode  string `yaml:"ssl_mode" env:"SSL_MODE" env-default:"disabled"`
}
