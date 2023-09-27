package client

type Config struct {
	Endpoint     string `mapstructure:"udoma.endpoint" description:"The endpoint of the udoma API" default:"https://udoma.co"`
	ApiKey       string `mapstructure:"udoma.apikey" description:"The API key to access the udoma API" default:""`
	ApiSecretKey string `mapstructure:"udoma.apisecretkey" description:"The API key secret to access the udoma API" default:"" hidden:"true"`
}
