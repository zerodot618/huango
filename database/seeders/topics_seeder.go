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

    seed.Add("SeedTopicsTable", func(db *gorm.DB) {

        topics  := factories.MakeTopics(10)

        result := db.Table("topics").Create(&topics)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}