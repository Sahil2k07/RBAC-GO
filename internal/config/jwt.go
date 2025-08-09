package config

type JWTConfig struct {
	CookieName string `toml:"cookie_name"`
	Secret     string `toml:"secret"`
}

func GetJWTConfig() JWTConfig {
	return JWTConfig{
		CookieName: globalConfig.JWT.CookieName,
		Secret:     globalConfig.JWT.Secret,
	}
}
