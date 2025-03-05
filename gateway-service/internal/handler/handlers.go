package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"

	"gateway/internal/service"
	"gateway/pkg/apperrors"
	"gateway/pkg/utils"
)

func HandleAuthService(w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ServeHTTP(w, r)
}

func HandlePostService(w http.ResponseWriter, r *http.Request) {
	authReq, err := service.AuthValidateRequest(r.Header)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	if authReq.StatusCode != http.StatusOK {
		var errMessage map[string]string
		err = json.NewDecoder(authReq.Body).Decode(&errMessage)
		if err != nil {
			utils.ResponseErrorJson(err, w)
			return
		}
		w.WriteHeader(authReq.StatusCode)
		utils.ResponseErrorJson(errors.New(errMessage["error"]), w)
		return
	}

	r.Header.Set("X-User-ID", authReq.Header.Get("X-User-ID"))

	targetURL, _ := url.Parse("http://localhost:8082")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ServeHTTP(w, r)
}

func HandleLikeService(w http.ResponseWriter, r *http.Request) {
	authReq, err := service.AuthValidateRequest(r.Header)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	if authReq.StatusCode != http.StatusOK {
		var errMessage map[string]string
		err = json.NewDecoder(authReq.Body).Decode(&errMessage)
		if err != nil {
			utils.ResponseErrorJson(err, w)
			return
		}
		w.WriteHeader(authReq.StatusCode)
		utils.ResponseErrorJson(errors.New(errMessage["error"]), w)
		return
	}

	r.Header.Set("X-User-ID", authReq.Header.Get("X-User-ID"))

	targetURL, _ := url.Parse("http://localhost:8083")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ServeHTTP(w, r)
}

func HandleCommentService(w http.ResponseWriter, r *http.Request) {
	authReq, err := service.AuthValidateRequest(r.Header)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	if authReq.StatusCode != http.StatusOK {
		var errMessage map[string]string
		err = json.NewDecoder(authReq.Body).Decode(&errMessage)
		if err != nil {
			utils.ResponseErrorJson(err, w)
			return
		}
		w.WriteHeader(authReq.StatusCode)
		utils.ResponseErrorJson(errors.New(errMessage["error"]), w)
		return
	}

	r.Header.Set("X-User-ID", authReq.Header.Get("X-User-ID"))

	targetURL, _ := url.Parse("http://localhost:8084")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ServeHTTP(w, r)
}

func HandleWallService(w http.ResponseWriter, r *http.Request) {
	authReq, err := service.AuthValidateRequest(r.Header)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	if authReq.StatusCode != http.StatusOK {
		var errMessage map[string]string
		err = json.NewDecoder(authReq.Body).Decode(&errMessage)
		if err != nil {
			utils.ResponseErrorJson(err, w)
			return
		}
		w.WriteHeader(authReq.StatusCode)
		utils.ResponseErrorJson(errors.New(errMessage["error"]), w)
		return
	}

	r.Header.Set("X-User-ID", authReq.Header.Get("X-User-ID"))

	targetURL, _ := url.Parse("http://localhost:8085")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ServeHTTP(w, r)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
