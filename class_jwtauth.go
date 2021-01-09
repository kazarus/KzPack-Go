package KzPack4Go

import "fmt"
import "github.com/dgrijalva/jwt-go"

type TKzCustomClaims struct {
	UserName string  `json:"userName"`
	EntityID string  `json:"entityID"`
	WhoBuild int64   `json:"whoBuild"`
	UserRoot int64   `json:"userRoot"`
	UserIndx int64   `json:"userIndx"`
	CreateAt float64 `json:"createAt"`
	jwt.StandardClaims
}

type TTokenUsr struct {
	UserName string  `json:"userName"`
	EntityID string  `json:"entityID"`
	WhoBuild int64   `json:"whoBuild"`
	UserRoot int64   `json:"userRoot"`
	UserIndx int64   `json:"userIndx"`
	CreateAt float64 `json:"createAt"`
}

type TErrorUsr struct {
	UserCode string
	ErrorCnt int64
}

type TBlackList map[string]TErrorUsr

func ParseToken1(tokenString string,tokenSecret string) (*TKzCustomClaims, error) {

	KzSecret := []byte(tokenSecret)

	var KzClaims = TKzCustomClaims{}
	MyClaims, eror := jwt.ParseWithClaims(tokenString, &KzClaims, func(token *jwt.Token) (interface{}, error) {
		return KzSecret, nil
	})

	if eror != nil {
		return nil, eror
	}

	if MyClaims != nil {
		if KzClaims, ok := MyClaims.Claims.(*TKzCustomClaims); ok && MyClaims.Valid {
			fmt.Println(KzClaims)
			return KzClaims, nil
		}
	}

	return nil, nil
}

func ParseToken2(tokenString string,tokenSecret string) (TTokenUsr, error) {

	KzSecret := []byte(tokenSecret)

	var KzClaims = TKzCustomClaims{}
	var TokenUsr = TTokenUsr{}
	MyClaims, eror := jwt.ParseWithClaims(tokenString, &KzClaims, func(token *jwt.Token) (interface{}, error) {
		return KzSecret, nil
	})

	if eror != nil {
		return TokenUsr, eror
	}

	if MyClaims != nil {
		if KzClaims, ok := MyClaims.Claims.(*TKzCustomClaims); ok && MyClaims.Valid {
			fmt.Println(KzClaims)

			TokenUsr.EntityID = KzClaims.EntityID
			TokenUsr.CreateAt = KzClaims.CreateAt
			TokenUsr.UserName = KzClaims.UserName
			TokenUsr.WhoBuild = KzClaims.WhoBuild
			TokenUsr.UserRoot = KzClaims.UserRoot
			TokenUsr.UserIndx = KzClaims.UserIndx

			return TokenUsr, nil
		}
	}

	return TokenUsr, nil
}
