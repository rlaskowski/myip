package email

import (
	"net/smtp"

	"github.com/rlaskowski/myip/config"
)

type SMTPServer struct {
}

func (s *SMTPServer) LoginAuth(config *config.Config) smtp.Auth {
	return LoginAuth(config.Email.Username, config.Email.Password)
}

func (s *SMTPServer) PlainAuth(config *config.Config) smtp.Auth {
	return smtp.PlainAuth("", config.Email.Username, config.Email.Password, config.Email.SMTP.Hostname)
}

func (s *SMTPServer) CRAMMD5Auth(config *config.Config) smtp.Auth {
	return smtp.CRAMMD5Auth(config.Email.Username, config.Email.Password)
}
