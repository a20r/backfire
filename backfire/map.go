package backfire

type Point struct {
    X   int `json:"x"`
    Y   int `json:"y"`
    Z   int `json:"z"`
}

type Configuration struct {
    Initial map[string]int `json:"initial"`
    Goal    map[string]int `json:"goal"`
}

type Obstacle struct {
    Dynamic  bool    `json:"dynamic"`
    Name     string  `json:"name"`
    Vertices []Point `json:"vertices"`
}

type Map struct {
    Name           string          `json:"name"`
    AuthorName     string          `json:"authorName"`
    AuthorEmail    string          `json:"authorEmail"`
    Width          int             `json:"width"`
    Height         int             `json:"height"`
    Depth          int             `json:"depth"`
    Configurations []Configuration `json:"configurations"`
    Obstacles      []Obstacle      `json:"obstacles"`
}
