package functions

import ( 
	"os"
	"encoding/json"
)

func GetJsonMapFromFile(path string) (map[string]interface{}, error) {
	// Reads a JSON file and converts it into a map
	jsonFile, err := os.ReadFile(path);
	if err != nil {
		// Return nil and the error if file reading fails
		return nil, err;
	}
	tempMap , jerr := UnJSON(jsonFile) ;
	if err != nil {
		// Return nil and the error if file reading fails
		return nil, jerr;
	}
	// Return the populated map
	return tempMap, nil;
}

// Function to convert data to JSON format
func ToJSON(data interface{}) string {
	jsonData, err := json.Marshal(data);
	if err != nil {
		return `{"error": "internal error"}` ;
	}
	return string(jsonData)
}

func UnJSON(data []byte) (map[string]interface{} , error){
	tempMap := make(map[string]interface{});
	err := json.Unmarshal(data, &tempMap);
	if err != nil {
		return nil , err;
	}
	return tempMap , nil;
}
