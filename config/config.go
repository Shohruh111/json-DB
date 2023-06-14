package config

type Config struct {
	DefaultOffset int
	DefaultLimit  int

	Path         string
	UserFileName string

	ProductFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Path = "./data"
	cfg.UserFileName = "/user.json"

	cfg.ProductFileName="/product.json"

	return cfg
}
