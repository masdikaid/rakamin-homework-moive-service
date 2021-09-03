package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/masdikaid-rakamin-homework-movie-service/common/responder"
	"github.com/masdikaid-rakamin-homework-movie-service/service"
)

func CreateMovie() http.HandlerFunc {
	var m service.Movie
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			fmt.Println(m)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		err = m.Create()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, m, nil)
		return
	}
}
