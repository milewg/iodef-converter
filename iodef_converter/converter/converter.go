package converter

import (
	"strings"
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	"os"
	log "github.com/sirupsen/logrus"

	iodef "iodef-converter/iodef_converter/go_Iodef20"
)

type Converter interface {
	LoadInputFile()
	Convert()
	Unmarshal()
	SaveOutputFile()
}

type FormatType string
const (
	XML FormatType = "xml"
	JSON FormatType = "json"
)

type ConverterBase struct {
	InputType FormatType
	InFilePath string
	Data []byte
	Root *IodefDocument
	OutFilePath string
}

func (c ConverterBase) LoadInputFile() *IodefDocument{
	format := strings.ToUpper(string(c.InputType))
	log.Debugf("Open %s file",format)

	// Open file
	inFile, err := os.Open(c.InFilePath)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer inFile.Close()

	// Read data from file
	log.Debugf("Read data from %s file", format)
	c.Data, _ = ioutil.ReadAll(inFile)

	// Unmarshal
	c.Root = c.Unmarshal()

	if err != nil{
		log.Error(err)
		os.Exit(1)
	}

	return c.Root
}

// General Unmarshal function
func (c ConverterBase) Unmarshal() *IodefDocument{
	// Check input format type
	switch (c.InputType) {
	case XML:
		return c.unmarshalXML();
	case JSON:
		return c.unmarshalJSON();
	default:
		log.Debug("Invalid input type")
		return nil
	}
}

// Unmarshal XML
func (c ConverterBase) unmarshalXML() *IodefDocument{
	log.Debug("Unmarshal XMl")
	err := xml.Unmarshal(c.Data, &c.Root)
	
	if err != nil{
		log.Error(err)
		os.Exit(1)
	}

	return c.Root
}

// Unmarshal JSON
func (c ConverterBase) unmarshalJSON() *IodefDocument{
	// Initialize struct with default value of xmlns:xsi (xmlns is implicitly initialized with JSON tag)
	c.Root = &IodefDocument{
		Xmlns: "urn:ietf:params:xml:ns:iodef-2.0",
		Xsi: "http://www.w3.org/2001/XMLSchema-instance", 
		TxsdIodefDocument: iodef.TxsdIodefDocument{},
	}

	// Unmarshal JSON data to struct
	log.Debug("Unmarshal JSON")
	err := json.Unmarshal(c.Data, &c.Root.TxsdIodefDocument)

	if err != nil{
		log.Error(err)
		os.Exit(1)
	}
	
	return c.Root
}

func (c ConverterBase) SaveOutputFile(data []byte) {
	// write to XML file
	log.Debug("Write output to XML file")
	file, err := os.Create(c.OutFilePath)

	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	file.Write(data)
	file.Close()

}

// Root struct for XML unmarshalling and JSON
type IodefDocument struct {
	XMLName xml.Name `xml:"IODEF-Document"`
	Xmlns   string   `xml:"xmlns,attr,omitempty"`
	Ds      string   `xml:"ds,attr,omitempty"`
	Xsi     string   `xml:"xsi,attr,omitempty"`
	SchemaLocation string `xml:"schemaLocation,attr,omitempty"`
	Iodef	string   `xml:"iodef,attr,omitempty"`
	iodef.TxsdIodefDocument
}

// Namespace workaround for root struct in XML marshaling
type XmlIodefDocument struct{
	XMLName xml.Name `xml:"IODEF-Document"`
	Xmlns   string   `xml:"xmlns,attr,omitempty"`
	Ds      string   `xml:"xmlns:ds,attr,omitempty"`
	Xsi     string   `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation string `xml:"xsi:schemaLocation,attr,omitempty"`
	Iodef	string   `xml:"xmlns:iodef,attr,omitempty"`
	iodef.TxsdIodefDocument
}