package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanmontilva/gambituser/models"
	"github.com/juanmontilva/gambituser/tools"
)

func Signup(sig models.Signup) error {
	fmt.Println("Comiensa Registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMysql() + "')"
	fmt.Println(sentencia)
	// NO ME INTERESA LAS FILAS AFECTADAS LO QUE ME INTERESA ES VER SI HAY ALGUN ERROR
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SIGNUP > EJECUCION EXITOSA")

	return nil

}
