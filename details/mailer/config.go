package mailer

type Config struct {
	Server string
}

func NewConfig() *Config {
	return &Config{
		Server: "smtp.gmail.com..",
	}
}


