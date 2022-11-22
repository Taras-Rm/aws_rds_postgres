package setup

import (
	"github.com/Taras-Rm/aws_rds/config"
	"github.com/go-pg/pg/v10"
)

func GetPostgresConnect() *pg.DB {
	config := config.Config

	address := config.Database.Host + ":" + config.Database.Port

	//connectUrl := fmt.Sprintf("postgres://%s:%s@%s/%s", config.Database.User, config.Database.Password, address, config.Database.Name)

	connect := pg.Connect(&pg.Options{
		Addr:     address,
		User:     config.Database.User,
		Password: config.Database.Password,
		Database: config.Database.Name,
	})

	return connect
}
