package config

type Config struct {
	Platform Platform `mapstructure:"platform"`
	Domain   Domain   `mapstructure:"domain"`
}

type Platform struct {
	Server Server `mapstructure:"web_server"`
}

type Domain struct{}

type Server struct {
	Port string `mapstructure:"port"`
}
