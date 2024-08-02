package shared

type User struct {
	ID                   uint64 `gorm:"primaryKey"`
	Name                 string
	Email                string
	Role                 string `gorm:"default:user"`
	Avatar               string
	Password             string
	Lang                 string  `gorm:"default:en-US"`
	ContainerID          *string `gorm:"column:containerId"`
	TaskID               *string `gorm:"column:taskId"`
	TwoFaSecret          string  `gorm:"column:twoFaSecret"`
	LastOnline           int64   `gorm:"column:lastOnline"`
	CreatedAt            int64   `gorm:"column:createdAt"`
	TwoFaEnabled         bool    `gorm:"column:twoFaEnabled;default:false"`
	EmailConfirm         bool    `gorm:"column:emailConfirm;default:false"`
	Verified             bool    `gorm:"default:false"`
	Referral             *uint64
	AssetID              *uint64 `gorm:"column:assetId"`
	AvailableTerminal    bool    `gorm:"column:avaliableTerminal;default:false"`
	AvailableAssist      bool    `gorm:"column:avaliableAssist;default:false"`
	AvailableLending     bool    `gorm:"column:avaliableLending;default:false"`
	AvailablePortfolio   bool    `gorm:"column:avaliablePortfolio;default:false"`
	AvailableMarketplace bool    `gorm:"column:avaliableMarketplace;default:false"`
}

func (User) TableName() string {
	return "users"
}

type AdminPermission struct {
	ID                       uint64 `gorm:"primaryKey"`
	UserId                   uint64
	User                     User `gorm:"foreignKey:ID;references:UserId"`
	IsAvaliableTools         bool
	IsAvaliableTerminals     bool
	IsAvaliableUsers         bool
	IsAvaliableBalances      bool
	IsAvaliableDocumentation bool
	IsAvaliableInsights      bool
	IsAvaliableBalancer      bool
	IsAvaliableNews          bool
	IsAvaliableTwitter       bool
	IsAvaliableForex         bool
	IsAvaliableLanguages	 bool
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