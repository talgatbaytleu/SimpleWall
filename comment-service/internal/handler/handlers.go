package handler

import (
	"fmt"
	"net/http"

	"poster/pkg/apperrors"
	"poster/pkg/utils"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
