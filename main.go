package main

import (
	"fmt"
	"github.com/gagliardetto/solana-go"
	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := telegrambot.NewBotAPI("7246445809:AAEKnIDGTTvumsHBZoroWtFZwGfL6mVpdvE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	//log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegrambot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Text == "/create" || update.Message.Text == "/Create" {
			account := solana.NewWallet()
			//fmt.Println("account private key:", account.PrivateKey)
			//fmt.Println("account public key:", account.PublicKey())

			publicKey := account.PublicKey()
			privateKey := account.PrivateKey

			msg := telegrambot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Wallet created 🎉 \n\nPublic Key: \n%s\n\nPrivate Key: \n%s\n\nCreate another wallet? Type /Create", publicKey, privateKey))

			log.Printf("Authorized on account %s", update.Message.Chat.ID)
			log.Printf("Authorized on account %s", update.Message.Text)
			log.Printf("Authorized on account %s", update.Message.MessageID)

			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		}

	}
}
