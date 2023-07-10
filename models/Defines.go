package models

var Missing_parameter = RPCResponse{Code: 1, Message: "Missing one of params"}
var Invalid_parameter = RPCResponse{Code: 2, Message: "Invalid parameter"}
var User_not_found = RPCResponse{Code: 3, Message: "User not found"}
var Insufficient_balance = RPCResponse{Code: 4, Message: "Insufficient balance"}
var Card_number_length_not_correct = RPCResponse{Code: 5, Message: "Card number should be 16 digits"}
var Payment_greater_than_balance = RPCResponse{Code: 6, Message: "Payment greater than installment balance"}
var Card_not_found = RPCResponse{Code: 7, Message: "Card not found"}
var Method_not_found = RPCResponse{Code: 8, Message: "Method not found"}
var Failde_to_deposite = RPCResponse{Code: 9, Message: "Failde to deposite"}
