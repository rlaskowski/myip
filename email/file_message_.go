package email

type FileMessage struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}
