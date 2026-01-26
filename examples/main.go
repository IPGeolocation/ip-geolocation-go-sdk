// IPGeolocation API Go SDK example for. For more examples, visit: https://ipgeolocation.io/documentation/ip-geolocation-api-go-sdk.html
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/sdk"
)

func main() {
	ctx := ipgeolocation.WithAPIKey(context.Background(), "your_api_key")
	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	respAstronomy, r, err := apiClient.AstronomyAPI.GetAstronomyDetails(ctx).Ip("1.1.1.1").Lang("fr").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AstronomyAPI.GetAstronomyDetails`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	responseJson, _ := json.MarshalIndent(respAstronomy, "", "  ")
	fmt.Println(string(responseJson))
}
