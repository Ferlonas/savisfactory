package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ferlonas/savisfactory/file"
)

const (
	signatureSize = 16
)

var saveSignature = []byte{0x05, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0xF9, 0x02, 0x01, 0x00, 0x11, 0x00, 0x00, 0x00}

type Parser struct {
	filename string
	saveFile *file.SaveFile
	fh       *os.File
}

func NewParser(file string) *Parser {
	return &Parser{
		filename: file,
	}
}

func (p *Parser) ParseFile() (*file.SaveFile, error) {
	p.OpenFile()
	defer p.CloseFile()

	p.verifySignature()
	p.parseHeader()

	return p.saveFile, nil
}

// File parsing methods below

func (p *Parser) verifySignature() error {
	data, length, err := p.readNextBlock(signatureSize)
	if err != nil {
		// errors.handleError(err)
		return err
	}
	if length < signatureSize {
		return fmt.Errorf("Signature was not the required size")
	}
	if bytes.Equal(saveSignature, data) {
		fmt.Println("Satisfactory save file confirmed.")
	}
	return nil
}

func (p *Parser) parseHeader() error {
	_, _, err := p.readNextBlock(headerSize)
	if err != nil {
		// errors.handleError(err)
		return err
	}
	return nil
}

func (p *Parser) parseFileInfo() error {
}

// File manipulation methods below

func (p *Parser) readNextBlock(blockSize int) ([]byte, int, error) {
	data := make([]byte, blockSize)
	count, err := p.fh.Read(data)
	if err != nil {
		fmt.Printf("Error while reading: %s\n", err.Error())
		return data, 0, err
	}
	fmt.Printf("Read %d bytes\n", count)
	return data, count, nil
}

func (p *Parser) OpenFile() error {
	fh, err := os.Open(p.filename)
	if err != nil {
		fmt.Printf("Error while opening file: %s\n", err.Error())
		return err
	}
	p.fh = fh
	p.saveFile = &file.SaveFile{}
	return nil
}

func (p *Parser) CloseFile() error {
	err := p.fh.Close()
	if err != nil {
		fmt.Printf("Error while closing file: %s\n", err.Error())
		return err
	}
	p.fh = nil
	return nil
}

func (p *Parser) DiscardSave() {
	p.saveFile = nil
}
