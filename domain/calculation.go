package domain

type UserLevel int

type CalcService interface {
	Calculate(int, int, int) (int, int, error)
}

type CalcStrategy interface {
	Calculate(int, int, int) (int, int)
}

type CalcResponse struct {
	Coin   int `json:"coin"`
	Points int `json:"points"`
}
