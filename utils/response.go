
package utils

import (
    "encoding/json"
    "log"
    "net/http"
    
)

type response struct {
    Status int  `json:"status"`
    Result interface{} `json:"result`
}

func newResponse(data interface{} ,status int) *response{
 return &response{
    Status:status,
    Result:data
 }
}



func (resp *response) bytes() []bytes{
    
}

func (resp *response) bytes() []bytes{
    
}
func (resp *response) string() []string{
    
}
func (resp *response) sendResponse(w http.ResponseWriter, r http.Request){
    w.Header().Set("content-type" , "application-json")
}

func StatusNoContent(){

}
func StatusNotFound(){
    
}
func StatusBadRequest(){
    
}
func MethodNotAllowed(){
    
}
func StatusConflict(){
    
}
func StatusInternalServerError(){

}
