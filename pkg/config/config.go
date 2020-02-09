package config

type (
	Config struct {
		API    API    `mapstructure:"api"`
		Logger Logger `mapstructure:"logger"`
		DB     DB     `mapstructure:"db"`
	}

	API struct {
		Address  string `mapstructure:"address"`
		WTimeout int    `mapstructure:"write_timeout"`
		RTimeout int    `mapstructure:"read_timeout"`
		GTimeout int    `mapstructure:"graceful_timeout"`
	}

	Logger struct {
		Encoding    string   `mapstructure:"encoding"`
		OutputPaths []string `mapstructure:"output_paths"`
	}

	DB struct {
		Provider        string `mapstructure:"provider"`
		Host            string `mapstructure:"host"`
		Port            int    `mapstructure:"port"`
		DB              string `mapstructure:"db_name"`
		User            string `mapstructure:"user"`
		Password        string `mapstructure:"password"`
		SSL             string `mapstructure:"ssl"`
		Timeout         string `mapstructure:"timeout"`
		Retry           int    `mapstructure:"retry"`
		RandomDataCount int    `mapstructure:"random_data_count"`
	}
)

func LoadConfig(path string) Config {
	return Config{
		API:    API{},
		Logger: Logger{},
		DB:     DB{},
	}
}
