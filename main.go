package main

import (
	"fmt"
	"net/http"
)

/*

1. network interface card --> NIC
2. socket received buffer
3. write buffer
3. electronic magnify
4. file descriptor --> 0 < 1, 2, 3, 4...........

6. router - wifi adaptar - NIC - write buffer - interapct kurnel - copy write buffer and all thing is stored in socket received buffer
 send buffer kurnel niye read buffer ar kase dai

 NIC  electromagnatic kore router ar kase patai



*/

func helloHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandlar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Habi.., a student")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandlar)

	mux.HandleFunc("/about", aboutHandlar)

	fmt.Println("server running on: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
