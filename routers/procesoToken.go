package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/monorozo/gotwettero/bd"
	"github.com/monorozo/gotwettero/models"
)

/*Email es el email devuelto el cual se usara en todos los endpoint*/
var Email string

/*IDUsuario es el ID devuelto el cual se usara en todos los endpoint*/
var IDUsuario string

/*ProcesoToken funcion que extrae sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersDesarrollo_grupodeFacebook")
	claims := &models.Claim{}
	/*Split separa en dos el dato recibido*/
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato del tocken invalido")
	}
	/*trimpace quita los espacion*/
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}
	return claims, false, string(""), err
}
