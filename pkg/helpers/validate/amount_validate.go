package validate

import (
	"net/http"
	"strconv"
	"strings"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"

	log "github.com/sirupsen/logrus"
)

func AmountValidate(w http.ResponseWriter, r *http.Request, amountS string) (correctAmount float64) {
	amountString := parseform.Pars(w, r, amountS)

	validAmount := strings.Split(amountString, ".")
	if len(validAmount) > 1 {
		if len(validAmount[1]) > 2 {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("The amount have more then 2 decimal places")
			jsonenc.JSONError(w, "The amount have more then 2 decimal places")
			return
		}
	}

	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Error with parcing amount")
		jsonenc.JSONError(w, "Error with parcing amount")
		return
	}

	if amount < 0.01 {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("The amount is negative")
		jsonenc.JSONError(w, "The amount is negative")
		return
	}

	return amount
}
