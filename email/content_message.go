package email

type ContentMessage struct {
	HTMLType bool   `json:"html_type"`
	Data     []byte `json:"data"`
}
