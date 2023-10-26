package user

type AssetType int

/**
 * Aggregate
 */
type Favorite interface {
	UserId() string

	Add(assetId string, assetType int) error
	Del(assetId string) error

	Count() int
}
