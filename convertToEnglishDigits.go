package functions

import (
	"strings"
	"unicode"
)
//Converts any digits in the input string to English digits
func ConvertToEnglishDigits(input string) string {
	//result to final string
	var result strings.Builder
	for _, char := range input {
		//find standard unicode digits chars
		if unicode.IsDigit(char) {
			//skip if char is english digit
			if '0' <= char && char <= '9' {
				result.WriteRune(char);
				continue ;
			}

			intChar := int(char) ; //get the char as Acci int code
			var zeroOfLocalDigit rune ; //define the base of the local digit
			englishDigit := '0' ;

			for i:= 1 ; i<11 ; i++ {
				if !unicode.IsDigit(rune(intChar - i)) {
					//find the zero acci of local digit char
					zeroOfLocalDigit = rune(intChar - i + 1) ;

					//Find the English digit equivalent of the input numeric character
					//based on the distance from the char '0'
					offset := char - zeroOfLocalDigit ;
					englishDigit += offset ;
					break;
				}
			}
			result.WriteRune(englishDigit);
		} else {
			//skip if char is not a digit
			result.WriteRune(char)
		}
	}
	// Return the resulting string with English digits
	return result.String()
}
