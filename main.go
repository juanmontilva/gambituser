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

// main es el punto de entrada de la aplicación Lambda
// Inicializa el manejador de eventos de AWS Lambda
func main() {
	lambda.Start(EjecutoLambda)
}

// EjecutoLambda es el manejador principal de eventos de Cognito
// Este método se ejecuta cuando un usuario confirma su registro en Cognito
//
// Parámetros:
//   - ctx: Contexto de la ejecución Lambda, necesario para manejar timeouts y cancelaciones
//   - event: Evento de Cognito que contiene la información del usuario que se registró
//
// Retorna:
//   - El mismo evento de Cognito (requerido por el contrato de AWS)
//   - Un error si algo falla durante el proceso
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// Inicializa la configuración de AWS necesaria para la Lambda
	awsgo.InicializoAws()

	// Valida que existan los parámetros de entorno necesarios
	if !ValidoParametros() {
		fmt.Println("ERROR EN LOS PARAMETROS. DEBE ENVIAR SECRETNAME")
		err := errors.New("error en los parametros debe enviar secretname")
		return event, err
	}

	// Estructura para almacenar los datos del usuario
	var datos models.Signup

	// Extrae la información relevante del usuario desde los atributos del evento
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

	// Lee los secretos necesarios para la conexión a la base de datos
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("error al leer el secret " + err.Error())
		return event, err
	}

	// Realiza el registro del usuario en la base de datos
	err = bd.Signup(datos)
	return event, err
}

// ValidoParametros verifica que existan las variables de entorno necesarias
// para el funcionamiento correcto de la Lambda
//
// Retorna:
//   - true si encuentra el parámetro SecretName
//   - false si no encuentra el parámetro
func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
