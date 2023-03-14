package model

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RPCResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}
