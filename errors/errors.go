package errors

import "installment_back/models"

var Success = models.RPCResponse{Code: 0, Message: "Success"}
var Missing_parameter = models.RPCResponse{Code: 1, Message: "Missing one of params"}
var Invalid_parameter = models.RPCResponse{Code: 2, Message: "Invalid parameter"}
var User_not_found = models.RPCResponse{Code: 3, Message: "User not found"}
var Insufficient_balance = models.RPCResponse{Code: 4, Message: "Insufficient balance"}
var Card_number_length_not_correct = models.RPCResponse{Code: 5, Message: "Card number should be 16 digits"}
var Payment_greater_than_balance = models.RPCResponse{Code: 6, Message: "Payment greater than installment balance"}
var Card_not_found = models.RPCResponse{Code: 7, Message: "Card not found"}
var Method_not_found = models.RPCResponse{Code: 8, Message: "Method not found"}
var Failed_to_deposite = models.RPCResponse{Code: 9, Message: "Failed to deposit"}
var Failed_to_add_card = models.RPCResponse{Code: 10, Message: "Failed to add card"}
var Installment_not_found = models.RPCResponse{Code: 11, Message: "Installment not found"}
var Failed_to_get_cards = models.RPCResponse{Code: 12, Message: "Failed to get cards"}
var Payment_not_found = models.RPCResponse{Code: 13, Message: "Payment not found"}
var Failed_to_add_payment = models.RPCResponse{Code: 14, Message: "Failed to add payment"}
