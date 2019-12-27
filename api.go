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
	// BaseURL is the URL to the REST API endpoint for a YouTrack Project. It should
	// end is a slash. For example: https://goyt.myjetbrains.com/youtrack/api/
	BaseURL *url.URL

	// Token is the permanent token used to make authenticated requests.
	// For more information, see:
	// https://www.jetbrains.com/help/youtrack/incloud/authentication-with-permanent-token.html
	Token string

	// EnableTracing turns on extra logging, including HTTP request/response logging.
	// NOTE that the authorization token will be logged when this is enabled.
	EnableTracing bool
}

func (api *Api) trace(v ...interface{}) {
	if api.EnableTracing {
		log.Println(v...)
	}
}

// DoRequest makes an authenticated HTTP request to the YouTrack API.
// jsonRequest and jsonResult are both optional, and depend on the request being made. A GET request,
// for example, should probably set jsonRequestBody to nil.
func (api *Api) DoRequest(ctx context.Context, resource *url.URL, method string, jsonRequest, jsonResult interface{}) error {
	api.trace(method, resource)

	resourceURL := api.BaseURL.ResolveReference(resource)
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

	req, err := http.NewRequestWithContext(ctx, method, resourceURL.String(), reqBody)
	if err != nil {
		log.Printf("NewRequestWithContext failed. %s", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+api.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	if api.EnableTracing {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			log.Fatalf("Failed to dump request: %s", err)
		}
		api.trace(fmt.Sprintf("Request Dump\n%s", string(dump)))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("%s %s request failed. %s", resource, method, err)
		return err
	}
	defer resp.Body.Close()

	if api.EnableTracing {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatalf("Failed to dump response: %s", err)
		}
		api.trace(fmt.Sprintf("Response Dump:\n%s", string(dump)))
	}

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
func (api *Api) Get(ctx context.Context, resource *url.URL, jsonResult interface{}) error {
	return api.DoRequest(ctx, resource, http.MethodGet, nil, jsonResult)
}

// Post makes an authenticated POST request to the YouTrack API.
func (api *Api) Post(ctx context.Context, resource *url.URL, jsonRequest, jsonResult interface{}) error {
	return api.DoRequest(ctx, resource, http.MethodPost, jsonRequest, jsonResult)
}
