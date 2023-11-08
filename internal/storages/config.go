package storages

type PostgresConfig struct {
	Host     string `yaml:"db_host" env-required:"true"`
	Port     int    `yaml:"db_port" env-required:"true"`
	User     string `yaml:"db_user" env-required:"true"`
	Password string `yaml:"db_password" env-required:"true"`
	DBname   string `yaml:"db_name" env-required:"true"`
}
