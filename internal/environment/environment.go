package environment

type EnvClient struct {
	Region          string `env:"AWS_REGION"`
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Port            string `env:"SERVER_PORT"`
}

var Env EnvClient
