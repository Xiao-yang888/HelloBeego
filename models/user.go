package models

type user struct {
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
    Password string `json:"password"`
}
