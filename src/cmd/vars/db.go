package vars

var (
	User     string
	Password string
	Host     string
	DBname   string
)

func SetUser(value string) {
	User = value
}

func SetPassword(value string) {
	Password = value
}

func SetHost(value string) {
	Host = value
}
func SetDBname(value string) {
	DBname = value
}
