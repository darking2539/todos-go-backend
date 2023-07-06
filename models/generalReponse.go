package models

type GeneralSucessResp struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Detail     string `json:"detail"`
}
