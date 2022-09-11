package models

type Cliente struct {
	Id_cliente int `json:"id_cliente" gorm:"primaryKey;auto_increment;not_null"`
	Nombre string `json:"nombre"`
	Contrasena string `json:"contrasena"`
}

func (p *Cliente) TableName() string {
    return "cliente"
}