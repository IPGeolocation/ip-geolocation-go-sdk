/*
IPGeolocation.io - IP intelligence products

Testing SecurityAPIService

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

func Test_ipgeolocationsdk_SecurityAPIService(t *testing.T) {
	apiKey := os.Getenv("IPGEO_SECURITY_API_KEY")
	require.NotEmpty(t, apiKey, "Please set IPGEO_SECURITY_API_KEY environment variable")

	ctx := ipgeolocation.WithAPIKey(context.Background(), apiKey)

	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	t.Run("Test SecurityAPIService GetBulkIpSecurityInfo", func(t *testing.T) {
		resp, httpRes, err := apiClient.IPSecurityAPI.GetBulkIpSecurityInfo(ctx).Ips([]string{"8.8.8.8"}).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityAPIService GetIpSecurityInfo", func(t *testing.T) {
		resp, httpRes, err := apiClient.IPSecurityAPI.GetIpSecurityInfo(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
