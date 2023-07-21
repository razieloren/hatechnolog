package models

import (
	"errors"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var (
	BasicPlan        = "basic"
	SupporterPlan    = "supporter"
	VIPPlan          = "vip"
	DefaultUserRoles = []string{"member"}
)

type User struct {
	gorm.Model
	Handle               string `gorm:"UNIQUE;NOT NULL"`
	Karma                uint32 `gorm:"NOT NULL"`
	AllowMarketingEmails bool   `gorm:"NOT NULL"`
	AllowWeeklyDigest    bool   `gorm:"NOT NULL"`
	TACAcceptedAt        *time.Time
	DiscordUserID        uint `gorm:"NOT NULL"`
	DiscordUser          DiscordUser
	GithubUserID         *uint
	GithubUser           GithubUser
	Session              Session
	PlanID               uint `gorm:"NOT NULL"`
	Plan                 Plan
	PlanGrantedAt        time.Time `gorm:"NOT NULL"`
	// Attached user's roles.
	Roles []Role `gorm:"many2many:api.user_role"`
	// User-specific permissions, this should be avoided.
	Permissions []Permission `gorm:"many2many:api.user_permission"`
}

func (user *User) TableName() string {
	return "api.user"
}

func (user *User) FromDiscordUser(dbConn *gorm.DB, discordUser *DiscordUser) error {
	if err := dbConn.Where(&User{
		DiscordUserID: discordUser.ID,
	}).Take(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("error fetching user: %w", err)
		}
		// User is created here, applying basic permissions & attaching discord user ID.
		user.DiscordUserID = discordUser.ID
		var roles []Role
		if err := dbConn.Where("name IN ?", DefaultUserRoles).Find(&roles).Error; err != nil {
			return fmt.Errorf("error fetching default roles: %w", err)
		}
		user.Roles = append(user.Roles, roles...)
	}
	// Sync handle with Discord's username.
	user.Handle = discordUser.Username
	var plan Plan
	planName := BasicPlan
	if discordUser.IsVIP {
		planName = VIPPlan
	} else if discordUser.IsSupporter {
		planName = SupporterPlan
	}
	if err := dbConn.Where(&Plan{Name: planName}).Take(&plan).Error; err != nil {
		return fmt.Errorf("error fetching plan type: %w", err)
	}
	if plan.ID != user.PlanID {
		user.PlanID = plan.ID
		user.PlanGrantedAt = time.Now().UTC()
	}
	if err := dbConn.Save(user).Error; err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}
	return nil
}

func (user *User) GetSession(dbConn *gorm.DB) (*Session, error) {
	var session Session
	if err := dbConn.Where(&Session{
		UserID: user.ID,
	}).Take(&session).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("error finding active session: %w", err)
		}
	} else {
		// An active session was found, but a new one requested, deleting it.
		if err := session.Invalidate(dbConn); err != nil {
			return nil, fmt.Errorf("error invalidating session: %w", err)
		}
	}
	// Creating a new, fresh session.
	session.UserID = user.ID
	sessionToken, err := gonanoid.New(SessionTokenLength)
	if err != nil {
		return nil, fmt.Errorf("error generating new session token")
	}
	session.Token = sessionToken
	session.Exipry = time.Now().UTC().Add(SessionLength)
	if err := dbConn.Save(&session).Error; err != nil {
		return nil, fmt.Errorf("error saving new session: %w", err)
	}
	return &session, nil
}
