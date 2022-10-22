package models

type Address struct{
  State string `json:"state"`
  City string `json:"city"`
  Pincode string `json:"pincode"`
}

type User struct{
  Name string `json:"name"`
  Age int `json:"age"`
  Address Address
}
