package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Geting all users!"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
