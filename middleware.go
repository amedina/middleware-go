package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {

	}
}