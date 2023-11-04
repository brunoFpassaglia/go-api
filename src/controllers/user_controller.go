package controllers

import "net/http"

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("criando usuarios"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios"))
}
func ShowUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("show usuarios"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("update usuarios"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("delete usuarios"))
}
