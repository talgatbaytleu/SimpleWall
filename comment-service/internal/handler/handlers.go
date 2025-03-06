package handler

import (
	"fmt"
	"net/http"

	"commenter/pkg/apperrors"
	"commenter/pkg/utils"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
