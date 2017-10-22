package config

type Config struct {
	Host        string `json:"host" yaml:"host,omitempty"`
	Port        int    `json:"port" yaml:"port,omitempty"`
	Environment string `json:"environment" yaml:"environment,omitempty"`
}
