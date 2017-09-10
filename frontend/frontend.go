// ProjectName:: RevAirQuality ver1.0
//
// ProjectTitle:: Revenue Oriented Air Quality Prediction Services
//
// ProjectDescription:: See doc folder for more informationj.
//
// ProjectFile:: frontend.go
//
// Author:: Shajulin Benedict shajulin@iiitkottayam.ac.in or benedict@in.tum.de
//
// PackageName: RevAirQualityService: main.go - main file for Reveneu Oriented Prediction Services
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany

package frontend

import (
  "fmt"
  "net/http"
  "log"
  "io"
  "os"
  "strconv"
  "gopkg.in/mgo.v2"
  "encoding/csv"
)

/// frontend info here.
var FrontEnd = `
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

<h1>Frontend for RevAirQualityService!!</h1>

<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Step 1: Collects sensor Data (Both training and testing)<p>
<p> Step 2: Invokes querylog service - collects testing data and prepares the frequency of invokation <p>
<p> Step 3: Invokes prediction services - for validation and graph generation <p>
<p> Step 4: Invokes revenue service - calculates the credits for the smartcity <p>

</body>
</html>
`



/// frontendhandler service - Revenue Oriented Air Quality Prediction service
func Frontendservice(w http.ResponseWriter, req *http.Request){

   //w.Write([]byte(Frontend))
   fmt.Fprintf(w, FrontEnd)

   //Time to time measurement from stattions
   //i) store csv to mongodb - https://gist.github.com/laddy/37fa5dcbbe7893167d8c
   type Mongo struct {
     StationCode   string
     Date   string
     State string
     City string
     Location string
     Agency string
     TypeofLocation string
     SO2 int
     NO2 int
     RSPM int
     PM string
    }

     session, err := mgo.Dial("localhost:27017")
     if err != nil {
       panic(err)
     }

     defer session.Close()
     session.SetMode(mgo.Monotonic, true)

     c := session.DB("RevAirQualityDB").C("entries")
     file, err := os.Open("/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-kerala.csv")

     //APdata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-arunachal.csv",sep=",",head=TRUE)
     //Pondydata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-puducherry.csv",sep=",",head=TRUE)
     //TNdata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-tamilnadu.csv",sep=",",head=TRUE)
     //Karnatakadata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-karnataka.csv",sep=",",head=TRUE)
     //Keraladata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-kerala.csv", sep=",", head=TRUE)



     if err != nil {
       panic(err)
     }
     defer file.Close()

     reader := csv.NewReader(file)

     for {
       record, err := reader.Read()
       if err == io.EOF {
         break
       } else if err != nil {
         panic(err)
       }
       so2int, err := strconv.Atoi(record[7])
       no2int, err := strconv.Atoi(record[8])
       rspm2int, err := strconv.Atoi(record[9])

       err = c.Insert(&Mongo{record[0], record[1], record[2], record[3], record[4], record[5], record[6], so2int, no2int, rspm2int, record[10] })

       if err != nil {
         panic(err)
       }
       log.Printf("%#v", record)
     }





}
