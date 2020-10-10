package reporter

import (
	"fmt"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/types"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/bradfitz/slice"
	"github.com/olekukonko/tablewriter"
)

type ResponseInfo = types.ResponseInfo

func ReportTrace(responses []ResponseInfo) {
	numberOfRequests := fmt.Sprintf("%v", len(responses))
	fastest, slowest := getFastestAndSlowestTime(responses)
	meanTime := mean(responses)
	medianTime := median(responses)
	succeeded, errorCodes := getResponseStatusAndCodes(responses)
	smallest, largest := getResponseSizes(responses)

	data := [][]string{
		[]string{
			numberOfRequests,
			fastest.String(),
			slowest.String(),
			meanTime.String(),
			medianTime.String(),
			fmt.Sprintf("%s %v", "%", succeeded),
			fmt.Sprintf("%v", errorCodes),
			fmt.Sprintf("%v", smallest),
			fmt.Sprintf("%v", largest),
		},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"# Requests",
		"Fastest Time",
		"Slowest Time",
		"Mean Times",
		"Median Times",
		"% Succeeded",
		"Error Codes",
		"Smallest Response",
		"Largest Response",
	})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func getFastestAndSlowestTime(responses []ResponseInfo) (time.Duration, time.Duration) {
	fastest := time.Duration(math.MaxInt32)
	slowest := time.Duration(0)

	for _, response := range responses {
		if response.Time > slowest {
			slowest = response.Time
		}

		if response.Time < fastest {
			fastest = response.Time
		}
	}

	return fastest, slowest
}

func mean(responses []ResponseInfo) time.Duration {
	sum := time.Duration(0)
	for _, response := range responses {
		sum += response.Time
	}
	return sum / time.Duration(len(responses))
}

func median(responses []ResponseInfo) time.Duration {
	slice.SortInterface(responses[:], func(i, j int) bool {
		return responses[i].Time < responses[j].Time
	})

	mNumber := len(responses) / 2

	if len(responses)%2 == 1 {
		return responses[mNumber].Time
	}

	return (responses[mNumber-1].Time + responses[mNumber].Time) / 2
}

func getResponseStatusAndCodes(responses []ResponseInfo) (int, []int) {
	succeeded := 0
	var errorCodes []int

	for _, response := range responses {
		if response.Succeed {
			succeeded += 1
		}

		if response.StatusCode != http.StatusOK {
			if contains(errorCodes, response.StatusCode) == false {
				errorCodes = append(errorCodes, response.StatusCode)
			}
		}
	}

	return succeeded / len(responses) * 100, errorCodes
}

func contains(arr []int, value int) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}

	return false
}

func getResponseSizes(responses []ResponseInfo) (int, int) {
	smallest := math.MaxInt32
	largest := 0

	for _, response := range responses {
		if response.Size > largest {
			largest = response.Size
		}

		if response.Size < smallest {
			smallest = response.Size
		}
	}

	return smallest, largest
}
