package shared

type OrganizedTickers struct {
	ID               uint64           `json:"id" gorm:"primaryKey"`
	Exchange         string           `json:"-" gorm:"uniqueIndex:idx_uniq_ticker"`
	ExchangeID       int64            `json:"exchangeId"`
	Symbol           string           `json:"symbol" gorm:"uniqueIndex:idx_uniq_ticker"`
	Timestamp        uint64           `json:"timestamp"`
	Datetime         string           `json:"datetime"`
	High             float64          `json:"high"`
	Low              float64          `json:"low"`
	Bid              float64          `json:"bid"`
	BidVolume        float64          `json:"bidVolume"`
	Ask              float64          `json:"ask"`
	AskVolume        float64          `json:"askVolume"`
	Vwap             float64          `json:"vwap"`
	Open             float64          `json:"open"`
	Close            float64          `json:"close"`
	Last             float64          `json:"last"`
	PreviousClose    float64          `json:"previousClose"`
	Change           float64          `json:"change"`
	Percentage       float64          `json:"percentage"`
	BaseVolume       float64          `json:"baseVolume"`
	QuoteVolume      float64          `json:"quoteVolume"`
	TickerAssetsInfo TickerAssetsInfo `json:"tickerAssetsInfo" gorm:"foreignKey:TickerId;references:ID"`
	PairStatus       PairStatus       `json:"pairStatus" gorm:"foreignKey:TickerId;references:ID"`
	CreatedAt        uint64           `json:"createdAt" gorm:"autoCreateTime:milli"`
}

type TickerAssetsInfo struct {
	TickerId       uint64 `gorm:"uniqueIndex:idx_uniq_ticker_id"`
	BaseAsset      string `json:"baseAsset"`
	BaseAssetId    uint64 `json:"baseAssetId"`
	QuoteAsset     string `json:"quoteAsset"`
	QuoteAssetId   uint64 `json:"quoteAssetId"`
	Symbol         string `json:"symbol"`
	Exclude        bool   `json:"exclude"`
	OnlyExcluded   bool   `json:"-" gorm:"-"`
	ExchangeSymbol string `json:"-" gorm:"-"`
}

type PairStatus struct {
	Status          string
	TickerId        uint64
	Exchange        string
	ExchangeId      int64  `gorm:"uniqueIndex:pair_status_idx"`
	BaseAssetId     uint64 `gorm:"uniqueIndex:pair_status_idx"`
	QuoteAssetId    uint64 `gorm:"uniqueIndex:pair_status_idx"`
	Symbol          string `gorm:"uniqueIndex:pair_status_idx"`
	MarketType      string `gorm:"uniqueIndex:pair_status_idx"`
	StatusUpdatedAt uint64 `gorm:"autoUpdateTime:milli"`
	CreatedAt       uint64
	Description     string
}

type Candle struct {
	Resolution   int
	Open         float64
	Close        float64
	High         float64
	Low          float64
	VolumeBase   float64
	VolumeQuote  float64
	Time         int64
	IsBuyerMaker bool
}