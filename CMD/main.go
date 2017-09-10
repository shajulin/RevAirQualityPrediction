// ProjectName:: RevAirQuality ver1.0
//
// ProjectTitle:: Revenue Oriented Air Quality Prediction Services
//
// ProjectDescription:: See doc folder for more information.
//
// Author:: Shajulin Benedict shajulin@iiitkottayam.ac.in or benedict@in.tum.de
//
// PackageName: RevAirQualityService: main.go - main file for Reveneu Oriented Prediction Services
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany
package main

import (
    _"encoding/json"
    "log"
    "fmt"
    _"os"
    _"strconv"
    _"io/ioutil"
    "net/http"
    _"gopkg.in/mgo.v2"
    _"gopkg.in/mgo.v2/bson"
    "github.com/gorilla/mux"
    "github.com/RevAirQualityService/frontend"
    "github.com/RevAirQualityService/querylog"
    "github.com/RevAirQualityService/revenue"
    "github.com/RevAirQualityService/rfpredictSO2"
    "github.com/RevAirQualityService/rfpredictNO2"
    "github.com/RevAirQualityService/rfpredictRSPM"
)


/// Main info here.
var mainInfo = `
<!DOCTYPE html>
<html>
<head>
  <title>EventInfo</title>
  <style>
    body {background-color: powderblue;}
    h1 {color: red;}
    p {color: blue;}
  </style>
</head>
<body>

<h1>Main:::RevAirQualityService!!</h1>

<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Welcome to Revenue Oriented Air Quality Prediction services! \n Options are i) /frontend ii) querylog iii) revenue iv) rfpredictionSO2 v) rfpredictionNO2 vi) rfpredictionRSPM" <p>

</body>
</html>
`


/// Infohandler service - Revenue Oriented Air Quality Prediction service
func infohandler(w http.ResponseWriter, req *http.Request){
   fmt.Fprintf(w, mainInfo)
}
func frontendhandler(w http.ResponseWriter, req *http.Request){
  frontend.Frontendservice(w, req)
  frontend.Preparedata()
}
func queryloghandler(w http.ResponseWriter, req *http.Request){
  querylog.Querylogservice(w, req)
}
func revenuehandler(w http.ResponseWriter, req *http.Request){
  revenue.Revenueservice(w, req)
}
func rfpredictSO2handler(w http.ResponseWriter, req *http.Request){
  rfpredictSO2.RFpredictSO2service(w, req)
}
func rfpredictNO2handler(w http.ResponseWriter, req *http.Request){
  rfpredictNO2.RFpredictNO2service(w, req)
}
func rfpredictRSPMhandler(w http.ResponseWriter, req *http.Request){
  rfpredictRSPM.RFpredictRSPMservice(w, req)
}


/**
  RevAirQualityService - main function
*/
func main() {
    fmt.Println("Initiating the RevAirQuality Prediction Services !!")
    fmt.Println("RevAirQualityService starts at localhost:8172")
    router := mux.NewRouter()

    router.HandleFunc("/", infohandler)
    router.HandleFunc("/frontend", frontendhandler)
    router.HandleFunc("/querylog", queryloghandler)
    router.HandleFunc("/revenue", revenuehandler)
    router.HandleFunc("/rfpredictSO2", rfpredictSO2handler)
    router.HandleFunc("/rfpredictNO2", rfpredictNO2handler)
    router.HandleFunc("/rfpredictRSPM", rfpredictRSPMhandler)

    log.Fatal(http.ListenAndServe(":8172", router))
}
