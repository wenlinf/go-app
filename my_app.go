package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func getMeasurement() (float64, error) {
	url := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local:80/get_measurement"

    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    measurementValue := 0.0
    _, err = fmt.Sscanf(string(body), "%f", &measurementValue)
    if err != nil {
        return 0, err
    }

    return measurementValue, nil
}

func main() {
    interval := 1 * time.Minute
    numMeasurements := 5

    var totalMeasurement float64

    for i := 0; i < numMeasurements; i++ {
        measurement, err := getMeasurement()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        fmt.Println("Measurement:", measurement)
        totalMeasurement += measurement
        time.Sleep(interval)
    }

    average := totalMeasurement / float64(numMeasurements)
    fmt.Println("Average:", average)
}
