package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	sortBy     = "by"
	sortByDesc = "desc"
	data       = "created_at"
)

func GetUserId(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()

	userId := validate.IdValidate(w, r, id)
	if userId < 1 {
		return
	}

	descriptionSlice := instance.GetUserIdDB(ctx, w, userId)
	if descriptionSlice == nil {
		return
	}

	for _, row := range descriptionSlice {
		jsonenc.JSONUGetUserIdDescription(w, row)
	}
}

func GetUserIdSort(w http.ResponseWriter, r *http.Request) {
	sortBy := parseform.Pars(w, r, sortBy)
	if sortBy != amount && sortBy != data && sortBy != nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		jsonenc.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	userId := validate.IdValidate(w, r, id)
	if userId < 1 {
		return
	}

	switch sortBy {
	case data:
		descriptionSlice := instance.GetUserIdSortDB(ctx, w, userId, data)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}

	case amount:
		descriptionSlice := instance.GetUserIdSortDB(ctx, w, userId, amount)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}
	}
}

func GetUserIdSortDesc(w http.ResponseWriter, r *http.Request) {
	sortByDesc := parseform.Pars(w, r, sortByDesc)
	if sortByDesc != amount && sortByDesc != data && sortByDesc != nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		jsonenc.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	userId := validate.IdValidate(w, r, id)
	if userId < 1 {
		return
	}

	switch sortByDesc {
	case data:
		descriptionSlice := instance.GetUserIdSortDB(ctx, w, userId, data)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}

	case amount:
		descriptionSlice := instance.GetUserIdSortDB(ctx, w, userId, amount)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}
	}
}
