package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/gorilla/mux"
    "github.com/slavcodev/rebilly"
    "github.com/slavcodev/rebilly/controllers"
    "github.com/slavcodev/rebilly/models"
)

func main() {
    ctr := controllers.CustomerController{}
    LoadTestData(&ctr);

    // Create router and configure routes
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index).Name("home")
    router.HandleFunc("/customers", ctr.ListCustomers).Name("customer-list")
    router.HandleFunc("/customers/{id:[a-zA-z0-9]+}", ctr.ViewCustomers).Name("customer")

    // Testing to generate a URL by name and params
    url, _ := router.Get("customer").URL("id", "abc")
    log.Print(url)

    // Start server
    log.Fatal(http.ListenAndServe(":3000", router))
}

// Controller action for root
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

// Function for loading fixtures
func LoadTestData(ctr *controllers.CustomerController) {
    ctr.Repository = make(models.Customers, 100)
    ctr.Repository.Add(rebilly.Schema{
        "ExternalId.Id":  {"A"},
        "Email":  {"user@mail.com"},
    })
    ctr.Repository.Add(rebilly.Schema{
        "ExternalId.Id":  {"B"},
        "Email":  {"user@mail.com"},
    })
    ctr.Repository.Add(&models.Customer{
        ExternalId:  rebilly.NewId("C"),
        Email:  "user@mail.com",
    })
}
