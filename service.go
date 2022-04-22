package myip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/rlaskowski/myip/config"
	"github.com/rlaskowski/myip/email"
)

var (
	addr   = "https://api.myip.com"
	rmu    = &sync.RWMutex{}
	ipInfo = make(map[string]*RemoteIP)
)

type Service struct {
	config *config.Config
	email  *email.Email
}

func NewService(config *config.Config) *Service {
	return &Service{
		config: config,
		email:  email.NewEmail(config),
	}
}

func (s *Service) Run() {
	for {
		<-time.After(time.Second * time.Duration(s.config.RefreshTime))

		rip, err := s.remoteIP()
		if err != nil {
			log.Printf("couldn't get remote IP")
			continue
		}

		if !s.checkIP(rip) {
			go func() {
				if err := s.sendEmail(rip); err != nil {
					log.Printf("Email wasn't send due to %s", err.Error())
				}
			}()
		}

	}
}

func (s *Service) remoteIP() (*RemoteIP, error) {
	res, err := http.Get(addr)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("body read error: %s", err.Error())
	}

	return ParseIP(body)
}

// Checking if the current IP was changed
func (s *Service) checkIP(rip *RemoteIP) bool {
	rmu.RLock()
	defer rmu.RUnlock()

	if _, ok := ipInfo[rip.IP]; !ok || len(ipInfo) == 0 {
		ipInfo[rip.IP] = rip
		return false
	}

	return true
}

func (s *Service) sendEmail(rip *RemoteIP) error {
	buff := &bytes.Buffer{}

	if _, err := fmt.Fprintf(buff, "Your IP address has been change to %s", rip.IP); err != nil {
		return err
	}

	content := &email.ContentMessage{
		Data: buff.Bytes(),
	}

	msg := email.NewSmtpMessage()

	for _, r := range s.config.Email.RecipientEmail {
		msg.AddRecipient(r)
	}

	msg.SetSubject(s.config.Email.Subject)
	msg.SetSender(s.config.Email.Sender.Name, s.config.Email.Sender.Email)
	msg.AddContent(content)

	return s.email.Send(msg)
}
