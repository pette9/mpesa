package mpesa

//AccountBalanceParameters ...
type AccountBalanceParameters struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	PartyB                 string `json:"PartyB"`
	ReceiverIdentifierType string `json:"ReceiverIdentifierType"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
	AccountType            string `json:"AccountType"`
}

//B2BQueryParameters ...
type B2BQueryParameters struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	Amount                 string `json:"Amount"`
	PartyA                 string `json:"PartyA"`
	SenderIdentifierType   string `json:"SenderIdentifierType"`
	PartyB                 string `json:"PartyB"`
	ReceiverIdentifierType string `json:"RecieverIdentifierType"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
	AccountReference       string `json:"AccountReference"`
}

//B2CQueryParameters ...
type B2CQueryParameters struct {
	InitiatorName      string `json:"InitiatorName"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	Amount             string `json:"Amount"`
	PartyA             string `json:"PartyA"`
	PartyB             string `json:"PartyB"`
	Remarks            string `json:"Remarks"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
	Occasion           string `json:"Occasion"`
}

//C2BRegisterURLParameters ...
type C2BRegisterURLParameters struct {
	ConfirmationURL string `json:"ConfirmationURL"`
	ResponseType    string `json:"ResponseType"`
	ShortCode       string `json:"ShortCode"`
}

//C2BSimulateParameters ...
type C2BSimulateParameters struct {
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	MSISDN        string `json:"MSISDN"`
	BillRefNumber string `json:"BillRefNumber"`
	ShortCode     string `json:"ShortCode"`
}

//LipaNaMpesaOnlinePaymentParameters ...
type LipaNaMpesaOnlinePaymentParameters struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

//LipaNaMpesaOnlineQueryParameters ...
type LipaNaMpesaOnlineQueryParameters struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}

//ReversalParameters ...
type ReversalParameters struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	ReceiverParty          string `json:"ReceiverParty"`
	ReceiverIdentifierType string `json:"ReceiverIdentifierType"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
	TransactionID          string `json:"TransactionID"`
	Occasion               string `json:"Occasion"`
}

//TransactionStatusParameters ...
type TransactionStatusParameters struct {
	CommandID          string `json:"CommandID"`
	ShortCode          string `json:"ShortCode"`
	IdentifierType     string `json:"IdentifierType"`
	Remarks            string `json:"Remarks"`
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
	TransactionID      string `json:"TransactionID"`
	Occasion           string `json:"Occasion"`
}
