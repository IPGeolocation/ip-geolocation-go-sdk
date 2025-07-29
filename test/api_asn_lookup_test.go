/*
IPGeolocation.io - IP intelligence products

Testing ASNLookupAPIService

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

func Test_ipgeolocationsdk_ASNLookupAPIService(t *testing.T) {

	apiKey := os.Getenv("IPGEO_API_KEY")
	require.NotEmpty(t, apiKey, "Please set IPGEO_API_KEY environment variable")

	ctx := ipgeolocation.WithAPIKey(context.Background(), apiKey)

	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	t.Run("Test ASNLookupAPIService GetAsnInfo", func(t *testing.T) {
		resp, httpRes, err := apiClient.ASNLookupAPI.GetAsnInfo(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
