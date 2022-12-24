package appConfig

type Config struct {
	Excludes []string
}

var AppConfig = Config{
	Excludes: []string{"./reportWebVitals"},
}
