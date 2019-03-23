package converter

import (
	"encoding/xml"
	"os"
	log "github.com/sirupsen/logrus"
)

type XMLConverter struct{
	ConverterBase
}

func (c XMLConverter) Convert() []byte{
	// Add xmlns definition in root tag by changing struct
	doc := (*XmlIodefDocument)(c.Root)

	// convert to XML
	log.Debug("Converting to XML format")
	data, err := xml.MarshalIndent(doc, "", "    ")

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Add XML header ("<?xml version="1.0" encoding="UTF-8"?>"") to xmlData
	log.Debug("Add XML header to output")
	data = []byte(xml.Header + string(data))

	return data
}