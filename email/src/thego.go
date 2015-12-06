package main

import (
	"fmt"
	"net/mail"
	"os"
	"github.com/jhillyerd/go.enmime"


)

func main(){


	file, err := os.Open("/home/snouto/Desktop/qp-utf8-header.raw")

	if  err != nil {
		fmt.Println("There was a problem opening the file")
		panic("Error opening the file")
		return
	}

	msg, _ := mail.ReadMessage(file)     // Read email using Go's net/mail

	mime, _ := enmime.ParseMIMEBody(msg) // Parse message body with enmime

	// Headers are in the net/mail Message
	fmt.Printf("From: %v\n", msg.Header.Get("From"))

	// enmime can decode quoted-printable headers
	fmt.Printf("Subject: %v\n", mime.GetHeader("Subject"))

	// The plain text body is available as mime.Text
	fmt.Printf("Text Body: %v chars\n", len(mime.Text))
	fmt.Printf("Reply-To: %v \n",mime.GetHeader("Reply-To"))

	// The HTML body is stored in mime.Html
	fmt.Printf("HTML Body: %v chars\n", len(mime.Html))

	// mime.Inlines is a slice of inlined attacments
	fmt.Printf("Inlines: %v\n", len(mime.Inlines))


	// mime.Attachments contains the non-inline attachments
	fmt.Printf("Attachments: %v\n", len(mime.Attachments))


	fmt.Printf("Content Type : %v\n",mime.GetHeader("Content-Type"))
	fmt.Printf("HTML Body :\n%v\n",mime.Html)

	// Output:
	// From: James Hillyerd <jamehi03@jamehi03lx.noa.com>
	// Subject: MIME UTF8 Test ¢ More Text
	// Text Body: 1300 chars
	// HTML Body: 1736 chars
	// Inlines: 0
	// Attachments: 0

}


