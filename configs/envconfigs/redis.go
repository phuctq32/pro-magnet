package envconfigs

type RedisConfig interface {
	Address() string
}

type redisConfig struct {
	env *envConfigs
}

func (cfg *redisConfig) Address() string {
	return cfg.env.RedisAddr
}
