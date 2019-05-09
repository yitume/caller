package ginsession

type Cfg struct {
	CallerGinSession CallerCfg
}

type CallerCfg struct {
	Name     string
	Size     int
	Network  string
	Addr     string
	Pwd      string
	Keypairs string
}
