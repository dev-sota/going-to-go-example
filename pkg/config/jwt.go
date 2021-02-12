package config

var JWT jwt

type jwt struct {
	Secret    string `env:"JWT_SECRET"`
	ExpireMin int    `env:"JWT_EXPIRE_MIN"`
}
