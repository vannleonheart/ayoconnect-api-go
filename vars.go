package ayoconnect

type Client struct {
	Config      *Config
	accessToken string
	requestId   string
	ipAddress   string
	phoneNumber string
}

type Config struct {
	BaseUrl      string `json:"base_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	MerchantCode string `json:"merchant_code"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	PhoneNumber  string `json:"phone_number"`
	IpAddress    string `json:"ip_address"`
}

type HttpErrorResponse struct {
	ErrorResponse
	ErrorText string `json:"Error"`
	ErrorCode string `json:"ErrorCode"`
}

type ErrorResponse struct {
	Code    int64                `json:"code"`
	Message string               `json:"message"`
	Errors  *[]ErrorResponseItem `json:"errors,omitempty"`
}

type ErrorResponseItem struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type AuthenticationResponse struct {
	ApiProductList   string `json:"apiProductList"`
	OrganizationName string `json:"organizationName"`
	DeveloperEmail   string `json:"developer.email"`
	TokenType        string `json:"tokenType"`
	ResponseTime     string `json:"responseTime"`
	ClientId         string `json:"clientId"`
	AccessToken      string `json:"accessToken"`
	ExpiresIn        string `json:"expiresIn"`
}

type GeneralResponse struct {
	Code            int64                `json:"code"`
	Message         string               `json:"message"`
	ResponseTime    string               `json:"responseTime"`
	TransactionId   string               `json:"transactionId,omitempty"`
	ReferenceNumber string               `json:"referenceNumber,omitempty"`
	CustomerId      string               `json:"customerId,omitempty"`
	MerchantCode    string               `json:"merchantCode,omitempty"`
	Errors          *[]ErrorResponseItem `json:"errors,omitempty"`
}

type BeneficiaryResponse struct {
	GeneralResponse
	BeneficiaryDetails struct {
		BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
		BeneficiaryBankCode      string `json:"beneficiaryBankCode"`
		BeneficiaryId            string `json:"beneficiaryId"`
		BeneficiaryName          string `json:"beneficiaryName"`
		AccountType              string `json:"accountType"`
	} `json:"beneficiaryDetails"`
}

type TransactionResponse struct {
	GeneralResponse
	Transaction struct {
		Amount          string `json:"amount"`
		Currency        string `json:"currency"`
		BeneficiaryId   string `json:"beneficiaryId"`
		Status          int64  `json:"status"`
		ReferenceNumber string `json:"referenceNumber"`
		Remark          string `json:"remark"`
	} `json:"transaction"`
}

type BalanceResponse struct {
	GeneralResponse
	AccountInfo []struct {
		AvailableBalance struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"availableBalance"`
	} `json:"accountInfo"`
}
