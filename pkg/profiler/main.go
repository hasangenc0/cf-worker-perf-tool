package profiler

import (
	"fmt"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/errors"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/logger"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/reporter"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/types"
	urlValidator "github.com/hasangenc0/cf-worker-perf-tool/pkg/url"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type ResponseInfo = types.ResponseInfo

func GetMetrics(profile int, url string) {
	if !urlValidator.IsValidUrl(url) {
		logger.Error(errors.InvalidUrlError)
		return
	}

	logger.Info(fmt.Sprintf("Profiling the %s ...", url))

	var responseInformation []ResponseInfo
	var wg sync.WaitGroup

	for i := 0; i < profile; i++ {
		wg.Add(1)
		go func() {
			info := makeRequest(url)
			defer wg.Done()
			responseInformation = append(responseInformation, info)
		}()
	}

	wg.Wait()
	reporter.ReportTrace(responseInformation)
}

func getSize(response *http.Response) int {
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return 0
	}

	return len(body)
}

func makeRequest(url string) ResponseInfo {
	responseInfo := ResponseInfo{
		Size:       0,
		Time:       0,
		Succeed:    false,
		StatusCode: 0,
	}

	req, _ := http.NewRequest("GET", url, nil)

	var start time.Time
	start = time.Now()
	response, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		logger.Error(errors.AnErrorOccurred)
	}

	if response != nil {
		responseInfo.Size = getSize(response)
		responseInfo.StatusCode = response.StatusCode
		responseInfo.Succeed = response.StatusCode == http.StatusOK
	}

	responseInfo.Time = time.Since(start)

	return responseInfo
}
