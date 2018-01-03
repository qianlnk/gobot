package gobot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetUUID(t *testing.T) {
	cfg := Load()
	fmt.Println(cfg)
	wx, err := NewWecat(cfg)
	if err != nil {
		panic(err)
	}

	wx.Start()
}

func TestTuling(t *testing.T) {
	params := make(map[string]interface{})
	params["userid"] = "123123123"
	params["key"] = "808811ad0fd34abaa6fe800b44a9556a"
	params["info"] = "你好"

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "http://www.tuling123.com/openapi/api", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Referer", WxReferer)
	req.Header.Add("User-agent", WxUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}
