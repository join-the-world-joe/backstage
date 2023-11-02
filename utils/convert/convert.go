package convert

import (
	"encoding/json"
	"strconv"
)

func Bytes2StringArray(slice []byte) string {
	l := len(slice)
	str := "[ "
	for k, v := range slice {
		val := int(v)
		if k < l-1 {
			str += strconv.Itoa(val) + " ,"
		} else {
			str += strconv.Itoa(val)
		}
	}
	str += " ]"
	return str
}

func BytesStringToString(bytesString string) string {
	// 1. to Bytes
	// 2. to string
	bytes := []byte{}
	err := json.Unmarshal([]byte(bytesString), &bytes)
	if err != nil {
		return ""
	}
	return string(bytes)
}
