package KzPack4Go


import "strconv"
import "time"

type KzUnicode string

func (s KzUnicode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(s))), nil
}

func ToDate() int64 {
	result, _ := strconv.ParseInt(time.Now().Format("20060102"), 10, 64)
	return result
}
func ToTime() float64 {
	result, _ := strconv.ParseFloat(time.Now().Format("20060102150405"), 64)
	return result
}

func ToKJQJ() int64 {
	result, _ := strconv.ParseInt(time.Now().Format("200601"), 10, 64)
	return result
}

func StrToInt(Value string) int64 {
	result, _ := strconv.ParseInt(Value, 10, 64)
	return result
}

func IntToStr(Value int64) string {
	result := strconv.Itoa(int(Value))
	return result
}

func Contain(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

