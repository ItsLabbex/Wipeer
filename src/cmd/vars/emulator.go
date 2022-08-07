package vars

var (
	Emulator       string
	UserMainDelete bool
)

func SetEmulatorType(value string) {
	Emulator = value
}

func SetUserMainIsDelete(value bool) {
	UserMainDelete = value
}
