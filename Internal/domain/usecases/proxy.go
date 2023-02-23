package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"proxy_map/Internal/domain/models"
	"proxy_map/Internal/infrastructure/repository"
	"proxy_map/Internal/infrastructure/repository/map_store"
	"proxy_map/pkg/events"
	"proxy_map/pkg/kafka"
)

type ProxyUseCase struct {
	httpClient *http.Client
	storage    repository.IProxyRepository
	conn       *kafka.Conn
}

func NewProxyUseCase(proxyMap *map_store.ProxyMap, conn *kafka.Conn) *ProxyUseCase {
	return &ProxyUseCase{
		httpClient: &http.Client{},
		storage:    proxyMap,
		conn:       conn,
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

	// TODO: prepare event
	//event := &events.CreateEventRequest{}

	// save response
	err = uc.storage.Save(proxyRequest, uc.httpResponseToProxy(resp))
	if err != nil {
		// TODO: event.Type = saveErrorType
		return err
	}

	byteEvent, err := json.Marshal(&events.CreateEventRequest{
		TypeTitle:   "UPDATED",
		CampaignID:  100000,
		InsertionID: 200000,
		UserID:      999,
		Cost: &events.Cost{
			Amount:   77777,
			Currency: "EUR",
		},
	})

	err = uc.conn.Send(kafka.Message{
		Topic: "quickstart",
		Key:   "event",
		Bytes: byteEvent,
	})
	if err != nil {
		return err
	}

	fmt.Println("OK")

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
