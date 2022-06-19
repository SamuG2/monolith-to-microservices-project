package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SamuG2/monolith-to-microservices-golang/pkg/common/cmd"
)

func createPaymentsMicroservices() amqp.paymentsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	paymentsService := payments_app.NewPaymentsService(
		payments_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)
	paymentsInterface, err := amqp.NewPaymentsInterface(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
		paymentsService,
	)
	if err != nil {
		panic(err)
	}
	return paymentsInterface

}

func main() {
	log.Println("Arrancando microservicio de pagos...")
	defer log.Println("Cerrando microservicio de pagos.")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicroservices()
	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}

}
