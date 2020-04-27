package api

import (
	"errors"
)

type Payload struct {
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
	AppId     string `json:"appId"`     // 小程序APP ID
	ExtId     string `json:"extId"`     // 小程序UUID
	Creator   string `json:"creator"`   // 创建者
	Role      string `json:"role"`      // 用户身份 U 用户 P主播 (可选)
	ProfileId string `json:"profileId"` // 主播union id
	RoomId    string `json:"roomId"`    // 房间ID
}

// 参数自动校验
func (p Payload) Valid() error  {
	if p.Creator == "" || p.RoomId == "" || p.ProfileId == "" {
		return errors.New("payload验证失败, 缺少必要参数")
	}
	return nil
}
