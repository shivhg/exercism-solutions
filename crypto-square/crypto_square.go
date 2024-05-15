package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	reg := regexp.MustCompile(`[^\p{L}\p{N}]`)
	continousString := reg.ReplaceAllString(pt, "")
	continousString = strings.ToLower(continousString)
	columns := int(math.Ceil(math.Sqrt(float64(len(continousString)))))
	rows := int(math.Floor(math.Sqrt(float64(len(continousString)))))
	if columns*rows < len(continousString) {
		rows++
	}

	result := ""
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			if j*columns+i >= len(continousString) {
				result += " "
				continue
			}
			result += continousString[j*columns+i : j*columns+i+1]
		}
		result += " "
	}

	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}
