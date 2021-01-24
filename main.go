package main
import (
	"log"
  "fmt"
	"github.com/signintech/gopdf"
)

func main() {
  makeRecipt("Kinyua","Ksh.30,000","hello")
}


func makeRecipt(name,amount,uid string){
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
	err = pdf.SetFont("mycoolfont", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	// pdf.Image("images/sign.png", 10, 10, nil
	// pdf.SetLineWidth(2)
	// pdf.SetLineType("dashed")
  // pdf.Line(15, 50, 585, 50)
  headerText := fmt.Sprintf("Receipt for payment done of amount %s",amount)
	rectFillColor(&pdf, "Nano", 100, 20, 20, 550, 100, 142, 170, 219, alignCenter, valignMiddle)
	pdf.SetX(20)
	pdf.SetY(200)
	pdf.Cell(nil, headerText)
	pdf.Cell(nil, "Hello Sir thanks for making me")
  outputfile := fmt.Sprintf("%s.pdf",uid)
	pdf.WritePdf(outputfile)
}


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

func rectFillColor(pdf *gopdf.GoPdf,
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

	pdf.SetX(x)

	if valign == valignMiddle {
		y = y + (h / 2) - (float64(fontSize) / 2)
	} else if valign == valignBottom {
		y = y + h - float64(fontSize)
	}

	pdf.SetY(y)
	pdf.Cell(nil, text)
	pdf.Br(20)
	pdf.SetX(30)
	pdf.Cell(nil,"Ksh.90,000")
}
