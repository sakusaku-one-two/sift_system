package database



import (
	_ "backend/util"
	"gorm.io/driver/postgress"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"fmt"
)


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 


//トランザクションを実施するためのコールバック関数の型定義
type TransactionCallBack func(tx *grom.Transaction) error


type DBManager interface {
	
	SetConfig(db_config DBConfig)
	GetConfig() DBConfig
	SetUpConnection()error
	NewConnection(dns string) (*gorm.DB,error)
	Migrate(arg_modles ...any) error
	NewTransaction(t_callback TransactionCallBack)
	Close() error
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


type DBConfig interface {
	func DnsSetup() string,bool
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type BaseConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
	TimeZone string
}



func (bc *BaseConfig) DnsSetup() (string,error) {
	
	config_values := map[string]any{
	    "host"     :bc.Host,
		"port"     :bc.Port,
		"user"     :bc.User,
		"passowrd" :bc.Password,
		"db_name"  :bc.DBName,
		"ssl_mode" :bc.SSLMode,
		"time_zone":bc.TimeZone,
	}
	
	for key,value := range config_values {
		if value == "" {
			return	"" ,fmt.Errorf("DB設定値に不備があります %v",key)
		}
	}


	return fmt.Sprintf("",config_values),nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~