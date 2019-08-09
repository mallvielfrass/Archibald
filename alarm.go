package main

import (
//	"fmt"
	"math/rand"
	"regexp"
//	"strings"
)

//func WordParse(messageText string) bool {
//	splittedMessage := []rune((strings.Split(messageText, "ент"))[0])
//	//	fmt.Println("splittedMessage[0]: ",splittedMessage)
//	lenSplittedMessage := len(splittedMessage)
//	fmt.Println("lenSplittedMessage:", lenSplittedMessage)
//	if lenSplittedMessage > 1 {
//		Shift := string(splittedMessage[lenSplittedMessage-2])
//		fmt.Println("lenSplittedMessage:", lenSplittedMessage, " |Shift:", Shift)
//		if Shift != " " {
//			return false
//		}
//	}
//
//	return true
//}
func GentooAlarm(messageText string) (bool, string) {
//	regexpRule, _ := regexp.Compile("\\B[Гг]ент")

regexpRule, _ := regexp.Compile("(?:[^A-Za-zА-Яа-яЁё0-9_]|^)[Гг]ент|[Gg]entoo")
regxpValue := regexpRule.MatchString(messageText)
	if regxpValue == true {
	//	ParseResult := WordParse(messageText)
	//	if ParseResult == true {
			ReplyLibrary := []string{"бан", "ловите гентушника", "бан за генту", "ребят, бьем его"}
			message := ReplyLibrary[rand.Intn(len(ReplyLibrary))]
			return true, message
	//	}
	}

	return false, ""
}
