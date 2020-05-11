package handlers

import (
    "fmt"
    "net/http"
    "log"
	"io/ioutil"
	"html/template"
	"path"
	"time"
	"net/url"

    "github.com/Jagrmi-C/gostarted/courses/lib"
)

type Profile struct {
	Name    string
	Address string
	Token	[]string
}

func addCookie(w http.ResponseWriter, name string, value string) {
    expire := time.Now().AddDate(1, 1, 1)
    cookie := http.Cookie{
        Name:    name,
        Value:   value,
		Expires: expire,
		MaxAge: 10,
	}
	fmt.Println("EX", expire)
    http.SetCookie(w, &cookie)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		expire := time.Now().AddDate(1, 1, 1)

		cookie.Name = "unused"
		cookie.Value = ""
		cookie.MaxAge = -1
		cookie.Path = "/api/v0/"
		cookie.HttpOnly = true
		cookie.Expires = expire

		http.SetCookie(w, cookie)
    }

	respText := []byte(`<a href="/api/v0/">HOME</a><br>old cookie del!\n`)

	if _, err := w.Write(respText); err != nil {
		log.Println(err)
	}
}

func GETRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tokenValue []string
	if val, cond := r.Header["Cookie"]; cond {
		tokenValue = val
		log.Println("Set token: ", tokenValue)
	}

	profile := Profile{"Alex", "Minsk", tokenValue}

	fp := path.Join("courses", "templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// log.Println("URL", r.URL)


	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func POSTRootHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		lib.ReturnInternalError(w)
		return
	}

	m, err:= url.ParseQuery(string(b))

	if err != nil {
		lib.ReturnInternalError(w)
		return
	}

	var profile Profile

	if val, ok := m["name"]; ok {
		profile.Name = val[0]
	}

	if val, ok := m["address"]; ok {
		profile.Address = val[0]
	}

	addCookie(w, profile.Name, profile.Address)

	log.Println(m, m["name"][0], r.URL)

	// fmt.Println()

	http.Redirect(w, r, "http://localhost:3003/api/v0/", http.StatusSeeOther)

	// if _, err := w.Write([]byte("POST")); err != nil {
	// 	log.Println(err)
	// }
}

// func TestHandler(w http.ResponseWriter, r *http.Request) {
// 	htmlBody := `<a href="/api/v0/">HOME</a>`
// 	if _, err := w.Write([]byte(htmlBody)); err != nil {
// 		log.Println(err)
// 	}
// }