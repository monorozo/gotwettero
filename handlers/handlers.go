package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/monorozo/gotwettero/middlew"
	"github.com/monorozo/gotwettero/routers"

	/*Captura el http, manejo al responde y request, valida si en el body hay informacion y en el header hay informacion ,
	envia respuesta al navegador, status, si se creo registro, logueo, tocken- envia y recibe informacion*/
	"github.com/rs/cors" /*permisos que se le otorgan al API*/ /*cors permite filtrar los accesos*/
)

func Manejadore() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	/*handler es un objeto que filtra y valida los permisos */
	/*Manejadores seteo mi puerto el handlrer y coloco a escuchar al servidor*/
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
