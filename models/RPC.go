package models

import "encoding/json"

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RPCResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetRaw(params interface{}) []byte {
	data, _ := json.Marshal(params)
	return data
}
