package youtube

type ChannelQueryParams struct {
	Name string `yaml:"name"`
}

type Config struct {
	ApiKey         string               `yaml:"apiKey"`
	TargetChannels []ChannelQueryParams `yaml:"targetChannels"`
}
