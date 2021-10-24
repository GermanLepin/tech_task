package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func WritingOff(w http.ResponseWriter, r *http.Request) {
	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	amount := validate.AmountValidate(w, r, amount)
	if amount < 0.01 {
		return
	}

	userId, balance := instance.BalanceInfoDB(ctx, w, id)
	if userId == 0 {
		return
	}

	if amount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, id, amount)
	jsonenc.JSONWritingOff(w, id, amount)
}
