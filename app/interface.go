package app

type ServerOptions struct {
	Host   string `json:"host" yaml:"host" toml:"host"`
	Port   int    `json:"port" yaml:"port" toml:"port"`
	Prefix string `json:"prefix" yaml:"prefix" toml:"prefix"`
	Mode   string `json:"mode" yaml:"mode" toml:"mode"`
}

type Options struct {
	ServerOptions
}

type Server interface {
	Start() error
	Shotdown() error
}
