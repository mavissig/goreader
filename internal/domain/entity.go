package domain

import "encoding/xml"

type Ingredient struct {
	ItemName  string `xml:"itemname" json:"ingredient_name"`
	ItemCount string `xml:"itemcount" json:"ingredient_count"`
	ItemUnit  string `xml:"itemunit" json:"ingredient_unit"`
}

type Cake struct {
	Name        string       `xml:"name" json:"name"`
	Time        string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type MainStruct struct {
	XMLName xml.Name `xml:"recipes"`
	Recipes []Cake   `xml:"cake" json:"cake"`
}
