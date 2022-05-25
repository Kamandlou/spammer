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
	repeat := flag.Int("repeat", 1, "Send several times")
	help := flag.Bool("help", false, "Display guides")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}
	if *from != "" && *password != "" && *to != "" && *msg != "" {
		sendEmail(*from, *password, *to, *msg, *repeat)
	}
}

func sendEmail(from string, password string, to string, msg string, repeat int) {
	auth := smtp.PlainAuth("", from, password, host)
	receivers := []string{to}
	for i := 0; i < repeat; i++ {
		err := smtp.SendMail(host+":"+port, auth, from, receivers, []byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Email send successfully %d :)\n", i+1)
	}
}
