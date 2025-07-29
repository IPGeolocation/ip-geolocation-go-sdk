/*
IPGeolocation.io - IP intelligence products

Testing UserAgentAPIService
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

func Test_ipgeolocationsdk_UserAgentAPIService(t *testing.T) {
	apiKey := os.Getenv("IPGEO_API_KEY")
	require.NotEmpty(t, apiKey, "Please set IPGEO_API_KEY environment variable")

	ctx := ipgeolocation.WithAPIKey(context.Background(), apiKey)

	configuration := ipgeolocation.NewConfiguration()
	apiClient := ipgeolocation.NewAPIClient(configuration)

	t.Run("Test UserAgentAPIService GetUserAgentDetails", func(t *testing.T) {
		resp, httpRes, err := apiClient.UserAgentAPI.GetUserAgentDetails(ctx).UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9").Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAgentAPIService ParseBulkUserAgentStrings", func(t *testing.T) {
		resp, httpRes, err := apiClient.UserAgentAPI.ParseBulkUserAgentStrings(ctx).UserAgents([]string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9"}).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
