package routers

import (
	"encoding/json"
	"net/http"

	"github.com/oscaralcalde/twitter/bd"
	"github.com/oscaralcalde/twitter/models"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(request.Body).Decode(&usuario)
	if err != nil {
		http.Error(writer, "Hubo un error con los datos recibidos: "+err.Error(), 400)
		return
	}

	if emailValidation(usuario.Email, writer) {
		return
	}

	if passwordValidation(usuario.Password, writer) {
		return
	}

	if passwordValidation(usuario.Password, writer) {
		return
	}

	if statusValidation(usuario, writer) {
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func emailValidation(email string, writer http.ResponseWriter) bool {
	if len(email) == 0 {
		http.Error(writer, "El email no puede estar vacio", 400)
		return true
	}
	return false
}

func passwordValidation(password string, writer http.ResponseWriter) bool {
	if len(password) <= 6 {
		http.Error(writer, "El password debe tener mas de 6 caracteres", 400)
		return true
	}
	return false
}

func repeatEmailValidation(email string, writer http.ResponseWriter) bool {
	_, existentEmail, _ := bd.RepeatEmailCheck(email)
	if existentEmail {
		http.Error(writer, "El email ya esta siendo usado", 400)
		return true
	}
	return false
}

func statusValidation(usuario models.Usuario, writer http.ResponseWriter) bool {
	_, status, err := bd.RegisterInsert(usuario)
	if err != nil {
		http.Error(writer, "Error al intentar de insertar el registro de usuario"+err.Error(), 400)
		return true
	}

	if status == false {
		http.Error(writer, "No se ha logrado instertar el registro del usuario", 400)
		return true
	}
	return false
}
