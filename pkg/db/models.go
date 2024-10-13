// pkg/db/models.go
package db

type WordPair struct {
	ID     uint   `gorm:"primaryKey"`
	UserID int64  `gorm:"index"` // To keep pairs separate for each user
	Word1  string `gorm:"not null"`
	Word2  string `gorm:"not null"`
}

type UserSettings struct {
	ID              uint  `gorm:"primaryKey"`
	UserID          int64 `gorm:"index"`
	PairsToSend     int   `gorm:"default:1"` // Default to sending 1 pair
	RemindersPerDay int   `gorm:"default:1"` // Default to 1 reminder per day
}
