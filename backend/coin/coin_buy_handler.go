package coin

import (
	"encoding/json"
	"fmt"
	"moon-base/common"
	"moon-base/logz"
	"moon-base/response"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type buyCoin struct {
	GetSupplyByIDFn     GetSupplyByIDFn
	UpdateSupplyByIDFn  UpdateSupplyByIDFn
	CreateHistoryLogsFn CreateHistoryLogsFn
}

func NewBuyCoin(getSupplyByIDFn GetSupplyByIDFn, updateSupplyByIDFn UpdateSupplyByIDFn, createHistoryLogsFn CreateHistoryLogsFn) http.Handler {
	return &buyCoin{
		GetSupplyByIDFn:     getSupplyByIDFn,
		UpdateSupplyByIDFn:  updateSupplyByIDFn,
		CreateHistoryLogsFn: createHistoryLogsFn,
	}
}

// Buy Coin
// @Summary Buy Coin
// @Description Method for buying coin.
// @Tags Coin
// @Accept json
// @Produce json
// @Param BuyCoinRequest body coin.BuyCoinRequest true "object body to create coin"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /buy [post]
func (s *buyCoin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl := logz.ContextLogger(r.Context())
	logger := rl.With(
		zap.String(common.Service, common.BuyCoinService),
	)
	defer logz.ExecutionTime(time.Now(), fmt.Sprintf("Start send to %s API", r.RequestURI), logger)

	var req BuyCoinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).BuyCoinValidateReq)
		return
	}
	defer r.Body.Close()

	if err := req.validate(); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).BuyCoinValidateReq)
		return
	}

	supply, err := s.GetSupplyByIDFn(r.Context(), req.CoinID)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).InternalDatabase)
		return
	}

	if err := s.UpdateSupplyByIDFn(r.Context(), req.CoinID, *supply-req.Moon); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).InternalDatabase)
		return
	}

	if err := s.CreateHistoryLogsFn(r.Context(), time.Now().Format(common.DateTimeFormat), req.UserID, fmt.Sprintf("1 MOON = %.0f THBT | %f", req.THBT/req.Moon, req.Moon), req.THBT, req.Moon); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).InternalDatabase)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).BuyCoinSuccess)
}
