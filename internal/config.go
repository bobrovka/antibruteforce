package internal

// Config базовый конфиг приложения
type Config struct {
	HTTPListen string `config:"host"` // ip и port на котором должен слушать web-сервер
	Redis      string

	RPMLogin    int64
	RPMPassword int64
	RPMIP       int64
}
