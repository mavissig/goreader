package app

import (
	"goreader_cli/internal/cli"
	"goreader_cli/internal/domain"
	"goreader_cli/internal/infrastructure/comparator"
	"goreader_cli/internal/infrastructure/printer"
	"goreader_cli/internal/infrastructure/reader"
)

type App struct {
	cli *cli.Cli
	uc  *domain.UseCase
	rd  *reader.Reader
	pr  *printer.Printer
	cp  *comparator.Comparator
}

func (a *App) Run() {
	a.rd = reader.New()
	a.pr = printer.New()
	a.cp = comparator.New()

	a.uc = domain.New(a.rd, a.pr, a.cp)

	a.cli = cli.New(a.uc)
}

func New() *App {
	return &App{}
}
