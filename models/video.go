package models

type Video struct{
  Title string `json:"title"` // json serlialization tag
  Description string `json:"description"`
  URL string `json:"url"`
}
