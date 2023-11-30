package reader

import (
	"encoding/xml"
	"fmt"
	"goreader_cli/internal/domain"
	"os"
)

type Xml struct {
	Reader
	dataBase domain.MainStruct
}

func (r *Xml) Mock() {
	fmt.Println("xml mock")
}

func (r *Xml) Read(path string) domain.MainStruct {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error read file: ", err)
		return domain.MainStruct{}
	}

	err = xml.Unmarshal(data, &r.dataBase)
	if err != nil {
		fmt.Println("Error outJson unmarshal: ", err)
		return domain.MainStruct{}
	}

	return r.dataBase
}
