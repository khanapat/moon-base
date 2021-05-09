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

type resetMoonCoin struct {
	ResetSupplyFn ResetSupplyFn
}

func NewResetMoonCoin(resetSupplyFn ResetSupplyFn) http.Handler {
	return &resetMoonCoin{
		ResetSupplyFn: resetSupplyFn,
	}
}

// Reset & Setup Supply Coin
// @Summary reset and setup supply coin
// @Description Method for setting up supply coin.
// @Tags Coin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /reset [get]
func (s *resetMoonCoin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl := logz.ContextLogger(r.Context())
	logger := rl.With(
		zap.String(common.Service, common.ResetMoonCoinService),
	)
	defer logz.ExecutionTime(time.Now(), fmt.Sprintf("Start send to %s API", r.RequestURI), logger)

	if err := s.ResetSupplyFn(r.Context()); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).InternalDatabase)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).ResetCoinSuccess)
}
