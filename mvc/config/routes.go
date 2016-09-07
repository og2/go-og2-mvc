package config

import (
    "mvc/models"
)
func GetRouteInfo(r string) models.Route {
    var route models.Route

    switch r {
    case "home":
        route = models.Route {
            Route: "/",
            RouteName: "home",
            RouteFilesName: []string{"header", "content", "footer"},
            RouteFilesDest: []string{"templates/header.html", "templates/content.html", "templates/footer.html"},
        }
        break
    case "login":
        break
    default:
        break
    }

    return route
}
