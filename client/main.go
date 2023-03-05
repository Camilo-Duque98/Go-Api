package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Productos struct {
	Data []Producto `json:"data"`
}
type Productos2 struct {
	Data Producto `json:"data"`
}
type Producto struct {
	Id_producto         int    `json:"id_producto"`
	Nombre              string `json:"nombre"`
	Cantidad_disponible int    `json:"cantidad_disponible"`
	Precio_unitario     int    `json:"precio_unitario"`
}

type CreateProduct struct {
	Nombre   string `json:"nombre"`
	Cantidad int    `json:"cantidad_disponible"`
	Precio   int    `json:"precio_unitario"`
}

type LoginStruct struct {
	Id_cliente int    `json:"id_cliente"`
	Contrasena string `json:"contrasena"`
}

type Result struct {
	Mas_vendido    int `json:"producto_mas_vendido"`
	Menos_vendido  int `json:"producto_menos_vendido"`
	Mas_ganancia   int `json:"producto_mas_ganancia"`
	Menos_ganancia int `json:"producto_menos_ganancia"`
}

type Compra struct {
	Id_cliente int       `json:"id_cliente"`
	Carro      []Carrito `json:"productos"`
}

type Carrito struct {
	Id_producto int `json:"id_producto"`
	Cantidad    int `json:"cantidad"`
}

var ID_Session int

func GetProducts() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:5000/api/productos", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Productos

	json.Unmarshal(bodyBytes, &responseObject)
	for _, array := range responseObject.Data {
		fmt.Println(array.Id_producto, ";", array.Nombre, ";", array.Precio_unitario, " por unidad;", array.Cantidad_disponible, " disponibles")
	}
}

// Borrar
// funcion estadistica
func GetStats() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:5000/api/estadisticas", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Result
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println("Producto mas vendido: ", responseObject.Mas_vendido)
	fmt.Println("Producto menos vendido: ", responseObject.Menos_vendido)
	fmt.Println("Producto con mas ganancia: ", responseObject.Mas_ganancia)
	fmt.Println("Producto con menos ganancia: ", responseObject.Menos_ganancia)
}

//----------------------------------------------------------------------------------------------------------------------------------------------

func CheckProduct(ID_PRODUCTO string) (flag bool, cantidad int, precio int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:5000/api/producto/"+ID_PRODUCTO, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var responseObject Productos2
	json.Unmarshal(bodyBytes, &responseObject)
	bodyString := string(bodyBytes)
	if bodyString != "{\"error\":\"RecordNotFound\"}" {
		cantidad = responseObject.Data.Cantidad_disponible
		precio := responseObject.Data.Precio_unitario
		flag = true
		return flag, cantidad, precio
	} else {
		flag = false
		return flag, 0, 0
	}
}

func BuyProducts() {

	var cantidadProductos int
	fmt.Print("Ingrese cantidad de productos a comprar: ")
	fmt.Scanln(&cantidadProductos)
	cont := 0
	var compra Compra
	compra.Id_cliente = ID_Session
	var producto string
	var carrito Carrito
	var gasto int
	cantidadTotal := 0

	var arrayIDS []int
	flagg := true

	for cont < cantidadProductos {

		fmt.Printf("Ingrese producto %d par id-cantidad: ", cont+1)
		fmt.Scan(&producto)
		separador := strings.Split(producto, "-")
		booleano, cantidadProductosStock, precio := CheckProduct(separador[0]) //booleano retorna true si los productos superan la cantidad
		if booleano == false {
			fmt.Printf("No existe el producto con id %s\n", separador[0])
		} else {

			num, _ := strconv.ParseInt(separador[0], 10, 0)
			cant, _ := strconv.ParseInt(separador[1], 10, 0)
			for _, element := range arrayIDS {
				if element == int(num) {
					flagg = false
				}
			}
			if flagg {
				if int(cant) < cantidadProductosStock {
					carrito.Id_producto = int(num)
					cantidadTotal += int(cant)
					carrito.Cantidad = int(cant)
					compra.Carro = append(compra.Carro, carrito)
					gasto += int(cant) * precio
				}
			} else {
				var aux Carrito
				var position int
				for pos, a := range compra.Carro {
					if a.Id_producto == int(num) {
						if a.Cantidad+int(cant) < cantidadProductosStock {
							a.Cantidad += int(cant)
							cantidadTotal += int(cant)
							gasto += int(cant) * precio
							position = pos
							aux = a
						} else {
							fmt.Println("Productos en stock insuficientes")
							break
						}
					}
				}
				if flagg == false {
					compra.Carro[position] = aux
				}
				flagg = true

			}
			//agregamos el arrayIDS
			arrayIDS = append(arrayIDS, int(num))
		}
		cont++
	}
	fmt.Println()
	jsonReq, err := json.Marshal(compra)
	resp, err := http.Post("http://localhost:5000/api/compras", "aplication/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if bodyString == "{\"error\":\"Producto no encontrado\"}" {
		fmt.Println("Ingrese el id de los productos correctamente")
	} else {
		if bodyString != "{\"error\":\"error\"}" && bodyString != "{\"data\":\"error\"}" && bodyString != "{\"data\":\"Key: 'CreateCompraInput.DetalleInputs' Error:Field validation for 'DetalleInputs' failed on the 'required' tag\"}" {
			fmt.Println("Monto total de la compra: ", gasto)
			fmt.Println("Cantidad de productos comprados: ", cantidadTotal)
		} else {
			fmt.Println("Producto/s sin stock")
		}
	}
}

// -------------------------------------------------------------------------------------------------------------------------------------------------------------------
func PostProduct() string {

	var name string
	var quantity int
	var price int

	fmt.Print("Ingrese el nombre: ")
	fmt.Scanln(&name)
	fmt.Print("Ingrese la disponibilidad: ")
	fmt.Scanln(&quantity)
	fmt.Print("Ingrese el precio unitario: ")
	fmt.Scanln(&price)

	product := CreateProduct{name, quantity, price}
	jsonReq, err := json.Marshal(product)
	resp, err := http.Post("http://localhost:5000/api/producto", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	return "Producto ingresado correctamente"
}
func DeleteProduct() {
	var id_producto int
	fmt.Print("Ingrese el id del producto a eliminar: ")
	fmt.Scanln(&id_producto)
	url := "http://localhost:5000/api/producto/" + strconv.Itoa(id_producto)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	fmt.Println("Producto Eliminado")
}
func Login() bool {
	var id int
	var contrasena string

	fmt.Print("Ingrese su id: ")
	fmt.Scanln(&id)
	fmt.Print("Ingrese su contraseña: ")
	fmt.Scanln(&contrasena)

	login := LoginStruct{id, contrasena}
	jsonReq, err := json.Marshal(login)
	resp, err := http.Post("http://localhost:5000/api/clientes/iniciar_sesion", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatalln(err)
		fmt.Println("Que paso aca")
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	if bodyString == "{\"acceso valido\":true}" {
		ID_Session = id
		return true
	} else {
		return false
	}
}
func ClientOption() {
	boolean := true

	for boolean == true {
		fmt.Println()
		fmt.Println("Opciones:")
		fmt.Println("1. Ver lista de productos")
		fmt.Println("2. Hacer compra")
		fmt.Println("3. Salir")
		fmt.Print("Ingrese una opción: ")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			GetProducts()
		case 2:
			BuyProducts()
		case 3:
			boolean = false
		}
	}
}
func ManagerOption() {
	boolean := true

	fmt.Println("Opciones:")
	for boolean == true {
		var option int
		fmt.Println("1. Ver lista de productos")
		fmt.Println("2. Crear producto")
		fmt.Println("3. Eliminar producto")
		fmt.Println("4. Ver estadísticas")
		fmt.Println("5. Salir")
		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			GetProducts()
		case 2:
			fmt.Println(PostProduct())
		case 3:
			DeleteProduct()
		case 4:
			GetStats()
		case 5:
			boolean = false
		default:
			fmt.Println("Ingrese una opción válida")
		}
	}
}

func ManagerSession() {
	boolean := Login()
	if boolean == true {
		fmt.Println("Inicio de sesión exitoso")
		ManagerOption()
	} else {
		fmt.Println("Error, no hay ninguna coincidencia con los datos ingresados.")
	}
}

func ClientSession() {
	//GetProducts()
	boolean := Login()
	if boolean == true {
		fmt.Println("Inicio de sesión exitoso")
		ClientOption()
	} else {
		fmt.Println("Error, no hay ninguna coincidencia con los datos ingresados.")
	}
}

func main() {
	boolean := true
	fmt.Println("Bienvenido")
	for boolean == true {
		fmt.Println()
		fmt.Println("Opciones:")
		fmt.Println("1. Iniciar sesión como cliente")
		fmt.Println("2. Iniciar sesión como administrador")
		fmt.Println("3. Salir")

		var option int
		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(&option)

		//fmt.Println(eleccion)
		switch option {
		case 1:
			ClientSession()
		case 2:
			ManagerSession()
		case 3:
			boolean = false
		default:
			fmt.Println("Ingrese una opción válida")
		}
	}
	fmt.Println("Hasta luego!")
}
