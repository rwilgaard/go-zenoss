package zenoss

import "net/http"

type Component struct {
    URL  string `json:"url,omitempty"`
    Text string `json:"text,omitempty"`
    UID  string `json:"uid,omitempty"`
    UUID string `json:"uuid,omitempty"`
}

type Event struct {
    ProdState         string                 `json:"prodState,omitempty"`
    Facility          interface{}            `json:"facility,omitempty"`
    EventClassKey     interface{}            `json:"eventClassKey,omitempty"`
    Agent             string                 `json:"agent,omitempty"`
    DedupID           string                 `json:"depupid,omitempty"`
    Location          []map[string]string    `json:"location,omitempty"`
    OwnerID           string                 `json:"ownerid,omitempty"`
    EventClass        map[string]string      `json:"eventClass,omitempty"`
    ID                string                 `json:"id,omitempty"`
    DevicePriority    string                 `json:"devicePriority,omitempty"`
    Monitor           string                 `json:"monitor,omitempty"`
    Priority          interface{}            `json:"priority,omitempty"`
    Details           map[string]interface{} `json:"details,omitempty"`
    DeviceClass       []map[string]string    `json:"DeviceClass,omitempty"`
    EventKey          string                 `json:"eventKey,omitempty"`
    EventClassMapping string                 `json:"eventClassMapping,omitempty"`
    ClearID           interface{}            `json:"clearid,omitempty"`
    DeviceGroups      []map[string]string    `json:"DeviceGroups,omitempty"`
    EventGroup        interface{}            `json:"eventGroup,omitempty"`
    Device            map[string]string      `json:"device,omitempty"`
    Message           string                 `json:"message,omitempty"`
    StateChange       float64                `json:"stateChange,omitempty"`
    IpAddress         []string               `json:"ipAddress,omitempty"`
    Systems           []map[string]string    `json:"Systems,omitempty"`
    Count             int                    `json:"count,omitempty"`
    Severity          int                    `json:"severity,omitempty"`
    EvID              string                 `json:"evid,omitempty"`
    Component         *Component             `json:"component,omitempty"`
    Summary           string                 `json:"summary,omitempty"`
    FirstTime         float64                `json:"firstTime,omitempty"`
    LastTime          float64                `json:"lastTime,omitempty"`
}

type QueryEventsList struct {
    Events     []Event `json:"events"`
    TotalCount int     `json:"totalCount"`
    Success    bool    `json:"success"`
    Asof       float64 `json:"asof"`
}

type QueryEventsResult struct {
    Response
    Result *QueryEventsList `json:"result"`
}

type QueryEventsQuery struct {
    UID     string                 `json:"uid,omitempty"`     // (optional) Context for the query (default: None)
    Start   int                    `json:"start,omitempty"`   // (optional) Min index of events to retrieve (default: 0)
    Limit   int                    `json:"limit,omitempty"`   // (optional) Max index of events to retrieve (default: 0)
    Sort    string                 `json:"sort,omitempty"`    // (optional) Key on which to sort the return results (default: 'lastTime')
    Dir     string                 `json:"dir,omitempty"`     // (optional) Sort order; can be either 'ASC' or 'DESC' (default: 'DESC')
    Params  map[string]interface{} `json:"params,omitempty"`  // (optional) Key-value pair of filters for this search. (default: None)
    History bool                   `json:"history,omitempty"` // (optional) True to search the event history table instead of active events (default: False)
    Keys    []string               `json:"keys,omitempty"`
}

func (a *API) QueryEvents(query QueryEventsQuery) (*QueryEventsResult, *http.Response, error) {
    r := request{
        Action: "EventsRouter",
        Method: "query",
        Data:   []interface{}{query},
        TID:    a.nextTID(),
    }

    path := "/zport/dmd/evconsole_router"
    req, err := a.NewRequest("POST", path, r)
    if err != nil {
        return nil, nil, err
    }

    var res QueryEventsResult
    resp, err := a.Do(req, &res)
    if err != nil {
        return nil, resp, err
    }

    return &res, resp, nil
}
