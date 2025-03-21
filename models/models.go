package models

// esto se hace para conectar al secrets manager
type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"Port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

// esto es para el evento recibir parametros, como son muchos se hara el de user email y el id, esto cuando se registran

type Signup struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
