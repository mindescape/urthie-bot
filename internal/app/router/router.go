package router

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Router struct {
	bot *tgbotapi.BotAPI
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot: bot,
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {

	switch {
	case update.Message != nil:
		log.Printf("Message from: %s", update.Message.From.UserName)
		r.HandleMessage(update.Message)
	}
}

func (r *Router) HandleMessage(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)

	r.bot.Send(reply)
}
