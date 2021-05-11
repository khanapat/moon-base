package response

import (
	"context"
	"moon-base/common"
)

var (
	EN = Global{
		AuthenBasicWeb:       ErrResponse{Code: ErrBasicAuthenticationCode, Title: ErrBasicAuthenticationMessageEN, Description: ErrAuthenticationDescEN},
		ResetCoinSuccess:     Response{Code: SuccessCode, Title: SuccessResetCoinMessageEN},
		GetSupplyCoinSuccess: Response{Code: SuccessCode, Title: SuccessGetSupplyCoinMessageEN},
		QueryHistorySuccess:  Response{Code: SuccessCode, Title: SuccessQueryHistoryMessageEN},
		BuyCoinSuccess:       Response{Code: SuccessCode, Title: SuccessBuyCoinMessageEN},
		BuyCoinValidateReq:   ErrResponse{Code: ErrInvalidRequestCode, Title: ErrBuyCoinMessageEN, Description: ErrContactAdminDescEN},
		InternalDatabase:     ErrResponse{Code: ErrDatabaseCode, Title: ErrInternalServerMessageEN, Description: ErrContactAdminDescEN},
	}

	TH = Global{
		AuthenBasicWeb:       ErrResponse{Code: ErrBasicAuthenticationCode, Title: ErrBasicAuthenticationMessageTH, Description: ErrAuthenticationDescTH},
		ResetCoinSuccess:     Response{Code: SuccessCode, Title: SuccessResetCoinMessageTH},
		GetSupplyCoinSuccess: Response{Code: SuccessCode, Title: SuccessGetSupplyCoinMessageTH},
		QueryHistorySuccess:  Response{Code: SuccessCode, Title: SuccessQueryHistoryMessageTH},
		BuyCoinSuccess:       Response{Code: SuccessCode, Title: SuccessBuyCoinMessageTH},
		BuyCoinValidateReq:   ErrResponse{Code: ErrInvalidRequestCode, Title: ErrBuyCoinMessageTH, Description: ErrContactAdminDescTH},
		InternalDatabase:     ErrResponse{Code: ErrDatabaseCode, Title: ErrInternalServerMessageTH, Description: ErrContactAdminDescTH},
	}

	Language = map[interface{}]Global{
		"en": EN,
		"th": TH,
	}
)

type Global struct {
	AuthenBasicWeb       ErrResponse
	ResetCoinSuccess     Response
	GetSupplyCoinSuccess Response
	QueryHistorySuccess  Response
	BuyCoinSuccess       Response
	BuyCoinValidateReq   ErrResponse
	InternalDatabase     ErrResponse
}

func ResponseContextLocale(ctx context.Context) *Global {
	v := ctx.Value(common.LocaleKey)
	if v == nil {
		return nil
	}
	l, ok := Language[v]
	if ok {
		return &l
	}
	return &EN
}
