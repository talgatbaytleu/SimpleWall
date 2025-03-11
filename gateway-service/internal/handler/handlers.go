package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"gateway/internal/service"
	"gateway/pkg/apperrors"
	"gateway/pkg/logger"
	"gateway/pkg/utils"
)

var (
	AuthServiceAddr         = os.Getenv("AUTH_SERVICE_ADDR")
	PostServiceAddr         = os.Getenv("POST_SERVICE_ADDR")
	LikeServiceAddr         = os.Getenv("LIKE_SERVICE_ADDR")
	CommentServiceAddr      = os.Getenv("COMMENT_SERVICE_ADDR")
	WallServiceAddr         = os.Getenv("WALL_SERVICE_ADDR")
	S3ServiceAddr           = os.Getenv("S3_SERVICE_ADDR")
	NotificationServiceAddr = os.Getenv("NOTIFICATION_SERVICE_ADDR")
)

func HandleAuthService(w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse(AuthServiceAddr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	logger.LogMessage("HandlerAuthService: " + r.URL.Path)

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

	targetURL, _ := url.Parse(PostServiceAddr)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	logger.LogMessage("HandlerPostService: " + r.URL.Path)

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

	targetURL, _ := url.Parse(LikeServiceAddr)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	logger.LogMessage("HandlerLikeService: " + r.URL.Path)

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

	targetURL, _ := url.Parse(CommentServiceAddr)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	logger.LogMessage("HandlerCommentService: " + r.URL.Path)

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

	targetURL, _ := url.Parse(WallServiceAddr)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	logger.LogMessage("HandlerWallService: " + r.URL.Path)

	proxy.ServeHTTP(w, r)
}

func HandleNotificationService(w http.ResponseWriter, r *http.Request) {
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

	targetURL, _ := url.Parse(NotificationServiceAddr)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	logger.LogMessage("HandlerNotificationService: " + r.URL.Path)

	proxy.ServeHTTP(w, r)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
