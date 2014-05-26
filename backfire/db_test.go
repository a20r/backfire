package backfire

import (
    "log"
    "testing"
)

func ProduceDb() *Db {

    var db_host string = "localhost"
    var db_port int = 28015
    var db_name string = "backfire"
    return MakeDb(db_host, db_port, db_name)

}

func TestInsert(t *testing.T) {

    log.Println("Testing Insert and MakeDb")
    db := ProduceDb()
    sample_map := Map{
        Name:           "test1",
        AuthorName:     "tester",
        AuthorEmail:    "test@example.com",
        Width:          640,
        Height:         480,
        Depth:          0,
        Configurations: make([]Configuration, 0),
        Obstacles:      make([]Obstacle, 0),
    }

    db.Insert(sample_map)

}

func TestGetByName(t *testing.T) {

    log.Println("Testing GetByName")

    test_name := "test1"
    db := ProduceDb()
    row, err := db.GetByName(test_name)

    if err != nil {
        t.Fatalf(err.Error())
    }

    log.Println(row)

}
