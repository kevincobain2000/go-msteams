package main

import (
	"flag"
	"log"

	"github.com/kevincobain2000/go-msteams/src"
)

var (
	hook     = flag.String("hook", "", "(REQUIRED) Description: MS Teams webhook URL")
	title    = flag.String("title", "My Title", "Description: Your title")
	subtitle = flag.String("subtitle", "My Summary", "Description: Your summary")
	subject  = flag.String("subject", "My Subject", "Description: Your subject")
	color    = flag.String("color", "", "Description: Your theme color") // unimplemented
	message  = flag.String("message", "My Message", "Description: Message body. HTML allowed.")
	proxy    = flag.String("proxy", "", "Description: Hit behind this proxy")
)

func main() {
	flag.Parse()

	log.Println("title:", *title)
	log.Println("subtitle:", *subtitle)
	log.Println("subject:", *subject)
	log.Println("color:", *color)
	log.Println("message:", *message)
	log.Println("hook:", *hook)
	log.Println("proxy:", *proxy)

	err := Send(*title, *subtitle, *subject, *message, *hook, *proxy)

	log.Println()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully sent!")
	}
}

func Send(title, subtitle, subject, message, hook, proxy string) (err error) {
	return src.Send(title, subtitle, subject, message, hook, proxy)
}
