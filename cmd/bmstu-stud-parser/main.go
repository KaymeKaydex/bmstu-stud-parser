package main

import (
	"context"

	jww "github.com/spf13/jwalterweatherman"

	"github.com/KaymeKaydex/bmstu-stud-parser/internal/pkg/app"
)

func main() {
	ctx := context.Background()

	jww.SetStdoutThreshold(jww.LevelInfo)

	app, err := app.NewApp()
	if err != nil {
		jww.INFO.Fatalln("Cant create app")
	}
	app.Run(ctx)

}
