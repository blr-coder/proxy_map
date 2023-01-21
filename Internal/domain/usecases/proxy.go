package usecases

import (
	"github.com/google/uuid"
	"net/http"
	"proxy_map/Internal/domain/models"
	"proxy_map/Internal/infrastructure/repository"
	"proxy_map/Internal/infrastructure/repository/map_store"
)

type ProxyUseCase struct {
	httpClient *http.Client
	storage    repository.IProxyRepository
}

func NewProxyUseCase(proxyMap *map_store.ProxyMap) *ProxyUseCase {
	return &ProxyUseCase{
		httpClient: &http.Client{},
		storage:    proxyMap,
	}
}

func (uc *ProxyUseCase) Proxy(proxyRequest *models.ProxyRequest) error {

	// prepare request
	req, err := uc.proxyRequestToHttp(proxyRequest)
	if err != nil {
		return err
	}

	// do request
	resp, err := uc.httpClient.Do(req)
	if err != nil {
		return err
	}

	// TODO: Add event webAPI
	// prepare event
	//event := &models.Event{}

	// save response
	err = uc.storage.Save(proxyRequest, uc.httpResponseToProxy(resp))
	if err != nil {
		//event.Type = saveErrorType
		return err
	}

	//event.Type = saveOkType
	/*go func(e *models.Event) {
		uc.webAPI.event.Push(e)
	}(event)*/

	return nil
}

func (uc *ProxyUseCase) ProxyMap() (map[*models.ProxyRequest]*models.ProxyResponse, error) {
	return uc.storage.All()
}

func (uc *ProxyUseCase) proxyRequestToHttp(pRequest *models.ProxyRequest) (*http.Request, error) {
	httpRequest, err := http.NewRequest(pRequest.Method, pRequest.URL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range pRequest.Headers {
		httpRequest.Header.Add(k, v)
	}

	return httpRequest, nil
}

func (uc *ProxyUseCase) httpResponseToProxy(httpResponse *http.Response) *models.ProxyResponse {
	return &models.ProxyResponse{
		ID:      uuid.New(),
		Status:  httpResponse.Status,
		Headers: httpResponse.Header,
		Length:  0,
	}
}
