package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/yeqown/go-qrcode/writer/standard"
	tele "gopkg.in/telebot.v3"
	"io"
	"os"
	"qr_generator/handlers"
	"time"
)

type Application struct {
	bot *tele.Bot
}

func newApp() *Application {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	token := InitCfg()
	settings := tele.Settings{
		Token: token.BotToken,
		Poller: &tele.LongPoller{
			Timeout: 30 * time.Second,
		},
	}
	if len(settings.Token) <= 1 {
		log.Fatal().Msg("Token is not set")
	}
	bot, err := tele.NewBot(settings)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("Successfully init")
	return &Application{
		bot: bot,
	}
}
func (app *Application) addHandler(path string, handler tele.HandlerFunc) {
	app.bot.Handle(path, handler)
}
func (app *Application) addHandlerMiddleware(path string, handler tele.HandlerFunc, middleware tele.MiddlewareFunc) {
	app.bot.Handle(path, handler)
}

func main() {
	file, err := os.Open("data.json")
	app := newApp()
	app.addHandler("/qr", handlers.QrGen)
	app.addHandler("/start", handlers.Start)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Msg("Can't close file - ." + err.Error())
		}
	}(file)

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
	}
	app.bot.Start()

}
