package mailer

type Mailer interface {
	Send(config *EmailConfig) error
}

type EmailConfig struct {
	From       string
	To         string
	Subject    string
	Content    string
	Template   interface{}
	TemplateId string
	Data       map[string]interface{}
}

func NewEmailConfigWithDynamicTemplate(from, to, subject, templateId string, data map[string]interface{}) *EmailConfig {
	return &EmailConfig{
		From:       from,
		To:         to,
		Subject:    subject,
		TemplateId: templateId,
		Data:       data,
	}
}
