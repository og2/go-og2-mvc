package models

type User struct {
    Name string
    Email string
}

type Route struct {
    Route string
    RouteName string
    RouteFilesName []string
    RouteFilesDest []string
}

type View struct {
    Route string
    Files []string
}

type Data struct {
    Destination string
    IData interface{}
}
