// ProjectName:: RevAirQuality ver1.0
//
// ProjectTitle:: Revenue Oriented Air Quality Prediction Services
//
// ProjectDescription:: See doc folder for more informationj.
//
// ProjectFile:: revenue.go
//
// Author:: Shajulin Benedict shajulin@iiitkottayam.ac.in or benedict@in.tum.de
//
// PackageName: RevAirQualityService: main.go - main file for Reveneu Oriented Prediction Services
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany


package revenue

import (
  "fmt"
  "net/http"
  "github.com/senseyeio/roger"
)

/// Revenue info here.
var Revenue = `
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

<h1>Revenue for RevAirQualityService!!</h1>

<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Step 1: Collect the list of testing data with number of queries <p>
<p> Step 2: If the number of queries is greater than 50, calculate the credit for city authorities a) 50*3*queries <p>
<p> Step 3: Otherwise, calculate the credit for city authorities as b) 63.3*3*queries <p>
<p> NOTE: we have utilized 3 air quality parameters for the measurements <p>

</body>
</html>
`



/// Revenue service - Revenue Oriented Air Quality Prediction service
func Revenueservice(w http.ResponseWriter, req *http.Request){

   //w.Write([]byte(rfpredictSO2))
   fmt.Fprintf(w, Revenue)

   //Pursue with the steps
   // connect to RServe using Roger
   rClient, err := roger.NewRClient("127.0.0.1", 6311)
   if err != nil {
   fmt.Printf("Failed to connect to RServe: %s", err.Error())
    return
   }
   // call prepareData R function, gathering the response
   returnVar, err := rClient.Eval("revenueData()")
   if err != nil {
     fmt.Printf("Revenue calcuation is not completed %s and %s", err.Error(), returnVar)
     return
   }

}
