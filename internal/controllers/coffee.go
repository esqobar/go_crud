package controllers

import (
	"coffee_app_crud/internal/helpers"
	"coffee_app_crud/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var coffee services.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee, err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffee": coffee})
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffee": coffeeCreated})
}

func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedCoffee, err := coffee.UpdateCoffee(id, coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffee": updatedCoffee})
}

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffee(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"status": "deleted"})
}
