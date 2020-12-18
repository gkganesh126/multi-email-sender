package models

type (
	Email struct {
		Provider []string  `json:"provider"`
		From     Contact   `json:"from"`
		To       []Contact `json:"to"`
		Cc       []Contact `json:"cc"`
		Bcc      []Contact `json:"bcc"`
		Subject  string    `json:"subject"`
		Message  string    `json:"message"`
	}
	Contact struct {
		Name    string `json:"name"`
		EmailID string `json:"emailID"`
	}
)
