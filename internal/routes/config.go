package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors" //chi is a lightweight, idiomatic and composable router for building Go HTTP services.
	// It's especially good at helping you write large REST API services that are kept maintainable as your project grows and changes. chi is built on the new context package introduced in Go 1.7 to handle signaling, cancelation and request-scoped values across a handler chain.
)

type Config struct {
	timeout time.Duration //to store multiple values of the same data type into a single variable
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           5,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSeconds int) *Config {
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c
}
func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
