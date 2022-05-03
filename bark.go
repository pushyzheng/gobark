package gobark

type OnFailed func(receiver string)

type OnSuccess func(receiver string)

type PushOptions struct {
	Title             string
	Content           string
	Url               string
	Receivers         []string
	CopyText          string
	Sound             SoundType
	Group             string
	Level             LevelType
	IconUrl           string // above iOS15
	AutomaticallyCopy bool
	IsArchive         bool
	OnSuccess         OnSuccess // The listener of successful event
	OnFailed          OnFailed  // The listener of failed event
}

func (options *PushOptions) SetTitle(title string) *PushOptions {
	options.Title = title
	return options
}

func (options *PushOptions) SetUrl(url string) *PushOptions {
	options.Url = url
	return options
}

func (options *PushOptions) SetReceiver(receiver string) *PushOptions {
	return options.SetReceivers([]string{receiver})
}

func (options *PushOptions) SetReceivers(receivers []string) *PushOptions {
	options.Receivers = receivers
	return options
}

func (options *PushOptions) SetCopyText(copyText string) *PushOptions {
	options.CopyText = copyText
	return options
}

func (options *PushOptions) SetSound(sound SoundType) *PushOptions {
	options.Sound = sound
	return options
}

func (options *PushOptions) SetGroup(group string) *PushOptions {
	options.Group = group
	return options
}

func (options *PushOptions) SetLevel(level LevelType) *PushOptions {
	options.Level = level
	return options
}

func (options *PushOptions) SetIconUrl(iconUrl string) *PushOptions {
	options.IconUrl = iconUrl
	return options
}

func (options *PushOptions) SetAutomaticallyCopy() *PushOptions {
	options.AutomaticallyCopy = true
	return options
}

func (options *PushOptions) SetArchived() *PushOptions {
	options.IsArchive = true
	return options
}

func (options *PushOptions) SetOnSuccess(OnSuccessFunc OnSuccess) *PushOptions {
	options.OnSuccess = OnSuccessFunc
	return options
}

func (options *PushOptions) SetOnFailed(OnFailedFunc OnFailed) *PushOptions {
	options.OnFailed = OnFailedFunc
	return options
}

func NewPushOptions(title string, content string) *PushOptions {
	return &PushOptions{
		Content: content,
		Title:   title,
		Sound:   DefaultSound,
		Level:   Active,
	}
}
