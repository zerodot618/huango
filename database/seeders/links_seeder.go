package seeders

import (
	"fmt"
	"zerodot618/huango/database/factories"
	"zerodot618/huango/pkg/console"
	"zerodot618/huango/pkg/logger"
	"zerodot618/huango/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedLinksTable", func(db *gorm.DB) {

		links := factories.MakeLinks(5)

		result := db.Table("links").Create(&links)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
