package main

import (
    "log"
    "github.com/go-resty/resty/v2"
    amqp "github.com/rabbitmq/amqp091-go"
)

// Función para manejar errores
func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func main() {
    // Crear cliente HTTP
    client := resty.New()

    // Conectar a RabbitMQ
	conn, err := amqp.Dial("amqp://kika:kikaokogi@44.194.169.99:5672")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    // Abrir un canal en RabbitMQ
    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // Declarar el exchange "logs"
    err = ch.ExchangeDeclare(
        "logs",   // Nombre del exchange
        "fanout", // Tipo del exchange
        true,     // Durable
        false,    // Auto-deleted
        false,    // Internal
        false,    // No-wait
        nil,      // Argumentos
    )
    failOnError(err, "Failed to declare an exchange")

    // Declarar la cola "reservations"
	q, err := ch.QueueDeclare(
		"reservation", // Nombre de la cola (cuidado con el nombre, en el error aparece "reservation" y no "reservations")
		true,          // ✅ Cambiar a `true` para que coincida con la configuración existente
		false,         // Eliminar cuando no se use
		false,         // Exclusivo
		false,         // No espera
		nil,           // Argumentos
	)
	
    failOnError(err, "Failed to declare a queue")

    // Vincular la cola "reservations" al exchange "logs"
    err = ch.QueueBind(
        q.Name,        // Nombre de la cola
        "",            // Clave de enrutamiento (vacío para fanout)
        "logs",        // Nombre del exchange
        false,
        nil,
    )
    failOnError(err, "Failed to bind the queue")

    // Consumir mensajes de la cola "reservations" (sin auto-ack)
    msgs, err := ch.Consume(
        q.Name,        // Nombre de la cola
        "",            // Nombre del consumidor
        false,         // No Auto-Acknowledge
        false,         // Exclusivo
        false,         // No local
        false,         // No esperar
        nil,           // Argumentos
    )
    failOnError(err, "Failed to register a consumer")

    // Canal para evitar que el programa termine
    forever := make(chan struct{})

    // Goroutine para procesar los mensajes
    go func() {
        for d := range msgs {
            log.Printf(" [x] Received message: %s", d.Body)

            // Verificar que el mensaje tenga contenido válido
            if len(d.Body) == 0 {
                log.Println("Mensaje vacío recibido, ignorando...")
                continue
            }

            // Imprimir información de la solicitud antes de enviarla
            log.Printf("Enviando solicitud POST a la API de reservaciones...")
            log.Printf("URL: http://127.0.0.1:8082/reservation")
            log.Printf("Headers: %v", map[string]string{"Content-Type": "application/json"})
            log.Printf("Body: %s", d.Body)

            // Intentos de reintento en caso de fallo
            for i := 0; i < 3; i++ {
                resp, err := client.R().
                    SetHeader("Content-Type", "application/json").
                    SetBody(d.Body).
                    Post("http://127.0.0.1:8082/reservation")

                // Si la solicitud es exitosa, salir del ciclo de reintentos
                if err == nil && resp.StatusCode() < 500 {
                    log.Printf("Solicitud exitosa. Código de estado: %d", resp.StatusCode())
                    break
                }

                log.Printf("Intento %d falló. Reintentando...", i+1)
            }

            // Mandar un ack explícito solo si todo fue exitoso
            err := d.Ack(false)
            failOnError(err, "Failed to acknowledge message")
        }
    }()

    log.Printf(" [*] Waiting for messages in 'reservations'. To exit press CTRL+C")
    <-forever
}
