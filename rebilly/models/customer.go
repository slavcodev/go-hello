package models

import (
    "time"
    "log"
    "github.com/gorilla/schema"
    "github.com/slavcodev/rebilly"
    "errors"
)

// ************************************
// Init model type and interface
// ************************************

type Customer struct {
    rebilly.ExternalId
    // Id          int         `json:"-"`
    Email       string      `json:"email"`
    CreateAt    time.Time   `json:"createAt"`
    CardsId     []int       `json:"-"`
}

// ************************************
// Init repository type and interface
// ************************************

type Customers map[string]Customer

// Method for add new customer to repository
// Usage `interface{}` meaning any type here
func (customers Customers) Add(customer interface{}) {
    // Check param type and cast to desired
    switch customer.(type) {
        case rebilly.Schema:
            customerObj := Customer{}
            // Cast `customer` var to type `rebilly.Schema`
            values, ok := customer.(rebilly.Schema)
            if ok {
                decoder := schema.NewDecoder();
                err := decoder.Decode(&customerObj, values)
                if err != nil {
                    log.Fatal(err)
                }
                customers[customerObj.Id] = customerObj
            }
            break
        case Customer:
            if values, ok := customer.(Customer); ok {
                customers[values.Id] = values
            }
            break
        case *Customer:
            if values, ok := customer.(*Customer); ok {
                customers[values.Id] = *values
            }
            break
        default:
            log.Fatal(errors.New("Invalid customer"))
    }
}

// Method for get repository items
func (customers *Customers) GetAll() ([]Customer) {
    v := make([]Customer, 0, len(*customers))

    for  _, value := range *customers {
        v = append(v, value)
    }

    return v
}

// Method for get customer by ExternalId
func (customers Customers) GetById(id string) (*Customer) {
    if customer, ok := customers[id]; ok {
        return &customer
    }
    return nil
}
