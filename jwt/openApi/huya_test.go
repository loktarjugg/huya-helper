package openApi

import (
	"testing"
	"time"
)

var (
	myJwt *HuYaJWT
)

func init()  {
	myJwt = &HuYaJWT{
		AppId:     "",
		AppSecret: "",
	}
}

func TestHuYaJWT_GetOpenAPIToken(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestHuYaJWT_GetOpenAPIToken",
			args:args{data:Data{
				Iat: time.Now().Unix(),
				Exp: time.Now().Add(time.Minute).Unix(),
			}},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := myJwt
			got, err := j.GetOpenAPIToken(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOpenAPIToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GetOpenAPIToken() got = %v", got)
			}
		})
	}
}