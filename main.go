package main

import (
    "flag"
    "github.com/wallarelvo/backfire/backfire"
)

func main() {

    var host_flag = flag.String(
        "host",
        "localhost",
        "Host the backfire server binds to",
    )

    var port_flag = flag.Int(
        "port",
        8000,
        "Port used for backfire server",
    )

    var db_host_flag = flag.String(
        "db_host",
        "localhost",
        "Host that the database uses",
    )

    var db_port_flag = flag.Int(
        "db_port",
        28015,
        "Port used for the database",
    )

    var db_name_flag = flag.String(
        "db_name",
        "backfire",
        "The name of the database being used",
    )

    flag.Parse()

    bfs := backfire.MakeServer(
        *host_flag,
        *port_flag,
        *db_host_flag,
        *db_port_flag,
        *db_name_flag,
    )

    bfs.Start()
}
