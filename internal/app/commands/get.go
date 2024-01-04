package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(arg)
	if err != nil {
		log.Printf("fail to get product idx %d: %v", arg, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)

	c.bot.Send(msg)
}
