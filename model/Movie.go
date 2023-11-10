package model
import "encoding/json"
import "net/http"

type Movie struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Actor string `json:"actor"`
	Actress string `json:"actress"`
}
func (movie *Movie) ToJson()([]byte,error){
	return  json.Marshal(movie)
}
func (movie *Movie) FromJson(re *http.Request){
	json.NewDecoder(re.Body).Decode(&movie)
}