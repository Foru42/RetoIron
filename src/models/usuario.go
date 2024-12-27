package models

type Usuario struct {
    //Modelo del usuario, con id, name, surname, email
    ID     int    `json:"id"`
    Name string `json:"name"`
    Surname string `json:"surname"`
    Email  string `json:"email"`
}
