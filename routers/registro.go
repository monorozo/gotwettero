package routers

import (
	"encoding/json"
	"net/http"

	"github.com/monorozo/gotwettero/bd"
	"github.com/monorozo/gotwettero/models"
)

/*registro es la funcion para crear el registro en la base de datos*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	/*Body es un stream, significa que solo se puede leer una vez y luego se destruye*/
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es un campo obligatorio", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener minimo 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "el email ingresado ya se encuentra registrado en el sistema", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al realizar el registro del usuario ingresado"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo realizar el registro del usuario ingrsado", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
