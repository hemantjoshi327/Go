package main

import(
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func main(){
pdf := gofpdf.New("P", "mm", "A4", "")
w, h := pdf.GetPageSize()
fmt.Printf("width=%v, Height=%v\n", w, h)
pdf.AddPage()

//Basic text writing
pdf.MoveTo(0, 0)
pdf.SetFont("Arial", "B", 16)
//pdf.Cell(40, 10, "Hello, world")
_, lineHight := pdf.GetFontSize()
pdf.SetTextColor(20, 200, 0)
pdf.Text(0, lineHight, "Hello Hemant")
pdf.MoveTo(0, lineHight*2.0)


err := pdf.OutputFileAndClose("Third-pdf.pdf")

if err != nil{
	panic(err)
}

}