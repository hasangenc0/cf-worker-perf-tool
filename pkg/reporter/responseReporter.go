package reporter

import (
	"fmt"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/errors"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/http"
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/logger"
	urlValidator "github.com/hasangenc0/cf-worker-perf-tool/pkg/url"
)

func ReportHttpResponse(url string) {
	logger.Info(fmt.Sprintf("Starting to analyse %s ...", url))
	if !urlValidator.IsValidUrl(url) {
		logger.Error(errors.InvalidUrlError)
		return
	}

	response, err := http.Get(url)

	if err != nil {
		logger.Error(errors.AnErrorOccurred)
	}

	fmt.Println(response)
}
