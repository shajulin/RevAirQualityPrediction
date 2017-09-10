// ProjectName:: RevAirQuality ver1.0
//
// ProjectTitle:: Revenue Oriented Air Quality Prediction Services
//
// ProjectDescription:: See doc folder for more informationj.
//
// ProjectFile:: querylog.go
//
// Author:: Shajulin Benedict shajulin@iiitkottayam.ac.in or benedict@in.tum.de
//
// PackageName: RevAirQualityService: main.go - main file for Reveneu Oriented Prediction Services
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany


package querylog

import (
  "fmt"
  "net/http"
  "github.com/senseyeio/roger"
)

/// Revenue info here.
var Querylog = `
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

<h1>Querylog for RevAirQualityService!!</h1>

<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Step 1: Create queries <p>
<p> Step 2: Create the list of air quality stations which were invoked for the analysis <p>
<p> Step 3: Identify the frequency of invokation <p>

</body>
</html>
`

/// Querylog service - Revenue Oriented Air Quality Prediction service
func Querylogservice(w http.ResponseWriter, req *http.Request){

   //w.Write([]byte(rfpredictSO2))
   fmt.Fprintf(w, Querylog)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   // call prepareData R function, gathering the response
   returnVar, err := rClient.Eval("queryData()")
   if err != nil {
     fmt.Printf("Testing data are not queried %s and %s", err.Error(), returnVar)
     return
   }

}
