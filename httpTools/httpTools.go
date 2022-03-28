package httpTools

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/coreservice-io/dns-client/respMsg"
	"github.com/imroc/req"
)

func Get(url string, token string, timeOutSec int, respResult interface{}) error {
	return request("GET", url, token, nil, timeOutSec, respResult)
}

func POST(url string, token string, postData interface{}, timeOutSec int, respResult interface{}) error {
	return request("POST", url, token, postData, timeOutSec, respResult)
}

func request(method string, url string, token string, postData interface{}, timeOutSec int, respResult interface{}) error {
	if respResult != nil {
		t := reflect.TypeOf(respResult).Kind()
		if t != reflect.Ptr && t != reflect.Slice && t != reflect.Map {
			return errors.New("value only support Pointer Slice and Map")
		}
	}

	r := req.New()
	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer " + token,
	}
	r.SetTimeout(time.Duration(timeOutSec) * time.Second)
	var resp *req.Resp
	var err error
	switch method {
	case "GET":
		resp, err = r.Get(url, authHeader)
	case "POST":
		resp, err = r.Post(url, authHeader, req.BodyJSON(postData))
	default:
		return fmt.Errorf("unsupported request method:%s", method)
	}
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 200 {
		return fmt.Errorf("network error, http code:%d", resp.Response().StatusCode)
	}
	respData := &respMsg.RespBody{
		Result: respResult,
	}
	err = resp.ToJSON(&respData)
	if err != nil {
		return err
	}
	if respData.Status <= 0 {
		return errors.New(respData.Msg)
	}
	return nil
}
