package main

import "net/http"

func main() {
	// Create a handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTP"))
	})

	// Create a router
	router := &http.ServeMux{}

	// Register the handler with the router
	router.Handle("/", handler)

	// Create a server and register the router
	server := http.Server{
		Addr:    ":8088",
		Handler: router,
	}
	// Start the HTTP Server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func LoggingMiddleware(inner http.Handler) http.Handler {
	return http.HandleFunc(){
		
	}
}
