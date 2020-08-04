package telegrambot

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

//BotStart function start the bot and send message
func BotStart(token string, ChatID int64, tableName string, id int, db *sql.DB) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	responce := tgbotapi.NewMessage(ChatID, readDB(tableName, id, db))
	bot.Send(responce)

}

type content struct {
	data      string
	source_ip string
	time      string
}

func readDB(tableName string, id int, db *sql.DB) string {
	str := "SELECT data, source_ip, time from " + tableName + " WHERE " + strconv.Itoa(id) + "= id"
	rows, err := db.Query(str)
	defer rows.Close()
	contents := []content{}
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		p := content{}
		err := rows.Scan(&p.data, &p.source_ip, &p.time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contents = append(contents, p)
	}

	return contents[0].data + " " + contents[0].source_ip + " " + contents[0].time

}
