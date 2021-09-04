package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masdikaid-rakamin-homework-movie-service/common/responder"
	"github.com/masdikaid-rakamin-homework-movie-service/service"
)

func CreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m service.Movie
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

		responder.NewHttpResponse(r, w, http.StatusCreated, m.Contract(), nil)
	}
}

func GetMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m *service.Movie
		var err error

		params := mux.Vars((r))
		m, err = service.GetMovie(params["slug"])
		if err != nil {
			fmt.Println(m)
			responder.NewHttpResponse(r, w, http.StatusNotFound, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, m.Contract(), nil)
	}
}

func UpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m *service.Movie
		var err error

		req := service.Movie{}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println(m)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		params := mux.Vars((r))
		m, err = service.GetMovie(params["slug"])
		if err != nil {
			fmt.Println(m)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		var newData *service.Movie
		newData, err = m.UpdateFromData(req)
		if err != nil {
			fmt.Println(m)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		responder.NewHttpResponse(r, w, http.StatusOK, newData.Contract(), nil)
	}
}

func DeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m *service.Movie
		var err error

		params := mux.Vars((r))
		m, err = service.GetMovie(params["slug"])
		if err != nil {
			fmt.Println(m)
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		err = m.Delete()
		if m.Delete() != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
		}

		responder.NewHttpResponse(r, w, http.StatusOK, "success", nil)
	}
}
