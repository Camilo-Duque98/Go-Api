package models

type Compra struct {
	Id_compra  int `json:"id_producto" gorm:"primaryKey;auto_increment;not_null"`
	Id_cliente int `json:"id_cliente"`
	//Id_cliente `json:"id_cliente"` //gorm:"foreignKey:Id_cliente"`
}

func (p *Compra) TableName() string {
	return "compra"
}

/*type Compra struct {
	Id_compra int `json:"id_compra" gorm:"primaryKey;auto_increment;not_null"`
	ClienteID int `json:"id_cliente" gorm:"foreignKey"`
}

func (p *Compra) TableName() string {
	return "compra"
}
*/
