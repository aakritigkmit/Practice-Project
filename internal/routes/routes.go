package routes

import "github.com/go-chi/chi"

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.getConfig().Timeout),
		router: chi.NewRouter(),

	}
}

func (r* Router) setRouters() *chi.Mux{

}

func (r* Router) setConfigsRouters() *chi.Mux{
	
}

func RouterHealth(){

}

func RouterProduct(){
	
}

func EnableTimeout(){
	
}

func EnableCORS(){
	
}


func EnableRecover(){
	
}

func EnableRequestID(){
	
}