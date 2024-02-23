package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/pborman/getopt/v2"
	"github.com/harkaitz/go-msmtp"
)

const help string =
`Usage: go-msmtp -v -t TO -f FROM -s SUBJECT -a ATTACH,... -m MESSAGE -T

Send mail using MSMTP.

Copyright (c) 2024 - Harkaitz Agirre - All rights reserved.`

func main() {
	var err     error
	var msg     msmtp.Message
	var arg     string
	var data  []byte
	var file   *os.File
	var mime    string
	
	// Error manager.
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "go-msmtp: error: %v\n", err.Error())
			os.Exit(1)
		}
	}()
	
	// Parse command line arguments.
	hFlag := getopt.BoolLong("help", 'h')
	vFlag := getopt.BoolLong("version", 'v')
	tFlag := getopt.StringLong("to", 't', "")
	fFlag := getopt.StringLong("from", 'f', "")
	sFlag := getopt.StringLong("subject", 's', "")
	aFlag := getopt.StringLong("attach", 'a', "")
	mFlag := getopt.StringLong("message", 'm', "")
	html  := getopt.BoolLong("html", 'T')
	getopt.SetUsage(func() { fmt.Println(help) })
	getopt.Parse()
	if *hFlag { getopt.Usage(); return }
	if *tFlag == "" { err = fmt.Errorf("Missing -t argument"); return }
	
	// Prepare message.
	msg.FromAccount = *fFlag
	msg.Subject  = *sFlag
	msg.Body = *mFlag
	msg.UseHTML = *html
	
	if *aFlag != "" {
		for _, arg = range strings.Split(*aFlag, ",") {
		
			file, err = os.Open(arg);
			if err != nil { return }
			defer file.Close()
			
			data, err = os.ReadFile(arg)
			if err != nil { return }
			
			mime, err = msmtp.Mime(arg)
			if err != nil { return }
			
			msg.Attachments = append(
				msg.Attachments,
				msmtp.Attachment{
					Name: arg,
					Mime: mime,
					Data: data,
				},
			)
			
		}
	}
	
	msmtp.Verbose = *vFlag
	err = msmtp.Send(*tFlag, msg)
	if err != nil { return }
	
	return
}
