package envconfigs

type S3Config interface {
	BucketName() string
	AccessKey() string
	SecretKey() string
	Region() string
	Domain() string
}

type s3Config struct {
	env *envConfigs
}

func (cfg *s3Config) BucketName() string {
	return cfg.env.S3BucketName
}

func (cfg *s3Config) AccessKey() string {
	return cfg.env.S3AccessKey
}

func (cfg *s3Config) SecretKey() string {
	return cfg.env.S3SecretKey
}

func (cfg *s3Config) Region() string {
	return cfg.env.S3Region
}

func (cfg *s3Config) Domain() string {
	return cfg.env.S3Domain
}
