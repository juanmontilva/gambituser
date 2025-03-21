package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/juanmontilva/gambituser/awsgo"
	"github.com/juanmontilva/gambituser/bd"
	"github.com/juanmontilva/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

// con esto iniciamos el lambda, esto es obligatorio
// el contexto funciona como un director de orquesta, permite majerar practicamente de todo, es obligatorio en las lambdas por poder de estructura y logica, el segundo parametro ojo es para cognito, devolvemos practicamente lo mismo incluyendo el error, esto es fundamental para el cloudwatch
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// con esto incializamos la lambda y recuerda que agregamos el paquete awsgo, ojo es importante, luego le agregamos las funciones que necesita
	awsgo.InicializoAws()

	if !ValidoParametros() {
		fmt.Println("ERROR EN LOS PARAMETROS. DEBE ENVIAR SECRETNAME")
		err := errors.New("error en los parametros debe enviar secretname")
		return event, err
	}

	var datos models.Signup

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("email= " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("sub = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("error al leer el secret " + err.Error())
		return event, err
	}

	err = bd.Signup(datos)
	return event, err

}

// si no encuentra el parametro retorna un false, muy util esto es bueno para mantener la lambda de manera correcta, buena practica
func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
