package functions

import (
	"strings" ;
	"regexp" ;
) 
var (
	// Define a regex pattern for phone number validation
	phoneRegex = regexp.MustCompile(`^\+?(\d{1,3})?[\s\-\.()]*(\d[\s\-\.()]*){10,30}$`);
	irPhoneRegex = regexp.MustCompile(`^(98|0)?9\d{9}$`);
);

func Is_IR_Phone(phone string) bool{
	return irPhoneRegex.MatchString(phone) ;
}

// IsEmail checks if the input string matches a standard email format
func IsEmail(email string) (string , bool) {
	if len(email) > 256 {
		return "" , false ;
	}
	email = ConvertToEnglishDigits(email);
	if rawMail,mailOk:= StrToOriginal(email , "email"); mailOk == nil {
		mail:= rawMail.(string) ;	

		//check if mail domain would have suffix
		parts := strings.Split(mail, "@");
		domain := parts[1] ;
		domainParts := strings.Split(domain, ".");
		if len(domainParts) >= 2 {
			return mail, true ;
		}
	}
	return "" , false ;
}

// isPhoneNumber checks if the input string matches a standard phone number format.
func IsPhone(phone string)  (string , bool) {
	phone = strings.ReplaceAll(phone, " ", "");
	if len(phone) > 30 {
		return "" , false ;
	}
	phone = ConvertToEnglishDigits(phone) ;
	if phoneRegex.MatchString(phone) {
		phone = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(phone, "(",""), ")", ""), "-",""), "+","") ;
		if len(phone) <= 15 {
			return phone , true ;
		}
	}
	return "" , false ;
}
/*
phones := []string{
	"+1 234 567 8900", // valid
	"+1(123)-456-7890",    // valid
	"(123) 456-7890",  // valid
	"1234567890",      // valid
	"+1234567890",     // valid
	"123 456 7890 123",   // invalid (too long)
	"123",             // invalid (too short)
	"123-abc-7890",    // invalid (contains letters)
	"   +1   234 567 8900", // valid after spaces removed
}
*/

func PhoneOrMail(data string)(string,string){
	phone := "";
	email := "";
	if isPhone , isok := IsPhone(data) ; isok {
		phone = isPhone ;
	}else if isEmail , isok := IsEmail(data) ; isok {
		email = isEmail ;
	}
	return phone , email ;
}
