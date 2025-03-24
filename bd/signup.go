package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juanmontilva/gambituser/models"
	"github.com/juanmontilva/gambituser/tools"
)

// Signup realiza el registro de un nuevo usuario en la base de datos
//
// Parámetros:
//   - sig: Estructura Signup que contiene la información del usuario a registrar
//
// Retorna:
//   - error: nil si el registro fue exitoso, error en caso contrario
//
// La función:
// 1. Establece conexión con la base de datos
// 2. Inserta el nuevo usuario con su email, UUID y fecha de registro
// 3. Maneja la desconexión de la base de datos
func Signup(sig models.Signup) error {
	fmt.Println("Comiensa Registro")

	// Establece conexión con la base de datos
	err := DbConnect()
	if err != nil {
		return err
	}

	// Asegura que la conexión se cierre al finalizar la función
	defer Db.Close()

	// Prepara la sentencia SQL para insertar el nuevo usuario
	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMysql() + "')"
	fmt.Println(sentencia)

	// Ejecuta la sentencia SQL
	// No nos interesa el número de filas afectadas, solo si hay algún error
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SIGNUP > EJECUCION EXITOSA")
	return nil
}
