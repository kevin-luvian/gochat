package setting

var ServerSetting = &Server{}

func Setup() {
	setupServer(ServerSetting)
}
