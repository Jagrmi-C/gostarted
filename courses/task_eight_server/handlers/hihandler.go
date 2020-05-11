package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
    "io/ioutil"

    "github.com/Jagrmi-C/gostarted/courses/lib"
)

func HiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func IncreseX2Handler(w http.ResponseWriter, r *http.Request) {
    b, err := ioutil.ReadAll(r.Body)

    if err != nil {
        lib.ReturnInternalError(w)
        return
    }

    var exampleObj lib.Test_struct

    if err := json.Unmarshal(b, &exampleObj); err != nil {
        lib.ReturnInternalError(w)
        return
    }

    // var res string
    res := lib.IsIntegerX2(exampleObj.ID)

    if _, err := w.Write([]byte(res)); err != nil {
        log.Println(err)
    }
}
