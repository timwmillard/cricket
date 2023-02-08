package grassroots

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) GetMatch(ctx context.Context, legacyMatchID string) (Match, error) {
	path := fmt.Sprintf("/scores/matches/00000000-0000-0000-0000-000000000000?LegacyMatchId=%s&ResponseModifier=IncludeScorecard&format=json&jsconfig=eccn", legacyMatchID)
	var match Match
	_, err := c.doRequest(ctx, http.MethodGet, path, nil, &match)
	if err != nil {
		return Match{}, err
	}
	return match, nil
}

const DefaultBaseURL = "https://grassrootsapi.cricket.com.au/"

const (
	defaultClientTimeout       = 10
	defaultConnTimeout         = 5
	defaultTLSHandshakeTimeout = 5
)

type Client struct {
	httpClient *http.Client

	BaseURL *url.URL
	authKey string
}

// NewClient create a new client
func NewClient(authKey string) *Client {
	client, err := NewClientWithBaseURL(DefaultBaseURL, authKey)
	if err != nil {
		panic(err)
	}
	return client
}

// NewClient create a new client, with default timeouts.
func NewClientWithBaseURL(baseURL, authKey string) (*Client, error) {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: defaultConnTimeout * time.Second,
		}).Dial,
		TLSHandshakeTimeout: defaultTLSHandshakeTimeout * time.Second}
	client := &http.Client{
		Timeout:   defaultClientTimeout * time.Second,
		Transport: netTransport,
	}
	return NewClientWithHTTPClient(client, baseURL, authKey)
}

// NewClientWithHTTPClient create a new client, with a http.Client.
func NewClientWithHTTPClient(client *http.Client, baseURL, authKey string) (*Client, error) {
	var c Client
	var err error

	c.httpClient = client
	c.BaseURL, err = url.Parse(baseURL)
	if err != nil {
		return &c, err
	}
	c.authKey = authKey
	return &c, nil
}

// doRequest is a convenient function to combine newRequest and do.
func (c *Client) doRequest(ctx context.Context, method string, path string, reqBody any, respBody any) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, path, reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, &respBody)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// newRequest creates a http.Request.  To be pass to the do method.
func (c *Client) newRequest(ctx context.Context, method string, path string, body any) (*http.Request, error) {
	url, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var bodyBuf io.ReadWriter
	if body != nil {
		bodyBuf = &bytes.Buffer{}
		encoder := json.NewEncoder(bodyBuf)
		err := encoder.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, url.String(), bodyBuf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("authkey", c.authKey)

	return req, nil
}

// do the actual request.
func (c *Client) do(req *http.Request, v any) (*http.Response, error) {

	// DEBUG: Uncomment to dump Request
	// dumpReq, _ := httputil.DumpRequest(req, true)
	// fmt.Println(string(dumpReq))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// DEBUG: Uncomment to dump Response
	// dumpResp, _ := httputil.DumpResponse(resp, true)
	// fmt.Println(string(dumpResp))

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, errorFromResponse(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type Error struct {
	StatusCode int    `json:"-"`     // http status code
	Title      string `json:"title"` // status message
}

func errorFromResponse(resp *http.Response) Error {
	var e Error
	err := json.NewDecoder(resp.Body).Decode(&e)
	e.StatusCode = resp.StatusCode
	if err != nil {
		e.Title = "error decoding response title"
	}
	if e.Title == "" {
		e.Title = http.StatusText(e.StatusCode)
	}
	return e
}

func (e Error) Error() string {
	return fmt.Sprintf("api error %d: %s", e.StatusCode, e.Title)
}
