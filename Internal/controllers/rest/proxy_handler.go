package rest

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"proxy_map/Internal/domain/models"
	"proxy_map/Internal/domain/usecases"
)

type ProxyHandler struct {
	useCase usecases.IProxyUseCase
}

func NewProxyHandler(uc *usecases.ProxyUseCase) *ProxyHandler {
	return &ProxyHandler{
		useCase: uc,
	}
}

func (h *ProxyHandler) Do(w http.ResponseWriter, r *http.Request) {

	proxyRequest := &models.ProxyRequest{}

	err := json.NewDecoder(r.Body).Decode(&proxyRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.useCase.Proxy(proxyRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *ProxyHandler) List(w http.ResponseWriter, r *http.Request) {
	m, err := h.useCase.ProxyMap()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	spew.Dump(m)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(proxyMapToDataResponse(m))
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

func proxyMapToDataResponse(m map[*models.ProxyRequest]*models.ProxyResponse) []*ResponseData {
	var data []*ResponseData
	for request, response := range m {
		data = append(data, &ResponseData{
			Req: request,
			Res: response,
		})
	}

	return data
}
