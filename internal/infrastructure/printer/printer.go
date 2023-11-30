package printer

import (
	"encoding/json"
	"encoding/xml"
	"goreader_cli/internal/domain"
	"log"
)

type Printer struct {
}

func (p *Printer) Print(suffix string, mainStruct domain.MainStruct) {
	switch suffix {
	case ".xml":
		p.PrintAsXml(mainStruct)
	case ".json":
		p.PrintAsJson(mainStruct)
	}
}

func (p *Printer) PrintAsJson(mainStruct domain.MainStruct) {
	jsonByte, err := json.MarshalIndent(mainStruct, "", "   ")
	if err != nil {
		log.Println(err)
	}

	println(string(jsonByte))
}

func (p *Printer) PrintAsXml(mainStruct domain.MainStruct) {
	xmlByte, err := xml.MarshalIndent(mainStruct, "", "   ")
	if err != nil {
		log.Println(err)
	}
	println(string(xmlByte))

}
func New() *Printer {
	return &Printer{}
}
