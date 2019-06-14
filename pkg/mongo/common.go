package mongo

type Cfg struct {
	CallerMongo map[string]CallerCfg
}

type CallerCfg struct {
	Debug bool

	URL      string
	Database string
}
