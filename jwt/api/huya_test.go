package api

import (
	"testing"
	"time"
)

var (
	myJwt *HuYaJWT
	appId string
	appSecret string
	extId string
)

var profileId = "" // 主播unionId
var roomId = "" // 主播房间号

func init()  {
	myJwt = NewJWT(appId, appSecret, extId)
}

func TestHuYaJWT_GetJWTToken(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestHuYaJWT_GetJWTToken",
			args: args{data:Data{
				Iat:       time.Now().Unix(),
				Exp:       time.Now().Add(time.Minute).Unix(),
				Creator:   "DEV",
				Role:      "P",
				ProfileId: profileId,
				RoomId:    roomId,
			}},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := myJwt.GetJWTToken(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GetJWTToken() got = %v", got)
			}
		})
	}
}