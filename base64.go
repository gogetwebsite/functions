package functions

import ( 
	"encoding/base64"
)

// Function to convert data to JSON format
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data);
}

func DecodeBase64(data string) ([]byte,error) {
	return base64.StdEncoding.DecodeString(data);
}
