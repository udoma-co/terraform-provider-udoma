package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

type UdomaClient struct {
	host         string
	scheme       string
	apiKey       string
	apiSecretKey string
}

var ErrNotFound = fmt.Errorf("not found")

func NewClient(cfg Config) (*UdomaClient, error) {

	u, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid Udoma endpoint provided (%v): %v. Expected format is `https://udoma.co", cfg.Endpoint, err)
	}

	client := UdomaClient{
		host:         u.Host,
		scheme:       u.Scheme,
		apiKey:       cfg.ApiKey,
		apiSecretKey: cfg.ApiSecretKey,
	}

	return &client, nil
}

// RoundTrip implements the http client `http.RoundTripper` interface. This
// generates the authorization token to be used when calling the UDOMA API
// with HMAC authentication.
func (c *UdomaClient) RoundTrip(r *http.Request) (res *http.Response, e error) {

	if r == nil {
		panic(fmt.Errorf("request is nil"))
	}

	b := []byte{}
	if r.Body != nil {
		b, _ = io.ReadAll(r.Body)
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewReader(b))
	}

	md5Sum := md5.Sum(b)
	contentType := r.Header.Get("Content-Type")
	validity := time.Now().Add(time.Second * 60).Unix()

	// create signature with correct format
	toSign := fmt.Sprintf("%s\n%s\n%d\n%x", r.URL.String(), contentType, validity, md5Sum)

	mac := hmac.New(sha256.New, []byte(c.apiSecretKey))
	if _, err := mac.Write([]byte(toSign)); err != nil {
		return nil, err
	}

	signature := hex.EncodeToString(mac.Sum(nil))

	// update request data and headers
	size := len(b)
	r.ContentLength = int64(size)
	r.Header.Set("Content-Length", strconv.Itoa(size))

	r.Header["X-ZEST-Validity"] = []string{strconv.FormatInt(validity, 10)}

	md5Checksum := fmt.Sprintf("%x", md5Sum)
	r.Header["Content-Md5"] = []string{md5Checksum}

	authString := fmt.Sprintf("ZEST %s:%s", c.apiKey, signature)
	r.Header["Authorization"] = []string{authString}

	// Send the request, get the response (or the error)
	return http.DefaultTransport.RoundTrip(r)
}

func (c *UdomaClient) GetApi() *api.DefaultAPIService {

	cfg := api.NewConfiguration()
	cfg.Host = c.host
	cfg.Scheme = c.scheme

	httpClient := &http.Client{Transport: c}

	cfg.HTTPClient = httpClient

	client := api.NewAPIClient(cfg)

	return client.DefaultAPI
}
