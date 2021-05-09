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

type getHistory struct {
	GetHistoryLogsFn GetHistoryLogsFn
}

func NewGetHistory(getHistoryLogsFn GetHistoryLogsFn) http.Handler {
	return &getHistory{
		GetHistoryLogsFn: getHistoryLogsFn,
	}
}

// Get History
// @Summary Get History Transaction
// @Description Method for searching history transaction.
// @Tags Coin
// @Accept json
// @Produce json
// @Param from query string false "History Date From"
// @Param to query string false "History Date To"
// @Success 200 {object} response.Response{data=coin.History} "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /history [get]
func (s *getHistory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl := logz.ContextLogger(r.Context())
	logger := rl.With(
		zap.String(common.Service, common.GetHistoryService),
	)
	defer logz.ExecutionTime(time.Now(), fmt.Sprintf("Start send to %s API", r.RequestURI), logger)

	v := r.URL.Query()

	m := make(map[string]interface{}, 0)
	if v.Get("from") != "" {
		m["history_date_from"] = v.Get("from")
	}
	if v.Get("to") != "" {
		m["history_date_to"] = v.Get("to")
	}
	order := "DESC"
	if v.Get("order") != "" {
		order = v.Get("order")
	}
	lists, err := s.GetHistoryLogsFn(r.Context(), m, order)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.ResponseContextLocale(r.Context()).InternalDatabase)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := response.ResponseContextLocale(r.Context()).QueryHistorySuccess
	resp.Data = &lists
	json.NewEncoder(w).Encode(&resp)
}
