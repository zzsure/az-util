package tool

import (
	"github.com/urfave/cli"
	"az-util/conf"
	"az-util/library/db"
	"az-util/library/log"
	"az-util/models"
)

var InitDB = cli.Command{
	Name:  "init",
	Usage: "az-util init db",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config.toml",
			Usage: "toml配置文件",
		},
		cli.StringFlag{
			Name:  "args",
			Value: "",
			Usage: "multi config cmd line args",
		},
	},
	Action: runInitDB,
}

func runInitDB(c *cli.Context) {
	conf.Init(c.String("conf"), c.String("args"))
	log.Init()
	db.Init()
	db.DB.LogMode(true)

	// TODO: 改为传参
	if true {
		models.MigrateTable()
	} else {
		models.CreateTable()
	}
}
