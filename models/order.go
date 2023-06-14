package models

type Order struct{
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Sum int `json:"sum"`
	SumCount int `json:"sum_count"`
	Status string `json:"status"`	
}

