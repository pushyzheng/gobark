package gobark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var githubIconUrl = "https://github.com/fluidicon.png"

func TestNewPushOptions(t *testing.T) {
	options := NewPushOptions("title", "content")
	options.SetReceiver("")
	options.SetGroup("test")
	options.SetIconUrl(githubIconUrl)
	options.SetSound(Calypso)
	options.SetLevel(Active)
	options.SetCopyText("hello")

	assert.Equal(t, options.Title, "title")
	assert.Equal(t, options.Content, "content")
	assert.Equal(t, options.Group, "test")
	assert.Equal(t, options.IconUrl, githubIconUrl)
	assert.Equal(t, options.Sound, Calypso)
	assert.Equal(t, options.Level, Active)
	assert.Equal(t, options.CopyText, "hello")
	assert.False(t, options.AutomaticallyCopy)
	assert.False(t, options.IsArchive)

	options.SetArchived()
	options.SetAutomaticallyCopy()
	assert.True(t, options.AutomaticallyCopy)
	assert.True(t, options.IsArchive)
}
