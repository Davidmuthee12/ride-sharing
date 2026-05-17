package main


func main() {
	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)
	svc.CreateTrip()
}