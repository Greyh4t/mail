package mail

import (
	"net/smtp"
	"strings"
)

type Mail struct {
	mailAddr string
	mailHost string
	auth     smtp.Auth
}

func NewMail(mailAddr, mailPass, mailHost string) *Mail {
	return &Mail{
		mailAddr: mailAddr,
		mailHost: mailHost,
		auth:     smtp.PlainAuth("", mailAddr, mailPass, strings.Split(mailHost, ":")[0]),
	}
}

func (m *Mail) Send(from string, to []string, subject, body string, isHtml bool) error {
	var contentType string
	if isHtml {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}

	if from != "" {
		from = from + "<" + m.mailAddr + ">"
	}

	msg := `To: ` + strings.Join(to, ";") + "\r\n" +
		`From: ` + from + "\r\n" +
		`Subject: ` + subject + "\r\n" +
		contentType + "\r\n\r\n" +
		body

	return smtp.SendMail(m.mailHost, m.auth, m.mailAddr, to, []byte(msg))
}
