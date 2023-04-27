package config

type Configuration struct {
	Server *Server `yaml:"server"`
	Auth   *Auth   `yaml:"auth"`
}

type Server struct {
	Port       string  `yaml:"port"`
	RateLimit  int     `yaml:"rate_limit"`
	RetryAfter float64 `yaml:"retry_after"`
	Https      Https   `yaml:"https,omitempty"`
}

type Https struct {
	Enable   bool   `yaml:"enable"`
	CertPath string `yaml:"cert_path,omitempty"`
	KeyPath  string `yaml:"key_path,omitempty"`
}

type Auth struct {
	Mode string `yaml:"mode"`
}
