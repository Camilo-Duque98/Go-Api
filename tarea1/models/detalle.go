package models

type Detalle struct {
	Id_compra   int `json:"id_compra"`
	Id_producto int `json:"id_producto"`
	Cantidad    int `json:"cantidad"`
}

func (p *Detalle) TableName() string {
	return "detalle"
}
