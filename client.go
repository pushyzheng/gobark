package gobark

import (
	"fmt"
	"github.com/asmcos/requests"
	logger "github.com/sirupsen/logrus"
)

type response struct {
	Code    int
	Data    interface{}
	Message string
}

type BarkClient struct {
	Domain string
}

func (client BarkClient) PushSimple(title string, content string, receivers []string) ([]string, error) {
	options := PushOptions{
		Title:     title,
		Content:   content,
		Receivers: receivers,
	}
	return client.Push(&options)
}

func (client BarkClient) Push(options *PushOptions) ([]string, error) {
	failed := make([]string, 0)

	if len(options.Receivers) == 0 {
		logger.Warn("No any receivers", options)
		return nil, nil
	}
	for _, key := range options.Receivers {
		url := client.buildUrl(*options, key)
		resp, err := requests.Get(url)
		if err != nil {
			logger.Errorf("request fail, url: %s, error: %s", url, err.Error())
			failed = fail(*options, failed, key)
			continue
		}
		data := response{}
		err = resp.Json(&data)
		if err != nil || data.Code != 200 {
			logger.Errorf("response stauts is't ok, url: %s, resp: %s", url, resp.Text())
			failed = fail(*options, failed, key)
			continue
		}
		if options.OnSuccess != nil {
			options.OnSuccess(key)
		}
	}
	if len(failed) == 0 {
		return nil, nil
	} else {
		return failed, fmt.Errorf("push failed: %d", len(failed))
	}
}

func fail(options PushOptions, failed []string, key string) []string {
	failed = append(failed, key)
	if options.OnFailed != nil {
		options.OnFailed(key)
	}
	return failed
}

func (client BarkClient) buildUrl(options PushOptions, key string) string {
	domain := client.Domain
	if len(domain) == 0 {
		domain = DefaultDomain
	}
	result := fmt.Sprintf("https://%s/%s", domain, key)
	if len(options.Title) != 0 {
		result = result + "/" + options.Title
	}
	result = fmt.Sprintf("%s/%s?automatically_copy=%t", result, options.Content, options.AutomaticallyCopy)
	if len(options.Url) != 0 {
		result = result + "&url=" + options.Url
	}
	if len(options.Sound) != 0 {
		result = result + "&sound=" + string(options.Sound)
	}
	if len(options.Group) != 0 {
		result = result + "&group=" + options.Group
	}
	if len(string(options.Level)) != 0 {
		result = result + "&level=" + string(options.Level)
	}
	if len(options.IconUrl) != 0 {
		result = result + "&icon=" + options.IconUrl
	}
	if options.IsArchive {
		result = result + "&isArchive=1"
	}
	return result
}
