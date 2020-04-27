package openApi

import (
	"errors"
)

type Payload struct {
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
	AppId     string `json:"appId"`
}

// 参数自动校验
func (p Payload) Valid() error  {
	if p.Iat == 0 || p.Exp == 0 || p.AppId == "" {
		return errors.New("虎牙token生成错误, 缺少必要参数")
	}
	return nil
}
