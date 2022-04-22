package email

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"sync"

	"github.com/rlaskowski/myip/config"
)

type Email struct {
	config *config.Config
	smtp   *SMTPServer
	mutex  *sync.Mutex
}

func NewEmail(config *config.Config) *Email {
	return &Email{
		config: config,
		smtp:   &SMTPServer{},
		mutex:  &sync.Mutex{},
	}
}

//Sending email
func (e *Email) Send(msg *SmtpMessage) error {
	sn := msg.SenderName()
	msg.SetSender(sn, e.config.Email.Sender.Email)

	return e.send(msg)
}

func (e *Email) send(msg *SmtpMessage) error {
	auth := e.smtp.LoginAuth(e.config)

	recipients := strings.Split(msg.Recipients(), ",")

	mb, err := msg.Bytes()
	if err != nil {
		return err
	}

	err = smtp.SendMail(fmt.Sprintf("%s:%d", e.config.Email.SMTP.Hostname, e.config.Email.SMTP.Port), auth, msg.SenderAddress(), recipients, mb)
	if err != nil {
		log.Printf("Error when try to send email due to: %s", err)
		return err
	}

	log.Printf("Email was sended successful to %s", msg.Recipients())

	return nil
}
