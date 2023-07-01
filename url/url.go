package url

import (
	"github.com/Adeilhampermadi14/presensi-mahasiswa/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.GetHome)
	page.Post("/insert", controller.InsertDafdir)
	page.Get("/dafdir", controller.GetDataDafdir)
	page.Get("/nilai", controller.GetDataNilai)
	page.Get("/pelanggar", controller.GetDafpel)
	page.Get("/bayar", controller.GetDataPembayaran)
}
