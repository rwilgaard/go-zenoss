package zenoss

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "sync"
    "time"
)

type API struct {
    endpoint *url.URL
    client   *http.Client
    username string
    password string
    tid      int
    mu       sync.Mutex
}

type request struct {
    Action string        `json:"action"`
    Method string        `json:"method"`
    Data   []interface{} `json:"data"`
    TID    int           `json:"tid"`
}

type Response struct {
    UUID   string      `json:"uuid"`
    Action string      `json:"action"`
    TID    int         `json:"tid"`
    Type   string      `json:"type"`
    Method string      `json:"method"`
}

func NewAPI(endpoint string, username string, password string) (*API, error) {
    if len(endpoint) == 0 {
        return nil, errors.New("url empty")
    }

    u, err := url.ParseRequestURI(endpoint)
    if err != nil {
        return nil, err
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
    }

    a := new(API)
    a.endpoint = u
    a.client = &http.Client{
        Transport: tr,
        Timeout:   10 * time.Second,
    }
    a.username = username
    a.password = password
    a.tid = 0

    return a, nil
}

func (a *API) nextTID() int {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.tid++
    return a.tid
}

func (a *API) NewRequest(method string, path string, body interface{}) (*http.Request, error) {
    u, err := url.ParseRequestURI(a.endpoint.String() + path)
    if err != nil {
        return nil, err
    }

    var buf *bytes.Buffer
    if body != nil {
        js, err := json.Marshal(body)
        if err != nil {
            return nil, err
        }
        buf = bytes.NewBuffer(js)
    }

    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    if (a.username != "") && (a.password != "") {
        req.SetBasicAuth(a.username, a.password)
    }

    return req, nil
}

func (a *API) Do(req *http.Request, i interface{}) (*http.Response, error) {
    resp, err := a.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return resp, fmt.Errorf("%v", resp)
    }

    if i == nil {
        return resp, nil
    }

    if err = json.NewDecoder(resp.Body).Decode(i); err != nil {
        return resp, err
    }

    return resp, nil
}
