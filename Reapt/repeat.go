
package reapt

func Repeat(character string, repeat int) string {
	result := ""
	for count := 0; count < repeat; count++ {
		result += character
	}
	return result
}