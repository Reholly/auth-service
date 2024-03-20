package main

import (
	_ "auth-service/docs"
	"auth-service/internal/application"
	"auth-service/internal/config"
)

func main() {
	config := config.LoadConfig()

	/*ms := service.NewMailService(*config)
	err := ms.SendMail(context.Background(), "atokadota2@yandex.ru", "HELLO FROM GOLANG", "hi )")
	if err != nil {
		return
	}*/
	app := application.NewApplication(config)

	app.Run()
}
