package config

type Template struct {
	WorkDirectory string `yaml:"work_directory"`

	Web      WebTemplate      `yaml:"web"`
	Database DatabaseTemplate `json:"database"`
}

type WebTemplate struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`

	Mode string `yaml:"mode"`
}

type DatabaseTemplate struct {
	MySQL MySQLTemplate `yaml:"mysql"`
}

type MySQLTemplate struct {
	Address  string `yaml:"address"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
