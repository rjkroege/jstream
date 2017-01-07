package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Namedscanner struct {
	name       string
	scanner    *bufio.Scanner
	moretoread bool
}

// Reads lines from standard in
// Args go like this: JSON key file
// TODO(rjk): Improve this. I could do something smarter about the
// labeling of the tags.
func main() {
	flag.Parse()

	data := make([]map[string]string, 0)
	scanners := make([]*Namedscanner, 0)

	args := flag.Args()
	for i := 0; i < len(args); i += 2 {
		fd, err := os.Open(args[i+1])
		if err != nil {
			log.Fatalln("couldn't open", args[i+1], "for reading")
		}

		ns := &Namedscanner{
			name:       args[i],
			scanner:    bufio.NewScanner(fd),
			moretoread: true,
		}
		scanners = append(scanners, ns)
	}

	if len(scanners) <= 0 {
		log.Fatalln("no files to read. exiting")
	}

	openfile := true
	for openfile {
		openfile = false
		obj := make(map[string]string)
		for _, ns := range scanners {
			if !ns.moretoread {
				continue
			}

			openfile = true
			if s := ns.scanner.Scan(); s {
				obj[ns.name] = ns.scanner.Text()
				data = append(data, obj)
			} else { // End or failure case
				ns.moretoread = false
				if err := ns.scanner.Err(); err != nil {
					log.Println("tool an error", err)
				}
			}
		}
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}
