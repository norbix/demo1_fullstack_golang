package handlers

import (
	"net/http"

	"github.com/norbix/demo1_fullstack_golang/backend/internal/services"
)

type AccountHandler interface {
	CreateAccount(http.ResponseWriter, *http.Request)
	GetAccounts(http.ResponseWriter, *http.Request)
}

type accountHandlerImpl struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) AccountHandler {
	return accountHandlerImpl{accountService: accountService}
}
