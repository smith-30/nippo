package config

type Config struct {
	HealthCheck HealthCheck
	Gops        Gops
	Logger      Logger
}

type HealthCheck struct {
	Enable bool
	Host   string
	Port   string
}

type Gops struct {
	Enable bool
}

type Logger struct {
	Development bool
	Level       string

	// "console" or "json"
	Encoding            string
	OutputPaths         []string
	AppErrorOutputPaths []string
	ErrorOutputPaths    []string
	EncoderConfig
}

type EncoderConfig struct {
	MessageKey    string
	LevelKey      string
	TimeKey       string
	NameKey       string
	CallerKey     string
	StacktraceKey string

	// "capital", "capitalColor" or "color". default is lowercase
	LevelEncoder  string
	CallerEncoder string
}

func NewDefaultConfig() *Config {
	return &Config{}
}
