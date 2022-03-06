package DB

import (
	"fita/model/entity"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func NewConnectionDB(driverDB string, database string, host string, user string, password string, port int) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
	)

	dialect := mysql.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(entity.Products{})
	_ = db.AutoMigrate(entity.Promotion{})
	_ = db.AutoMigrate(entity.Order{})
	

	// Create Product

	// product := []entity.Products{
	// 	{
	// 		SKU: "120P90",
	// 		Name: "Google Home",
	// 		Price: 49.99,
	// 		Qty: 10,
	// 	},
	// 	{
	// 		SKU: "43N23P",
	// 		Name: "Macbook Pro",
	// 		Price: 5399.99,
	// 		Qty: 5,
	// 	},
	// 	{
	// 		SKU: "A304SD",
	// 		Name: "Alexa Speaker",
	// 		Price: 109.50,
	// 		Qty: 10,
	// 	},
	// 	{
	// 		SKU: "120P90",
	// 		Name: "Raspberry Pi B",
	// 		Price: 30.00,
	// 		Qty: 2,
	// 	},	
	// }

	// promo := []entity.Promotion{
	// 	{
	// 		ProductID: "43N23P",
	// 		TypePromo: "item",
	// 		Item: "120P90",
	// 		Amount: 0,
	// 		Qty: 1,
	// 	},
	// 	{
	// 		ProductID: "120P90",
	// 		TypePromo: "amount",
	// 		Item: "",
	// 		Amount: 49.99,
	// 		Qty: 3,
	// 	},
	// 	{
	// 		ProductID: "A304SD",
	// 		TypePromo: "percentage",
	// 		Item: "",
	// 		Amount: 0,
	// 		Qty: 3,
	// 		Percentage: 10,
	// 	},
	// }
	// db.Create(&product)
	// db.Create(&promo)


	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	//pool time
	tm := time.Minute * time.Duration(20)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(tm)

	return db, nil

}
func DbInit() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}

	DB, err = NewConnectionDB(viper.GetString("database.driver"), viper.GetString("database.schema"),
		viper.GetString("database.hostname"), viper.GetString("database.username"), viper.GetString("database.password"),
		viper.GetInt("database.port"))
	if err != nil {
		return err
	}

	return nil
}