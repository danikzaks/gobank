package main

func main() {
	server := NewAPIServer(":9120")
	server.Run()
}
