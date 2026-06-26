package contracts

type API interface {
	ValidateURL(url string) error
}
