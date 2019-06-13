package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/zip", zipHandler)
	http.ListenAndServe(":8080", nil)

	//GenerateZip()
}

// GenerateZip -
func GenerateZip() (string, *bytes.Buffer, error) {
	record := []string{"test1", "test2", "test3"} // just some test data to use for the wr.Writer() method below.

	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	file, err := w.Create("test.csv")

	rows := GenerateCSV(record)

	_, err = file.Write(rows)

	err = w.Close()

	//write the zipped file to the disk
	//ioutil.WriteFile("compressed_test.zip", buf.Bytes(), 0777)

	return "compressed_test.zip", buf, err
}

// GenerateCSV -
func GenerateCSV(record []string) []byte {

	b := &bytes.Buffer{}       // creates IO Writer
	wr := csv.NewWriter(b)     // creates a csv writer that uses the io buffer.
	for i := 0; i < 100; i++ { // make a loop for 100 rows just for testing purposes
		wr.Write(record) // converts array of string to comma seperated values for 1 row.
	}
	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	return b.Bytes()
}

func zipHandler(w http.ResponseWriter, r *http.Request) {

	filename, buf, _ := GenerateZip()

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Write(buf.Bytes())
}
