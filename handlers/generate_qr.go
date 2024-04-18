package handlers

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	tele "gopkg.in/telebot.v3"
	"os"
	"strconv"
	"time"
)

func QrGen(c tele.Context) error {

	if len(c.Message().Text) < 5 {
		log.Warn().Msg("Error create QR. Name file is 0")
		return c.Send("Error create QR. Name file is 0")
	}
	qr, err := getQrImage(context.TODO(), c.Message().Text[4:], strconv.Itoa(int(c.Message().Sender.ID)))
	defer func() {
		err := os.Remove(qr)
		if err != nil {
			log.Error().Msg("Can't delete file - " + err.Error())
		}
		log.Info().Msg("Delete qr code - " + qr)
	}()
	if err != nil {
		c.Sender()
	}
	p := &tele.Photo{File: tele.FromDisk(qr)}
	return c.SendAlbum(tele.Album{p})
}
func getQrImage(ctx context.Context, text string, fileName string) (string, error) {
	urlFile := fmt.Sprintf("./%s-%v.jpg", fileName, time.Now().Unix())
	ctx.Done()
	code, err := qrcode.New(text)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	writer, err := standard.New(urlFile)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	if err = code.Save(writer); err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	log.Info().Msg(fmt.Sprintf("The Qr code was generated successfully"))
	return urlFile, nil
}
