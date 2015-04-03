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

type Customers []Customer

// Method for add new customer to repository
func (customers *Customers) Add(customer interface{}) (*Customers) {
    switch customer.(type) {
        case rebilly.Schema:
            customerObj := Customer{}
            values, ok := customer.(rebilly.Schema)
            if ok {
                decoder := schema.NewDecoder();
                err := decoder.Decode(&customerObj, values)
                if err != nil {
                    log.Fatal(err)
                }
                *customers = append(*customers, customerObj)
            }
            break
        case Customer:
            if values, ok := customer.(Customer); ok {
                *customers = append(*customers, values)
            }
            break
        case *Customer:
            if values, ok := customer.(*Customer); ok {
                *customers = append(*customers, *values)
            }
            break
        default:
            log.Fatal(errors.New("Invalid customer"))
    }

    return customers
}

// Method for get repository items
func (customers *Customers) GetAll() (Customers) {
    return *customers
}

// Method for get customer by ExternalId
func (customers *Customers) GetById(id string) (*Customer) {
    for _, customer := range *customers {
        if customer.ExternalId.Id == id {
            return &customer
        }
    }
    return nil
}
