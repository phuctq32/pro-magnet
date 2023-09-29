package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"pro-magnet/components/mailer"
)

type sendGridProvider struct {
	ApiKey string
}

func NewSendGridProvider(key string) *sendGridProvider {
	return &sendGridProvider{ApiKey: key}
}

func (sg *sendGridProvider) Send(config *mailer.EmailConfig) error {
	client := sendgrid.NewSendClient(sg.ApiKey)

	fromEmail := mail.NewEmail("ProMagnet", config.From)
	toEmail := mail.NewEmail("", config.To)
	mailData := mail.NewV3Mail()
	mailData.SetFrom(fromEmail)
	mailData.SetTemplateID(config.TemplateId)

	personalization := mail.NewPersonalization()
	personalization.AddTos(toEmail)
	mailData.AddPersonalizations(personalization)
	config.Data["subject"] = config.Subject
	personalization.DynamicTemplateData = config.Data

	_, err := client.Send(mailData)
	if err != nil {
		return err
	}

	return nil
}
