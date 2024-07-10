package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

var TuturuBuffer = make([][]byte, 0)
var GnomeBuffer = make([][]byte, 0)
var YasouUltBuffer = make([][]byte, 0)
var ZedBuffer = make([][][]byte, 0)

func LoadSounds() {
	loadSound("tuturu.dca", &TuturuBuffer)
	loadSound("gnome.dca", &GnomeBuffer)
	loadSound("yasou_ult.dca", &YasouUltBuffer)

	tempBuffer := make([][]byte, 0)
	loadSound("sounds/zed/0.dca", &tempBuffer)
	ZedBuffer = append(ZedBuffer, tempBuffer)
	tempBuffer = make([][]byte, 0)
	loadSound("sounds/zed/1.dca", &YasouUltBuffer)
	ZedBuffer = append(ZedBuffer, tempBuffer)
	tempBuffer = make([][]byte, 0)
	loadSound("sounds/zed/2.dca", &YasouUltBuffer)
	ZedBuffer = append(ZedBuffer, tempBuffer)
	tempBuffer = make([][]byte, 0)
	loadSound("sounds/zed/3.dca", &YasouUltBuffer)
	ZedBuffer = append(ZedBuffer, tempBuffer)
	tempBuffer = make([][]byte, 0)
	loadSound("sounds/zed/4.dca", &YasouUltBuffer)
	ZedBuffer = append(ZedBuffer, tempBuffer)
}

func loadSound(filePath string, buffer *[][]byte) error {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		*buffer = append(*buffer, InBuf)
	}
}
