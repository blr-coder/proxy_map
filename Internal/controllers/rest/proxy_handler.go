package rest

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"net/http"
	"proxy_map/Internal/domain/models"
	"proxy_map/Internal/infrastructure/repository/map_store"
)

type ProxyHandler struct {
	httpClient *http.Client
	storage    *map_store.ProxyMap
}

func NewProxyHandler(storage *map_store.ProxyMap) *ProxyHandler {
	return &ProxyHandler{
		httpClient: &http.Client{},
		storage:    storage,
	}
}

func (h *ProxyHandler) Do(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("REQUEST>>> \n %+v\n", r)

	proxyRequest := &models.ProxyRequest{}

	// TODO #1: Move Decode, NewRequest and adding headers to models - like "func (pr *ProxyRequest) PrepareToSending() ((*http.Request, error))"

	err := json.NewDecoder(r.Body).Decode(&proxyRequest)
	if err != nil {
		// TODO: Maybe can create better error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(">>>>>>>KEY")
	fmt.Printf("BODY: %+v\n", proxyRequest)
	fmt.Printf("PROXY_URL: %+v\n", proxyRequest.URL)
	fmt.Printf("PROXY_HEADERS: %+v\n", proxyRequest.Headers)
	fmt.Println(">>>>>>>KEY")

	//req, err := http.NewRequest(http.MethodGet, "https://www.axxonsoft.com/", nil)
	req, err := http.NewRequest(proxyRequest.Method, proxyRequest.URL, nil)
	if err != nil {
		// TODO: Maybe can create better error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for k, v := range proxyRequest.Headers {
		req.Header.Add(k, v)
		fmt.Printf("HEADER: %s, VALUE: %s\n", k, v)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		// TODO: Maybe can create better error
		http.Error(w, err.Error(), resp.StatusCode)
		return
	}

	fmt.Printf("RESP>>> \n %+v\n", resp)
	fmt.Println(">>>>>>>>VAL")
	fmt.Println("RESP_STATUS:", resp.Status)
	fmt.Println("RESP_HEADER:", resp.Header)
	fmt.Println(">>>>>>>>VAL")

	proxyResponse := &models.ProxyResponse{
		ID:      uuid.New(),
		Status:  resp.Status,
		Headers: resp.Header,
		Length:  0,
	}

	go func(request *models.ProxyRequest, response *models.ProxyResponse) {
		// TODO #2 Use UseCase
		err := h.storage.Save(request, response)
		if err != nil {
			fmt.Println("EEEERRRRR:", err)
			return
		}
	}(proxyRequest, proxyResponse)

	w.WriteHeader(http.StatusOK)
}

func (h *ProxyHandler) List(w http.ResponseWriter, r *http.Request) {
	m, err := h.storage.All()
	if err != nil {
		// TODO: Maybe can create better error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("MAP:", m)

	/*mreq := models.ProxyRequest{
		Method: "TEST",
		URL:    "TEST",
	}

	mres := models.ProxyResponse{
		ID:     uuid.New(),
		Status: "OK",
	}

	newMap := make(map[models.ProxyRequest]models.ProxyResponse)
	newMap[mreq] = mres*/

	var list []*ResponseData
	for request, response := range m {
		list = append(list, &ResponseData{
			Req: request,
			Res: response,
		})
	}

	spew.Dump(list)

	//data, err := json.Marshal(list)
	/*data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		// TODO: Maybe can create better error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}

type ResponseData struct {
	Req *models.ProxyRequest  `json:"req"`
	Res *models.ProxyResponse `json:"res"`
}
