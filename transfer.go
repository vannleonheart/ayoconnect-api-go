package ayoconnect

import (
	"fmt"
)

func (c *Client) AddBeneficiary(transactionId, accountNo, bankCode string) (*BeneficiaryResponse, error) {
	if err := c.prepareAccessToken(); err != nil {
		return nil, err
	}

	phoneNumber := c.Config.PhoneNumber
	if len(c.phoneNumber) > 0 {
		phoneNumber = c.phoneNumber
	}

	ipAddress := c.Config.IpAddress
	if len(c.ipAddress) > 0 {
		ipAddress = c.ipAddress
		c.ipAddress = ""
	}

	targetUrl := fmt.Sprintf("%s/api/v1/bank-disbursements/beneficiary", c.Config.BaseUrl)
	requestId := c.requestId

	c.requestId = ""

	requestHeaders := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		"Authorization":    fmt.Sprintf("Bearer %s", c.accessToken),
		"A-Correlation-ID": requestId,
		"A-Merchant-Code":  c.Config.MerchantCode,
		"A-Latitude":       c.Config.Latitude,
		"A-Longitude":      c.Config.Longitude,
	}

	requestData := map[string]interface{}{
		"transactionId": transactionId,
		"phoneNumber":   phoneNumber,
		"customerDetails": map[string]string{
			"ipAddress": ipAddress,
		},
		"beneficiaryAccountDetails": map[string]string{
			"accountNumber": accountNo,
			"bankCode":      bankCode,
		},
	}

	var result BeneficiaryResponse

	if err := sendHttpPost(targetUrl, &requestData, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Disburse(transactionId, customerId, beneficiaryId, amount, currency, remark string) (*TransactionResponse, error) {
	if err := c.prepareAccessToken(); err != nil {
		return nil, err
	}

	targetUrl := fmt.Sprintf("%s/api/v1/bank-disbursements/disbursement", c.Config.BaseUrl)
	requestId := c.requestId

	c.requestId = ""

	requestHeaders := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		"Authorization":    fmt.Sprintf("Bearer %s", c.accessToken),
		"A-Correlation-ID": requestId,
		"A-Merchant-Code":  c.Config.MerchantCode,
		"A-Latitude":       c.Config.Latitude,
		"A-Longitude":      c.Config.Longitude,
		"A-IP-Address":     c.Config.IpAddress,
	}

	requestData := map[string]interface{}{
		"transactionId": transactionId,
		"customerId":    customerId,
		"beneficiaryId": beneficiaryId,
		"amount":        amount,
		"currency":      currency,
		"remark":        remark,
	}

	var result TransactionResponse

	if err := sendHttpPost(targetUrl, &requestData, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDisbursementStatusByCorrelationId(correlationId, transactionId, transactionReferenceNumber, beneficiaryId, customerId string) (*TransactionResponse, error) {
	if err := c.prepareAccessToken(); err != nil {
		return nil, err
	}

	targetUrl := fmt.Sprintf("%s/api/v1/bank-disbursements/status/%s", c.Config.BaseUrl, correlationId)
	requestId := c.requestId

	c.requestId = ""

	ipAddress := c.Config.IpAddress
	if len(c.ipAddress) > 0 {
		ipAddress = c.ipAddress
		c.ipAddress = ""
	}

	requestHeaders := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		"Authorization":    fmt.Sprintf("Bearer %s", c.accessToken),
		"A-Correlation-ID": requestId,
		"A-Merchant-Code":  c.Config.MerchantCode,
		"A-Latitude":       c.Config.Latitude,
		"A-Longitude":      c.Config.Longitude,
		"A-IP-Address":     ipAddress,
	}

	requestData := map[string]interface{}{
		"transactionId":              transactionId,
		"transactionReferenceNumber": transactionReferenceNumber,
		"beneficiaryId":              beneficiaryId,
		"customerId":                 customerId,
	}

	var result TransactionResponse

	if err := sendHttpGet(targetUrl, &requestData, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetBalance() (*BalanceResponse, error) {
	if err := c.prepareAccessToken(); err != nil {
		return nil, err
	}

	targetUrl := fmt.Sprintf("%s/api/v1/merchants/balance", c.Config.BaseUrl)
	requestId := c.requestId
	transactionId := RandomString(32)

	c.requestId = ""

	requestHeaders := map[string]string{
		"Content-Type":     "application/json",
		"Accept":           "application/json",
		"Authorization":    fmt.Sprintf("Bearer %s", c.accessToken),
		"A-Correlation-ID": requestId,
		"A-Merchant-Code":  c.Config.MerchantCode,
	}

	requestData := map[string]interface{}{
		"transactionId": transactionId,
	}

	var result BalanceResponse

	if err := sendHttpGet(targetUrl, &requestData, &requestHeaders, &result); err != nil {
		return nil, err
	}

	if result.Code != 200 {
		return nil, ErrorResponse{
			Code:    result.Code,
			Message: result.Message,
			Errors:  result.Errors,
		}
	}

	return &result, nil
}
