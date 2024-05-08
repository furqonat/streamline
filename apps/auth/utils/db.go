package utils

import "apps/auth/db"

type Database struct {
	*db.PrismaClient
}

func NewDatabase(logger Logger) Database {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		logger.Fatalf("unable connect to database %s", err)
	}

	return Database{
		client,
	}
}
