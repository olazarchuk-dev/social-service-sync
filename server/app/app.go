package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"social-service-sync/server/app/device/controllers"
	"social-service-sync/server/app/device/events/ws"
)

func Run() {

	app := fiber.New()
	db := DbConn()
	m := ws.NewMongo(db)
	hub := ws.NewHub()
	go hub.Run(m)

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	controllers.Init(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello internet")
	})

	//
	//newSocialSetting := entity.NewSocialSetting("", "", false, 0, 0, 1647937982625, entity.NewDevice("Firefox", "98.0"))
	//socialSettingId, err := ws.RepositoryCreate(newSocialSetting)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("New SocialSetting-ID = '%v'", socialSettingId)
	//

	app.Listen(":3005")

}
