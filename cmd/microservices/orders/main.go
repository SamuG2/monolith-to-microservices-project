package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/SamuG2/monolith-to-microservices-golang/pkg/common/cmd"

)

func createOrderMicroService()(router *chi.Mux, closeFn func()){
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r := cmd.CreateRouter()
	orders_public_http.AddRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)
	return r, func() {

	}
}

func main() {
	log.Println("Arrancando microservicio de ordenes...")
	ctx := cmd.Context()
	r, closeFn :=createOrderMicroService()
	defer closeFn()

	server := &http.Server(Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r)

	go func(){
		if err := server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}()

		<-ctx.Done()
		log.Println("Cerrando microservicio de ordenes.")
		if err:= server.Close(); err != nil{
			panic(err)
		}
	}
}
