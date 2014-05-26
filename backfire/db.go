package backfire

import (
    r "github.com/christopherhesse/rethinkgo"
    "strconv"
)

type Db struct {
    Host string
    Port int
    Name string
}

func MakeDb(host string, port int, name string) *Db {

    return &Db{host, port, name}

}

func (db *Db) Connect() (*r.Session, error) {

    session, conn_err := r.Connect(db.Host+":"+strconv.Itoa(db.Port), db.Name)
    return session, conn_err

}

func (db *Db) GetTableSpec() r.TableSpec {

    return r.TableSpec{
        Name:       "maps",
        PrimaryKey: "id",
    }

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
            db.GetTableSpec(),
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
