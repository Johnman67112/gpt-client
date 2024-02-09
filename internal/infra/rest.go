package infra

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Johnman67112/gpt-client/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-retryablehttp"
)

var GeneralClient *Client

type Config struct {
	Endpoint string `env:"CHATGPT_URL" envDefault:"http://localhost:8000"`
	ApiKey   string `env:"CHATGPT_APIKEY" envDefault:"1234"`
}

type Client struct {
	Config *Config
	Client *retryablehttp.Client
}

type Response struct {
	*http.Response
}

func NewClient(config *Config, retryClient *retryablehttp.Client) *Client {
	return &Client{
		Config: config,
		Client: retryClient,
	}
}

func (c *Client) NewRequest(ctx *gin.Context, method, path string, body io.Reader) (*retryablehttp.Request, error) {
	r, err := retryablehttp.NewRequestWithContext(ctx, method, c.Config.Endpoint+path, body)
	if err != nil {
		return &retryablehttp.Request{}, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", c.Config.ApiKey)

	return r, nil
}

func (c *Client) DoRequest(ctx *gin.Context, req *retryablehttp.Request) (*Response, error) {
	r, err := c.Client.Do(req)
	if err != nil {
		return &Response{}, err
	}

	return &Response{Response: r}, err
}

func (r *Response) Decode(payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func (c *Client) ChatRequest(ctx *gin.Context, request domain.ChatRequest) (*domain.ChatResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return &domain.ChatResponse{}, err
	}

	req, err := c.NewRequest(ctx, http.MethodPost, "", bytes.NewBuffer(requestBody))
	if err != nil {
		return &domain.ChatResponse{}, err
	}

	res, err := c.DoRequest(ctx, req)
	if err != nil {
		return &domain.ChatResponse{}, err
	}

	defer res.Body.Close()

	var response domain.ChatResponse
	if err := res.Decode(&response); err != nil {
		return &domain.ChatResponse{}, err
	}

	return &response, nil
}
