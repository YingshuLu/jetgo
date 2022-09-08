package vm

import (
	"go/token"
	"strconv"
	"strings"
)

const (
	trueToken  = "true"
	falseToken = "false"
	nullToken  = "nil"
	arrayToken = "array"
	forkToken  = "fork"
	tableToken = "table"
)

// logicToken check if op is logic token
func logicToken(op token.Token) bool {
	if op == token.LAND || op == token.LOR {
		return true
	}
	return false
}

// varToken check if variant is internal token
func varToken(name string) bool {
	ok := true
	switch name {
	case trueToken:
	case falseToken:
	case nullToken:
	case forkToken:
	case tableToken:
	default:
		ok = false
	}
	return ok
}

// varShouldLocal check if variant is local
func varShouldLocal(name string) bool {
	return name[0] < 'A' || name[0] > 'Z'
}

// varRenamed check if name renamed
func varRenamed(name string) bool {
	return name[0] == 'x' && len(name) == 4 && name[2] == '_' && name[3] == '_'
}

// varRenaming rename variant with index
func varRenaming(index byte) string {
	return renaming([]byte("x __"), index)
}

func constRenamed(name string) bool {
	return name[0] == 'c' && len(name) == 4 && name[2] == '_' && name[3] == '_'
}

func constRenaming(index byte) string {
	return renaming([]byte("c __"), index)
}

func renaming(template []byte, index byte) string {
	template[1] = index
	return string(template)
}

func indexing(name string) int {
	return int(name[1])
}

func stringIntNumber(s string) (interface{}, error) {
	return strconv.ParseInt(s, 0, 64)
}

func stringFloatNumber(s string) (interface{}, error) {
	return strconv.ParseFloat(s, 64)
}

func stringTrim(s string) (interface{}, error) {
	return strings.Trim(s, "\""), nil
}
