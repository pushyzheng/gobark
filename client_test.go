package gobark

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var myKey = os.Getenv("MAC_BARK_KEY")

var wrongKey = "wrong key"

func TestBarkClient_PushSimple(t *testing.T) {
	client := BarkClient{Domain: DefaultDomain}
	failed, err := client.PushSimple("title", "content", []string{myKey})
	if err != nil {
		panic(err)
	}
	assert.Equal(t, len(failed), 0)
}

func TestBarkClient_PushSuccess(t *testing.T) {
	client := BarkClient{Domain: DefaultDomain}

	options := NewPushOptions("title", "content")
	options.SetReceivers([]string{myKey, "1892912"})
	options.SetIconUrl("https://github.com/fluidicon.png")
	options.SetGroup("test")
	options.SetOnSuccess(func(key string) {
		assert.Equal(t, myKey, key)
	})
	options.SetLevel(TimeSensitive)
	options.SetArchived()

	failed, err := client.Push(options)
	if err != nil {
		panic(err)
	}
	fmt.Println(failed)
}

func TestBarkClient_PushFailed(t *testing.T) {
	client := BarkClient{Domain: DefaultDomain}

	options := NewPushOptions("title", "content")
	options.SetReceiver(wrongKey)
	options.SetOnFailed(func(receiver string) {
		assert.Equal(t, wrongKey, receiver)
	})

	failed, err := client.Push(options)
	assert.NotNil(t, err)
	assert.Equal(t, failed[0], wrongKey)
}
