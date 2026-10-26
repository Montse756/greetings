# Saludos en Go
Este paquete porporciona una forma simple de obtener saludos personalizados

## Instalacion
Ejecuta el siguiente comando para instalar el paquete:
```bash
    go get -u github.com/Montse756/greetings
```

##Uso
```go 
    package main
    import (
            "fmt"
            "github.com/Montse756/greetings.git"
            )
    func main(){
        msg,err := greeting.Hello("Montse")
        if err != nil{
            fmt.Println("Ocurrio un error: ",err)
            return
        }
        fmt.Println(msg)
    }
```