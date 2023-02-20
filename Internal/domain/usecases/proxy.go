package usecases

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"proxy_map/Internal/domain/models"
	"proxy_map/Internal/infrastructure/repository"
	"proxy_map/Internal/infrastructure/repository/map_store"
	"proxy_map/pkg/events"
)

type ProxyUseCase struct {
	httpClient  *http.Client
	storage     repository.IProxyRepository
	eventSender events.EventSender
}

func NewProxyUseCase(proxyMap *map_store.ProxyMap, sender events.EventSender) *ProxyUseCase {
	return &ProxyUseCase{
		httpClient:  &http.Client{},
		storage:     proxyMap,
		eventSender: sender,
	}
}

func (uc *ProxyUseCase) Proxy(proxyRequest *models.ProxyRequest) error {

	ctx := context.TODO()

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

	// TODO: prepare event
	//event := &events.CreateEventRequest{}

	// save response
	err = uc.storage.Save(proxyRequest, uc.httpResponseToProxy(resp))
	if err != nil {
		// TODO: event.Type = saveErrorType
		return err
	}

	err = uc.eventSender.Send(ctx, &events.CreateEventRequest{
		TypeTitle:   "FIRST",
		CampaignID:  1,
		InsertionID: 2,
		UserID:      3,
		Cost: &events.Cost{
			Amount:   100500,
			Currency: "EUR",
		},
	})
	if err != nil {
		return err
	}

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
