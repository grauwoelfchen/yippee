package main

import (
	_ "fmt"
	"log"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	title := "Test"
	// Meta
	pdf.SetTitle(title, false)
	// Header
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "B", 14)
		pdf.Image(filepath.Join("./src", "logo.png"),
			10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		//pdf.Ln()
	})
	// Footer
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
	})
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello")
	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
