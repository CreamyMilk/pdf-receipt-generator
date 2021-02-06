package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

const (
	valignTop    = 1
	valignMiddle = 2
	valignBottom = 3
)

const (
	alignLeft   = 4
	alignCenter = 5
	alignRight  = 6
)

func main() {
	makeRecipt("Kinyua", "Ksh.30,000", "hello")
}

func makeRecipt(name, amount, uid string) {
	//Create pdf doc and add Page
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("mycoolfont", "fonts/SourceSansPro-Light.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("mycoolfont", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}
	//Header
	lightheaderColor := map[string]uint8{"r": 142, "g": 170, "b": 219}
	darkheaderColor := map[string]uint8{"r": 91, "g": 126, "b": 215}
	lightCords := map[string]float64{"x": 20, "y": 20}
	darkCords := map[string]float64{"x": 460, "y": 20}
	lightBlueHeader(&pdf, "RECEIPT", 0, lightCords["x"], lightCords["y"], 550.0, 100.0, lightheaderColor["r"], lightheaderColor["g"], lightheaderColor["b"], alignLeft, valignMiddle)
	addPaymentBlock(&pdf, "$223023.00", 0, darkCords["x"], darkCords["y"], 110.0, 100.0, darkheaderColor["r"], darkheaderColor["g"], darkheaderColor["b"], alignCenter, valignMiddle)
	//Signature
	drawSignature(&pdf)
	//Generate PDFs with a file name
	outputfile := fmt.Sprintf("%s.pdf", uid)
	pdf.WritePdf(outputfile)
}

func drawSignature(pdf *gopdf.GoPdf) {
	box := gopdf.Rect{}
	box.W = 200
	box.H = 200
	pdf.Image("images/sign.png", 380, 650, &box)
}
func lightBlueHeader(pdf *gopdf.GoPdf,
	text string,
	fontSize int,
	x, y, w, h float64,
	r, g, b uint8,
	align, valign int,
) {
	pdf.SetLineWidth(0)
	pdf.SetFillColor(r, g, b)
	pdf.SetStrokeColor(r, g, b) //setup fill color
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0, 0, 0)

	if align == alignCenter {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + (w / 2) - (textw / 2)
	} else if align == alignRight {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + w - textw
	}

	pdf.SetX(x + 10)

	if valign == valignMiddle {
		y = y + (h / 2) - (float64(fontSize) / 2)
	} else if valign == valignBottom {
		y = y + h - float64(fontSize)
	}

	pdf.SetY(y - 20)
	err := pdf.SetFont("mycoolfont", "", 40)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetTextColor(255, 255, 255)
	pdf.Cell(nil, text)
	pdf.Br(20)
}
func addPaymentBlock(pdf *gopdf.GoPdf,
	text string,
	fontSize int,
	x, y, w, h float64,
	r, g, b uint8,
	align, valign int,
) {
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(r, g, b) //setup fill color
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0, 0, 0)
	if align == alignCenter {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + (w / 2) - (textw / 2)
	} else if align == alignRight {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + w - textw
	}

	if valign == valignMiddle {
		y = y + (h / 2) - (float64(fontSize) / 2)
	} else if valign == valignBottom {
		y = y + h - float64(fontSize)
	}

	pdf.SetX(x + 70)
	pdf.SetY(y - 35)
	_ = pdf.SetFont("mycoolfont", "", 10)
	pdf.Cell(nil, "Amount")

	pdf.SetTextColor(255, 255, 255)
	//pdf.Cell(nil, text)
	pdf.SetX(x + 44)
	pdf.SetY(y - 8)
	_ = pdf.SetFont("mycoolfont", "", 21)

	pdf.Cell(nil, text)
}
