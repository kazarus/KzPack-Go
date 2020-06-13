package KzPack4Go

import (
	"os"
	"strconv"
	"time"
)

type KzUnicode string

func (s KzUnicode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(s))), nil
}

func ToDate() int64 {
	result, _ := strconv.ParseInt(time.Now().Format("20060102"), 10, 64)
	return result
}

func DatePlus(aStartDay int) int64 {
	result, _ := strconv.ParseInt(time.Now().AddDate(0, 0, aStartDay).Format("20060102"), 10, 64)
	return result
}

func ToTime() float64 {
	result, _ := strconv.ParseFloat(time.Now().Format("20060102150405"), 64)
	return result
}

func TimePlus(aStartDay int) float64 {
	result, _ := strconv.ParseFloat(time.Now().AddDate(0, 0, aStartDay).Format("20060102150405"), 64)
	return result
}

func ToKJQJ() int64 {
	result, _ := strconv.ParseInt(time.Now().Format("200601"), 10, 64)
	return result
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
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

func FileExist(aFilePath string) (bool, error) {
	_, eror := os.Stat(aFilePath)
	if eror == nil {
		return true, nil
	}
	if os.IsNotExist(eror) {
		return false, nil
	}
	return true, eror
}
