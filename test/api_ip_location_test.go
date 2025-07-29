/*
IPGeolocation.io - IP intelligence products

Testing IPLocationAPIService

*/

package ipgeolocationsdk

import (
	"context"
	"os"
	"testing"

	ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ipgeolocationsdk_IPLocationAPIService(t *testing.T) {
	apiKey := os.Getenv("IPGEO_API_KEY")
	require.NotEmpty(t, apiKey, "Please set IPGEO_API_KEY environment variable")

	ctx := ipgeolocation.WithAPIKey(context.Background(), apiKey)

	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	t.Run("Test IPLocationAPIService GetBulkIpGeolocation", func(t *testing.T) {
		resp, httpRes, err := apiClient.IPGeolocationAPI.GetBulkIpGeolocation(ctx).Ips([]string{"8.8.8.8"}).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IPLocationAPIService GetIpGeolocation", func(t *testing.T) {
		resp, httpRes, err := apiClient.IPGeolocationAPI.GetIpGeolocation(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
