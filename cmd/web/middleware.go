package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)


func NoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
