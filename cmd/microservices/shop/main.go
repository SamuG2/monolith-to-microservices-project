package main

import (
	"log"
	"net/http"
	"os"
	"github.com/SamuG2/monolith-to-microservices-golang/pkg/common/cmd"

)

func createShopMicroservice() *chi.Mux{
	shopProduct := shop_infra_product.NewMemoryRepository()
	r := cmd.CreateRouter()
	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo)

	return r
}

func main(){
	log.Println("Arrancando microservicio de tienda...")

	ctx := cmd.Context()

	r := createShopMicroservice()

	server := &http.Server(Addr: os.Getenv("SHOP_PRODUCT_SERVICE_BIND_ADDRESS"), Handler: r)
	go func(){
		if err := server.ListenAndServe(); err != http.ErrServerClosed{
			panic(err)
		}
	}()

		<-ctx.Done()
		log.Println("Cerrando microservicio de tienda")
		if err := server.Close(); err != nil{
			panic(err)
		}
	
}