package db

type Config struct {
	User string
	Pass string
	Adds string
	Name string
}

func NewConfig() *Config {
	return &Config{
		User: "root",
		Pass: "toor",
		Adds: "127.0.0.1",
		Name: "oficina",
	}
}
