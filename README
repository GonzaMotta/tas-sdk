# Proyecto Go con TAS SDK

Este proyecto utiliza el paquete `tas-sdk` para realizar auditorías y guardar trazas en un servicio. A continuación, se explican los pasos para descargar, renombrar e implementar las funciones básicas de auditoría y trazabilidad.

## Requisitos previos

- Tener instalado Go en tu sistema.
- Contar con acceso a GitHub y la capacidad de descargar paquetes desde él.

## Instalación

1. Descarga el paquete `tas-sdk` ejecutando el siguiente comando:

    ```bash
    go get github.com/GonzaMotta/tas-sdk
    ```

2. Renombra el paquete en tu proyecto para facilitar su uso. Esto se hace agregando un alias en tu archivo de código Go:

    ```go
    import sdkTas "github.com/GonzaMotta/tas-sdk"
    ```

## Uso

### Crear una Auditoría

Para crear una nueva auditoría, usa la función `CreateAudit` proporcionada por el SDK. Este método devuelve un `auditId` único que podrás utilizar para asociarlo a otras operaciones de seguimiento.

```go
auditId := sdkTas.CreateAudit() // 673228f4f0579b84116e0b6f
```


### Guardar una Traza tipo string 

Una vez que tengas un auditId, puedes guardar una traza en tu sistema utilizando SaveTraceV2. Esta función toma tres parámetros: una descripción de la traza, el nombre del servicio y el auditId.

```go
data := "here your data in string format " 
sdkTas.SaveStringTrace("Servicio fake de api", data, auditId) // return trace id 
```


### Guardar una Traza tipo Objeto 

Una vez que tengas un auditId, puedes guardar una traza en tu sistema utilizando SaveTraceV2. Esta función toma tres parámetros: una descripción de la traza, el nombre del servicio y el auditId.

```go
service := "nombre_del_servicio" 

objecto := map[string]interface{}{
		"campo_!":   "value-1",
		"campo_2":     1,
		"campo_3": "value_3",
	}

sdkTas.SaveObjectTrace("Servicio fake de api", data interface{}, auditId) // return trace id 
```

