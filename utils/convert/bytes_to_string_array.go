package convert

import "strconv"

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
