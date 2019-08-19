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
	lenArch:=len(archz)-1
	for j := 0; j < lenArch; j++ {
	arch:=strings.Split(archz[j],"|")
	fmt.Println(arch[0],arch[1],arch[2],arch[3])
}
}
