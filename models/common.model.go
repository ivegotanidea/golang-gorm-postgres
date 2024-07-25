package models

//
//import (
//	"time"
//
//	"github.com/google/uuid"
//)
//
//type User2 struct {
//	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	Phone          string    `gorm:"unique;size:50"`
//	TelegramUserId int64     `gorm:"not null"`
//	Password       string    `gorm:"not null"`
//	Verified       bool      `gorm:"default:false"`
//	CreatedAt      time.Time `gorm:"not null"`
//	UpdatedAt      time.Time `gorm:"not null"`
//	Avatar         *string   `gorm:"type:varchar"`
//}
//
//type Profile struct {
//	ID     uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	UserID uuid.UUID `gorm:"foreignKey:UserID"
//`
//	Active    bool    `gorm:"default:false"`
//	Phone     string  `gorm:"unique;size:50"`
//	Name      string  `gorm:"type:varchar(30)"`
//	Age       int     `gorm:"not null"`
//	Height    int     `gorm:"not null"`
//	Weight    int     `gorm:"not null"`
//	Bust      float64 `gorm:"not null"`
//	Ethnos    string  `gorm:"type:varchar(30)"`
//	Bio       string  `gorm:"type:varchar(1000)"`
//	Moderated bool    `gorm:"default:false"`
//	Verified  bool    `gorm:"default:false"`
//
//	PriceInHouseContact int
//	PriceInHouseHour    int
//	PriceSaunaContact   int
//	PriceSaunaHour      int
//	PriceVisitContact   int
//	PriceVisitHour      int
//	PriceCarContact     int
//	PriceCarHour        int
//
//	Photos       []Photo `gorm:"foreignKey:ProfileID"`
//	ContactPhone string  `gorm:"type:varchar(30)"`
//	ContactWA    string  `gorm:"type:varchar(30)"`
//	ContactTG    string  `gorm:"type:varchar(30)"`
//
//	Tags []ProfileTag `gorm:"many2many:profile_tags"`
//}
//
//// Photo model definition
//type Photo struct {
//	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	ProfileID uuid.UUID `gorm:"type:uuid"` // Foreign key to Profile
//	URL       string    `gorm:"not null"`
//	Disabled  bool      `gorm:"default:false"`
//	CreatedAt time.Time `gorm:"not null"`
//}
//
//type Service struct {
//	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	CreatedAt time.Time `gorm:"not null"`
//
//	// client id
//	ClientUserID    uuid.UUID `gorm:"foreignKey:UserID"`
//	ClientLatitude  float32
//	ClientLongitude float32
//
//	// profile id
//	ProfileID        uuid.UUID `gorm:"foreignKey:ProfileID"`
//	ProfileLatitude  float32
//	ProfileLongitude float32
//
//	// user's score
//	ProfileRatingID uuid.UUID     `gorm:"type:uuid"`
//	ProfileRating   ProfileRating `gorm:"foreignKey:ProfileRatingID"`
//
//	// profile author's score
//	UserRatingID uuid.UUID  `gorm:"type:uuid"`
//	UserRating   UserRating `gorm:"foreignKey:UserRatingID"`
//}
//
//// rename to review
//type UserRating struct {
//	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	ServiceID uuid.UUID `gorm:"type:uuid;not null"` // Foreign key to Service
//	UserID    uuid.UUID `gorm:"foreignKey:UserID"`
//
//	Tags []UserTag `gorm:"many2many:rating_user_tags"`
//
//	Review string `gorm:"type:text"`
//
//	Score int `gorm:"not null;check:score >= 1 AND score <= 5"`
//
//	CreatedAt time.Time `gorm:"not null"`
//}
//
//type ProfileRating struct {
//	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	ServiceID uuid.UUID `gorm:"type:uuid;not null"` // Foreign key to Service
//	ProfileID uuid.UUID `gorm:"foreignKey:ProfileID"`
//
//	// todo add type: who rated
//
//	Tags []ProfileTag `gorm:"many2many:rating_profile_tags"`
//
//	Review string `gorm:"type:text"`
//
//	Score int `gorm:"not null;check:score >= 1 AND score <= 5"`
//
//	CreatedAt time.Time `gorm:"not null"`
//}
//
//type RatingProfileTags struct {
//	// following props are both primary keys
//	RatingID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
//	ProfileTagID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
//	Type         string    // like or dislike
//}
//
//type RatingUserTags struct {
//	// following props are both primary keys
//	RatingID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
//	UserTagID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
//	Type      string    // like or dislike
//}
//
//type ProfileTags struct {
//	// following props are both primary keys
//	ProfileID    uuid.UUID // foreign key
//	ProfileTagID uuid.UUID // foreign key
//	Price        int64     // nullable
//}
//
//type ProfileTag struct {
//	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	Name string    `gorm:"type:varchar(30);unique;not null"`
//}
//
//type UserTag struct {
//	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
//	Name string    `gorm:"type:varchar(30);unique;not null"`
//}
