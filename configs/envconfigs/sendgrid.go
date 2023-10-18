package envconfigs

type SendgridConfig interface {
	ApiKey() string
	FromEmail() string
	VerifyEmailTemplateId() string
	VerificationURL() string
}

type sendgridConfig struct {
	env *envConfigs
}

func (cfg *sendgridConfig) ApiKey() string {
	return cfg.env.SendGridApiKey
}

func (cfg *sendgridConfig) FromEmail() string {
	return cfg.env.SendGridFromEmail
}

func (cfg *sendgridConfig) VerifyEmailTemplateId() string {
	return cfg.env.SendGridVerifyEmailTemplateId
}

func (cfg *sendgridConfig) VerificationURL() string {
	return cfg.env.VerificationURL
}
