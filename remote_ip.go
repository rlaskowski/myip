package myip

import (
	"encoding/json"
	"fmt"
)

type RemoteIP struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	CC      string `json:"cc"`
}

func ParseIP(data []byte) (*RemoteIP, error) {
	rip := &RemoteIP{}

	if err := json.Unmarshal(data, rip); err != nil {
		return nil, fmt.Errorf("couldn't parse RemoteIP due to: %s", err.Error())
	}

	return rip, nil
}
