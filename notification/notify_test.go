package notification

import (
	"fmt"
	"github.com/loktarjuge/huyaHelper/jwt/api"
	"testing"
	"time"
)

// 运行测试要求在IP在虎牙的白名单内
// 运行测试前需要注意更改下列设置
var notify = Notify{
	AppId:     "", // 开发者 app id
	AppSecret: "", // 开发者 app secret
	ExtId:     "", // 小程序ID
}

var profileId = "" // 主播unionId
var roomId = "" // 主播房间号

func TestNotify_PushDeliverByProfileId(t *testing.T) {
	type args struct {
		event     string
		message   string
		profileId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNotify_PushDeliverByProfileId",
			args:args{
				event:     "test event",
				message:   "test",
				profileId: profileId,
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &notify
			if err := n.PushDeliverByProfileId(tt.args.event, tt.args.message, tt.args.profileId); (err != nil) != tt.wantErr {
				t.Errorf("PushDeliverByProfileId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotify_PushDeliverByUserId(t *testing.T) {
	type args struct {
		event     string
		message   string
		profileId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"TestNotify_PushDeliverByUserId",
			args:args{
				event:     "TestNotify_PushDeliverByUserId",
				message:   "TestNotify_PushDeliverByUserId",
				profileId: profileId,
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &notify
			if err := n.PushDeliverByUserId(tt.args.event, tt.args.message, tt.args.profileId); (err != nil) != tt.wantErr {
				t.Errorf("PushDeliverByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotify_PushDeliverRoomByProfileId(t *testing.T) {
	type args struct {
		event     string
		message   string
		profileId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:	"TestNotify_PushDeliverRoomByProfileId",
			args:args{
				event:     "TestNotify_PushDeliverRoomByProfileId",
				message:   "TestNotify_PushDeliverRoomByProfileId",
				profileId: profileId,
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &notify
			if err := n.PushDeliverRoomByProfileId(tt.args.event, tt.args.message, tt.args.profileId); (err != nil) != tt.wantErr {
				t.Errorf("PushDeliverRoomByProfileId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func init()  {
	jwt := api.NewJWT(notify.AppId, notify.AppSecret, notify.ExtId)
	token, err := jwt.GetJWTToken(api.Data{
		Iat:       time.Now().Unix(),
		Exp:       time.Now().Add(time.Minute).Unix(),
		Creator:   "DEV", // 创建者（token生成方：SYS平台，DEV开发者）
		Role:      "P", // 用户身份：U用户，P主播（可选）
		ProfileId: profileId,
		RoomId:    roomId,
	})
	if err != nil {
		panic(fmt.Sprintf("token生成失败 : %s", err.Error()))
	}
	notify.Token = token
}