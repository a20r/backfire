package backfire

import (
    "errors"
    r "github.com/christopherhesse/rethinkgo"
    "strconv"
)

type Db struct {
    Host string
    Port int
    Name string
}

func MakeDb(host string, port int, name string) *Db {

    db := &Db{host, port, name}
    db.Init()
    return db

}

func (db *Db) GetTable(table_name string) r.Exp {

    return r.Db(db.Name).Table(table_name)

}

func (db *Db) Connect() (*r.Session, error) {

    session, conn_err := r.Connect(db.Host+":"+strconv.Itoa(db.Port), db.Name)
    return session, conn_err

}

func (db *Db) Create() error {

    var session *r.Session
    var err error

    session, err = db.Connect()

    if err != nil {
        return err
    } else {
        err = r.DbCreate(db.Name).Run(session).Exec()
        err = r.Db(db.Name).TableCreateWithSpec(
            r.TableSpec{
                Name:       "maps",
                PrimaryKey: "name",
            },
        ).Run(session).Exec()

        if err != nil {
            return err
        }
    }

    return nil

}

func (db *Db) Init() (bool, error) {

    var session *r.Session
    var err error

    session, err = db.Connect()

    if err != nil {
        return false, err
    }

    var dbs []string

    err = r.DbList().Run(session).All(&dbs)

    for _, db_name := range dbs {
        if db.Name == db_name {
            return false, nil
        }
    }

    err = db.Create()

    return true, err
}

func (db *Db) Insert(cmap Map) (*r.WriteResponse, error) {

    var session *r.Session
    var err error

    session, err = db.Connect()

    if err != nil {
        return nil, err
    }

    var response r.WriteResponse

    err = db.GetTable("maps").Insert(cmap).Run(session).One(&response)

    if err != nil {
        return nil, err
    } else {
        return &response, nil
    }

}

func (db *Db) GetByName(name string) (*Map, error) {

    var session *r.Session
    var err error

    session, err = db.Connect()

    if err != nil {
        return nil, err
    }

    var cmap Map

    err = db.GetTable("maps").Get(name).Run(session).One(&cmap)

    if cmap.Name == "" {
        return nil, errors.New("Name not found")
    }

    if err != nil {
        return nil, err
    }

    return &cmap, nil

}

func (db *Db) GetByAuthorName(name string) ([]Map, error) {

    var session *r.Session
    var err error

    session, err = db.Connect()

    if err != nil {
        return nil, err
    }

    var maps []Map

    err = db.GetTable("maps").Filter(
        r.Row.Attr("authorName").Eq(name),
    ).Run(session).All(&maps)

    if err != nil {
        return nil, err
    }

    return maps, nil

}
