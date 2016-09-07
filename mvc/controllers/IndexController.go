package controllers

import (
    "mvc/models"
    "mvc/views"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    users := []models.User {
        //User{"martin", "martin@email.com"}
    }

    ids := []string{}
    pass := []string{}

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/gamersgg_gggdb")
    if err != nil {
        log.Fatal(err)
    } else {
        rows, err := db.Query("SELECT id, username, email, password FROM tblusers")
        checkErr(err)
        defer rows.Close()
        for rows.Next() {
            var username string
            var email string
            var id string
            var password string
            err = rows.Scan(&username, &email, &id, &password)
            userRow := models.User{Name: username, Email: email}
            users = append(users, userRow)
            ids = append(ids, id)
            pass = append(pass, password)
            checkErr(err)
        }
        rows.Close()
    }
    defer db.Close()

    data := models.Data{
        Destination: "content",
        IData: users,
    }

    data2 := models.Data{
        Destination: "header",
        IData: ids,
    }

    data3 := models.Data{
        Destination: "content",
        IData: pass,
    }

    data4 := []models.Data{}
    data4 = append(data4, data)
    data4 = append(data4, data3)
    data5 := []models.Data{}
    data5 = append(data5, data2)

    views.LoadView(w, "home", data4, data5)

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
