package KzPack4Go

import "encoding/json"
import "log"
import "io"
import "crypto/md5"
import "crypto/rand"
import "encoding/base64"
import "encoding/hex"
import "strings"

type TEROR struct {
	ERORCODE string
	ERORMEMO string
	LISTDATA interface{}
}

var eror error

func (self TEROR) MarshalJSON() ([]byte, error) {
	var result = make(map[string]interface{})
	result["ERORCODE"] = KzUnicode(self.ERORCODE)
	result["ERORMEMO"] = KzUnicode(self.ERORMEMO)
	result["LISTDATA"] = self.LISTDATA

	return json.Marshal(result)
}
func ToEROR(aERORMEMO string) TEROR {
	var result TEROR
	result.ERORCODE = "EROR"
	result.ERORMEMO = aERORMEMO
	return result
}

func ToMEMO(aERORMEMO string) TEROR {
	var result TEROR
	result.ERORCODE = "EROR"
	result.ERORMEMO = aERORMEMO
	return result
}

func ToTRUE(aListData interface{}) TEROR {
	var result TEROR
	result.ERORCODE = "TRUE"
	result.ERORMEMO = "TRUE"
	result.LISTDATA = aListData
	return result
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s)) //使用zhifeiya名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}

func ToGUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(GetMd5String(base64.URLEncoding.EncodeToString(b)))
}

func ChEror(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
