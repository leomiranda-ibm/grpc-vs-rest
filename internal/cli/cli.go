package cli

import (
	"errors"
	"flag"
	"rest-vs-grpc/internal/constants"
)

func HandleParams() (int, error) {
	qtyRequests := flag.Int("requests", constants.QtyRequests, "inform how many request do you want")

	flag.Parse()

	if *qtyRequests < 1 {
		return 0, errors.New("inform the number of requests")
	}

	return *qtyRequests, nil
}
