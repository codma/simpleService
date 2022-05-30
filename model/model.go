package model

//모델 정의

type Store struct {
	StoreId   int    `json:"storeId"`
	StoreName string `json:"storeName"`
	PlanCode  string `json:"planCode"`
	Domain    string `json:"domain"`
	Activate  string `json:"activate"`
}

type AddOneStore struct {
	StoreName string `json:"storeName"`
	PlanCode  string `json:"planCode"`
	Domain    string `json:"domain"`
}
