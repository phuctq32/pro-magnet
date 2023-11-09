package envconfigs

type MongoConfig interface {
	ConnectionString() string
	DBName() string
}

type mongoConfig struct {
	env *envConfigs
}

func (cfg *mongoConfig) ConnectionString() string {
	return cfg.env.MongoConnStr
}

func (cfg *mongoConfig) DBName() string {
	return cfg.env.MongoDBName
}
