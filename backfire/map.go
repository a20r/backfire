package backfire

type Point struct {
    X   int `json:x`
    Y   int `json:y`
    Z   int `json:z`
}

type Configuration struct {
    Initial Point `json:initial`
    Goal    Point `json:goal`
}

type Obstacle struct {
    Dynamic  bool    `json:dynamic`
    Name     string  `json:name`
    Vertices []Point `json:vertices`
}

type Map struct {
    Name        string `json:name`
    AuthorName  string `json:author_name`
    AuthorEmail string `json:author_email`
    Id          string `json:id`
    Map         struct {
        Width          int             `json:width`
        Height         int             `json:height`
        Depth          int             `json:depth`
        Configurations []Configuration `json:configurations`
        Obstacles      []Obstacle      `json:obstacles`
    }   `json:map`
}
