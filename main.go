package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

//http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/all/xml/3840?res=3hourly&key

type SiteList struct {
    Locations struct {
        Location []struct {
            Elevation       string `json:"elevation,omitempty"`
            ID              string `json:"id"`
            Latitude        string `json:"latitude"`
            Longitude       string `json:"longitude"`
            Name            string `json:"name"`
            Region          string `json:"region,omitempty"`
            UnitaryAuthArea string `json:"unitaryAuthArea,omitempty"`
            ObsSource       string `json:"obsSource,omitempty"`
            NationalPark    string `json:"nationalPark,omitempty"`
        } `json:"Location"`
    } `json:"Locations"`
}

type weatherStations struct {
    id     string
    Name   string
    Region string
    UA     string
}

// https://tutorialedge.net/golang/consuming-restful-api-with-go/
func siteList() map[string]weatherStations {
    //response, err := http.Get("http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/all/json/sitelist?key=429155dc-b1d0-402c-a200-2d155160cf52")
    response, err := http.Get("http://datapoint.metoffice.gov.uk/public/data/val/wxobs/all/json/sitelist?key=429155dc-b1d0-402c-a200-2d155160cf52")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject SiteList
    err = json.Unmarshal(responseData, &responseObject)

    wsMap := make(map[string]weatherStations)
    for _,v := range responseObject.Locations.Location {
        wsMap[v.Name] = weatherStations{id: v.ID, Name: v.Name, Region: v.Region, UA: v.UnitaryAuthArea}
    }

    return wsMap
}

func main() {
    var siteList = siteList()
    for k,v := range siteList {
        fmt.Printf("%v %v\n", k,v)
    }
}