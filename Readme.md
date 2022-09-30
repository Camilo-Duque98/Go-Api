# Tarea 1 sistemas distribuidos
### **Integrantes**
| Nombre y Apellido | Rol         |
|-------------------|-------------|
| Javier Aguayo     | 201873056-4 |
| Felipe Cruz       |             |
| Camilo Villar     | 201873096-3 |  

# Pequeña explicación
Ántes de ejecutar el programa, queremos informar que el programa lo dividimos en dos carpetas, servidor y cliente:  
* Servidor: Esta carpeta, contiene a todo lo relacionado a la API, y como su nombre lo dice, está centrada en el lado del servidor. Nos basamos en el patrón Model-View-Controller para la creación de esta API. Utilizamos ``Gin`` para realizar las peticiones a la base de datos, y ocupamos un orm para hacer los mapeos a los datos entrantes, llamado ``Gorm``.
* Cliente: Esta carpeta esta centrada en el lado del cliente. contiene dos archivos, que es main.go y su  modulo correspondiente. 

Ahora, ¿Por que realizamos esta separación de archivos?, esta separación de archivos la realizamos para que no allá colisiones por la función ``main`` definida por golang, ya que generaría problemas si dos archivos que se encuentran en el mismo módulo contienen la función main.  

# Ejecución del programa  
Para ejecutar el programa, necesitaremos dos terminales abiertas, una que estará centrada en el lado del servidor, y otra que estará centrada en el lado del cliente.  
Para la consola que se encarga del servidor, utilizamos el siguiente comando para compilarlo:  
```go  
/> go run server.go
```  
Para la consola del cliente realizamos algo similar a lo mencionado anteriormente:  
```go  
/> go run main.go
```