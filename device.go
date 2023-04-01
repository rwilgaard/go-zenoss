package gozenoss

import (
    "net/http"
)

type Device struct {
    Name            string                       `json:"name,omitempty"`
    UID             string                       `json:"uid,omitempty"`
    ProductionState int                          `json:"productionState,omitempty"`
    Systems         []map[string]string          `json:"systems,omitempty"`
    Groups          []map[string]string          `json:"groups,omitempty"`
    Collector       string                       `json:"collector,omitempty"`
    IpAddress       string                       `json:"ipAddress,omitempty"`
    IpAddressString string                       `json:"ipAddressString,omitempty"`
    SerialNumber    string                       `json:"serialNumber,omitempty"`
    PythonClass     string                       `json:"pythonClass,omitempty"`
    HwManufacturer  map[string]string            `json:"hwManufacturer,omitempty"`
    OsModel         map[string]string            `json:"osmodel,omitempty"`
    Priority        int                          `json:"priority,omitempty"`
    HwModel         map[string]string            `json:"hwModel,omitempty"`
    TagNumber       string                       `json:"tagNumber,omitempty"`
    OsManufacturer  map[string]string            `json:"osManufacturer,omitempty"`
    Location        string                       `json:"location,omitempty"`
    Events          map[string]map[string]string `json:"events,omitempty"`
}

type GetDevicesResult struct {
    TotalCount int      `json:"totalCount"`
    Hash       string   `json:"hash"`
    Success    bool     `json:"success"`
    Devices    []Device `json:"devices"`
}

type GetDevicesQuery struct {
    UID    string            `json:"uid,omitempty"`    // Unique identifier of the organizer to get devices from
    Start  int               `json:"start,omitempty"`  // (optional) Offset to return the results from; used in pagination (default: 0)
    Params map[string]string `json:"params,omitempty"` // (optional) Key-value pair of filters for this search. Can be one of the following: name, ipAddress, deviceClass, or productionState (default: None)
    Keys   []string          `json:"keys,omitempty"`
    Limit  int               `json:"limit,omitempty"` // (optional) Number of items to return; used in pagination (default: 50)
    Sort   string            `json:"sort,omitempty"`  // (optional) Number of items to return; used in pagination (default: 50)
    Dir    string            `json:"dir,omitempty"`   // (optional) Sort order; can be either 'ASC' or 'DESC' (default: 'ASC')
}

func (a *API) GetDevices(query GetDevicesQuery) (*Response, *http.Response, error) {
    r := request{
        Action: "DeviceRouter",
        Method: "getDevices",
        Data:   []interface{}{query},
        Tid:    a.nextTid(),
    }

    path := "/zport/dmd/device_router"
    req, err := a.NewRequest("POST", path, r)
    if err != nil {
        return nil, nil, err
    }

    res := Response{Result: GetDevicesResult{}}
    resp, err := a.Do(req, &res)
    if err != nil {
        return nil, resp, err
    }

    return &res, resp, nil
}
