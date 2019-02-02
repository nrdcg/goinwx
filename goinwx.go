package goinwx

import (
	"net/url"

	"github.com/kolo/xmlrpc"
)

const (
	APIBaseUrl        = "https://api.domrobot.com/xmlrpc/"
	APISandboxBaseUrl = "https://api.ote.domrobot.com/xmlrpc/"
	APILanguage       = "eng"
)

// Client manages communication with INWX API.
type Client struct {
	// HTTP client used to communicate with the INWX API.
	RPCClient *xmlrpc.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// API username
	Username string

	// API password
	Password string

	// User agent for client
	APILanguage string

	// Services used for communicating with the API
	Account     AccountService
	Domains     DomainService
	Nameservers NameserverService
	Contacts    ContactService
}

type ClientOptions struct {
	Sandbox bool
}

type Request struct {
	ServiceMethod string
	Args          map[string]interface{}
}

// NewClient returns a new INWX API client.
func NewClient(username, password string, opts *ClientOptions) *Client {
	var useSandbox bool
	if opts != nil {
		useSandbox = opts.Sandbox
	}

	var baseURL *url.URL

	if useSandbox {
		baseURL, _ = url.Parse(APISandboxBaseUrl)
	} else {
		baseURL, _ = url.Parse(APIBaseUrl)
	}

	rpcClient, _ := xmlrpc.NewClient(baseURL.String(), nil)

	client := &Client{RPCClient: rpcClient,
		BaseURL:  baseURL,
		Username: username,
		Password: password,
	}

	client.Account = &AccountServiceOp{client: client}
	client.Domains = &DomainServiceOp{client: client}
	client.Nameservers = &NameserverServiceOp{client: client}
	client.Contacts = &ContactServiceOp{client: client}

	return client
}

// NewRequest creates an API request.
func (c *Client) NewRequest(serviceMethod string, args map[string]interface{}) *Request {
	if args != nil {
		args["lang"] = APILanguage
	}

	return &Request{ServiceMethod: serviceMethod, Args: args}
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req Request) (*map[string]interface{}, error) {
	var resp Response
	err := c.RPCClient.Call(req.ServiceMethod, req.Args, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.ResponseData, CheckResponse(&resp)
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *Response) error {
	if c := r.Code; c >= 1000 && c <= 1500 {
		return nil
	}

	return &ErrorResponse{Code: r.Code, Message: r.Message, Reason: r.Reason, ReasonCode: r.ReasonCode}
}
