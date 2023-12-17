package config

type config struct {
	BaseUrl          string `yaml:"base_url"`
	ServerPort       string `yaml:"server_port"`
	MySqlDsn         string `yaml:"mysql_dsn"`
	AllowedOrigins   string `yaml:"allowed_origins"`
	AllowAutoMigrate bool   `yaml:"allow_auto_migrate"`
}
