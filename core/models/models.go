package models

import "github.com/google/uuid"

type ProxyRequest struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	//Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      uuid.UUID `json:"id"`
	Status  string    `json:"status"`
	Headers []string  `json:"headers"`
	Length  uint64    `json:"length"`
}
