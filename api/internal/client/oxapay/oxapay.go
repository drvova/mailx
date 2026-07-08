package oxapay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	MerchantKey string
	HTTP        *http.Client
}

func New(merchantKey string) *Client {
	return &Client{MerchantKey: merchantKey, HTTP: http.DefaultClient}
}

type InvoiceReq struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency,omitempty"`
	Lifetime    int     `json:"lifetime,omitempty"`
	CallbackURL string  `json:"callback_url,omitempty"`
	ReturnURL   string  `json:"return_url,omitempty"`
	Email       string  `json:"email,omitempty"`
	OrderID     string  `json:"order_id,omitempty"`
	Description string  `json:"description,omitempty"`
}

type InvoiceResp struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
	Data   struct {
		TrackID    string `json:"track_id"`
		PaymentURL string `json:"payment_url"`
		ExpiredAt  int64  `json:"expired_at"`
		Date       int64  `json:"date"`
	} `json:"data"`
	Error struct {
		Type    string `json:"type"`
		Key     string `json:"key"`
		Message string `json:"message"`
	} `json:"error"`
}

func (c *Client) CreateInvoice(req InvoiceReq) (*InvoiceResp, error) {
	body, _ := json.Marshal(req)
	httpReq, err := http.NewRequest("POST", "https://api.oxapay.com/v1/payment/invoice", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("merchant_api_key", c.MerchantKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result InvoiceResp
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	if result.Status != 200 {
		return &result, fmt.Errorf("oxapay error: %s", result.Error.Message)
	}

	return &result, nil
}

type PaymentInfo struct {
	TrackID string `json:"track_id"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Amount  float64 `json:"amount"`
	OrderID string `json:"order_id"`
	Email   string `json:"email"`
}

func (c *Client) GetPaymentInfo(trackID string) (*PaymentInfo, error) {
	httpReq, err := http.NewRequest("GET", fmt.Sprintf("https://api.oxapay.com/v1/payment/%s", trackID), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("merchant_api_key", c.MerchantKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var wrapper struct {
		Status int         `json:"status"`
		Data   PaymentInfo `json:"data"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(respBody, &wrapper); err != nil {
		return nil, err
	}

	if wrapper.Status != 200 {
		return nil, fmt.Errorf("oxapay error: %s", wrapper.Error.Message)
	}

	return &wrapper.Data, nil
}
