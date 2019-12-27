package youtrack

import (
	"errors"
	"flag"
	"net/url"
	"os"
)

var NotFoundError = errors.New("Not Found")

func defaultValue(flagValue *string, environmentalVariableName string) (string, error) {
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

var permanentToken *string = flag.String("yt-token", "", "YouTrack permanent token")

func defaultToken() (string, error) {
	return defaultValue(permanentToken, "YOUTRACK_TOKEN")
}

var baseURL *string = flag.String("yt-url", "", "YouTrack API base URL")

func defaultURL() (*url.URL, error) {
	v, err := defaultValue(baseURL, "YOUTRACK_URL")
	if err != nil {
		return nil, err
	}

	return url.Parse(v)
}
