package bootstrap

import "event-booking-api/storage"

type Application struct {
	Env	*Env
}

func App() Application {
	app := &Application{}
	app.Env = LoadEnv()

	storage.InitDb()

	return *app
}
