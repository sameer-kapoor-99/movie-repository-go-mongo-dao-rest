package factory

import (
	"../interfaces"
	"../mongodb"
	"log"
)

func FactoryDao(e string) interfaces.UserDao {
	var i interfaces.UserDao
	switch e {
	case "mongodb":
		i = mongodb.UserImplMongodb{}
	default:
		log.Fatalf("Unsupported %s database", e)
		return nil
	}

	return i
}
