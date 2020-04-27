package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

var (
	DeliverByUserId = "https://apiext.huya.com/message/deliverByUserId" // 观众端用户单播接口
	DeliverRoomByProfileId = "https://apiext.huya.com/message/deliverRoomByProfileId" // 直播间广播接口
	DeliverByProfileId = "https://apiext.huya.com/message/deliverByProfileId" // 主播单播接口
	MaxReCount = int8(3) // 消息重试次数
	ResponseStatus = map[string]string{
		"0": "操作成功",
		"1001": "参数错误",
		"1002": "调用失败",
		"1003": "服务器异常",
		"1005": "参数验证失败",
		"1010": "超出消息大小",
	}
)

type Notify struct {
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	ExtId string `json:"ext_id"`
	Token string `json:"token"`
}

type Data struct {
	Event string `json:"event"`
	Message string `json:"message"`
	ProfileId string `json:"profile_id"`
	Token string `json:"token"`
}

// 观众端用户单播接口
func (n *Notify) PushDeliverByUserId(event, message, profileId string) error {
	data := Data{
		Event:     event,
		Message:   message,
		ProfileId: profileId,
	}
	return n.sendMessage(DeliverByUserId, data, 0)
}

// 直播间广播接口
func (n *Notify) PushDeliverRoomByProfileId(event, message, profileId string) error {
	data := Data{
		Event:     event,
		Message:   message,
		ProfileId: profileId,
	}
	return n.sendMessage(DeliverRoomByProfileId, data, 0)
}

// 主播单播接口
func (n *Notify) PushDeliverByProfileId(event, message, profileId string) error {
	data := Data{
		Event:     event,
		Message:   message,
		ProfileId: profileId,
	}
	return  n.sendMessage(DeliverByProfileId, data, 0)
}

// 发送消息
func (n *Notify) sendMessage(url string, data Data, reCount int8) error {
	data.Token = n.Token
	if data.Token == "" {
		return errors.New("请传递正确的TOKEN")
	}

	newData := map[string]string{
		"profileId": data.ProfileId,
		"event":     data.Event,
		"message":   data.Message,
	}

	formData, err := json.Marshal(newData)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST",
		url+fmt.Sprintf("?appId=%s&extId=%s", n.AppId, n.ExtId),
		bytes.NewBuffer(formData))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", data.Token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	status := gjson.Get(string(body), "status")

	if status.Exists() {
		s := status.String()
		if s == "0" {
			return nil
		}

		if reCount >= MaxReCount {
			if e, ok := ResponseStatus[s]; ok {
				return errors.New(e)
			}
			return errors.New(fmt.Sprintf("发送事件失败，重试次数过多. 当前重试次数: %d 次 Event: %s Message: %s ProfileId: %s Token: %s",
				reCount, data.Event, data.Message, data.ProfileId, data.Token))
		}

		count := reCount + 1

		return n.sendMessage(url, data, count)
	}

	return errors.New(fmt.Sprintf("解析虎牙返回的json错误，没有找到key为status的项, body: %s", string(body)))
}
