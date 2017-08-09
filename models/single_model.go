package nbModels

// Result model of Single check API
type SingleCheckResponseModel struct {
	GenericResponseModel
	VerificationModel
	CreditsInfo CreditsInfoModel `json:"credits_info"`
}

// Request
type SingleCheckRequestModel struct {
	GenericRequestModel
	Email       string `json:"email"`
	AddressInfo bool `json:"address_info"`
	CreditInfo  bool `json:"credit_info"`
	Timeout     int `json:"timeout"`
}
