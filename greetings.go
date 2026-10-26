package greeting

import (
	"errors"
	"fmt"
	"math/rand"

	
)

//endpoint
func Hello(name string) (string, error) {
	//entrada de los datos
	if name == ""{
		return "",errors.New("El nombre esta vacio")
	}
	message := fmt.Sprintf(randomFormats(),name)
	return message,nil
}

//endpoint que utiliza a otro endpoint(recibe envia info)
func Hellos(names []string)(map[string]string, error){
	messages := make(map[string]string)
	for _,name := range names{
		message,err := Hello(name)
		if err != nil{
			return nil,err
		}
		messages[name]=message
	}
	return messages,nil
}

func randomFormats()string{
	formats := []string{
		"Hola!, %v Bienvenido!",
		"Que bueno verte, %v!",
		"Saludo, %v, encantado de concocerte!",
	}
	return formats[rand.Intn(len(formats))]
}

