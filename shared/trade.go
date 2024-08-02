package shared

import (
	"errors"

	"gorm.io/gorm"
)

func (ticker *OrganizedTickers) BeforeCreate(tx *gorm.DB) (err error) {
	var existingTicker OrganizedTickers
	result := tx.Where("symbol = ? AND exchange = ?", ticker.Symbol, ticker.Exchange).First(&existingTicker)

	if result.Error == nil || result.RowsAffected > 0 {
		ticker.ID = existingTicker.ID
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tx.Raw("SELECT nextval('organized_tickers_id_seq')").Scan(&ticker.ID).Error
	} else {
		return result.Error
	}

	return nil
}

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
type AggregatedAsset struct {
	Symbol     string  `json:"symbol"`
	SymbolId   uint64  `json:"symbolId" gorm:"uniqueIndex:idx_uniq_symbol_id"`
	Price      float64 `json:"price"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Percentage float64 `json:"percentage"`
	Volume     float64 `json:"volume"`
	MarketCap  float64 `json:"marketCap"`
}
type ExchangesHistory struct {
	Exchange   string  `json:"exchange"`
	ExchangeID uint64  `gorm:"uniqueIndex:exchange_history_unique" json:"exchangeId"`
	Resolution uint64  `gorm:"uniqueIndex:exchange_history_unique" json:"resolution"`
	Timestamp  uint64  `gorm:"uniqueIndex:exchange_history_unique" json:"timestamp"`
	Volume     float64 `json:"volume"`
}
type GeneralDataHistory struct {
	Marketcap                  float64 `json:"marketcap"`
	MarketcapChange            float64 `json:"marketcapChange"`
	MarketcapChangePerc        float64 `json:"marketcapChangePerc"`
	Volume                     float64 `json:"volume"`
	VolumeChange               float64 `json:"volumeChange"`
	VolumeChangePerc           float64 `json:"volumeChangePerc"`
	DefiVolume                 float64 `json:"defiVolume"`
	DefiVolumeChange           float64 `json:"defiVolumeChange"`
	DefiVolumeChangePerc       float64 `json:"defiVolumeChangePerc"`
	StableVolume               float64 `json:"stableVolume"`
	StableVolumeChange         float64 `json:"stableVolumeChange"`
	StableVolumeChangePerc     float64 `json:"stableVolumeChangePerc"`
	BitcoinMarketcap           float64 `json:"bitcoinMarketcap"`
	BitcoinMarketcapChange     float64 `json:"bitcoinMarketcapChange"`
	BitcoinMarketcapChangePerc float64 `json:"bitcoinMarketcapChangePerc"`
	BitcoinDominance           float64 `json:"bitcoinDominance"`
	BitcoinDominanceChange     float64 `json:"bitcoinDominanceChange"`
	Timestamp                  uint64  `json:"timestamp"`
}
type DexOrganizedTickers struct {
	ID            uint64  `json:"id" gorm:"primaryKey"`
	Exchange      string  `json:"-" gorm:"uniqueIndex:idx_uniq_ticker"`
	ExchangeID    int64   `json:"exchangeId"`
	Symbol        string  `json:"symbol" gorm:"uniqueIndex:idx_uniq_ticker"`
	Timestamp     uint64  `json:"timestamp"`
	Datetime      string  `json:"datetime"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Bid           float64 `json:"bid"`
	BidVolume     float64 `json:"bidVolume"`
	Ask           float64 `json:"ask"`
	AskVolume     float64 `json:"askVolume"`
	Vwap          float64 `json:"vwap"`
	Open          float64 `json:"open"`
	Close         float64 `json:"close"`
	Last          float64 `json:"last"`
	PreviousClose float64 `json:"previousClose"`
	Change        float64 `json:"change"`
	Percentage    float64 `json:"percentage"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	CreatedAt     uint64  `json:"createdAt" gorm:"autoCreateTime:milli"`
}