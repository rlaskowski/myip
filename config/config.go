package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RefreshTime int         `yaml:"refreshTime"`
	Email       EmailConfig `yaml:"email"`
}

type EmailConfig struct {
	Sender         Sender     `yaml:"sender"`
	Subject        string     `yaml:"subject"`
	RecipientEmail []string   `yaml:"recipientEmail"`
	SMTP           ServerInfo `yaml:"smtp"`
	POP3           ServerInfo `yaml:"pop3"`
	Username       string     `yaml:"username"`
	Password       string     `yaml:"password"`
}

type Sender struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type ServerInfo struct {
	Hostname   string `yaml:"hostname"`
	Port       int    `yaml:"port"`
	Encryption bool   `yaml:"encryption"`
}

func ConfigFile(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("config file doesn't exists")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file error: %s", err.Error())
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("data config file error: %s", err.Error())
	}

	cf := &Config{}

	if err := DeserializeFromYaml(cf, data); err != nil {
		return nil, fmt.Errorf("deserialize config file errror: %s", err.Error())
	}

	return cf, nil
}

func DeserializeFromYaml(value interface{}, buff []byte) error {
	return yaml.Unmarshal(buff, value)
}
