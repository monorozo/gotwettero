package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/monorozo/gotwettero/bd"
	"github.com/monorozo/gotwettero/jwt"
	"github.com/monorozo/gotwettero/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Por favor ingrese el usuario"+err.Error(), 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWt(documento)
	if err != nil {
		http.Error(w, "Error al generar el tocken"+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*como se graba una cookie desde el backend*/
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "jwtKey",
		Expires: expirationTime,
	})
}
