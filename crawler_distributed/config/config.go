package config

const (
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ProfileParser"
	NilParser     = "NilParser"

	// Service ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// es
	ElasticIndex = "dating_profile"

	// RPC EndPoint
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
