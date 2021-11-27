package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
)

func main() {
	token := ""
	if token == "" {
		if len(os.Args) > 1 {
			token = os.Args[1]
		} else {
			println("No Token Provided!")
			os.Exit(0)
		}
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message != nil && update.Message.IsCommand() { // ignore any non-Message Commands Updates

			args := strings.Trim(update.Message.CommandArguments(), " ")
			command := update.Message.Command()

			switch command {
			case "code":
				if update.Message.CommandArguments() == "" {
					SendMessage(update, bot, "No Args!")
					continue
				}
				lineEntity, err := GetICDLinEntityByCode(args)
				entity, err2 := GetICDFoundationByID(lineEntity.ID())
				lineEntity, err = GetICDLinEntityByID(entity.ID())

				if err2 != nil {
					PrintErr(err)
					break
				}
				if err != nil {
					PrintErr(err)
					break
				}
				SendResult(update, entity, lineEntity, bot)
				break

			case "code10":
				if update.Message.CommandArguments() == "" {
					SendMessage(update, bot, "No Args!")
					continue
				}
				entity10, err := GetICD10ByCode(args)
				if err != nil {
					PrintErr(err)
					break
				}
				entity10.Code = args
				SendResult10(update, entity10, bot)
				break

			case "search":
				if update.Message.CommandArguments() == "" {
					SendMessage(update, bot, "No Args!")
					continue
				}
				lineEntityList, err := SearchICDLinEntity(args)
				if err != nil {
					PrintErr(err)
					break
				}
				if len(lineEntityList) == 0 {
					SendMessage(update, bot, "Nothing Found!")
				} else {
					lineEntity := lineEntityList[0]
					entity, err2 := GetICDFoundationByID(lineEntity.ID())
					lineEntity, err = GetICDLinEntityByID(entity.ID())
					if err2 != nil {
						PrintErr(err)
						break
					}
					SendResult(update, entity, lineEntity, bot)
				}
				break
			case "help":
				text := "Commands:\n\n " +
					"/search {term}\n " +
					"/code {ICD-11 code}\n " +
					"/code10 {ICD-10 code}\n "
				SendMessage(update, bot, text)
			case "start":
				text := "Hey there! - " +
					"This bot allows you to search the ICD10 and ICD11 diagnostic tools through Telegram.\n\n" +
					"Hit /help to find out about my commands.\n\n" +
					"The source code can be found <a href=\"https://w3y.cc/icdbottgrefer\">here</a>."
				SendMessage(update, bot, text)

			}
		}
	}
}

func SendMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "html"
	_, err := bot.Send(msg)
	if err != nil {
		PrintErr(err)
	}
}

func SendResult(update tgbotapi.Update, entity Entity, lineEntity LineEntity, bot *tgbotapi.BotAPI) {
	result := "<b>Title: </b>" + entity.Title.Value
	if lineEntity.BrowserUrl == "" {
		result += "<b> \nCode: </b>" + lineEntity.GetCode()
	} else {
		result += "<b> \nCode: </b><a href='" + lineEntity.BrowserUrl + "'>" + lineEntity.GetCode() + "</a>"
	}
	if entity.Definition.Value != "" {
		result += "<b>\n\nDescription: </b>	" + entity.Definition.Value
	}
	children := entity.GetChildren()
	parents := entity.GetParent()
	TextChildren := TextParentChildFromEntity(children)
	TextParent := TextParentChildFromEntity(parents)
	if TextChildren != "" {
		result += "\n\n<b>children:</b>\n" + TextChildren
	}
	if TextParent != "" {
		result += "\n\n<b>Parents:</b>\n" + TextParent
	}
	SendMessage(update, bot, result)
}

func TextParentChildFromEntity(ChildorParent []Entity) string {
	textResult := ""
	for _, element := range ChildorParent {
		elementLine, err := GetICDLinEntityByID(element.ID())
		if err != nil {
			PrintErr(err)
			continue
		}
		textResult += "<a href='" + elementLine.BrowserUrl + "'>" + elementLine.GetCode() + "</a> - " + element.Title.Value + "\n"
	}
	return textResult
}
func SendResult10(update tgbotapi.Update, entity10 Entity10, bot *tgbotapi.BotAPI) {

	result := "<b> \nCode: </b><a href='https://icd.who.int/browse10/2019/en#/" + entity10.Code + "'>" + entity10.Code + "</a><b> \nTitle: </b>" + entity10.Name + "<b>\n\nDescription: </b>	" + entity10.Description
	SendMessage(update, bot, result)
}
