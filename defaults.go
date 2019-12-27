package youtrack

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
)

var NotFoundError = errors.New("Not Found")

var baseURLFlag *string = flag.String("yt-url", "", "YouTrack API base URL. Should end with api/.")
var permanentTokenFlag *string = flag.String("yt-token", "", "YouTrack permanent token")
var enableTracingFlag *bool = flag.Bool("yt-trace", false, "Enable YouTrack API log tracing")

// NewDefaultApi returns a new API using a baseURL and permanent authorization token
// from either flags or environment.
func NewDefaultApi() (*Api, error) {
	baseURLString, err := defaultString(baseURLFlag, "YOUTRACK_URL")
	if err != nil {
		log.Print("Could not determine default YouTrack URL.")
		return nil, err
	}
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		log.Printf("Could not parse base URL %s. Error: %s", baseURLString, err)
		return nil, err
	}

	token, err := defaultString(permanentTokenFlag, "YOUTRACK_TOKEN")
	if err != nil {
		log.Print("Could not determine default YouTrack Token.")
		return nil, err
	}

	enableTracing := false
	if enableTracingFlag != nil {
		enableTracing = *enableTracingFlag
	}

	return &Api{
		BaseURL:       baseURL,
		Token:         token,
		EnableTracing: enableTracing,
	}, nil
}

func defaultString(flagValue *string, environmentalVariableName string) (string, error) {
	var val string

	if flagValue != nil {
		val = *flagValue
	}

	if val == "" {
		val = os.Getenv(environmentalVariableName)
	}

	if val == "" {
		return "", NotFoundError
	}

	return val, nil
}
