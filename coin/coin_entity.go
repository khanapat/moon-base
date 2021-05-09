package coin

type Coin struct {
	ID       *int     `json:"id" example:"1"`
	CoinName *string  `json:"coinName" example:"MOON"`
	Supply   *float64 `json:"supply" example:"1000"`
}

type History struct {
	Number   *int     `json:"number" example:"1"`
	DateTime *string  `json:"dateTime" example:"2021-04-01 10:00"`
	UserID   *string  `json:"userId" example:"AAA"`
	THBT     *float64 `json:"thbt" example:"100"`
	Moon     *float64 `json:"moon" example:"0.2"`
	Rate     *string  `json:"rate" example:"1 MOON = 55 THBT | 0.01818181"`
}
