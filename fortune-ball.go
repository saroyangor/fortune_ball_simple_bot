package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"strings"
)

const TOKEN = "6627300550:AAEGreyTT9QpcwuDNRHTe3NugxpEXibfd9M"

var bot *tgbotapi.BotAPI
var chatId int64

var fortuneTellerNames = [3]string{"кицунэ", "моя госпожа", "солнышко"}

var answers = []string{
	"Да",
	"Нет",
	"Однозначно",
	"Ни в коем случае",
	"Иногда нужно просто сказать 'Да' и идти вперед.",
	"Иногда нужно признать, что момент благоприятен, и действовать.",
	"Смотрите на огонь, который горит внутри, и следуй за ним.",
	"Иногда лучшее решение - подождать, как ветка, пока не упадет плод.",
	"Иногда нужно посмотреть на мир глазами детей, чтобы найти ответ.",
	"Как дерево прогибается под ветром, так и решение может потребовать гибкости.",
	"Иногда нужно позволить событиям развернуться сами по себе.",
	"Иногда ответ находится в молчании.",
	"Слушай шепот природы, и она подскажет, следует ли вмешиваться.",
	"Перед тем как действовать, прислушайся к своему сердцу.",
	"Подумай, какие следы ты оставишь на своем пути.",
	"Иногда мудрость заключается в том, чтобы ничего не делать.",
	"Есть время для быстрых решений и время для тщательных размышлений.",
	"Искусство жизни заключается в том, чтобы знать, когда остановиться.",
	"Лучшие действия рождаются из спокойного размышления.",
	"Соблюдай баланс между тем, что хочешь, и тем, что нужно.",
	"Созреешь ли, как плод, зависит от того, когда сорвать.",
	"Иногда тишина гораздо могущественнее слов.",
	"Пусть твои действия будут как следы водопада — красивыми и несомненными.",
	"Решения, принятые с любовью, всегда правильные.",
	"Не спеши, дай своим мыслям пройти через сердце, прежде чем действовать.",
	"Внутренний мир — твой лучший советчик.",
	"Сначала освети свой собственный путь, а потом помоги другим.",
	"Смотрите на светлое будущее, но не забывайте учиться из прошлого.",
	"Доверься интуиции, она знает, какой путь для тебя правильный.",
	"Сердце может видеть то, что глаза не могут увидеть.",
	"Иногда трудные решения приносят самые красивые плоды.",
	"Слушай мудрость старших — они прошли многое на своем пути.",
	"Следуй за своей страстью, она покажет твой истинный путь.",
	"Помни, что твои действия могут влиять на мир вокруг тебя.",
	"Сконцентрируйся на моменте, чтобы понять, что следует делать дальше.",
	"Любовь к жизни и всему живому будет твоим надежным проводником.",
	"Позволь своей душе мурлыкать в гармонии с миром, прежде чем решить, стоит ли что-то делать.",
	"Сомнение — это знак, что вам нужно подумать глубже.",
	"Честность всегда является лучшей стратегией.",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}

	return false
}

func getFortuneTellerAnswer() string {
	index := rand.Intn(len(answers))

	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getFortuneTellerAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID

			sendMessage("Задай свой вопрос, обращяясь ко мне в вежливой форме. " +
				"Ответом на вопрос должен быть \"Да\" либо \"Нет\". " +
				"Например, \"Кицунэ, я готов сменить работу?\" " +
				"или \"Моя госпожа, я действительно хочу пойти на эту вечеринку?\"",
			)
		}

		if isMessageForFortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
