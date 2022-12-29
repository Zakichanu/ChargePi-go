//go:build linux
// +build linux

package reader

import (
	"context"
	"encoding/hex"
	"github.com/clausecker/nfc/v2"
	log "github.com/sirupsen/logrus"
	"github.com/warthog618/gpiod"
	"time"
)

var (
	Iso14443A  = nfc.Modulation{Type: nfc.ISO14443a, BaudRate: nfc.Nbr106}
	ISO14443b  = nfc.Modulation{Type: nfc.ISO14443b, BaudRate: nfc.Nbr106}
	Felica1    = nfc.Modulation{Type: nfc.Felica, BaudRate: nfc.Nbr212}
	Felica2    = nfc.Modulation{Type: nfc.Felica, BaudRate: nfc.Nbr424}
	Jewel      = nfc.Modulation{Type: nfc.Jewel, BaudRate: nfc.Nbr106}
	ISO14443bi = nfc.Modulation{Type: nfc.ISO14443biClass, BaudRate: nfc.Nbr106}

	modulations = []nfc.Modulation{
		Iso14443A, ISO14443b, Felica1, Felica2, Jewel, ISO14443bi,
	}
)

type TagReader struct {
	TagChannel    chan string
	reader        *nfc.Device
	DeviceAddress string
	ResetPin      int
}

// init Initialize the NFC/RFID tag reader. Establish the connection and set up the reader.
func (reader *TagReader) init() {
	dev, err := nfc.Open(reader.DeviceAddress)
	if err != nil {
		log.Fatalf("Cannot communicate with reader")
	}

	reader.reader = &dev
	err = reader.reader.InitiatorInit()
	if err != nil {
		log.Fatal("Failed to initialize reader")
	}
}

// ListenForTags poll the reader for NFC/RFID tags. Uses multiple modulations for different standards.
// Send the ID of the detected card through the TagChannel. If there is a problem with the reader,
// hardware Reset the device.
func (reader *TagReader) ListenForTags(ctx context.Context) {
	reader.init()
	var (
		err    error
		count  int
		target nfc.Target
		UID    string
		UIDLen int
	)

Listener:
	for {
		select {
		case <-ctx.Done():
			break Listener
		default:
			count, target, err = reader.reader.InitiatorPollTarget(modulations, 1, 300*time.Millisecond)
			if err != nil {
				log.WithError(err).Errorf("Error polling the reader")
				reader.Reset()
				continue
			}

			if count > 0 {
				switch target.Modulation() {
				case Iso14443A:
					var card = target.(*nfc.ISO14443aTarget)
					UIDLen = card.UIDLen
					var ID = card.UID
					UID = hex.EncodeToString(ID[:UIDLen])

				case ISO14443b:
					var card = target.(*nfc.ISO14443bTarget)
					UIDLen = len(card.ApplicationData)
					var ID = card.ApplicationData
					UID = hex.EncodeToString(ID[:UIDLen])

				case Felica1, Felica2:
					var card = target.(*nfc.FelicaTarget)
					var UIDLen = card.Len
					var ID = card.ID
					UID = hex.EncodeToString(ID[:UIDLen])

				case Jewel:
					var card = target.(*nfc.JewelTarget)
					var ID = card.ID
					UIDLen = len(ID)
					UID = hex.EncodeToString(ID[:UIDLen])

				case ISO14443bi:
					var card = target.(*nfc.ISO14443biClassTarget)
					var ID = card.UID
					UIDLen = len(ID)
					UID = hex.EncodeToString(ID[:UIDLen])
				default:
					log.Warn("NFC modulation unknown")
					continue
				}

				reader.TagChannel <- UID
			}

			time.Sleep(time.Millisecond * 200)
		}
	}
}

func (reader *TagReader) GetTagChannel() <-chan string {
	return reader.TagChannel
}

func (reader *TagReader) GetType() string {
	return PN532
}

// Cleanup Close the reader device connection.
func (reader *TagReader) Cleanup() {
	close(reader.TagChannel)
	reader.reader.Close()
}

// Reset Implements the hardware reset by pulling the ResetPin low and then releasing.
func (reader *TagReader) Reset() {
	log.Infof("Resetting the reader")
	// Refer to gpiod docs
	c, err := gpiod.NewChip("gpiochip0")
	pin, err := c.RequestLine(reader.ResetPin, gpiod.AsOutput(1))
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	time.Sleep(time.Millisecond * 300)
	err = pin.SetValue(0)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	time.Sleep(time.Millisecond * 300)
	err = pin.SetValue(1)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
}
