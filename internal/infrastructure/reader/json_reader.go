package reader

import (
	"encoding/json"
	"fmt"
	"goreader_cli/internal/domain"
	"os"
)

type Json struct {
	Reader
	dataBase domain.MainStruct
}

func (r *Json) Mock() {
	fmt.Println("json mock")
}

func (r *Json) Read(path string) domain.MainStruct {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error read file: ", err)
		return domain.MainStruct{}
	}

	err = json.Unmarshal(data, &r.dataBase)
	if err != nil {
		fmt.Println("Error outJson unmarshal: ", err)
		return domain.MainStruct{}
	}

	return r.dataBase
}
