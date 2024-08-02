package shared

type NewsFeed struct {
	Id				string
	Tags			string
	PredictedTags	string `gorm:"column:predictedTags"`
	SourceID		uint64 `gorm:"column:sourceId"`
	Title			string
	Content			string
	Link			string
	Preview			string
	Timestamp		uint64
	IsNew			bool   `gorm:"column:isNew"`
	ShowTags		string `gorm:"column:showTags"`
	Source			Source
}

func (NewsFeed) TableName() string {
	return "news_feeds"
}

type Source struct {
	Localisations 		string
	Sources 			string
	SourcesLink			string `gorm:"column:sourcesLink"`
	RssLink				string `gorm:"column:rssLink"`
	TagsCategories		string `gorm:"column:tagsCategories"`
	Contact				string
	Comment				string
	Logo				string
	Timestamp			uint64
	Status				string
	ID					uint64
	AssetID				string `gorm:"column:asset"`
}

func (Source) TableName() string {
	return "sources"
}