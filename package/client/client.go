package client

import (
	"encoding/json"
	"fmt"
	"training/package/crypto"
	"training/package/log"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

const dataHash = "DATA-HASH"
const dataSignature = "DATA-SIGNATURE"
const systemName = "SYSTEM-NAME"

type Client struct {
	client     resty.Client
	privateKey []byte
	system     string
}

type Response struct {
	Code    int         `json:"code"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func New(url string, privateKey []byte, system string) (*Client, error) {
	c := &Client{
		client:     *resty.New(),
		privateKey: privateKey,
		system:     system,
	}
	c.client.BaseURL = url
	return c, nil
}
func (c *Client) SetHeader(key, value string) {
	c.client.Header.Add(key, value)
}

func (c *Client) Post(ctx echo.Context, endPoint string, data interface{}) (*Response, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Errorf("error marshalling data: %v", err)
		return nil, err
	}
	hash, signature, err := crypto.Sign(bytes, c.privateKey)
	if err != nil {
		log.Errorf("error signing message: %v", err)
		return nil, err
	}
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader(dataHash, hash).
		SetHeader(dataSignature, signature).
		SetHeader(systemName, c.system).
		SetBody(data).Post(c.client.BaseURL + endPoint)
	if err != nil {
		log.Errorf("error getting response: %v", err)
		return nil, err
	}
	var response Response
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Errorf("error unmarshalling response: %v", err)
		return nil, err
	}
	return &response, nil
}

func (c *Client) Get(ctx echo.Context, endPoint string) (*Response, error) {
	url := c.client.BaseURL + endPoint
	bytes := []byte("get")
	hash, signature, err := crypto.Sign(bytes, c.privateKey)
	if err != nil {
		log.Errorf("error signing message: %v", err)
		return nil, err
	}
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader(dataHash, hash).
		SetHeader(dataSignature, signature).
		SetHeader(systemName, c.system).
		Get(url)

	if err != nil {
		log.Errorf("error geting response: %v", err)
		return nil, err
	}

	var response Response
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Errorf("error unmarshalling response: %v", err)
		return nil, err
	}
	return &response, nil
}

func convertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int8, int32, int64, float32, float64, bool:
		return fmt.Sprintf("%v", v)
	case []interface{}:
		var strValues []string
		for _, num := range v {
			strValues = append(strValues, fmt.Sprintf("%v", num))
		}
		return strings.Join(strValues, ",")
	default:
		return fmt.Sprintf("%v", v)
	}
}
