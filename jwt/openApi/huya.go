package openApi


import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type HuYaJWT struct {
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type Data struct {
	Iat int64 `json:"iat"`
	Exp       int64  `json:"exp"`
}

func NewJWT(appId, appSecret string) *HuYaJWT {
	return &HuYaJWT{
		AppId:     appId,
		AppSecret: appSecret,
	}
}

// 获取开放平台API Token
func (j *HuYaJWT) GetOpenAPIToken(data interface{}) (string, error) {

	newData, ok := data.(Data)

	if !ok {
		return "", errors.New("data decode fail")
	}
	payload := Payload{
		Iat:   newData.Iat,
		Exp:   newData.Exp,
		AppId: j.AppId,
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