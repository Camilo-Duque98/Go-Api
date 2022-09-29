package models

type Compra struct {
	Id_compra  int       `json:"id_producto" gorm:"primaryKey;auto_increment;not_null"`
	Id_cliente int       `json:"id_cliente"`
	Detalles   []Detalle `json:"productos"`
}

func (p *Compra) TableName() string {
	return "compra"
}
