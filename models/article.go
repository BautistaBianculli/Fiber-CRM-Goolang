package models

type Article struct {
	ID          string  `json:"id"`
	Codart      string  `json:"codart"`
	Descripcion string  `json:"descripcion"`
	PrecioLista float64 `json:"precio_lista"`
	Bonif1      float64 `json:"bonif1"`
	Bonif2      float64 `json:"bonif2"`
	Prventa     float64 `json:"prventa"`
}
