package handler

import "net/http"

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()
	// userDal := dal.NewUserDal(dal.MainDB)
	// userLogic := logic.NewUserLogic(userDal)
	// userHandler := NewUserHandler(userLogic)
	//
	// mux.HandleFunc("POST /registrate", userHandler.RegistrateUser)
	// mux.HandleFunc("POST /login", userHandler.LoginUser)
	// mux.HandleFunc("GET /validate", userHandler.CheckToken)
	mux.HandleFunc("/", NotFoundHandler)

	return mux
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
