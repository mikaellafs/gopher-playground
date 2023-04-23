package config

type Configuration struct {
	Server *Server `yaml:"server"`
	Auth   *Auth   `yaml:"auth"`
}

type Server struct {
	Port       string  `yaml:"port"`
	RateLimit  int     `yaml:"rate_limit"`
	RetryAfter float64 `yaml:"retry_after"`
}

type Auth struct {
	Mode string `yaml:"mode"`
}
