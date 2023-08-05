package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	user_messages "backend/modules/api/endpoints/messages/user"
	"backend/x/identity"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var (
	DefaultUserRoles = []string{"member"}
)

type User struct {
	gorm.Model
	Handle        string `gorm:"UNIQUE;NOT NULL"`
	Karma         uint32 `gorm:"NOT NULL"`
	State         int32  `gorm:"NOT NULL"`
	DiscordUserID uint   `gorm:"NOT NULL"`
	DiscordUser   DiscordUser
	GithubUserID  *uint
	GithubUser    GithubUser
	Session       Session
	PlanID        uint `gorm:"NOT NULL"`
	Plan          Plan
	PlanGrantedAt time.Time `gorm:"NOT NULL"`
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
			return fmt.Errorf("take user: %w", err)
		}
		// User is created here, applying basic permissions & attaching discord user ID.
		user.DiscordUserID = discordUser.ID
		var roles []Role
		if err := dbConn.Where("name IN ?", DefaultUserRoles).Find(&roles).Error; err != nil {
			return fmt.Errorf("find roles: %w", err)
		}
		user.Roles = append(user.Roles, roles...)
	}
	// Sync handle with Discord's username.
	user.Handle = discordUser.Username
	var plan Plan
	planName := BasicPlan
	if discordUser.IsVIP {
		planName = VIPPlan
	}
	if err := dbConn.Where(&Plan{Name: planName}).Take(&plan).Error; err != nil {
		return fmt.Errorf("take plan: %w", err)
	}
	if plan.ID != user.PlanID {
		user.PlanID = plan.ID
		user.PlanGrantedAt = time.Now().UTC()
	}
	// Determine if we have enough info for the user.
	if !discordUser.HatechnologMember || !discordUser.EmailVerified {
		user.State = int32(user_messages.UserState_IN_CREATION)
	} else {
		user.State = int32(user_messages.UserState_CREATED)
	}
	if err := dbConn.Save(user).Error; err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}

func (user *User) FromHandle(dbConn *gorm.DB, handle string) error {
	return dbConn.Where(&User{
		Handle: handle,
	}).Preload("DiscordUser").Preload("GithubUser").Preload("Plan").Take(user).Error
}

func (user *User) GetSession(dbConn *gorm.DB, identity *identity.Identity) (*Session, []byte, error) {
	var session Session
	if err := dbConn.Where(&Session{
		UserID: user.ID,
	}).Take(&session).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("take session: %w", err)
		}
	} else {
		// An active session was found, but a new one requested, deleting it.
		if err := session.Invalidate(dbConn); err != nil {
			return nil, nil, fmt.Errorf("invalidate: %w", err)
		}
	}
	// Creating a new, fresh session.
	session.UserID = user.ID
	sessionPublicID, err := gonanoid.New(SessionTokenLength)
	if err != nil {
		return nil, nil, fmt.Errorf("new nanoid: %w", err)
	}
	session.PublicID = sessionPublicID
	sessionToken := make([]byte, 32)
	if _, err := rand.Read(sessionToken); err != nil {
		return nil, nil, fmt.Errorf("read: %w", err)
	}
	iv, encData, err := identity.GCMAESEncrypt(sessionToken)
	if err != nil {
		return nil, nil, fmt.Errorf("GCMAESEncrypt: %w", err)
	}
	session.IV = iv
	session.EncToken = encData
	session.Exipry = time.Now().UTC().Add(SessionLength)
	if err := dbConn.Save(&session).Error; err != nil {
		return nil, nil, fmt.Errorf("save session: %w", err)
	}
	return &session, sessionToken, nil
}
