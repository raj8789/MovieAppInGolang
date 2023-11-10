package main

import ("fmt"
"log"
"net/http"
"os"
"github.com/gorilla/mux"
 "handler"
)

func main(){
	fmt.Println("Movie App Started Successfully")
	logger := log.New(os.Stdout, "Movie-API", log.LstdFlags)
	ph := handler.NewMovieHandler(logger)
	sm:=mux.NewRouter()
	
	getsubRouter:=sm.Methods(http.MethodGet).Subrouter()
	getsubRouter.HandleFunc("/movies",ph.GetMovies)
	getsubRouter.HandleFunc("/movies/{id:[0-9]+}",ph.GetMovie)

	putsubRouter:=sm.Methods(http.MethodPut).Subrouter()
	putsubRouter.HandleFunc("/movies/{id:[0-9]+}",ph.UpdateMovie)


	postRouter:=sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/movies",ph.CreateMovie)

	deleteRouter:=sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/movies/{id:[0-9]+}",ph.DeleteMovie)

	http.ListenAndServe(":8080", sm)
}
