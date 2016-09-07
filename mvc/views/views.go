package views

import (
    "fmt"
    "html/template"
    "net/http"
    "mvc/config"
    "mvc/models"
)

func LoadView(w http.ResponseWriter, routeName string, data ...[]models.Data) {

    routeInfo := config.GetRouteInfo(routeName)

    view := models.View{
        Route: routeInfo.RouteName,
        Files: routeInfo.RouteFilesDest,
    }

    templates := template.Must(template.ParseFiles(view.Files...))

    if len(data) <= 0 {
        for i := 0; i < len(view.Files); i++ {
            templates.ExecuteTemplate(w, routeInfo.RouteFilesName[i], nil)
        }
    } else {
        dataDest := []string{}
        var x, y int
        for x = 0; x < len(data); x++ {
            for y = 0; y < len(data[x]); y++ {
                dataDest = append(dataDest, data[x][y].Destination)
            }
            fmt.Println("strings: ",dataDest[x])
        }
        fmt.Println("strings: ",dataDest[x])
        fmt.Println("strings: ",len(dataDest))
        fmt.Println(" ")
        fmt.Println(" ")
        fmt.Println(" ")
        loadDataFlag := 0
        var cont int
        var j, k, z int
        for j = 0; j < len(view.Files); j++ {
            fmt.Println("length: ",len(data))
            for k = 0; k < len(data); k++ {
                for z = 0; z < len(data[k]); z++ {
                    if routeInfo.RouteFilesName[j] == data[k][z].Destination {
                        cont++;
                    }
                    if cont >= len(data[k]) {
                        templates.ExecuteTemplate(w, routeInfo.RouteFilesName[j], data[k])
                        loadDataFlag = 1
                        cont = 0
                    }
                }
            }
            if loadDataFlag == 0 {
                fmt.Println("Load Template: nil")
                templates.ExecuteTemplate(w, routeInfo.RouteFilesName[j], nil)
            } else {
                loadDataFlag = 0
            }
        } //EACH CYCLE LOADS A FILE
    }
}
