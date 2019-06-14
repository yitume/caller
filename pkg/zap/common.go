package zap

type Cfg struct {
	CallerZap map[string]CallerCfg
}

type CallerCfg struct {
	Debug bool
	Level string
	Path  string
}
