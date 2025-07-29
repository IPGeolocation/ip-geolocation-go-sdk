/*
IPGeolocation.io - IP intelligence products

Testing AbuseContactAPIService

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

func Test_ipgeolocationsdk_AbuseContactAPIService(t *testing.T) {

	apiKey := os.Getenv("IPGEO_API_KEY")
	require.NotEmpty(t, apiKey, "Please set IPGEO_API_KEY environment variable")

	ctx := ipgeolocation.WithAPIKey(context.Background(), apiKey)

	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	t.Run("Test AbuseContactAPIService GetAbuseContactInfo", func(t *testing.T) {
		resp, httpRes, err := apiClient.AbuseContactAPI.GetAbuseContactInfo(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
