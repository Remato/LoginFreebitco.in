package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sclevine/agouti"
)

var reconnect = false

func click(username string, password string, profile string) {

	fmt.Printf("[Carregou Driver]::%s\n", username)

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{profile, "--no-first-run", "--no-sandbox", "--headless"}),
	)
	//RETIRE O --HEADLESS CASO QUEIRA VER A EXECUÇÃO DA AUTOMAÇÃO DO BROWSER

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()

	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	if err := page.Navigate("https://freebitco.in/"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	fmt.Printf("[Entrou no Site]::%s\n", username)
	time.Sleep(2 * time.Second)
	page.FindByLink("LOGIN").Click()
	time.Sleep(2 * time.Second)
	page.Find("#login_form_btc_address").Fill(username)
	page.Find("#login_form_password").Fill(password)
	page.FindByID("login_button").Click()
	time.Sleep(2 * time.Second)
	page.FindByButton("ROLL!").Click()

	for {
		time.Sleep(3605 * time.Second)
		page.FindByButton("ROLL!").Click()
	}
}

func clickWithCaptcha(username string, password string, profile string) {

	fmt.Printf("[Carregou Driver]::%s\n", username)

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{profile, "--no-first-run", "--no-sandbox", "--headless"}),
	)

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()

	if err != nil {
		log.Fatal("Failed to open page:", err)
	}

	if err := page.Navigate("https://freebitco.in/"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	fmt.Printf("[Entrou no Site]::%s\n", username)
	time.Sleep(2 * time.Second)
	page.FindByLink("LOGIN").Click()
	time.Sleep(2 * time.Second)
	page.Find("#login_form_btc_address").Fill(username)
	page.Find("#login_form_password").Fill(password)
	page.FindByID("login_button").Click()
	time.Sleep(2 * time.Second)
	page.FindByID("play_without_captchas_button").Click()
	page.FindByButton("ROLL!").Click()

	for {
		time.Sleep(3605 * time.Second)
		page.FindByButton("ROLL!").Click()
	}
}

func main() {
	//aqui são 2 threads rodando em contas diferentes, pra fazer isso basta setar um profile de 'user' diferente do chromeDrive
	go click("username", "password", "user-data-dir=/home/pi/.config/chromium/profile") //use o diretorio do profile de seu browser
	clickWithCaptcha("username", "password", "user-data-dir=/home/pi/.config/chromium/profile")
}
