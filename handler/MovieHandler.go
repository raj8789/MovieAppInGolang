package handler

import (
	"fmt"
	"log"
	"model"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"
)

var movies []*model.Movie
type MovieHandler struct {
	logger *log.Logger
}
func NewMovieHandler(logger *log.Logger)*MovieHandler{
	createDefaultMovie()
	mh:=&MovieHandler{logger: logger}
	return mh
}
func createDefaultMovie(){
	movies=append(movies,&model.Movie{1,"Singham","A Policeman Fight as Warrior in The Movie","Ajay Devgan","Kajal AggarWal"})
	movies=append(movies,&model.Movie{2,"Tare Jammen Par","A Teacher Teaches a Very Weak Student and Makes him Comfortable for Study","Amir Khan","Tisca Chopra"})
}
func (moviehandler *MovieHandler) GetMovie(rw http.ResponseWriter, re *http.Request){
	vars:=mux.Vars(re)
	id, err := strconv.Atoi(vars["id"])
	if err!=nil{
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw,"Id is Invalid")
		return
	}
	movie:=getMovieById(id)
	if movie==nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw,"Movie is not Present")
		return
	}
	jsondata,err:=movie.ToJson()
	if err!=nil{
		rw.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(rw,"Movie is Invalid")
		return
	}else{
		rw.WriteHeader(http.StatusOK)
	}
	rw.Header().Set("Content-Type", "application/json")
    rw.Write(jsondata)
}
func (moviehandler *MovieHandler) GetMovies(rw http.ResponseWriter, re *http.Request){
	jsonData, err := json.Marshal(movies)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    rw.Header().Set("Content-Type", "application/json")
    rw.Write(jsonData)
}
func (moviehandler *MovieHandler) UpdateMovie(rw http.ResponseWriter, re *http.Request){
	vars:=mux.Vars(re)
	id, err := strconv.Atoi(vars["id"])
	if err!=nil{
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw,"Id is Invalid")
		return
	}
	movietoUpdate:=getMovieById(id)
	if movietoUpdate==nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw,"Movie is not Present")
		return
	}else{
		if  re.Header.Get("name") != ""{
			name:=re.Header.Get("name")
			movietoUpdate.Name=name
		}
		if  re.Header.Get("description") != ""{
			description:=re.Header.Get("description")
			movietoUpdate.Description=description
		}
		if  re.Header.Get("actor") != ""{
			actor:=re.Header.Get("actor")
			movietoUpdate.Actor=actor
		}
		if  re.Header.Get("actress") != ""{
			actress:=re.Header.Get("actress")
			movietoUpdate.Actress=actress
		}
		result:=getIndexOfTheMovieInList(id)
		if result==-1{
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(rw,"Movie is not Present")
			return
		}
		movies[result]=movietoUpdate
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw,"Movie is Updated SuccessFully")
	}
}

func (moviehandler *MovieHandler) CreateMovie(rw http.ResponseWriter, re *http.Request){
	user:=&model.Movie{}
	user.FromJson(re)
	if user==nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw,"Movie is Invalid Not Created")
	}else{
		movies=append(movies,user)
		rw.WriteHeader(http.StatusOK)
		fmt.Fprint(rw,"Movie is  Created")
	}
}
func (moviehandler *MovieHandler) DeleteMovie(rw http.ResponseWriter, re *http.Request){
	vars:=mux.Vars(re)
	id, err := strconv.Atoi(vars["id"])
	if err!=nil{
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw,"Id is Invalid")
		return
	}
	result:=getIndexOfTheMovieInList(id)
	if result==-1{
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw,"Movie is not Present")
		return
	}else{
		prevlen:=len(movies)
		movies=append(movies[:result],movies[result+1:]...)
		newlen:=len(movies)
		if prevlen-newlen == 1{
			rw.WriteHeader(http.StatusOK)
			fmt.Fprint(rw,"Movie is  Deleted")
		}
	}
}
func getIndexOfTheMovieInList(id int)int{
	for ind,value:= range movies{
		if id==value.ID{
			return ind
		}
	}
	return -1
}
func getMovieById(id int)*model.Movie{
	for _,value :=range movies{
		if value.ID==id{
			return value
		}
	}
	return nil
}


