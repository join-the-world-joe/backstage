package phonedata

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	cmcc               byte = iota + 0x01 //中国移动
	cucc                                  //中国联通
	ctcc                                  //中国电信
	ctcc_v                                //电信虚拟运营商
	cucc_v                                //联通虚拟运营商
	cmcc_v                                //移动虚拟运营商
	int_len            = 4
	char_len           = 1
	head_length        = 8
	phone_index_length = 9
)

type PhoneRecord struct {
	PhoneNum string
	Province string
	City     string
	ZipCode  string
	AreaZone string
	CardType string
}

var (
	content     []byte
	CardTypemap = map[byte]string{
		cmcc:   "中国移动",
		cucc:   "中国联通",
		ctcc:   "中国电信",
		ctcc_v: "中国电信虚拟运营商",
		cucc_v: "中国联通虚拟运营商",
		cmcc_v: "中国移动虚拟运营商",
	}
	total_len, firstoffset int32
)

type PhoneData struct{}

func NewPhoneData(file string) (*PhoneData, error) {
	var err error
	content, err = ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	total_len = int32(len(content))
	firstoffset = get4(content[int_len : int_len*2])

	return new(PhoneData), nil
}

func (p *PhoneData) Find(phone string) (*PhoneRecord, error) {
	return find(phone)
}

func Debug() {
	fmt.Println(version())
	fmt.Println(totalRecord())
	fmt.Println(firstRecordOffset())
}

func (pr PhoneRecord) String() string {
	return fmt.Sprintf("PhoneNum: %s\nAreaZone: %s\nCardType: %s\nCity: %s\nZipCode: %s\nProvince: %s\n", pr.PhoneNum, pr.AreaZone, pr.CardType, pr.City, pr.ZipCode, pr.Province)
}

func get4(b []byte) int32 {
	if len(b) < 4 {
		return 0
	}
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}

func getN(s string) (uint32, error) {
	var n, cutoff, maxVal uint32
	i := 0
	base := 10
	cutoff = (1<<32-1)/10 + 1
	maxVal = 1<<uint(32) - 1
	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		default:
			return 0, errors.New("invalid syntax")
		}
		if v >= byte(base) {
			return 0, errors.New("invalid syntax")
		}

		if n >= cutoff {
			// n*base overflows
			n = (1<<32 - 1)
			return n, errors.New("value out of range")
		}
		n *= uint32(base)

		n1 := n + uint32(v)
		if n1 < n || n1 > maxVal {
			// n+v overflows
			n = (1<<32 - 1)
			return n, errors.New("value out of range")
		}
		n = n1
	}
	return n, nil
}

func version() string {
	return string(content[0:int_len])
}

func totalRecord() int32 {
	return (int32(len(content)) - firstRecordOffset()) / phone_index_length
}

func firstRecordOffset() int32 {
	return get4(content[int_len : int_len*2])
}

// 二分法查询phone数据
func find(phone_num string) (pr *PhoneRecord, err error) {
	if len(phone_num) < 7 || len(phone_num) > 11 {
		return nil, errors.New("illegal phone length")
	}

	var left int32
	phone_seven_int, err := getN(phone_num[0:7])
	if err != nil {
		return nil, errors.New("illegal phone number")
	}
	phone_seven_int32 := int32(phone_seven_int)
	right := (total_len - firstoffset) / phone_index_length
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		offset := firstoffset + mid*phone_index_length
		if offset >= total_len {
			break
		}
		cur_phone := get4(content[offset : offset+int_len])
		record_offset := get4(content[offset+int_len : offset+int_len*2])
		card_type := content[offset+int_len*2 : offset+int_len*2+char_len][0]
		switch {
		case cur_phone > phone_seven_int32:
			right = mid - 1
		case cur_phone < phone_seven_int32:
			left = mid + 1
		default:
			cbyte := content[record_offset:]
			end_offset := int32(bytes.Index(cbyte, []byte("\000")))
			data := bytes.Split(cbyte[:end_offset], []byte("|"))
			card_str, ok := CardTypemap[card_type]
			if !ok {
				card_str = "未知电信运营商"
			}
			pr = &PhoneRecord{
				PhoneNum: phone_num,
				Province: string(data[0]),
				City:     string(data[1]),
				ZipCode:  string(data[2]),
				AreaZone: string(data[3]),
				CardType: card_str,
			}
			return
		}
	}
	return nil, errors.New("phone's data not found")
}
