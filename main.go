package main

import (
	"flag"
	"fmt"
	"net/smtp"
)

const host = "smtp.gmail.com"
const port = "587"

func main() {
	from := flag.String("from", "", "Sender email")
	password := flag.String("password", "", "Sender email password")
	to := flag.String("to", "", "Receiver email")
	msg := flag.String("msg", "", "Email message")
	help := flag.Bool("help", false, "Display guides")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}
	if *from != "" && *password != "" && *to != "" && *msg != "" {
		sendEmail(*from, *password, *to, *msg)
	}
}

func sendEmail(from string, password string, to string, msg string) {
	auth := smtp.PlainAuth("", from, password, host)
	receivers := []string{to}
	err := smtp.SendMail(host+":"+port, auth, from, receivers, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email send successfully :)")
}
