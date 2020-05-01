package config

import (
	"bytes"
	"io/ioutil"

	v "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

// Config application config structure, use viper for config and .env file or global env
type Config struct {
	Version       string
	ListenAddress string
	Production    bool

	Postgres              PostgresConfig
	Rabbit                RabbitConfig
	TaxCalculationService ExternalService
	TaxCalculationParams  ExternalParams
}

// PostgresConfig configuration structure for postgres database
type PostgresConfig struct {
	Addr     string
	User     string
	Pass     string
	Database string
	Debug    bool
}

type RabbitConfig struct {
	Host  string
	Port  int
	User  string
	Pass  string
	VHost string

	ConsumerExchangeName string
	ConsumerQueueName    string
	ConsumerRouteKey     string

	ProducerQueueName string
	ProducerRouteKey  string
}

type ExternalService struct {
	Address   string
	AuthToken string
	Retry     int
}

type ExternalParams struct {
	CompanyName       string
	RegistrationAU    map[string]struct{}
	RegistrationUS    map[string]struct{}
	RegistrationEU    map[string]struct{}
	RegistrationIsoAU string
	RegistrationIsoUS string
	RegistrationIsoEU string
	TaxIdUS           string
	TaxIdAU           string
	TaxIdEU           string
	ProductClass      string
	TransactionType   string
}

// Validate validate config structure
func (c Config) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.Version, v.Required),
		v.Field(&c.ListenAddress, v.Required),
		v.Field(&c.Postgres),
		v.Field(&c.TaxCalculationService),
	)
}

// Validate postgres config structure
func (p PostgresConfig) Validate() error {
	return v.ValidateStruct(&p,
		v.Field(&p.Addr, v.Required),
		v.Field(&p.User, v.Required),
		v.Field(&p.Pass, v.Required),
		v.Field(&p.Database, v.Required),
	)
}

// Validate rabbitMQ config structure
func (r RabbitConfig) Validate() error {
	return v.ValidateStruct(&r,
		v.Field(&r.Host, v.Required),
		v.Field(&r.Port, v.Required),
		v.Field(&r.User, v.Required),
		v.Field(&r.Pass, v.Required),
		//v.Field(&r.VHost, v.Required),
		v.Field(&r.ConsumerExchangeName, v.Required),
		v.Field(&r.ConsumerQueueName, v.Required),
		v.Field(&r.ConsumerRouteKey, v.Required),
		v.Field(&r.ProducerQueueName, v.Required),
		v.Field(&r.ProducerRouteKey, v.Required),
	)
}

// Validate external service config structure
func (p ExternalService) Validate() error {
	return v.ValidateStruct(&p,
		v.Field(&p.Address, v.Required),
		v.Field(&p.Retry, v.Required),
	)
}

// InitConfig initialize application configuration and validate all fields
func InitConfig() *Config {
	body, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("can't init config, err %s", err.Error())
	}

	return initConfig(body)
}

func initConfig(body []byte) *Config {
	vip := viper.New()
	vip.SetConfigType("yml")
	if err := vip.ReadConfig(bytes.NewBuffer(body)); err != nil {
		log.Fatal(err)
	}
	vip.AutomaticEnv()

	c := &Config{}

	c.Version = vip.GetString("VERSION")
	c.ListenAddress = vip.GetString("API_LISTEN_ADDRESS")
	c.Production = vip.GetBool("PRODUCTION")

	c.Postgres.Addr = vip.GetString("PG_ADDR")
	c.Postgres.User = vip.GetString("PG_USER")
	c.Postgres.Pass = vip.GetString("PG_PASS")
	c.Postgres.Database = vip.GetString("PG_NAME")
	c.Postgres.Debug = vip.GetBool("PG_DEBUG")

	c.Rabbit.Host = vip.GetString("RABBIT_HOST")
	c.Rabbit.Port = vip.GetInt("RABBIT_PORT")
	c.Rabbit.User = vip.GetString("RABBIT_USER")
	c.Rabbit.Pass = vip.GetString("RABBIT_PASS")
	c.Rabbit.VHost = vip.GetString("RABBIT_VHOST")

	c.Rabbit.ConsumerExchangeName = vip.GetString("RABBIT_CONSUMER_EXCHANGE_NAME")
	c.Rabbit.ConsumerQueueName = vip.GetString("RABBIT_CONSUMER_QUEUE_NAME")
	c.Rabbit.ConsumerRouteKey = vip.GetString("RABBIT_CONSUMER_ROUTE_KEY")

	c.Rabbit.ProducerQueueName = vip.GetString("RABBIT_PRODUCER_QUEUE_NAME")
	c.Rabbit.ProducerRouteKey = vip.GetString("RABBIT_PRODUCER_ROUTE_KEY")

	c.TaxCalculationService.Address = vip.GetString("API_ADDRESS")
	c.TaxCalculationService.AuthToken = vip.GetString("API_AUTHORIZATION_TOKEN")
	c.TaxCalculationService.Retry = vip.GetInt("ARI_RETRY")

	c.TaxCalculationParams.CompanyName = vip.GetString("COMPANY_NAME")
	c.TaxCalculationParams.ProductClass = vip.GetString("PRODUCT_CLASS")
	c.TaxCalculationParams.TransactionType = vip.GetString("TRANSACTION_TYPE")

	c.TaxCalculationParams.RegistrationIsoAU = vip.GetString("AU_ISO_REGISTRATION")
	c.TaxCalculationParams.RegistrationIsoUS = vip.GetString("US_ISO_REGISTRATION")
	c.TaxCalculationParams.RegistrationIsoEU = vip.GetString("EU_ISO_REGISTRATION")

	c.TaxCalculationParams.TaxIdUS = vip.GetString("TAX_ID_US")
	c.TaxCalculationParams.TaxIdAU = vip.GetString("TAX_ID_AU")
	c.TaxCalculationParams.TaxIdEU = vip.GetString("TAX_ID_EU")

	RegistrationAU := vip.GetStringSlice("AU_REGISTRATION")
	RegistrationUS := vip.GetStringSlice("US_REGISTRATION")
	RegistrationEU := vip.GetStringSlice("EU_REGISTRATION")

	c.TaxCalculationParams.RegistrationAU = map[string]struct{}{}
	c.TaxCalculationParams.RegistrationUS = map[string]struct{}{}
	c.TaxCalculationParams.RegistrationEU = map[string]struct{}{}

	for _, division := range RegistrationAU {
		c.TaxCalculationParams.RegistrationAU[division] = struct{}{}
	}
	for _, division := range RegistrationUS {
		c.TaxCalculationParams.RegistrationUS[division] = struct{}{}
	}
	for _, division := range RegistrationEU {
		c.TaxCalculationParams.RegistrationEU[division] = struct{}{}
	}

	if err := c.Validate(); err != nil {
		log.Fatalf("can't validate config, err %s", err.Error())
	}

	return c
}

func (p PostgresConfig) GetAddr() string {
	return p.Addr
}
func (p PostgresConfig) GetUser() string {
	return p.User
}
func (p PostgresConfig) GetPass() string {
	return p.Pass
}
func (p PostgresConfig) GetDB() string {
	return p.Database
}
func (p PostgresConfig) GetDebug() bool {
	return p.Debug
}
