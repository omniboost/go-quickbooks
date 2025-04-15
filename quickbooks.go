package quickbooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/omniboost/go-quickbooks/sdk"
	"github.com/omniboost/go-quickbooks/sdk/consts"
)

// Quickbooks client type
type Quickbooks struct {
	RealmID      string
	AccessToken  string
	baseURL      string
	minorVersion int

	// http client to use
	http *http.Client
	// Debugging flag
	debug bool
}

// NewClient creates a new client to work with Quickbooks
func NewClient(httpClient *http.Client, realmID string) *Quickbooks {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	q := Quickbooks{}
	q.http = httpClient
	q.RealmID = realmID

	q.SetBaseURL(sdk.ProductionURL)
	return &q
}

// makeGetRequest makes a GET request to Quickbooks API.
// endpoint should start with a leading '/'
func (q *Quickbooks) makeGetRequest(endpoint string) (*http.Response, error) {
	urlStr := q.baseURL + endpoint

	if q.MinorVersion() != 0 {
		u, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}

		v, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			return nil, err
		}
		v.Add("minorversion", strconv.Itoa(q.MinorVersion()))
		u.RawQuery = v.Encode()
		urlStr = u.String()
	}

	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	// headers
	request.Header.Set("accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+q.AccessToken)

	if q.debug == true {
		dump, _ := httputil.DumpRequestOut(request, true)
		log.Println(string(dump))
	}

	response, err := q.http.Do(request)
	if q.debug == true {
		dump, _ := httputil.DumpResponse(response, true)
		log.Println(string(dump))
	}
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, handleError(response)
	}

	return response, nil
}

// PostRequest makes a POST request to Quickbooks API.
// endpoint should start with a leading '/'
func (q *Quickbooks) makePostRequest(endpoint string, body interface{}) (*http.Response, error) {
	urlStr := q.baseURL + endpoint

	if q.MinorVersion() != 0 {
		u, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}

		v, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			return nil, err
		}
		v.Add("minorversion", strconv.Itoa(q.MinorVersion()))
		u.RawQuery = v.Encode()
		urlStr = u.String()
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", urlStr, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	// headers
	request.Header.Set("accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+q.AccessToken)

	if q.debug == true {
		dump, _ := httputil.DumpRequestOut(request, true)
		log.Println(string(dump))
	}

	response, err := q.http.Do(request)
	if q.debug == true && response != nil {
		dump, _ := httputil.DumpResponse(response, true)
		log.Println(string(dump))
	}
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, handleError(response)
	}

	return response, nil
}

func (q *Quickbooks) SetBaseURL(URL string) {
	q.baseURL = URL
}

func (q *Quickbooks) SetDebug(debug bool) {
	q.debug = debug
}

func (q *Quickbooks) SetMinorVersion(minorVersion int) {
	q.minorVersion = minorVersion
}

func (q *Quickbooks) MinorVersion() int {
	return q.minorVersion
}

func handleError(response *http.Response) error {
	switch response.StatusCode {
	case 400:
		// b, errz := ioutil.ReadAll(response.Body)
		// log.Println(string(b))
		// log.Println(errz)
		// os.Exit(2)
		qbError := ErrorObject{}
		err := json.NewDecoder(response.Body).Decode(&qbError)
		if err != nil {
			return err
		}

		return qbError
	case 401:
		sdkError := SDKError{}
		return sdkError.New(consts.QBAuthorizationFault, consts.QBAuthenticationFaultCode, consts.QBAuthorizationFaultMessage)
	}

	return fmt.Errorf("unexpected response code: %s", response.Status)
}
