package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// HTTPError is an error returned by a failed API call
type HTTPError struct {
	StatusCode int
	RequestURL *url.URL
	Message    string
}

func (err HTTPError) Error() string {
	if msgs := strings.SplitN(err.Message, "\n", 2); len(msgs) > 1 {
		return fmt.Sprintf("HTTP %d: %s (%s)\n%s", err.StatusCode, msgs[0], err.RequestURL, msgs[1])
	} else if err.Message != "" {
		return fmt.Sprintf("HTTP %d: %s (%s)", err.StatusCode, err.Message, err.RequestURL)
	}
	return fmt.Sprintf("HTTP %d (%s)", err.StatusCode, err.RequestURL)
}

// Client allows the HTTP requests
type Client struct {
	http *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		http: httpClient,
	}
}

func (c Client) GetCurrentUser() {

}

// Request does the HTTP Requests
func (c Client) Request(hostname string, method string, p string, body io.Reader, data interface{}) error {

	url := fmt.Sprintf("https://%s/%s", hostname, p)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// TODO MOVE OUTSIDE
	req.Header.Set("zsessionid", os.Getenv("RALLY_API_TOKEN"))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !success {
		return HandleHTTPError(resp)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	return nil
}

func HandleHTTPError(resp *http.Response) error {
	httpError := HTTPError{
		StatusCode: resp.StatusCode,
		RequestURL: resp.Request.URL,
	}

	if !jsonTypeRE.MatchString(resp.Header.Get("Content-Type")) {
		httpError.Message = resp.Status
		return httpError
	}

	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		httpError.Message = err.Error()
		return httpError
	}

	return httpError
}

var jsonTypeRE = regexp.MustCompile(`[/+]json($|;)`)

func inspectableMIMEType(t string) bool {
	return strings.HasPrefix(t, "text/") || jsonTypeRE.MatchString(t)
}
