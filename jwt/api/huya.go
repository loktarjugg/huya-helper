package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type HuYaJWT struct {
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	ExtId string `json:"ext_id"`
}

type Data struct {
	Iat int64 `json:"iat"`
	Exp       int64  `json:"exp"`
	Creator string `json:"creator"`
	Role string `json:"role"`
	ProfileId string `json:"profile_id"`
	RoomId string `json:"room_id"`
}

func NewJWT(appId, appSecret, extId string) *HuYaJWT {
	return &HuYaJWT{
		AppId:     appId,
		AppSecret: appSecret,
		ExtId: extId,
	}
}

func (j *HuYaJWT) GetJWTToken(data interface{}) (string, error) {

	newData, ok := data.(Data)

	if !ok {
		return "", errors.New("data decode fail")
	}

	payload := Payload {
		Iat:       newData.Iat,
		Exp:       newData.Exp,
		AppId:     j.AppId,
		ExtId:     j.ExtId,
		Creator:   newData.Creator,
		Role:      newData.Role,
		ProfileId: newData.ProfileId,
		RoomId:    newData.RoomId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(j.getSecret())

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *HuYaJWT) getSecret() interface{} {
	return []byte(j.AppSecret)
}