// cmd/yourbot/main.go
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	reminderBot "github.com/smith3v/tg-word-reminder/pkg/bot"
	"github.com/smith3v/tg-word-reminder/pkg/config"
	"github.com/smith3v/tg-word-reminder/pkg/db"
)

var logger = slog.Default()

func main() {
	config.LoadConfig("config.json")
	if err := db.InitDB(config.AppConfig.Database); err != nil {
		logger.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(reminderBot.DefaultHandler),
	}
	b, err := bot.New(config.AppConfig.Telegram.Token, opts...)
	if err != nil {
		logger.Error("failed to create bot", "error", err)
		os.Exit(1)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, reminderBot.HandleStart)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/clear", bot.MatchTypeExact, reminderBot.HandleClear)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/setnum", bot.MatchTypePrefix, reminderBot.HandleSetNumOfPairs)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/setfreq", bot.MatchTypePrefix, reminderBot.HandleSetFrequency)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/getpair", bot.MatchTypeExact, reminderBot.HandleGetPair)

	go reminderBot.StartPeriodicMessages(ctx, b)

	logger.Info("Starting bot...")
	b.Start(ctx)
}
