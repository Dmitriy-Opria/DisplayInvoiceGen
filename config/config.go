package config

import (
	"bytes"
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"io/ioutil"
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
	SalesForce            SalesForceConfig
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
	Host         string
	Port         int
	User         string
	Pass         string
	VHost        string
	ExchangeName string

	ConsumerInvoiceQueueName string
	ConsumerInvoiceRouteKey  string

	ConsumerSFQueueName string
	ConsumerSFRouteKey  string

	ProducerPDFQueueName string
	ProducerPDFRouteKey  string

	ProducerSFQueueName string
	ProducerSFRouteKey  string
}

type ExternalService struct {
	Address   string
	AuthToken string
	Retry     int
}

type SalesForceConfig struct {
	ApiVersion    string
	ClientID      string
	ClientSecret  string
	UserName      string
	Pass          string
	SecurityToken string
	Retry         int
	RecordTypeId  string

	TickerPeriod     int64
	ValidationPeriod int64
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
		v.Field(&c.SalesForce),
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
		v.Field(&r.ExchangeName, v.Required),
		v.Field(&r.ConsumerInvoiceQueueName, v.Required),
		v.Field(&r.ConsumerInvoiceRouteKey, v.Required),
		v.Field(&r.ConsumerSFQueueName, v.Required),
		v.Field(&r.ConsumerSFRouteKey, v.Required),
		v.Field(&r.ProducerPDFQueueName, v.Required),
		v.Field(&r.ProducerPDFRouteKey, v.Required),
		v.Field(&r.ProducerSFQueueName, v.Required),
		v.Field(&r.ProducerSFRouteKey, v.Required),
	)
}

// Validate external service config structure
func (p ExternalService) Validate() error {
	return v.ValidateStruct(&p,
		v.Field(&p.Address, v.Required),
		v.Field(&p.Retry, v.Required),
	)
}

// Validate external service config structure
func (p SalesForceConfig) Validate() error {
	return v.ValidateStruct(&p,
		v.Field(&p.ApiVersion, v.Required),
		v.Field(&p.ClientID, v.Required),
		v.Field(&p.ClientSecret, v.Required),
		v.Field(&p.UserName, v.Required),
		v.Field(&p.Pass, v.Required),
		v.Field(&p.SecurityToken, v.Required),
		v.Field(&p.RecordTypeId, v.Required),
		v.Field(&p.TickerPeriod, v.Required),
		v.Field(&p.ValidationPeriod, v.Required),
	)
}

// InitConfig initialize application configuration and validate all fields
func InitConfig() *Config {
	body, err := ioutil.ReadFile("config.yaml")
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
	c.Rabbit.ExchangeName = vip.GetString("RABBIT_EXCHANGE_NAME")

	c.Rabbit.ConsumerInvoiceQueueName = vip.GetString("RABBIT_CONSUMER_INVOICE_QUEUE_NAME")
	c.Rabbit.ConsumerInvoiceRouteKey = vip.GetString("RABBIT_CONSUMER_INVOICE_ROUTE_KEY")

	c.Rabbit.ConsumerSFQueueName = vip.GetString("RABBIT_CONSUMER_SF_QUEUE_NAME")
	c.Rabbit.ConsumerSFRouteKey = vip.GetString("RABBIT_CONSUMER_SF_ROUTE_KEY")

	c.Rabbit.ProducerPDFQueueName = vip.GetString("RABBIT_PRODUCER_PDF_QUEUE_NAME")
	c.Rabbit.ProducerPDFRouteKey = vip.GetString("RABBIT_PRODUCER_PDF_ROUTE_KEY")

	c.Rabbit.ProducerSFQueueName = vip.GetString("RABBIT_PRODUCER_SF_QUEUE_NAME")
	c.Rabbit.ProducerSFRouteKey = vip.GetString("RABBIT_PRODUCER_SF_ROUTE_KEY")

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

	c.SalesForce.ApiVersion = vip.GetString("SALES_FORCE_VERSION")
	c.SalesForce.ClientID = vip.GetString("SALES_FORCE_CLIENT_ID")
	c.SalesForce.ClientSecret = vip.GetString("SALES_FORCE_CLIENT_SECRET")
	c.SalesForce.UserName = vip.GetString("SALES_FORCE_USER_NAME")
	c.SalesForce.Pass = vip.GetString("SALES_FORCE_PASSWORD")
	c.SalesForce.SecurityToken = vip.GetString("SALES_FORCE_SECURITY_TOKEN")
	c.SalesForce.RecordTypeId = vip.GetString("RECORD_TYPE_ID")
	c.SalesForce.TickerPeriod = vip.GetInt64("TICKER_PERIOD_SEC")
	c.SalesForce.ValidationPeriod = vip.GetInt64("VALIDATION_PERIOD_MIN")

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
