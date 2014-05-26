package backfire

import (
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "strconv"
)

type BackfireServer struct {
    Host     string
    Port     int
    Database *Db
}

type ErrorMessage struct {
    Err  string `json:"error"`
    Code int    `json:"code"`
}

func MakeServer(
    host string,
    port int,
    db_host string,
    db_port int,
    db_name string,
) *BackfireServer {

    return &BackfireServer{host, port, MakeDb(db_host, db_port, db_name)}

}

func (bfs *BackfireServer) Start() {

    handler := rest.ResourceHandler{}
    handler.SetRoutes(
        &rest.Route{"GET", "/maps/name/:name", bfs.GetMapByName},
        &rest.Route{"GET", "/maps/authorName/:name", bfs.GetMapByAuthorName},
    )

    http.ListenAndServe(bfs.Host+":"+strconv.Itoa(bfs.Port), &handler)

}

func (bfs *BackfireServer) GetMapByAuthorName(
    w rest.ResponseWriter,
    req *rest.Request,
) {

    name := req.PathParam("name")
    maps, err := bfs.Database.GetByAuthorName(name)

    if err != nil {
        empty_list := make([]Map, 0)
        w.WriteJson(empty_list)
    }

    w.WriteJson(maps)

}

func (bfs *BackfireServer) GetMapByName(
    w rest.ResponseWriter,
    req *rest.Request,
) {

    name := req.PathParam("name")
    maps, err := bfs.Database.GetByName(name)

    if err != nil {
        w.WriteJson(&ErrorMessage{
            Err:  err.Error(),
            Code: 1,
        })
        return
    }

    w.WriteJson(maps)

}
