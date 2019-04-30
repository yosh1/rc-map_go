package models

// model definition
type RailroadCrossing struct {
	Name      string `json:"name" binding:"required"`
	TelNumber int    `json:"tel_number" binding:"required"`
	Mail      string `json:"mail" binding:"required"`
	SaleName  string `json:"sale_name" binding:"required"`
	Type      int    `json:"type" binding:"required"`
	IsSuccess int    `json:"is_success" binding:"required"`
	Date      string `json:"date" binding:"required"`
}

type RailroadCrossingRepository struct {
}

func NewRailroadCrossingRepository() RailroadCrossingRepository {
	return RailroadCrossingRepository{}
}
