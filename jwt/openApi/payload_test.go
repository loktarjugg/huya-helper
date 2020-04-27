package openApi

import (
	"testing"
	"time"
)

var appId = "" // app id

func TestPayload_Valid(t *testing.T) {
	type fields struct {
		Iat   int64
		Exp   int64
		AppId string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "TestPayload_Valid",
			fields:  fields{
				Iat:   time.Now().Unix(),
				Exp:   time.Now().Add(time.Minute).Unix(),
				AppId: appId,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Payload{
				Iat:   tt.fields.Iat,
				Exp:   tt.fields.Exp,
				AppId: tt.fields.AppId,
			}
			if err := p.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}