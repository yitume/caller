package echosession

type Cfg struct {
	CallerEchoSession CallerCfg
}

type CallerCfg struct {
	Name     string
	Size     int
	Network  string
	Addr     string
	Pwd      string
	Keypairs string
}
