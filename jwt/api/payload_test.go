package api

import (
	"testing"
)

func TestPayload_Valid(t *testing.T) {
	type fields struct {
		Iat       int64
		Exp       int64
		AppId     string
		ExtId     string
		Creator   string
		Role      string
		ProfileId string
		RoomId    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestPayload_Valid",
			fields:fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Payload{
				Iat:       tt.fields.Iat,
				Exp:       tt.fields.Exp,
				AppId:     tt.fields.AppId,
				ExtId:     tt.fields.ExtId,
				Creator:   tt.fields.Creator,
				Role:      tt.fields.Role,
				ProfileId: tt.fields.ProfileId,
				RoomId:    tt.fields.RoomId,
			}
			if err := p.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}