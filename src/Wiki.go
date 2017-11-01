package main

import (
	"fmt"
	"net/http"
)

var Name string = "Test";

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if Name != "Test" {

		fmt.Fprintf(w, "<p> Hello, %s </p>", Name)

	}

	if Name == "Test" {

		fmt.Fprintf(w, "<p> Hello, %s </p>", Name)
	}
}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", loginPageHandler)
	http.ListenAndServe(":8080", nil)
}

const LoginPage = `
<h1>Login</h1>
<form method="post" action="/">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

func loginPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, LoginPage);

	if Name != "Test" {

		fmt.Fprintf(response, "Hello, %s", Name)

	}

	if Name == "Test"{
		fmt.Fprintf(response, "<p> Hello, %s </p>", Name)
	}
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	Name = request.FormValue("name")
	//pass := request.FormValue("password")
	redirectTarget := "/"
	http.Redirect(response, request, redirectTarget, 302)
}

