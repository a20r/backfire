package backfire

import (
    "github.com/ant0ine/go-json-rest/rest"
    r "github.com/christopherhesse/rethinkgo"
    "log"
    "net/http"
)

type Message struct {
    Body string
}

func Run() {

    session, _ := r.Connect("localhost:28015", "test")

    var dbs []string
    r.DbList().Run(session).All(&dbs)
    log.Println(dbs)

    handler := rest.ResourceHandler{}
    handler.SetRoutes(
        &rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
            w.WriteJson(&Message{
                Body: "Hello World!",
            })
        }},
    )
    http.ListenAndServe(":8000", &handler)
}
