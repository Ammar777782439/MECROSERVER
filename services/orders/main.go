package main

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()
		
	grpcServer := NewGRPCServer(":9091")
	grpcServer.Run()
}
