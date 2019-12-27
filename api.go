// youtrack package manipulates YouTrack issues via the REST API as described by
// https://www.jetbrains.com/help/youtrack/incloud/Resources-for-Developers.html
package youtrack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Api is the YouTrack API context.
// BaseURL and Token must be set.
type Api struct {
	BaseURL       *url.URL
	Token         string
	EnableTracing bool
}

// NewDefaultApi returns a new API using a baseURL and permanent authorization token
// from either flags or environment.
func NewDefaultApi() (*Api, error) {
	baseURL, err := defaultURL()
	if err != nil {
		log.Print("Could not determine default YouTrack URL.")
		return nil, err
	}

	token, err := defaultToken()
	if err != nil {
		log.Print("Could not determine default YouTrack Token.")
		return nil, err
	}

	return &Api{
		BaseURL:       baseURL,
		Token:         token,
		EnableTracing: true,
	}, nil
}

func (api *Api) trace(v ...interface{}) {
	if api.EnableTracing {
		log.Println(v...)
	}
}

// DoRequest makes an authenticated HTTP request to the YouTrack API.
// jsonRequest and jsonResult are both optional, and depend on the request being made. A GET request,
// for example, should probably set jsonRequestBody to nil.
func (api *Api) DoRequest(ctx context.Context, resource, method string, jsonRequest, jsonResult interface{}) error {
	api.trace(method, resource)

	resourceURL := api.BaseURL.ResolveReference(&url.URL{Path: resource})
	api.trace("Resolved URL", resourceURL)

	reqBody := new(bytes.Buffer)

	if jsonRequest != nil {
		enc := json.NewEncoder(reqBody)
		err := enc.Encode(jsonRequest)
		if err != nil {
			log.Printf("Failed to encode json reqBody. %s", err)
			return err
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, resourceURL.String(), reqBody)
	if err != nil {
		log.Printf("NewRequestWithContext failed. %s", err)
		return err
	}

	if api.EnableTracing {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			log.Fatalf("Failed to dump request: %s", err)
		}
		api.trace("Request Dump\n", string(dump))
	}

	req.Header.Add("Authorization", "Bearer "+api.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("%s %s request failed. %s", resource, method, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading body of HTTP response with status %d. %s", resp.StatusCode, err)
		}
		return fmt.Errorf("GET %s failed with status code %d. Body: %s", resource, resp.StatusCode, string(respBody))
	}

	if jsonResult != nil {
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(jsonResult)
		if err != nil {
			log.Printf("Failed to decode json result. %s", err)
			return err
		}
	}

	return nil
}

// Get makes an authenticated GET request to the YouTrack API.
func (api *Api) Get(ctx context.Context, resource string, jsonResult interface{}) error {
	return api.DoRequest(ctx, resource, http.MethodGet, nil, jsonResult)
}

// Post makes an authenticated POST request to the YouTrack API.
func (api *Api) Post(ctx context.Context, resource string, jsonRequest, jsonResult interface{}) error {
	return api.DoRequest(ctx, resource, http.MethodPost, jsonRequest, jsonResult)
}
