package main

import "github.com/wallarelvo/backfire/backfire"

var DbHost string = "localhost"
var DbPort int = 28015
var DbName string = "backfire"

var BackfireHost string = "localhost"
var BackfirePort int = 8000

func main() {
    bfs := backfire.MakeServer(
        BackfireHost,
        BackfirePort,
        DbHost,
        DbPort,
        DbName,
    )

    bfs.Start()
}
