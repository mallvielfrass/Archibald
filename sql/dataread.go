package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)




func main() {
	library, err := ioutil.ReadFile("archlinux_ru.txt")
	if err != nil {
		log.Fatalln(err)
	}
	archz := strings.Split(string(library), "\n")
	arch:=strings.Split(archz[0],"|")
	fmt.Println(arch[0])
}
