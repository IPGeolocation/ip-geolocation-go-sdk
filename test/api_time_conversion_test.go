/*
IPGeolocation.io - IP intelligence products

Testing TimeConversionAPIService

*/

package ipgeolocationsdk

import (
	"context"
	"os"
	"testing"

	ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ipgeolocationsdk_TimeConversionAPIService(t *testing.T) {
	apiKey := os.Getenv("IPGEO_API_KEY")
	require.NotEmpty(t, apiKey, "Please set IPGEO_API_KEY environment variable")

	ctx := ipgeolocation.WithAPIKey(context.Background(), apiKey)

	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	t.Run("Test TimeConversionAPIService ConvertTimeBetweenTimezones", func(t *testing.T) {
		resp, httpRes, err := apiClient.TimeConversionAPI.ConvertTimeBetweenTimezones(ctx).TzFrom("Europe/London").TzTo("America/New_York").Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
