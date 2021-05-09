package coin

import (
	"fmt"
	"moon-base/response"
	"unicode/utf8"

	"github.com/pkg/errors"
)

type BuyCoinRequest struct {
	CoinID   int     `json:"coinId" example:"1"`
	UserID   string  `json:"userId" example:"XXX"`
	THBT     float64 `json:"thbt" example:"100"`
	Moon     float64 `json:"moon" example:"2"`
	Slippage float64 `json:"slippage" example:"5"`
}

func (req *BuyCoinRequest) validate() error {
	if req.CoinID < 1 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'coinId' must be REQUIRED field but the input is '%v'.", req.CoinID)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.UserID) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'userId' must be REQUIRED field but the input is '%v'.", req.UserID)), response.ValidateFieldError)
	}
	if req.THBT < 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'thbt' must be REQUIRED field but the input is '%v'.", req.THBT)), response.ValidateFieldError)
	}
	if req.Moon < 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'moon' must be REQUIRED field but the input is '%v'.", req.Moon)), response.ValidateFieldError)
	}
	if req.Slippage < 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'slippage' must be REQUIRED field but the input is '%v'.", req.Slippage)), response.ValidateFieldError)
	}
	return nil
}
