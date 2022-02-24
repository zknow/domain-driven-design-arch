package mail

import (
	"strings"

	"github.com/zknow/my-arch/config"
	"gopkg.in/gomail.v2"
)

// Setting Security:
// https://www.google.com/settings/security/lesssecureapps

type Options struct {
	MailTo  string
	Subject string
	Body    string
}

func Send(o *Options) error {
	cfg := config.GetConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.Mail.Account)
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)
	m.SetHeader("Subject", o.Subject)
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.Account, cfg.Mail.Passwd)

	return d.DialAndSend(m)
}
