package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getMatrix() ([]float64, error) {
    // send request to get response matrix from microplate reader
	url := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local:80/get_measurement"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

    // convert the reponse to an array of floats and return the array
	lines := strings.Split(string(body), "\n")
	var matrix []float64

	for _, line := range lines {
		fields := strings.Fields(line)
		for _, field := range fields {
			value, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return nil, err
			}
			matrix = append(matrix, value)
		}
	}

	return matrix, nil
}

func main() {
    // set the interval to call the microplate reader
	interval := 1 * time.Minute

	for {
		matrix, err := getMatrix()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
        // calculate the average of the returned values and print it out
		var totalMeasurement float64
		numMeasurements := len(matrix)

		for _, value := range matrix {
			totalMeasurement += value
		}

		average := totalMeasurement / float64(numMeasurements)
		fmt.Println("Average:", average)

		time.Sleep(interval)
	}
}
