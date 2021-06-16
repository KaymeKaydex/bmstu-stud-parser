package app

import (
	"context"
	"github.com/KaymeKaydex/bmstu_stud_parser/internal/app/parser"
	"github.com/geziyor/geziyor"
)

type App struct {
	p parser.Parser
	g *geziyor.Geziyor
}

func NewApp() (*App,error){
	parser := parser.Parser{}
	return &App{
		p: parser,
		g: geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: parser.GetURLs("http://studsovet.bmstu.ru/page/"),
			ParseFunc: parser.Parse,
			Exporters: nil,
		}),
	}, nil
}
func (a *App) Run(ctx context.Context)  {
	a.g.Start()
}