package main

import (
	"log"
	"net/http"
	"strings"
	"fmt"
	"math/rand"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
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
         // log.Printf("%+v\n", updates)
         log.Printf("Message: [%s] %s %s", update.Message.From.UserName, update.Message.Text, update.Message.Sticker)

	 mess:=update.Message.From.UserName+":\n "+update.Message.Text
	 gentlen:=len(strings.Split(mess,"гент"))
	 fmt.Println("glen",gentlen)
	 banword:=[]string{"бан","ловите гентушника","бан за генту"}
	 if gentlen>1 {
		 mess=banword[rand.Intn(len(banword))]
	 }
		  fmt.Println("mess:",mess)	  
                 // messtext:="test"
                  log.Printf("Message: [%s] %s %s", update.Message.From.UserName, update.Message.Text)                  
                  //bot.GetFileDirectURL(
                  text:= mess
                  msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
                  //msg.ReplyToMessageID = update.Message.MessageID
                  bot.Send(msg)
          }
  }

