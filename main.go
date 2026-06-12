package main

import (
	"fmt"
	"os"
	"os/exec"
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
	for i := saniye; i >= 0; i-- {
		saat := i / 3600
		dakika := (i % 3600) / 60
		saniye := i % 60
		fmt.Printf("%s %02d:%02d:%02d\r", "Kalan süre", saat, dakika, saniye)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\n%s", "TAMAMLANDI!")
	fmt.Println("\a")
	exec.Command("kdialog", "--title", "Timer", "--msgbox", "Süre doldu!").Run()
}
