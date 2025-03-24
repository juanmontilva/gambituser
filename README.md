# GambitUser - Sistema de Gestión de Usuarios

Este proyecto es un sistema de gestión de usuarios desarrollado en Go, que utiliza servicios de AWS para proporcionar una solución escalable y segura para el manejo de usuarios en una tienda online.

## Descripción del Proyecto

GambitUser es un microservicio que maneja la autenticación y gestión de usuarios para una tienda online. El sistema está diseñado para ser desplegado como funciones Lambda en AWS, proporcionando una arquitectura serverless.

## Estructura del Proyecto

```
gambituser/
├── main.go           # Punto de entrada principal de la aplicación
├── awsgo/           # Configuraciones y utilidades específicas de AWS
├── models/          # Definiciones de estructuras de datos y modelos
├── secretm/         # Manejo de secretos y configuraciones sensibles
├── tools/           # Herramientas y utilidades generales
└── bd/              # Operaciones y conexiones con la base de datos
```

## Componentes Principales

1. **main.go**: Punto de entrada de la aplicación que inicializa los servicios y configura las rutas.
2. **awsgo/**: Contiene la configuración y utilidades para interactuar con servicios AWS.
3. **models/**: Define las estructuras de datos y modelos utilizados en la aplicación.
4. **secretm/**: Maneja la gestión segura de secretos y configuraciones sensibles.
5. **tools/**: Proporciona utilidades generales y herramientas de soporte.
6. **bd/**: Gestiona las operaciones de base de datos y conexiones.

## Requisitos

- Go 1.x o superior
- Cuenta de AWS con acceso a:
  - AWS Lambda
  - AWS DynamoDB
  - AWS Secrets Manager
  - AWS Cognito

## Configuración

1. Clonar el repositorio:
```bash
git clone [URL_DEL_REPOSITORIO]
```

2. Instalar dependencias:
```bash
go mod download
```

3. Configurar variables de entorno:
```bash
export AWS_REGION=tu-region
export AWS_ACCESS_KEY_ID=tu-access-key
export AWS_SECRET_ACCESS_KEY=tu-secret-key
```

## Despliegue

El proyecto está diseñado para ser desplegado como dos funciones Lambda en AWS:

1. **Lambda de Registro**: Maneja el proceso de registro de nuevos usuarios
2. **Lambda de Autenticación**: Gestiona la autenticación y validación de usuarios

## Documentación Técnica

### Flujo de Registro
1. El usuario envía sus datos de registro
2. La Lambda de Registro valida los datos
3. Se crea el usuario en DynamoDB
4. Se genera un token de confirmación
5. Se envía un email de confirmación

### Flujo de Autenticación
1. El usuario envía sus credenciales
2. La Lambda de Autenticación valida las credenciales
3. Se genera un token JWT
4. Se devuelve el token al usuario

## Seguridad

- Todas las contraseñas se almacenan con hash seguro
- Se utilizan tokens JWT para la autenticación
- Los secretos se manejan a través de AWS Secrets Manager
- Se implementan políticas de CORS y rate limiting

## Mantenimiento

Para mantener el proyecto actualizado y seguro:
1. Revisar regularmente las dependencias
2. Actualizar las políticas de seguridad de AWS
3. Monitorear los logs de CloudWatch
4. Realizar pruebas de integración periódicas

## Contribución

Para contribuir al proyecto:
1. Fork el repositorio
2. Crear una rama para tu feature
3. Realizar tus cambios
4. Enviar un Pull Request

## Licencia

[Especificar la licencia del proyecto]