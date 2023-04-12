package config

type Configuration struct {
	Server *Server `yaml:"server"`
	Auth   *Auth   `yaml:"auth"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Auth struct {
	AuthMode string `yaml:"auth_mode"`
}
