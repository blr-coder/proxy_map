package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (req ProxyRequest) MarshalBinary() ([]byte, error) {
	return json.Marshal(req)
}

type ProxyResponse struct {
	ID      uuid.UUID           `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  uint64              `json:"length"`
}

func (res ProxyResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(res)
}
