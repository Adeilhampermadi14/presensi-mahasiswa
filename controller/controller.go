package controller

import (
	"github.com/Adeilhampermadi14/presensi-mahasiswa/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetHome(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}

func InsertDafdir(c *fiber.Ctx) error {
	dafdir := new(xxx.Dafdir)
	ps := xxx.InsertDafdir(config.MongoConn,
		dafdir.Keterangan,
		dafdir.Kehadiran,
	)
	return c.JSON(ps)
}

func GetDataDafdir(c *fiber.Ctx) error {
	getket := xxx.GetDataDafdir("Masuk")
	return c.JSON(getket)
}

func GetDataNilai(c *fiber.Ctx) error {
	getlai := xxx.GetDataNilai("Matematika, Biologi")
	return c.JSON(getlai)
}

func GetDafpel(c *fiber.Ctx) error {
	getpel := xxx.GetDataDafpel("Terlambat")
	return c.JSON(getpel)
}

func GetDataPembayaran(c *fiber.Ctx) error {
	getyar := xxx.GetDataPembayaran("Terbayar")
	return c.JSON(getyar)
}
