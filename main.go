package main

import (
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	if len(os.Args[1:]) != 2 {
		log.Fatalf("Usage: %s template outconfig\n", os.Args[0])
	}
	funcMap := template.FuncMap{
		"random": RandStringBytes,
	}
	vars := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		vars[pair[0]] = pair[1]
	}
	base := path.Base(os.Args[1])
	tmpl, err := template.New(base).Funcs(funcMap).ParseFiles(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to Parse File: %s", err)
	}
	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("Unable To Open File For Writing", err)
	}
	err = tmpl.ExecuteTemplate(f, base, vars)
	if err != nil {
		log.Fatalf("Unable to Parse File: %s", err)
	}
}
