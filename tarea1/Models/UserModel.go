package Models

type Cliente struct {
	id_cliente      uint   `json:"id_cliente"`
	nombre    string `json:"nombre"`
	contrasena   string `json:"contrasena"`
}

func (b *Cliente) TableName() string {
	return "cliente"
}