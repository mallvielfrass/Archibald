package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	//	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	//	"math/rand"
	"net/http"
	//	"os"
	"strings"
)

func main() {
	//	pconf, err := ioutil.ReadFile("proxy.conf")
	//	if err != nil {
	//		log.Fatalln("proxy config read:", err)
	//	}
	//	proxyD := strings.Split(string(pconf), "\n")

	//	auth := proxy.Auth{
	//		User:     proxyD[0],
	//		Password: proxyD[1],
	//	}

	//	dialer, err := proxy.SOCKS5("tcp", proxyD[2], &auth, proxy.Direct)
	//	if err != nil {
	//		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
	//	}
	//
	//	tr := &http.Transport{Dial: dialer.Dial}
	//	client := &http.Client{
	//		Transport: tr,
	//	}
	ready, err := ioutil.ReadFile("token.pem")
	if err != nil {
		log.Fatalln("Не удалось считать ключ из файла:", err)
	}
	key := (strings.Split(string(ready), "\n"))[0]
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://vielfrassx.tk:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
	for update := range updates {
		log.Printf("update")
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		fmt.Println(tgbotapi.Chat.ChatConfig(update.Message.Chat.ID))
		log.Printf("Message: [%s] %s %s", update.Message.From.UserName, update.Message.Text, update.Message.Sticker)
//if update.Message.Caption != "" { // ignore any non-Message Updates
//	                         fmt.Println("update.Message.Caption: ",update.Message.Caption)
  //       }

		mess := update.Message.From.UserName + ":\n " + update.Message.Text
		if update.Message.Caption != "" { // ignore any non-Message Updates
                                   fmt.Println("update.Message.Caption: ",update.Message.Caption)
           mess=mess+"\n"+"update.Message.Caption: "+update.Message.Caption
			   }
	//	fmt.Println(bot.ChatConfig())
		fmt.Println("mess:", mess)
		retbool, rettext := GentooAlarm(mess)
		fmt.Println("retbool: ", retbool, " | rettext: ", rettext)
		if retbool == true {
			truemsg := tgbotapi.NewMessage(update.Message.Chat.ID, rettext)
			bot.Send(truemsg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mess)
			bot.Send(msg)
		}
		//	fmt.Println("return text bool: ",rettext,retbool)
		//	log.Printf("Message: [%s] %s %s", update.Message.From.UserName, update.Message.Text)
		///	//bot.GetFileDirectURL(
		//	text := mess
		//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		//msg.ReplyToMessageID = update.Message.MessageID
		//	bot.Send(msg)
	}
}
