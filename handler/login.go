package handler

import (
	"encoding/json"
	"net/http"

	"github.com/saalcazar/first-api-echo/authorization"
	"github.com/saalcazar/first-api-echo/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "Estructura no valida", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "usuario o contraseña no validos", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "no se pudo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
	return

}

func isLoginValid(data *model.Login) bool {
	return data.Email == "info@saalcazar.org" && data.Password == "123456"
}
