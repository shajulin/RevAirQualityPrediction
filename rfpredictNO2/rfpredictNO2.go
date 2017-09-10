// ProjectName:: RevAirQuality ver1.0
//
// ProjectTitle:: Revenue Oriented Air Quality Prediction Services
//
// ProjectDescription:: See doc folder for more informationj.
//
// ProjectFile:: rfpredictNO2.go
//
// Author:: Shajulin Benedict shajulin@iiitkottayam.ac.in or benedict@in.tum.de
//
// PackageName: RevAirQualityService: main.go - main file for Reveneu Oriented Prediction Services
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany


package rfpredictNO2

import (
  "fmt"
  "net/http"
  "github.com/senseyeio/roger"
)

/// rfpredictNO2 info here.
var RFpredictNO2 = `
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

<h1>rfpredictNO2 for RevAirQualityService!!</h1>

<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Step 1: Collects training and testing data <p>
<p> Step 2: Models training data using RF algorithm <p>
<p> Step 3: Predicts the testing data using RF algorithm <p>
<p> Step 4: Output the testing and training data in graph form <p>

</body>
</html>
`

/// rfpredictRSPMhandler service - Revenue Oriented Air Quality Prediction service
func RFpredictNO2service(w http.ResponseWriter, req *http.Request){

   //w.Write([]byte(rfpredictSO2))
   fmt.Fprintf(w, RFpredictNO2)
   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   // call predictSO2 R function, gathering the response
   returnVar, err := rClient.Eval("predictNO2()")
   if err != nil {
     fmt.Printf("PredictNO2 is not completed %s and %s", err.Error(), returnVar)
     return
   }

}
