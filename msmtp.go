package msmtp

import (
	"encoding/base64"
	"os/exec"
	"strings"
	"path"
	"fmt"
	"log"
	"github.com/google/uuid"
)

// Message is a struct that contains the information needed to send an email
// using MSMTP.
type Message struct {
	FromAccount string       // Optional, otherwise default account is used.
	Subject     string       // Optional.
	Body        string       // Optional, by default an empty string.
	UseHTML     bool         // Optional, by default false.
	Attachments []Attachment // Optional.
}

// Attachment is a struct that contains the information needed to attach a file
// to an email.
type Attachment struct {
	Name   string // Required, some file name.
	Mime   string // Required, you can calculate it with Mime().
	Data   []byte // Optional
}

// Verbose is a global variable that can be set to true to enable verbose mode.
// In verbose mode, the program will print debug information to the standard
// output.
var Verbose bool = false

// Send sends an email using MSMTP.
func Send(to string, m Message) (err error) {
	var proc    *exec.Cmd
	var code     string
	var boundary string = uuid.New().String()
	var i        int
	
	code += "To: " + to + "\n"
	if m.Subject != "" {
		code += "Subject: " + m.Subject + "\n"
	}
	code += "MIME-Version: 1.0" + "\n"
	code += "Content-Type: multipart/mixed; boundary=\"" + boundary + "\"\n\n"
	code += "\n"
	code += "--" + boundary + "\n"
	if m.UseHTML {
		code += "Content-Type: text/html; charset=\"UTF-8\"\n"
	} else {
		code += "Content-Type: text/plain; charset=\"UTF-8\"\n"
	}
	code += "\n"
	if m.UseHTML {
		code += "<html>\n"
		code += "  <body>\n"
		code += "    <table>\n"
		code += "      <tbody>\n"
		code += m.Body
		code += "      </tbody>\n"
		code += "    </table>\n"
		code += "  </body>\n"
		code += "</html>\n"
	} else {
		code += m.Body
	}
	code += "\n\n"
	for i = range m.Attachments {
		var mime, name string
		mime = m.Attachments[i].Mime
		name = m.Attachments[i].Name
		if mime == "" {
			err = fmt.Errorf("Missing MIME type for attachment %d", i)
			return
		}
		if name == "" {
			err = fmt.Errorf("Missing name for attachment %d", i)
			return
		}
		
		code += "--" + boundary + "\n"
		code += "Content-Disposition: attachment; filename=\"" + strings.TrimSpace(path.Base(name)) +"\"\n"
		code += "Content-Type: " + mime + "\n"
		code += "Content-Transfer-Encoding: base64\n"
		code += "\n"
		if m.Attachments[i].Data != nil {
			code += base64.StdEncoding.EncodeToString(m.Attachments[i].Data)
		}
	}
	code += "--" + boundary + "--\n"
	
	if Verbose {
		log.Printf("go-msmtp: Sending email to %s ...\n", to)
	}
	if m.FromAccount != "" {
		proc = exec.Command("msmtp", "-t", "-a", m.FromAccount)
	} else {
		proc = exec.Command("msmtp", "-t")
	}
	proc.Stdin = strings.NewReader(code)
	err = proc.Run()
	if err != nil { return err }
	if Verbose {
		log.Printf("go-msmtp: Email sent to %s\n", to)
	}
	
	return
}

// Mime returns the MIME type of a file using file(1).
func Mime(file string) (mime string, err error) {
	var proc *exec.Cmd
	var data []byte
	proc = exec.Command("file", "--mime-type", file)
	data, err = proc.Output()
	if err != nil { return }
	mime = strings.Split(string(data), ": ")[1]
	mime = strings.TrimSpace(mime)
	return
}
