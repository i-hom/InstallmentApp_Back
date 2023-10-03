package controllers

import (
	"encoding/json"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/services"
)

type InstallmentController struct {
	installmentService *services.InstallmentService
}

func NewInstallmentController(installmentService *services.InstallmentService) *InstallmentController {
	return &InstallmentController{installmentService: installmentService}
}

func (ic *InstallmentController) Pay(params interface{}) models.RPCResponse {
	var installmentData models.InstallmentPayment
	if err := json.Unmarshal(models.GetRaw(params), &installmentData); err != nil {
		return errors.Invalid_parameter
	}

	if installmentData.InstallmentID.IsZero() || installmentData.CardID.IsZero() {
		return errors.Missing_parameter
	}

	return ic.installmentService.Pay(installmentData)
}
