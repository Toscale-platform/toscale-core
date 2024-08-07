package shared

import "database/sql"

type User struct {
	ID                   uint64 `gorm:"primaryKey"`
	Name                 string
	Email                string
	Role                 string
	Avatar               string
	Password             string
	CreatedAt            uint64 `gorm:"column:createdAt"`
	Referral             sql.Null[uint64]
	Lang                 string
	EmailConfirmed       bool             `gorm:"column:emailConfirm"`
	TwoFaSecret          string           `gorm:"column:twoFaSecret"`
	ContainerId          sql.NullString   `gorm:"column:containerId"`
	TaskId               sql.NullString   `gorm:"column:taskId"`
	TwoFaEnabled         bool             `gorm:"column:twoFaEnabled"`
	LastOnline           uint64           `gorm:"column:lastOnline"`
	Verified             bool             `gorm:"column:verified"`
	NewsLang             string           `gorm:"column:newslang"`
	AssetId              sql.Null[uint64] `gorm:"column:assetId"`
	AvaliableMarketplace bool             `gorm:"column:avaliableMarketplace"`
	AvaliablePortfolio   bool             `gorm:"column:avaliablePortfolio"`
	AvaliableTerminal    bool             `gorm:"column:avaliableTerminal"`
	AvaliableLending     bool             `gorm:"column:avaliableLending"`
	AvaliableAssist      bool             `gorm:"column:avaliableAssist"`
}

func (User) TableName() string {
	return "users"
}

type AdminPermission struct {
	ID                        uint64 `gorm:"primaryKey"`
	UserId                    uint64
	User                      User `gorm:"foreignKey:ID;references:UserId"`
	IsAvaliableTools          bool
	IsAvaliableTerminals      bool
	IsAvaliableUsers          bool
	IsAvaliableBackendTesting bool
	IsAvaliableDocumentation  bool
	IsAvaliableInsights       bool
	IsAvaliableBalancer       bool
	IsAvaliableNews           bool
	IsAvaliableTwitter        bool
	IsAvaliableForex          bool
	IsAvaliableLanguages      bool
}

type InstalledService struct {
	ID        uint64 `gorm:"primaryKey"`
	Position  int
	UserID    uint64 `gorm:"column:userId"`
	ServiceID string `gorm:"column:serviceId"`
}

func (InstalledService) TableName() string {
	return "installed_services"
}

type NotificationSettings struct {
	Name     string `gorm:"primaryKey"`
	Category string
	Event    string
}

func (NotificationSettings) TableName() string {
	return "notification_settings"
}

type NotificationUsersSettings struct {
	ID                        string `gorm:"primaryKey"`
	UserID                    uint64 `gorm:"column:userId"`
	NotificationsSettingsName string `gorm:"column:notificationsSettingName"`
	ServiceID                 string `gorm:"column:serviceId"`
	Web                       bool
	Telegram                  bool
	Mobile                    bool
}

func (NotificationUsersSettings) TableName() string {
	return "notifications_users_settings"
}

type NotificationTokens struct {
	ID       uint64 `gorm:"primaryKey"`
	Platform string
	Token    string
	State    bool
	UserID   uint64 `gorm:"column:userId"`
}

func (NotificationTokens) TableName() string {
	return "notification_tokens"
}

// toscale unified API apikey
type ApiKey struct {
	ID                uint64 `gorm:"primaryKey;autoIncrement:true"`
	Key               string
	UserID            uint64 `gorm:"uniqueIndex:idx_uniq_api_user"`
	User              User   `gorm:"foreignKey:ID;references:UserID"`
	TokensUsedMinute  uint64 `gorm:"default:0"`
	TotalTokensMinute uint64 `gorm:"default:30"`
	TokensUsedMonth   uint64 `gorm:"default:0"`
	TotalTokensMonth  uint64 `gorm:"default:10000"`
	ExpirityDate      uint64
}
