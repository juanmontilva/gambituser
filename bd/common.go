package bd

import (
	// database/sql es un estandar en golang pero la otra importacion que son los drivers no, hay que hacer un go get
	"database/sql"
	"fmt"
	"os"

	// este es el driver que se importa, es recomendado poner _ porque no me importa, solamente que funcione el driver de sql, no me la se toda, tengo qeu averiguar, esto esta en la documentacion, muy interesante, si no se pone el _ da error
	_ "github.com/go-sql-driver/mysql"
	"github.com/juanmontilva/gambituser/models"
	"github.com/juanmontilva/gambituser/secretm"
)

var SecretModel models.SecretRDSJson
var err error
// la variable Db va en mayuscula porque se leera de distintos lugares, tiene que ser de tipo puntero, todo lo que tiene que ver con manejo de base de datos se maneja de esa manera por velocidad, esto va a estar muy presente en el desarrollo
var Db *sql.DB

func ReadSecret() error {

	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}


func DbConnect() error{
	Db, err = sql.Open("mysql",  ConnStr(SecretModel))
	if err != nil{
		fmt.Println(err.Error())
		return err
	}
	// esto funciona como doble chequeo, muy recomendado ping funciona para apuntar el error, esto es la conexion basica de la base de datos
	err = Db.Ping()
	if err!= nil{
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("conexion exitosa de la base de datos")
	return nil

}


func ConnStr(claves models.SecretRDSJson) string{
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	// nombre de la base de datos para conectar
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	// cuando entre en produccion no se recomienda para el cloudwacht el fmt
	fmt.Println(dsn)
	return dsn

}