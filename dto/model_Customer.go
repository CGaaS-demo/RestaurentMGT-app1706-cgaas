package dto



type Customer struct {
    CustomerId   string ` json:"CustomerId" `
    FirstName   string ` json:"FirstName" `
    LastName   string ` json:"LastName" `
    Email   string ` json:"Email" `
    ProfileImage   string ` json:"ProfileImage" `
    CoverImage   string ` json:"CoverImage" `
    Country   string ` json:"Country" `
    Phone   string ` json:"Phone" `
    Address   string ` json:"Address" `
    City   string ` json:"City" `
    RestaurentId   string ` json:"RestaurentId" `
    Deleted bool `json:"deleted"`}

