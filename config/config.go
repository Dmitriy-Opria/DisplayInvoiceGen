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

type ExternalService struct {
	Address   string
	AuthToken string
}

type ExternalParams struct {
	CompanyName     string
	DivisionAU      map[string]struct{}
	DivisionUS      map[string]struct{}
	DivisionEU      map[string]struct{}
	TaxIdUS         string
	TaxIdAU         string
	TaxIdEU         string
	ProductClass    string
	TransactionType string
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

// Validate external service config structure
func (p ExternalService) Validate() error {
	return v.ValidateStruct(&p,
		v.Field(&p.Address, v.Required),
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

	c.TaxCalculationService.Address = vip.GetString("API_ADDRESS")
	c.TaxCalculationService.AuthToken = vip.GetString("API_AUTHORIZATION_TOKEN")

	c.TaxCalculationParams.CompanyName = vip.GetString("COMPANY_NAME")
	c.TaxCalculationParams.ProductClass = vip.GetString("PRODUCT_CLASS")
	c.TaxCalculationParams.TransactionType = vip.GetString("TRANSACTION_TYPE")

	c.TaxCalculationParams.TaxIdUS = vip.GetString("TAX_ID_US")
	c.TaxCalculationParams.TaxIdAU = vip.GetString("TAX_ID_AU")
	c.TaxCalculationParams.TaxIdEU = vip.GetString("TAX_ID_EU")

	divisionAU := vip.GetStringSlice("DIVISION_AU")
	divisionUS := vip.GetStringSlice("DIVISION_US")
	divisionEU := vip.GetStringSlice("DIVISION_EU")

	c.TaxCalculationParams.DivisionAU = map[string]struct{}{}
	c.TaxCalculationParams.DivisionUS = map[string]struct{}{}
	c.TaxCalculationParams.DivisionEU = map[string]struct{}{}

	for _, division := range divisionAU {
		c.TaxCalculationParams.DivisionAU[division] = struct{}{}
	}
	for _, division := range divisionUS {
		c.TaxCalculationParams.DivisionUS[division] = struct{}{}
	}
	for _, division := range divisionEU {
		c.TaxCalculationParams.DivisionEU[division] = struct{}{}
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
