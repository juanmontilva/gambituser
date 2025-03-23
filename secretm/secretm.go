package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/juanmontilva/gambituser/awsgo"
	"github.com/juanmontilva/gambituser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {

	var datosSecret models.SecretRDSJson
	fmt.Println("> Pido Secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	// estructura para procesar el getsecretvalue en este sentido se hace el puntero
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(">Lectura secret ok" + nombreSecret)

	return datosSecret, nil

	// todo resumido como funcion de secrets manager
}
