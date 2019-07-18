package config

const (
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ProfileParser"
	NilParser     = "NilParser"

	// Service ports
	ItemSaverPort = 1234

	// es
	ElasticIndex = "dating_profile"

	// RPC EndPoint
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Rate limiting
	Qps = 20
)
