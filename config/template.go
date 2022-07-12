package config

type Template struct {
	WorkDirectory string `yaml:"work_directory"`

	Web WebTemplate `yaml:"web"`
}

type WebTemplate struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`

	Mode string `yaml:"mode"`
}
