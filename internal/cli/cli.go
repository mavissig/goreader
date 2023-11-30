package cli

import (
	"flag"
	"fmt"
)

type UseCase interface {
	ConvertAndPrintFile(string)
	CompareFiles(string, string)
}

type Cli struct {
	usecase UseCase
}

func (c *Cli) Do() {
	file := flag.String("f", "", "")

	oldFile := flag.String("old", "", "")
	newFile := flag.String("new", "", "")

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "f" {
			c.usecase.ConvertAndPrintFile(*file)
		} else if f.Name == "old" || f.Name == "new" {
			fmt.Printf("flag used: %v\n file old: %v\nfile new: %v", f.Name, *oldFile, *newFile)
		}
	})

}

func New(uc UseCase) *Cli {
	return &Cli{
		usecase: uc,
	}
}
