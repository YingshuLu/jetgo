package vm

import "github.com/fatih/structs"

type iFact struct {
	Fact interface{}
}

// Fact return standard fact
func Fact(data interface{}) interface{} {
	fact := &iFact{
		Fact: data,
	}
	return structs.Map(fact)
}
