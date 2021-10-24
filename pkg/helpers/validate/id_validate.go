package validate

import (
	"net/http"
	"strconv"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"

	log "github.com/sirupsen/logrus"
)

func IdValidate(w http.ResponseWriter, r *http.Request, idAccount string) (id int64) {
	userId := parseform.Pars(w, r, idAccount)

	id, err := strconv.ParseInt(userId, 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Error with parcing id")
		jsonenc.JSONError(w, "Error with parcing id")
		return
	}

	if id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect value id user")
		jsonenc.JSONError(w, "Incorrect value id user")
		return
	}

	return
}
