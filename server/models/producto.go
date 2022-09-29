package models

type Producto struct {
	Id_producto         int    `json:"id_producto" gorm:"primaryKey;auto_increment;not_null"`
	Nombre              string `json:"nombre"`
	Cantidad_disponible int    `json:"cantidad_disponible"`
	Precio_unitario     int    `json:"precio_unitario"`
}

func (p *Producto) TableName() string {
	return "producto"
}