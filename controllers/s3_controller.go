package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"s3-microservice/common"
	"s3-microservice/services"
)

type S3Controller struct {
	s3Service services.S3Service
}

func S3ControllerInstance() *S3Controller {
	return &S3Controller{}
}

func (s3Controller *S3Controller) Upload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	eTag, err := s3Controller.s3Service.UploadFile(handler.Filename, file)
	if err != nil {
		apiErr := common.NewAPIError(err.Error(), http.StatusInternalServerError)
		common.WriteErrorResponse(apiErr, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eTag)
}

func (s3Controller *S3Controller) Download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["filename"]

	fileBytes, err := s3Controller.s3Service.DownloadFile(key)
	if err != nil {
		apiErr := common.NewAPIError(err.Error(), http.StatusInternalServerError)
		common.WriteErrorResponse(apiErr, w)
		return
	}

	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, key))
	w.Write(fileBytes)
}
