package comparator

import (
	"fmt"
	"goreader_cli/internal/domain"
)

type Comparator struct {
}

func (c Comparator) Eq(db1 domain.MainStruct, db2 domain.MainStruct) {
	for indOld, cakeOld := range db1.Recipes {
		for indNew, cakeNew := range db2.Recipes {
			if cakeOld.Name == cakeNew.Name && indOld == indNew {
				if cakeNew.Time != cakeOld.Time {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s min\" instead of \"%s min\"\n", cakeOld.Name, cakeOld.Time, cakeNew.Time)
				}
				for indIngredientOld, ingredientOld := range cakeOld.Ingredients {
					for indIngredientNew, ingredientNew := range cakeNew.Ingredients {
						if ingredientOld.ItemName == ingredientNew.ItemName && indIngredientOld == indIngredientNew {
							if ingredientNew.ItemUnit == "" && ingredientOld.ItemUnit != ingredientNew.ItemUnit {
								fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", ingredientOld.ItemUnit, ingredientOld.ItemName, cakeOld.Name)
							} else if ingredientOld.ItemUnit != ingredientNew.ItemUnit {
								fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", ingredientOld.ItemName, cakeOld.Name, ingredientOld.ItemUnit, ingredientNew.ItemUnit)
							} else {
								if ingredientOld.ItemCount != ingredientNew.ItemCount {
									fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", ingredientOld.ItemName, cakeOld.Name, ingredientOld.ItemCount, ingredientNew.ItemCount)
								}
							}
						} else if ingredientNew.ItemName != ingredientOld.ItemName && indIngredientOld == indIngredientNew {
							fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", ingredientNew.ItemName, cakeOld.Name)
							fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", ingredientOld.ItemName, cakeOld.Name)
							break
						}
					}
				}
			} else if cakeOld.Name != cakeNew.Name && indOld == indNew {
				fmt.Printf("ADDED cake %s\n", cakeNew.Name)
				fmt.Printf("REMOVED cake %s\n", cakeOld.Name)
				break
			}
		}
	}
}

func New() *Comparator {
	return &Comparator{}
}
