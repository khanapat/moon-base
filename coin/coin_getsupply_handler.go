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

type getSupplyCoin struct {
	GetSupplyFn GetSupplyFn
}

func NewGetSupplyCoin(getSupplyFn GetSupplyFn) http.Handler {
	return &getSupplyCoin{
		GetSupplyFn: getSupplyFn,
	}
}

// Get Coin Supply
// @Summary Get Coin Supply
// @Description Method for getting coin supply.
// @Tags Coin
// @Accept json
// @Produce json
// @Param id query string false "Coin ID"
// @Param coinName query string false "Coin Name"
// @Success 200 {object} response.Response{data=coin.Coin} "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /supply [get]
func (s *getSupplyCoin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl := logz.ContextLogger(r.Context())
	logger := rl.With(
		zap.String(common.Service, common.GetSupplyCoinService),
	)
	defer logz.ExecutionTime(time.Now(), fmt.Sprintf("Start send to %s API", r.RequestURI), logger)

	v := r.URL.Query()

	m := make(map[string]interface{}, 0)
	if v.Get("id") != "" {
		m["id"] = v.Get("id")
	}
	if v.Get("coinName") != "" {
		m["coin_name"] = v.Get("coinName")
	}
	lists, err := s.GetSupplyFn(r.Context(), m)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).InternalDatabase)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := response.ResponseContextLocale(r.Context()).GetSupplyCoinSuccess
	resp.Data = &lists
	json.NewEncoder(w).Encode(&resp)
}
