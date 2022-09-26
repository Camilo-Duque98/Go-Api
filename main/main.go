package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Productos struct {
	Data []Producto `json:"data"`
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

func GetProducts() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/productos", nil)
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

		fmt.Println(array.Id_producto, ";", array.Nombre, " por unidad;", array.Cantidad_disponible, " disponibles")
	}

}

// funcion estadistica
func GetStats() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/estadisticas", nil)
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

	fmt.Print(bodyBytes)
	//json.Unmarshal(bodyBytes, &responseObject)
	/*var responseObject Productos

	json.Unmarshal(bodyBytes, &responseObject)
	for _, array := range responseObject.Data {

		fmt.Println(array.Id_producto, ";", array.Nombre, " por unidad;", array.Cantidad_disponible, " disponibles")
	}*/

}

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
	resp, err := http.Post("http://localhost:8080/api/producto", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	//bodyBytes, _ := ioutil.ReadAll(resp.Body)

	//bodyString := string(bodyBytes)

	//fmt.Println(bodyString)

	return "Producto ingresado correctamente"
}
func DeleteProduct() {
	var id_producto int
	fmt.Print("Ingrese el id del producto a eliminar: ")
	fmt.Scanln(&id_producto)
	url := "http://localhost:8080/api/producto/" + strconv.Itoa(id_producto)
	//jsonReq, err:= json.Marshal(id_producto)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
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
	resp, err := http.Post("http://localhost:8080/api/clientes/iniciar_sesion", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if err != nil {
		log.Fatalln(err)
		fmt.Println("Que paso aca")
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	if bodyString == "{\"acceso valido\":true}" {
		return true
	} else {
		return false
	}
}
func ClientOption() {
	boolean := true

	for boolean == true {
		fmt.Println("Opciones:")
		fmt.Println("1. Ver lista de productos")
		fmt.Println("2. Hacer compra")
		fmt.Println("3. Salir")
		fmt.Println("Ingrese una opción: ")
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			GetProducts()
		case 2:
			fmt.Println("aqui supuestamente va hacer compra")
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
			fmt.Println("Aquí va a ver estadísticas")
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
	//presentacion

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
