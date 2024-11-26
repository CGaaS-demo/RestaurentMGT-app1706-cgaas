package dto



type Restaurent struct {
    RestaurentId   string ` json:"RestaurentId" `
    Name   string ` json:"Name" `
    Location   string ` json:"Location" `
    Rating   float64 ` json:"Rating" `
    Cuisine   string ` json:"Cuisine" `
    DeliveryAvailable   bool ` json:"DeliveryAvailable" `
    MenuItems   string ` json:"MenuItems" `
    Menu   []string ` json:"Menu" `
    Deleted bool `json:"deleted"`}

