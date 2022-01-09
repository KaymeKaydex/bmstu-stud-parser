package app

import (
	"context"

	"github.com/geziyor/geziyor"

	"github.com/KaymeKaydex/bmstu-stud-parser/internal/app/parser"
)

type App struct {
	p parser.Parser
	g *geziyor.Geziyor
}

func NewApp() (*App, error) {
	parser := parser.Parser{}
	return &App{
		p: parser,
		g: geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: parser.GetURLs("http://studsovet.bmstu.ru/old-version/page/"),
			ParseFunc: parser.Parse,
			Exporters: nil,
		}),
	}, nil
}
func (a *App) Run(ctx context.Context) {
	a.g.Start()

}
