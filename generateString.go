package functions

import ( 
	"strings"
	"crypto/rand"
	"math/big"

	"github.com/google/uuid"
)
var (
	charset string = `0123456789abcdefghijklmnopqrstuvwxyz!@#$%^&*()!@#$%^&*()ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()0123456789!@#$%^&*()`;
	otpSet string = `0123456789` ;

	lenCharset int64 ;
	lenOtpSet int64 ;
	OTP_Length int = 6 ;
)

func init(){
	lenCharset = int64(len(charset));
	lenOtpSet = int64(len(otpSet));
}

func GenerateRandomKey(length int) (string, error) {
	// Generates a random key of specified length
	var result strings.Builder
	result.Grow(length);
	for i := 0; i < length; i++ {
		// Generate a random index to select a character from the charset
		num, err := rand.Int(rand.Reader, big.NewInt(lenCharset))
		if err != nil {
			return "", err
		}
		// Append the randomly selected character to the result
		result.WriteByte(charset[num.Int64()])
	}
	// Return the generated random key
	return result.String(), nil
}

func GenerateOTP() (string, error) {
	// Generates a random key of specified length
	var result strings.Builder
	result.Grow(OTP_Length);
	for i := 0; i < OTP_Length ; i++ {
		// Generate a random index to select a character from the charset
		num, err := rand.Int(rand.Reader, big.NewInt(lenOtpSet))
		if err != nil {
			return "", err
		}
		// Append the randomly selected character to the result
		result.WriteByte(otpSet[num.Int64()])
	}  
	// Return the generated random key
	return result.String(), nil
}

func GenerateUUID()string{
	return uuid.NewString();
}
