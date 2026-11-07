# Changelog

## 2.1.0

- Add Arabic (`ar`) for single and bulk lookups.
- Add detailed lookup methods for CDN, proxy, bot, and corporate gateway data.
- Keep redirects from sending API keys, headers, and bulk requests to another server. Hide API keys in printed settings and request errors.
- Keep the HTTP status when an error response cannot be read. Check that typed bulk results match the number of inputs.
- Fix closing a client during a lookup and using a custom HTTP transport.

## 2.0.0

- Replaced the old generated multi-API SDK with the new Go SDK for the IPGeolocation.io IP Location API.
- Added typed and raw single lookup methods for `/v3/ipgeo`.
- Added typed and raw bulk lookup methods for `/v3/ipgeo-bulk`.
- Added request validation, response metadata, clearer error handling, and updated public docs.
- Added unit tests, live integration tests, and live field parity tests.
