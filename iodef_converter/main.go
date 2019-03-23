package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"path"
	log "github.com/sirupsen/logrus"

	"iodef-converter/iodef_converter/converter"
)

// Flags
var (
	flagFileName = flag.String("file", "", "Input file's name in /infiles/ from project directory")
	flagOption   = flag.String("option", "", "Conversion options:\n\t+ xml: convert to xml\n\t+ json: convert to json\n\t")
)

func main() {
	// Parse flags
	flag.Parse()

	if len(*flagFileName) > 0 {
        // Get extension from file
		ext := path.Ext(*flagFileName)

		// declare input and output paths
		inputPath := os.Getenv("GOPATH")+ "/src/iodef-converter/infiles/" + *flagFileName
		xmlOutputPath := os.Getenv("GOPATH")+ "/src/iodef-converter/outfiles/" + strings.TrimSuffix(*flagFileName, ext) + ".xml"
		jsonOutputPath := os.Getenv("GOPATH")+ "/src/iodef-converter/outfiles/" + strings.TrimSuffix(*flagFileName, ext) + ".json"

		if len(*flagOption) > 0{
			switch strings.ToLower(*flagOption){
			case "xml":
				// convert xml to xml
				var c converter.XMLConverter
				c.InFilePath = inputPath
				c.OutFilePath = xmlOutputPath
				switch ext{
				case ".xml":
					c.InputType = converter.XML
				case ".json":
					c.InputType = converter.JSON
				default:
					// Unknown extension in file name
					log.WithField("file",*flagFileName).Error("File is not a XML/JSON file")
				}	
				c.Root = c.LoadInputFile()
				c.Data = c.Convert()
				c.SaveOutputFile(c.Data)
			case "json":
				// convert xml to json
				var c converter.JsonConverter
				c.InFilePath = inputPath
				c.OutFilePath = jsonOutputPath
				switch ext{
				case ".xml":
					c.InputType = converter.XML
				case ".json":
					c.InputType = converter.JSON
				default:
					// Unknown extension in file name
					log.WithField("file",*flagFileName).Error("File is not a XML/JSON file")
				}
				c.Root = c.LoadInputFile()
				c.Data = c.Convert()
				c.SaveOutputFile(c.Data)
			case "cbor","cddl":
				// CBOR work-in-progress
				log.Debug("Do convert XML to CBOR/CDDL")
				fmt.Println("CBOR conversion is not supported yet.")
			default:
				// Other input => imform invalid command
				log.WithField("option",*flagOption).Error("Invalid command")
			}
		} else {
			// No input found for -option
			log.Debug("No option command provided.")
		}
	} else {
		// No input found for -file
		log.WithField("file",*flagFileName).Error("File is not a XML/JSON file")
	}
}
