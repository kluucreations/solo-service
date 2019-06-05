package solo

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kluucreations/solo-service/gen/restapi/operations"
)

func Configure(api *operations.SoloServiceAPI, service Service) {
	api.GetV1LoanLoanIDPaymentsHandler = operations.GetV1LoanLoanIDPaymentsHandlerFunc(func(p operations.GetV1LoanLoanIDPaymentsParams) middleware.Responder {
		return middleware.NotImplemented("GetV1LoanLoanIDPaymentsHandler has not yet been implemented")
	})
	api.GetV1LoansHandler = operations.GetV1LoansHandlerFunc(func(p operations.GetV1LoansParams) middleware.Responder {
		return middleware.NotImplemented("GetV1LoansHandler has not yet been implemented")
	})
}
