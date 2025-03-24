package client

type FetcherServiceClient interface {
	FetchBills() error
}
