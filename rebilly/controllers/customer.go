package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/slavcodev/rebilly/models"
)

type CustomerController struct {
    Repository models.Customers
}

func (context *CustomerController) ListCustomers(w http.ResponseWriter, r *http.Request) {
    customers := context.Repository.GetAll();
    json.NewEncoder(w).Encode(customers)
}

func (context *CustomerController) ViewCustomers(w http.ResponseWriter, r *http.Request) {
    if id, ok := mux.Vars(r)["id"]; ok {
        customer := context.Repository.GetById(id);

        if customer == nil {
            http.NotFound(w, r)
        }

        json.NewEncoder(w).Encode(&customer)
    }
}
