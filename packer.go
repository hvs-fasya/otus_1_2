package otus_1_2

import (
	"strconv"
	"strings"
	"unicode"
)

//Packer common packer interface
type Packer interface {
	Unpack(string) string
}

//PackerSvc packer service structure
type PackerSvc struct {
	EscapeSymbol string
}

//NewPackerSvc init packer service
func NewPackerSvc(escape string) *PackerSvc {
	return &PackerSvc{
		EscapeSymbol: escape,
	}
}

//Unpack unpack given string
func (svc *PackerSvc) Unpack(str string) string {
	var unpacked string
	var runes = []rune(str)
	var digitalStr = ""
	var l = len(runes)
	if l == 0 || unicode.IsDigit(runes[0]) {
		return unpacked
	}
	var prev = []rune(``)
	var escapeNext bool
	for i := 0; i <= l-1; i++ {
		if !unicode.IsDigit(runes[i]) {
			unpacked = unpacked + prevRepeated(prev, digitalStr)
			digitalStr = ""
			if string(runes[i]) == svc.EscapeSymbol {
				escapeNext = !escapeNext
				if string(prev) == svc.EscapeSymbol {
					unpacked = string([]rune(unpacked)[:len(unpacked)-1])
				}
			}
			prev = []rune{runes[i]}
			continue
		}
		if escapeNext {
			escapeNext = false
			prev = []rune{runes[i]}
			continue
		}
		digitalStr = digitalStr + string(runes[i])

	}
	//handle last rune
	if !unicode.IsDigit(runes[l-1]) {
		unpacked = unpacked + string(runes[l-1])
	} else {
		unpacked = unpacked + prevRepeated(prev, digitalStr)
	}
	return unpacked
}

func prevRepeated(prev []rune, digitalStr string) string {
	count, _ := strconv.Atoi(digitalStr)
	if count == 0 {
		count = 1
	}
	return strings.Repeat(string(prev), count)
}
