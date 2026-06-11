package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("timer <süre>")
	} else if len(os.Args) == 2 {
		dakika := os.Args[1]
		dondurulenSaniye := zamanAl(dakika)
		zamanYazdir(dondurulenSaniye)
	}
}
func zamanAl(sure string) (saniye int) {
	dakika, err := strconv.Atoi(sure)
	if err != nil {
		fmt.Println("Süre çevirilemedi")
		os.Exit(1)
	}
	saniye = dakika * 60
	return
}
func zamanYazdir(saniye int) {

}
