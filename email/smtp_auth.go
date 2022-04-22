package email

import (
	"errors"
	"net/smtp"
)

type SmtpAuth struct {
	username, password string
}

func LoginAuth(username, password string) *SmtpAuth {
	return &SmtpAuth{username, password}
}

func (a *SmtpAuth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *SmtpAuth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unrecognized server in authlogin kind")
		}
	}
	return nil, nil
}
