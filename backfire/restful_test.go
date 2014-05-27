package backfire

import (
    "github.com/ant0ine/go-json-rest/rest/test"
    "io/ioutil"
    "net/http"
    "testing"
)

var base_url string = "http://localhost:8000"

func TestGetMapByAuthorName(t *testing.T) {

    author_name := "Alex Wallar"
    resp, err := http.Get(base_url + "/maps/authorName/" + author_name)

    if err != nil {
        t.Fatalf(err.Error())
        return
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    t.Log(string(body))

}

func TestGetMapByName(t *testing.T) {

    map_name := "test1"
    resp, err := http.Get(base_url + "/maps/name/" + map_name)

    if err != nil {
        t.Fatalf(err.Error())
        return
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    t.Log(string(body))

}

func TestPostMap(t *testing.T) {

    form := &Map{
        Name:           "PostTest",
        AuthorEmail:    "wallarelvo@gmail.com",
        AuthorName:     "Alex Wallar",
        Width:          640,
        Height:         480,
        Depth:          0,
        Configurations: make([]Configuration, 0),
        Obstacles:      make([]Obstacle, 0),
    }

    client := &http.Client{}

    req := test.MakeSimpleRequest("POST", base_url+"/maps", form)

    resp, err := client.Do(req)

    if err != nil {
        t.Fatalf(err.Error())
        return
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    t.Log(string(body))

}
