package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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
	for i := saniye; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s %2d\r", "Kalan süre", i)
	}
}
