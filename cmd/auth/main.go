package auth

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// "версия" в контекст
const (
	version = "0.0.1"
)

func main() {

	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//
	defer cancel()

	// Закидываем в переменные среды "версию"
	os.Setenv(version, version)

	app, err := config.GetConfig(ctx, os.Args)
	if err != nil {
		panic(err)
	}
	app.Run(ctx)

}
