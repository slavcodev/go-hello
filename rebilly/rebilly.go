package rebilly

// Register alias for type used for cast map with values to struct
type Schema map[string][]string;

type ExternalId struct {
    Id  string  `json:"id"`
}

func NewId(id string) ExternalId {
    return ExternalId {Id: id}
}

// Init function runs before main
// func init() {}
