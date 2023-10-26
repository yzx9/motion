package recommender

// Aggrete root
type Recommender interface {
	List(limit int) []Recommendation
	ListByKeyword(limit int, keyword string) []Recommendation
	ListByVideo(limit int, videoId string) []Recommendation
}

// valobj
type Recommendation struct {
	AssetKind string
	AssetId   string
}

// valobj
type AssetType int

const (
	AssetVideo AssetType = 10
)

func GetUserRecommender(userId string) Recommender {
	return nil // TODO: not implement
}
