package mail

import (
	"log"
	"net/smtp"
	"net/textproto"

	"github.com/pulsar-go/pulsar/config"

	"github.com/jordan-wright/email"
)

// Mail exposes the Mail type.
type Mail struct {
	Message *email.Email
	// Reserved for future things.
}

// Create creates a new empty mail.
func Create() *Mail {
	mail := &Mail{Message: email.NewEmail()}
	return mail.From(config.Settings.Mail.From)
}

// ReplyTo determines where the message will get replied to.
func (m *Mail) ReplyTo(replyTo []string) *Mail {
	m.Message.ReplyTo = replyTo
	return m
}

// From determines who the message is from (There is already a default value provided in the config here).
func (m *Mail) From(from string) *Mail {
	m.Message.From = from
	return m
}

// To determines who the message is for.
func (m *Mail) To(to ...string) *Mail {
	m.Message.To = to
	return m
}

// Bcc determines the Blind carbon copy.
func (m *Mail) Bcc(bcc ...string) *Mail {
	m.Message.Bcc = bcc
	return m
}

// Cc determines the Carbon copy.
func (m *Mail) Cc(cc ...string) *Mail {
	m.Message.Cc = cc
	return m
}

// Subject determines the Carbon copy.
func (m *Mail) Subject(subject string) *Mail {
	m.Message.Subject = subject
	return m
}

// Text determines the message plaintext (optional).
func (m *Mail) Text(text string) *Mail {
	m.Message.Text = []byte(text)
	return m
}

// HTML determines the message HTML (optional).
func (m *Mail) HTML(html string) *Mail {
	m.Message.HTML = []byte(html)
	return m
}

// Sender override From as SMTP envelope sender (optional).
func (m *Mail) Sender(sender string) *Mail {
	m.Message.Sender = sender
	return m
}

// Headers determine the message headers.
func (m *Mail) Headers(headers textproto.MIMEHeader) *Mail {
	m.Message.Headers = headers
	return m
}

// AttachFile attaches a file to the mail.
func (m *Mail) AttachFile(filename string) *Mail {
	_, err := m.Message.AttachFile(filename)
	if err != nil {
		log.Println(err)
	}
	return m
}

// Send sends the mail.
func (m *Mail) Send() error {
	// Copy to keep the code length considerable.
	s := &config.Settings.Mail
	return m.Message.Send(s.Host+":"+s.Port, smtp.PlainAuth(s.Identity, s.Username, s.Password, s.Host))
}
