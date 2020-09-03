package KzPack4Go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func SubStr(str string, start int, end int) string {

	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])

}

func TimeToStr(Value float64) string {
	//#20200618105004 -> 2020-06-18 10:50:04
	result := fmt.Sprintf("%.0f", Value)

	if len(result) != 14 {
		return "-1"
	}

	a := result[0:4]
	b := result[4:6]
	c := result[6:8]
	d := result[8:10]
	e := result[10:12]
	f := result[12:14]

	return fmt.Sprintf("%s-%s-%s %s:%s:%s", a, b, c, d, e, f)

}

func DateToStr(Value int64) string {
	//#20200618 -> 2020-06-18
	result := fmt.Sprintf("%d", Value)

	if len(result) != 8 {
		return "-1"
	}

	a := result[0:4]
	b := result[4:6]
	c := result[6:8]

	return fmt.Sprintf("%s-%s-%s", a, b, c)

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

func FormatFloat(Value float64) float64 {

	result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", Value), 64)
	return result

}

func IfThenI(Value bool, Result0, Result1 int64) int64 {

	Result := Result1
	if Value {
		Result = Result0
	}

	return Result
}

func IfThenS(Value bool, Result0, Result1 string) string {

	Result := Result1
	if Value {
		Result = Result0
	}

	return Result
}

func IfThenF(Value bool, Result0, Result1 float64) float64 {

	Result := Result1
	if Value {
		Result = Result0
	}

	return Result
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
	if eror != nil {
		return false, eror
	}
	if os.IsNotExist(eror) {
		return false, nil
	}

	return true, nil

}

func IsFile(aFilePath string) bool {

	fi, e := os.Stat(aFilePath)
	if e != nil {
		return false
	}

	return !fi.IsDir()

}

func ToFile(aFilePath string, v interface{}) error {

	var eror error
	var source []byte

	/*
		if !IsFile(aFilePath) {
			return errors.New("Error: The FilePath is not a FileName;")
		}
	*/

	source, eror = json.Marshal(v)
	if eror != nil {
		return eror
	}

	eror = ioutil.WriteFile(aFilePath, source, 0666)
	if eror != nil {
		return eror
	}

	return nil

}

func InFile(aFilePath string, v interface{}) error {

	var eror error
	var didExist bool

	/*
		if !IsFile(aFilePath) {
			return errors.New("Error: The FilePath is not a FileName;")
		}
	*/

	didExist, _ = FileExist(aFilePath)
	if !didExist {
		eror = ToFile(aFilePath, v)
		if eror != nil {
			return eror
		}
		return nil
	}

	source, eror := ioutil.ReadFile(aFilePath)
	if eror != nil {
		return eror
	}

	eror = json.Unmarshal(source, &v)
	if eror != nil {
		return eror
	}

	return nil

}

func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}