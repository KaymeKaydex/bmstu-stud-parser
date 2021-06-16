package main

import (
	"context"
	"github.com/KaymeKaydex/bmstu_stud_parser/internal/pkg/app"
	jww "github.com/spf13/jwalterweatherman"
)



func main() {
	ctx := context.Background()

	jww.SetStdoutThreshold(jww.LevelInfo)

	app,err := app.NewApp()
	if err!= nil {
		jww.INFO.Fatalln("Cant create app")
	}
	app.Run(ctx)

}


