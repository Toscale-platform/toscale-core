package shared

import "database/sql"

type Tabler interface {
	TableName() string
}

type Asset struct {
	ID                      uint64 `gorm:"primaryKey;autoIncrement:true"`
	Symbol                  string `gorm:"default:''"`
	Name                    string `gorm:"default:''"`
	Slug                    string `gorm:"uniqueIndex:toscale_asset_slug"`
	CryptorankName          string `gorm:"uniqueIndex:crypto_unique"`
	DescriptionRu           string `gorm:"default:''"`
	DescriptionEn           string `gorm:"default:''"`
	Whitepaper              string `gorm:"default:''"`
	Website                 string `gorm:"default:''"`
	Watchlist               uint64 `gorm:"default:0"`
	CreatedAt               uint64 `gorm:"deault:0"`
	Price                   float64
	Marketcap               float64
	Volume24h               float64
	High24h                 float64
	Low24h                  float64
	PercentChange24h        float64
	PercentChange7d         float64
	Category                string `gorm:"default:''"`
	Rank                    uint64
	CirculatingSupply       float64
	MaxSupply               float64
	TotalSupply             float64
	FDV                     float64
	CryptorankId            uint64                   `gorm:"uniqueIndex:cryptorank_unique"`
	CoinmarketcapId         uint64                   `gorm:"uniqueIndex:coinmarketcap_unique"`
	CoingeckoId             sql.NullString           `gorm:"uniqueIndex:coingecko_unique"`
	BuyLink                 string                   `gorm:"default:''"`
	CountryTags             []CountryTag             `gorm:"foreignKey:EntityID;references:ID"`
	Image60                 string                   `gorm:"default:''"`
	IsActive                bool                     `gorm:"default:true"`
	AssetAlternateNames     []AssetAlternateName     `gorm:"foreignKey:AssetID;references:ID"`
	TypeTags                []TypeTag                `gorm:"foreignKey:EntityID;references:ID"`
	ConsensusAlgorithmTags  []ConsensusAlgorithmTag  `gorm:"foreignKey:EntityID;references:ID"`
	EncryptionAlgorithmTags []EncryptionAlgorithmTag `gorm:"foreignKey:EntityID;references:ID"`
	SelfReportedTags        []SelfReportedTag        `gorm:"foreignKey:EntityID;references:ID"`
	SocialMediaTags         []SocialMediaTag         `gorm:"foreignKey:EntityID;references:ID"`
	AuditTags               []AuditTag               `gorm:"foreignKey:EntityID;references:ID"`
	IndustryTags            []IndustryTag            `gorm:"foreignKey:EntityID;references:ID"`
	ContractTags            []ContractTag            `gorm:"foreignKey:EntityID;references:ID"`
}

func (Asset) TableName() string {
	return "currencies"
}

type CountryTag struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID   uint64 `gorm:"uniqueIndex:idx_country_tag"`
	EntityType string `gorm:"uniqueIndex:idx_country_tag"`
	Country    Country
	CountryID  uint64 `gorm:"uniqueIndex:idx_country_tag"`
	IsActive   bool   `gorm:"default:true"`
}

type AssetAlternateName struct {
	ID                    uint64 `gorm:"primaryKey;autoIncrement:true"`
	AssetID               uint64 `gorm:"uniqueIndex:alt_symbol_unique"`
	Exchange              string `gorm:"uniqueIndex:alt_symbol_unique"`
	AltName               string `gorm:"uniqueIndex:alt_symbol_unique"`
	IsActive              bool   `gorm:"default:true"`
	ExcludeFromAggregates bool   `gorm:"default:false"`
}

type TypeTag struct {
	ID                 uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID           uint64 `gorm:"uniqueIndex:idx_asset_type"`
	EntityType         string `gorm:"uniqueIndex:idx_asset_type"`
	TypeExtendedInfo   TypeExtendedInfo
	TypeExtendedInfoID uint64 `gorm:"uniqueIndex:idx_asset_type"`
	IsActive           bool   `gorm:"default:true"`
}

type ConsensusAlgorithmTag struct {
	ID                      uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID                uint64 `gorm:"uniqueIndex:idx_asset_consensus"`
	EntityType              string `gorm:"uniqueIndex:idx_asset_consensus"`
	ConsensusExtendedInfoID uint64 `gorm:"uniqueIndex:idx_asset_consensus"`
	ConsensusExtendedInfo   ConsensusExtendedInfo
	IsActive                bool `gorm:"default:true"`
}

type EncryptionAlgorithmTag struct {
	ID                       uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID                 uint64 `gorm:"uniqueIndex:idx_asset_enc"`
	EntityType               string `gorm:"uniqueIndex:idx_asset_enc"`
	EncryptionExtendedInfoID uint64 `gorm:"uniqueIndex:idx_asset_enc"`
	EncryptionExtendedInfo   EncryptionExtendedInfo
	IsActive                 bool `gorm:"default:true"`
}

type SelfReportedTag struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID   uint64 `gorm:"uniqueIndex:idx_asset_self"`
	EntityType string `gorm:"uniqueIndex:idx_asset_self"`
	Name       string `gorm:"uniqueIndex:idx_asset_self"`
	IsActive   bool   `gorm:"default:true"`
}

type SocialMediaTag struct {
	ID                        uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID                  uint64 `gorm:"uniqueIndex:idx_asset_socm"`
	EntityType                string `gorm:"uniqueIndex:idx_asset_socm"`
	SocialMediaExtendedInfoID uint64 `gorm:"uniqueIndex:idx_asset_socm"`
	SocialMediaExtendedInfo   SocialMediaExtendedInfo
	Link                      string `gorm:"uniqueIndex:idx_asset_socm"`
	IsActive                  bool   `gorm:"default:true"`
	SourceID                  *string
}

type AuditTag struct {
	ID                  uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID            uint64 `gorm:"uniqueIndex:idx_asset_audit"`
	EntityType          string `gorm:"uniqueIndex:idx_asset_audit"`
	AuditExtendedInfoID uint64 `gorm:"uniqueIndex:idx_asset_audit"`
	AuditExtendedInfo   AuditExtendedInfo
	AuditReport         string `gorm:"uniqueIndex:idx_asset_audit"`
	IsActive            bool   `gorm:"default:true"`
}

type IndustryTag struct {
	ID                     uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID               uint64 `gorm:"uniqueIndex:idx_asset_icat"`
	EntityType             string `gorm:"uniqueIndex:idx_asset_icat"`
	CategoryExtendedInfoId uint64 `gorm:"uniqueIndex:idx_asset_icat"`
	CategoryExtendedInfo   CategoryExtendedInfo
	IsActive               bool `gorm:"default:true"`
}

type ContractTag struct {
	ID                       uint64 `gorm:"primaryKey;autoIncrement:true"`
	EntityID                 uint64
	EntityType               string `gorm:"uniqueIndex:idx_asset_cont"`
	BlockchainExtendedInfoID uint64 `gorm:"uniqueIndex:idx_asset_cont"`
	BlockchainExtendedInfo   BlockchainExtendedInfo
	ContractAddress          string `gorm:"uniqueIndex:idx_asset_cont"`
	IsActive                 bool   `gorm:"default:true"`
}

type BlockchainExtendedInfo struct {
	ID                      uint64 `gorm:"primaryKey;autoIncrement:true"`
	Rank                    uint64
	Name                    string         `gorm:"uniqueIndex:idx_uniq_blockchain"`
	Slug                    *string        `gorm:"uniqueIndex:idx_uniq_blockchain_slug"`
	CoinmarketcapId         *uint64        `gorm:"uniqueIndex:idx_uniq_blockchain_coinmarketcap"`
	CoinmarketcapName       sql.NullString `gorm:"uniqueIndex:idx_uniq_blockchain_coinmarketcap_name"`
	CoingeckoId             sql.NullString `gorm:"uniqueIndex:idx_uniq_blockchain_coingecko"`
	LamaName                sql.NullString `gorm:"uniqueIndex:idx_uniq_lama_name"`
	Symbol                  string
	NativeToken             string
	Description             string
	DescriptionEn           string
	DescriptionRu           string
	Algorithm               string
	Protocols               string
	Layer                   string
	ActiveUsers             uint64
	TVL                     float64
	Stables                 float64
	Volume24h               float64
	Fees24h                 float64
	McapTVL                 float64
	LogoLink                string
	ExplorerLink            string
	IsActive                bool `gorm:"default:true"`
	IgnoreCase              bool `gorm:"default:false"`
	AssetID                 sql.NullInt64
	Asset                   Asset
	ConsensusAlgorithmTags  []ConsensusAlgorithmTag  `gorm:"foreignKey:EntityID;references:ID"`
	EncryptionAlgorithmTags []EncryptionAlgorithmTag `gorm:"foreignKey:EntityID;references:ID"`
}

func (BlockchainExtendedInfo) TableName() string {
	return "blockchain_extended_info"
}

type ProtocolExtendedInfo struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	Rank        uint64
	Name        string `gorm:"uniqueIndex:idx_uniq_protocol"`
	Symbol      string
	NativeToken string
	URL         string
	Description string
	Category    string
	Chains      string
	ChainTvls   string
	Audits      string
	TVL         float64
	Change1d    float64
	Change7d    float64
	Change1m    float64
	Mcap        float64
	LogoLink    string
}

func (ProtocolExtendedInfo) TableName() string {
	return "protocol_extended_info"
}

type Forex struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Ticker string `gorm:"uniqueIndex:idx_uniq_ticker"`
	Bid    float64
	Ask    float64
	Open   float64
	Low    float64
	High   float64
	Close  float64
	Date   int64
	Type   string `gorm:"-" json:"type"`
}

type ForexAsset struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	Symbol      string `gorm:"uniqueIndex:idx_uniq_forex_symbol"`
	Name        string
	Logo        string
	LanguageID  uint64       `gorm:"column:lang_id"`
	Language    Language     `gorm:"foreignKey:LanguageID;references:ID"`
	CountryTags []CountryTag `gorm:"foreignKey:EntityID;references:ID"`
	IsActive    bool
	Type        string
	Grapheme    string `gorm:"default:''"`
}

type Language struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"uniqueIndex:idx_uniq_lang" json:"name"`
	Lang     string `json:"lang"`
	Iso      string `json:"iso"`
	LogoLink string `json:"logo"`
}

type SocialMediaExtendedInfo struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name     string `gorm:"uniqueIndex:idx_uniq_social_media"`
	LogoLink string
}

func (SocialMediaExtendedInfo) TableName() string {
	return "social_media_extended_info"
}

type AuditExtendedInfo struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name     string `gorm:"uniqueIndex:idx_uniq_audit"`
	Link     string
	LogoLink string
}

func (AuditExtendedInfo) TableName() string {
	return "audit_extended_info"
}

type ConsensusExtendedInfo struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name   string
	Symbol string `gorm:"uniqueIndex:idx_uniq_consensus"`
}

func (ConsensusExtendedInfo) TableName() string {
	return "consensus_extended_info"
}

type EncryptionExtendedInfo struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name   string
	Symbol string `gorm:"uniqueIndex:idx_uniq_encryption"`
}

func (EncryptionExtendedInfo) TableName() string {
	return "encryption_extended_info"
}

type TypeExtendedInfo struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"uniqueIndex:idx_uniq_type"`
}

func (TypeExtendedInfo) TableName() string {
	return "type_extended_info"
}

type IndustryExtendedInfo struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	IndustryName string `gorm:"uniqueIndex:idx_uniq_industry"`
}

func (IndustryExtendedInfo) TableName() string {
	return "industry_extended_info"
}

type CategoryExtendedInfo struct {
	ID                     uint64               `gorm:"primaryKey;autoIncrement:true"`
	IndustryExtendedInfoID uint64               `gorm:"uniqueIndex:idx_uniq_industry"`
	IndustryExtendedInfo   IndustryExtendedInfo `gorm:"foreignKey:IndustryExtendedInfoID;references:ID"`
	CategoryName           string               `gorm:"uniqueIndex:idx_uniq_category"`
}

func (CategoryExtendedInfo) TableName() string {
	return "category_extended_info"
}

type Exchange struct {
	ID               uint64 `gorm:"primaryKey;autoIncrement:true"`
	CryptorankName   string `gorm:"default:'';uniqueIndex:exchanges_crypto_unique"`
	GroupId          string
	Name             string `gorm:"uniqueIndex:exchanges_name_unique"`
	Slug             string `gorm:"uniqueIndex:toscale_exchange_slug"`
	Volume24h        float64
	Volume7d         float64
	Volume30d        float64
	SpotPairs        uint64 `gorm:"default:0"`
	MarginPairs      uint64 `gorm:"default:0"`
	FuturesPairs     uint64 `gorm:"default:0"`
	DeliveryPairs    uint64 `gorm:"default:0"`
	IsHidden         bool   `gorm:"default:false"`
	ParentExchangeID uint64
	InstrumentID     uint64           `gorm:"default:1"`
	DescriptionRu    string           `gorm:"default:''"`
	DescriptionEn    string           `gorm:"default:''"`
	Website          string           `gorm:"default:''"`
	FeesLink         string           `gorm:"default:''"`
	ReferalLink      string           `gorm:"default:''"`
	AuditTags        []AuditTag       `gorm:"foreignKey:EntityID;references:ID"`
	SocialMediaTags  []SocialMediaTag `gorm:"foreignKey:EntityID;references:ID"`
	AssetID          uint64
	CountryTags      []CountryTag `gorm:"foreignKey:EntityID;references:ID"`
	CreatedAt        uint64       `gorm:"deault:0"`
	Technology       string       `gorm:"default:'CeFi'"`
	Logo             string       `gorm:"default:''"`
	Timezone         string
	Proof            float64 `gorm:"deault:0"`
	Instrument       Instrument
	Asset            Asset
	Currencies       int         `gorm:"deault:0"`
	Children         []*Exchange `gorm:"foreignKey:ParentExchangeID"`
}

type Instrument struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"uniqueIndex:idx_name"`
}

type WatchList struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID   uint64 `gorm:"uniqueIndex:idx_watch_list_uniq"`
	EntityID uint64 `gorm:"uniqueIndex:idx_watch_list_uniq"`
	Type     string `sql:"type:ENUM('asset', 'pair')" gorm:"column:type"`
	Flag     uint64 `gorm:"default:0"`
}

type HistoricTvlAll struct {
	Timestamp         uint64 `gorm:"uniqueIndex:idx_historic_name"`
	Name              string `gorm:"uniqueIndex:idx_historic_name"`
	Type              string
	TotalUsdLiquidity float64
}

type HistoricTvl struct {
	Timestamp         uint64 `gorm:"primaryKey"`
	TotalUsdLiquidity float64
}

type Country struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name    string `gorm:"uniqueIndex:idx_name_uniq"`
	Picture string
}

type AssetVote struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserId    uint64 `gorm:"uniqueIndex:idx_vote_uniq"`
	AssetId   uint64 `gorm:"uniqueIndex:idx_vote_uniq"`
	Positive  bool
	Timestamp int64
}

type Resolution struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Timeframe  string `gorm:"uniqueIndex:idx_reso_uniq" json:"timeframe"`
	Resolution uint64 `gorm:"uniqueIndex:idx_reso_uniq" json:"resolution"`
}

type ResearchPair struct {
	Symbol         string   `json:"symbol" gorm:"uniqueIndex:pair_idx"`
	BaseAssetId    uint64   `json:"baseAssetId" gorm:"uniqueIndex:pair_idx"`
	QuoteAssetId   uint64   `json:"quoteAssetId" gorm:"uniqueIndex:pair_idx"`
	BaseAssetType  string   `json:"baseAssetType" gorm:"uniqueIndex:pair_idx"`
	QuoteAssetType string   `json:"qouteAssetType" gorm:"uniqueIndex:pair_idx"`
	BaseAssetSlug  string   `json:"baseAssetSlug"`
	QuoteAssetSlug string   `json:"quoteAssetSlug"`
	MarketType     string   `json:"marketType"`
	ExchangeID     uint64   `gorm:"uniqueIndex:pair_idx"`
	Exchange       Exchange `gorm:"foreignKey:ExchangeID;references:ID"`
	Active         bool     `json:"active"`
}
