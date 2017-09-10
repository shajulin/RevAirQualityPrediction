// ProjectName:: RevAirQuality ver1.0
//
// ProjectTitle:: Revenue Oriented Air Quality Prediction Services
//
// ProjectDescription:: See doc folder for more informationj.
//
// ProjectFile:: prepareData.go
//
// Author:: Shajulin Benedict shajulin@iiitkottayam.ac.in or benedict@in.tum.de
//
// PackageName: RevAirQualityService: main.go - main file for Reveneu Oriented Prediction Services
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany

package frontend

import (
 "fmt"
 "github.com/senseyeio/roger"
)
 func Preparedata() {
    // connect to RServe using Roger
    rClient, err := roger.NewRClient("127.0.0.1", 6311)
    if err != nil {
    fmt.Printf("Failed to connect to RServe: %s", err.Error())
     return
    }
    // call prepareData R function, gathering the response
    returnVar, err := rClient.Eval("prepareData()")
    if err != nil {
    fmt.Printf("Training and testing data are not prepared %s and %s", err.Error(), returnVar)
    return
    }
}
