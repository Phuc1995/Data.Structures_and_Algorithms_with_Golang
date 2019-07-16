package main

import "fmt"

func main()  {
/*	http.HandleFunc("/api/v1/message", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "hello",
			"created_at": time.Now(),
		})
	})
	http.ListenAndServe(":8080",nil)*/

	n := 5
	x,y := 0,1
	fmt.Println(x)
	for i := 0; i< n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}
