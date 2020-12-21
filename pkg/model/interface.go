package model

type OrderCtx interface {
	ID_() int
	OrderNo_() string
	UserName_() string
	Amount_() float64
	Status_() string
	FileUrl_() string
}
