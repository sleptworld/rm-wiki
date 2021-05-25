package DB

import (
	"context"
	"github.com/sleptworld/test/Config"
	"github.com/sleptworld/test/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	once   sync.Once
	DBtool *dbt
	Db     *gorm.DB
	errH   error
)

type dbt struct {
	Dsn     string
	Config  gorm.Config
	Debug   bool
	Plugins []gorm.Plugin
	Context	context.Context
}

func GetDbInstance() *dbt {
	once.Do(func() {
		DBtool = &dbt{
			Dsn:     Config.Dsn,
			Config:  gorm.Config{},
			Debug:   Config.DBDebug,
			Plugins: Config.DBPlugins,
			Context: nil,
		}

	})

	return DBtool
}

func (m *dbt) InitDBPool() (*gorm.DB, error) {
	dsn := postgres.Open(m.Dsn)
	Db, errH = gorm.Open(dsn, &m.Config)

	if errH != nil {
		return nil, errH
	} else {
		for _,p := range m.Plugins{
			err := Db.Use(p)
			if err != nil {
				return nil, err
			}
		}
		if m.Context != nil{
			Db = Db.WithContext(m.Context)
		}
		if m.Debug {
			Db = Db.Debug()
		}
		return Db, nil
	}
}

func InitDB	() error{
	err := Db.AutoMigrate(&Lang{}, &UserGroup{},&User{}, &Draft{}, &Cat{}, &Entry{}, &History{}, &Tag{})

	Db.Exec("CREATE INDEX path_gist_idx ON cats USING gist(path)")
	Db.Exec("CREATE INDEX path_idx ON cats USING btree(path)")

	// init
	langs := []Lang{
		{Lang: "zh-CN"},
		{Lang: "zh-TW"},
	}

	groups := []UserGroup{
		{GroupName: "Admin",Level: 3},
		{GroupName: "Anonymous",Level: 0},
	}

	for _,l := range langs{
		Db.Where(l).FirstOrCreate(&Lang{})
	}

	for _,g := range groups{
		Db.Where(g).FirstOrCreate(&UserGroup{})
	}

	CreateCat(Db,"root",nil)
	anonymous := User{
		Name:        "Anonymous",
		Email:       "empty",
		UserGroupID: 2,
	}
	UserPretreatment(&anonymous,tools.GetRandomString(16))

	_,err = RegisterUser(Db,&anonymous,&User{})
	return err
}