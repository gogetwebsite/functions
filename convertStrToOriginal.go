package functions

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	"net/mail"
	"net/url"
	"regexp"
)

// Converts string values to the specified target type
func StrToOriginal(value string, targetType string) (interface{}, error) {
	value = strings.TrimSpace(value);
	switch targetType {
		case "int":
			return strconv.ParseInt(value,10,64);
		case "hex":
			return strconv.ParseInt(value,16,64);			
		case "bool","checkbox","radio":
			return strconv.ParseBool(value)
		case "float" , "range":
			return strconv.ParseFloat(value,64)
		case "email":
			mail , err := mail.ParseAddress(value)
			if err != nil {
				return nil, err
			}
			return mail.Address , nil
		case "date":
			return time.Parse("2006-01-02", value)
		case "time":
			return time.Parse("15:04:05", value)
		case "datetime":
			return time.Parse("2006-01-02T15:04:05", value)
		case "datetime-local","datetime-utc","RFC3339":
			return time.Parse(time.RFC3339, value)
		case "month":
			return time.Parse("2006-01", value);

		case "[]int":
			// Convert string to slice of integers
			strValues := strings.Split(value, ",")
			intValues := make([]int, len(strValues))
			for i, str := range strValues {
				intVal, err := strconv.Atoi(strings.TrimSpace(str))
				if err != nil {
					return nil, err
				}
				intValues[i] = intVal
			}
			return intValues, nil
		case "[]float":
			// Convert string to slice of floats
			strValues := strings.Split(value, ",")
			floatValues := make([]float64, len(strValues))
			for i, str := range strValues {
				floatVal, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
				if err != nil {
					return nil, err
				}
				floatValues[i] = floatVal
			}
			return floatValues, nil
		case "[]string":
			// Convert string to slice of strings
			strValues := strings.Split(value, ",")
			for i := range strValues {
				strValues[i] = strings.TrimSpace(strValues[i])
			}
			return strValues, nil
		case "json":
			// Convert string to JSON object
			var jsonData map[string]interface{}
			if err := json.Unmarshal([]byte(value), &jsonData); err != nil {
				return nil, err
			}
			return jsonData, nil

		case "url":
			// Validate URL format
			_, err := url.ParseRequestURI(value)
			if err != nil {
				return nil, err
			}
			return value, nil
		case "color":
			// Validate color value
			value = strings.ReplaceAll(strings.ToLower(value) , " " , "") ;
			lenV := len(value) ;
			if lenV >= 4 && lenV <= 9 && value[0] == '#' {
				if _ , isInt := StrToOriginal(string(value[1:]) , "hex") ; isInt == nil{
					return value, nil
				}
			}else if strings.HasPrefix(value , "rgb") && lenV <= 21 {
				rgbRegex := `^rgba?\((25[0-5]|2[0-4]\d|1\d{2}|\d{1,2}),(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2}),(25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})(?:,(0|0?\.\d+|1(\.0+)?))?\)$`;
				if regexp.MustCompile(rgbRegex).MatchString(value) {
					return value, nil ;
				}
			}
			return nil, errors.New("invalid color value")
		case "checkbox-group", "radio-group":
			// Convert string to slice for checkbox or radio groups
			return strings.Split(value, ","), nil
		default:
			return nil, errors.New("unknown or unsupported input type");
	}
}
