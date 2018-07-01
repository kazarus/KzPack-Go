package KzPack4Go


import "encoding/json"
import "log"

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

func ChEror(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
