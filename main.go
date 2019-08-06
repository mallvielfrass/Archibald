package main

import (
	"log"
	"net/http"
	"strings"
	"fmt"
	"math/rand"
	"io/ioutil"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)
func alarm(word string)(string,bool){
	x:=false
	
	gentlen:=len(strings.Split(word,"гент"))
           fmt.Println("glen",gentlen)
           banword:=[]string{"бан","ловите гентушника","бан за генту"}
           if gentlen>1 {
            x=true	
		   word=banword[rand.Intn(len(banword))]
		   fmt.Println(word)
           }



return	word,x
}
func main() {
	pconf, err := ioutil.ReadFile("proxy.conf")
            if err != nil {
		    log.Fatalln("proxy config read:", err)
            }
          proxyD := (strings.Split(string(pconf), "\n"))[0]

	auth := proxy.Auth{
	    User:     proxyD[0],
	    Password: proxyD[1],
	}

	dialer, err := proxy.SOCKS5("tcp", proxyD[2], &auth, proxy.Direct)
		if err != nil {
		    fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		}

	tr := &http.Transport{Dial: dialer.Dial}
	myClient := &http.Client{
	    Transport: tr,
	}
	ready, err := ioutil.ReadFile("token.pem")
          if err != nil {
                  log.Fatalln("Не удалось считать ключ из файла:", err)
          }
        key := (strings.Split(string(ready), "\n"))[0]
        bot, err := tgbotapi.NewBotAPI(key, client *http.Client)
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
	// gentlen:=len(strings.Split(mess,"гент"))
	// fmt.Println("glen",gentlen)
	 //banword:=[]string{"бан","ловите гентушника","бан за генту"}
	// if gentlen>1 {
//		 mess=banword[rand.Intn(len(banword))]
//	 }
		  fmt.Println("mess:",mess)
go fmt.Println(alarm(update.Message.Text))                 // messtext:="test"
                  log.Printf("Message: [%s] %s %s", update.Message.From.UserName, update.Message.Text)
                  //bot.GetFileDirectURL(
                  text:= mess
                  msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
                  //msg.ReplyToMessageID = update.Message.MessageID
                  bot.Send(msg)
          }
  }


