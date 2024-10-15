package setting

type Config struct {
	Server        ServerSetting `mapstructure:"server"`
	Jwt           JwtSetting    `mapstructure:"jwt"`
	DB            DBSetting     `mapstructure:"db"`
	Logger        LoggerSetting `mapstructure:"logger"`
	Redis         RedisSetting  `mapstructure:"redis"`
	Smtp          SmtpSetting   `mapstructure:"smtp"`
	GoogleSetting GoogleSetting `mapstructure:"google"`
	AsynqSetting  AsynqSetting  `mapstructure:"asynq"`
	AWS           AWSSetting    `mapstructure:"aws"`
	S3            S3Setting     `mapstructure:"s3"`
}

type ServerSetting struct {
	Name      string `mapstructure:"name"`
	Port      int    `mapstructure:"port"`
	Mode      string `mapstructure:"mode"`
	WebURL    string `mapstructure:"webURL"`
	ServerURL string `mapstructure:"serverURL"`
}

type JwtSetting struct {
	Secret            string `mapstructure:"secret"`
	RefreshSecret     string `mapstructure:"refreshSecret"`
	Expiration        int    `mapstructure:"expiration"`
	RefreshExpiration int    `mapstructure:"refreshExpiration"`
}

type DBSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbname"`
	SslMode         string `mapstructure:"sslmode"`
	Timezone        string `mapstructure:"timezone"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

type SmtpSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Sender   string `mapstructure:"sender"`
}

type GoogleSetting struct {
	ClientID     string `mapstructure:"clientId"`
	ClientSecret string `mapstructure:"clientSecret"`
}

type AsynqSetting struct {
	DelayInSeconds       int `mapstructure:"delayInSeconds"`
	MaxConcurrentWorkers int `mapstructure:"maxConcurrentWorkers"`
}

type AWSSetting struct {
	Region string `mapstructure:"region"`
}

type S3Setting struct {
	BucketName string `mapstructure:"bucketName"`
}
