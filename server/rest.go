package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	//"log"

	"github.com/julienschmidt/httprouter"
)

const (
	contentTypeHeader     = "Content-Type"
	applicationJSONHeader = "application/json"
)

func APIStartJob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	startJobReq := APIStartJobReq{}

	w.Header().Set(contentTypeHeader, applicationJSONHeader)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(body, &startJobReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIError{Error: err.Error()})
		return
	}

	jobID, err := StartJobOnWorker(startJobReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIError{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIStartJobRes{JobID: jobID})
}

func createRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/start", APIStartJob)
	// router.POST("/stop", apiStopJob)
	// router.POST("/query", apiQueryJob)

	return router
}

func API() {
	srv := &http.Server{
		Addr:    config.HTTPServer.Addr,
		Handler: createRouter(),
	}

	log.Println("HTTP Server listening on", config.HTTPServer.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
