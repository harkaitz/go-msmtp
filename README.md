GO MSMTP
========

Small library for sending mails using msmtp(1) and file(1) in Go.

## Go documentation

    package msmtp // import "github.com/harkaitz/go-msmtp"
    
    var Verbose bool = false
    func Mime(file string) (mime string, err error)
    func Send(to string, m Message) (err error)
    type Attachment struct{ ... }
    type Message struct{ ... }

## Go type Message

    package msmtp // import "."
    
    type Message struct {
        FromAccount string       // Optional, otherwise default account is used.
        Subject     string       // Optional.
        Body        string       // Optional, by default an empty string.
        UseHTML     bool         // Optional, by default false.
        Attachments []Attachment // Optional.
    }
        Message is a struct that contains the information needed to send an email
        using MSMTP.
    

## Go type Attachment

    package msmtp // import "."
    
    type Attachment struct {
        Name string // Required, some file name.
        Mime string // Required, you can calculate it with Mime().
        Data []byte // Optional
    }
        Attachment is a struct that contains the information needed to attach a file
        to an email.
    

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
