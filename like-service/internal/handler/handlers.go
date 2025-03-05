package handler

import (
	"fmt"
	"net/http"

	"liker/pkg/apperrors"
	"liker/pkg/utils"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
