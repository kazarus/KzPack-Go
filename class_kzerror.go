package KzPack4Go

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/satori/go.uuid"
)

type TEROR struct {
	ONSTATUS bool
	ERORCODE string
	ERORMEMO string
	LISTDATA interface{}
}

var eror error

func (self TEROR) MarshalJSON() ([]byte, error) {

	var result = make(map[string]interface{})

	result["ONSTATUS"] = self.ONSTATUS
	result["ERORCODE"] = KzUnicode(self.ERORCODE)
	result["ERORMEMO"] = KzUnicode(self.ERORMEMO)
	result["LISTDATA"] = self.LISTDATA

	return json.Marshal(result)
}

func ToEROR(aERORMEMO string) TEROR {
	var result TEROR
	result.ONSTATUS = false
	result.ERORCODE = "EROR"
	result.ERORMEMO = aERORMEMO
	return result
}

func ToMEMO(aERORMEMO string) TEROR {
	var result TEROR
	result.ONSTATUS = false
	result.ERORCODE = "EROR"
	result.ERORMEMO = aERORMEMO
	return result
}

func  ToDATA(aOnStatus bool,aErorCode string,aErorMemo string,aListData interface{})  TEROR {
	var result TEROR
	result.ONSTATUS = aOnStatus
	result.ERORCODE = aErorCode
	result.ERORMEMO = aErorMemo
	result.LISTDATA = aListData
	return result
}

func ToTRUE(aListData interface{}) TEROR {
	var result TEROR
	result.ONSTATUS = true
	result.ERORCODE = "TRUE"
	result.ERORMEMO = "TRUE"
	result.LISTDATA = aListData
	return result
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func ToGUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(GetMd5String(base64.URLEncoding.EncodeToString(b)))
}

func ToUUID() string {
	uuid, eror := uuid.NewV4()
	if eror != nil {
		fmt.Printf("KzPackGo.ToUUID GET WRONG.%s", eror)
		return ""
	}
	return fmt.Sprintf("{%s}", uuid)
}

func ChEror(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
