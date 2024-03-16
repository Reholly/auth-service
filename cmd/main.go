package main

import (
	"auth-service/internal/config"
	"auth-service/internal/service"
	"context"
)

func main() {
	config := config.LoadConfig()

	ms := service.NewMailService(*config)
	err := ms.SendMail(context.Background(), "atokadota2@yandex.ru", "HELLO FROM GOLANG", "hi )")
	if err != nil {
		return
	}
}
