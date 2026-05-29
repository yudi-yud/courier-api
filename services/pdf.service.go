package services

import (
	"bytes"
	"courier-api/models"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/go-pdf/fpdf"
)

type PDFService interface {
	GenerateAirwayBill(shipment *models.Shipment) ([]byte, error)
}

type pdfService struct{}

func NewPDFService() PDFService {
	return &pdfService{}
}

func (s *pdfService) GenerateAirwayBill(shipment *models.Shipment) ([]byte, error) {

	pdf := fpdf.New("P", "mm", "A6", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 5)

	// =========================
	// HEADER
	// =========================
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "COURIER-API EXPEDITION", "", 0, "C", false, 0, "")
	pdf.Ln(10)

	pdf.Line(5, pdf.GetY(), 100, pdf.GetY())
	pdf.Ln(3)

	// =========================
	// RESI
	// =========================
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 12, shipment.ResiNumber, "", 0, "C", false, 0, "")
	pdf.Ln(10)

	// =========================
	// BARCODE
	// =========================
	// generate barcode
	bar, err := code128.Encode(shipment.ResiNumber)
	if err != nil {
		return nil, err
	}

	// scale barcode
	scaledBar, err := barcode.Scale(bar, 300, 70)
	if err != nil {
		return nil, err
	}

	// simpan sementara
	tmpFile := "barcode.png"

	file, err := os.Create(tmpFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = png.Encode(file, scaledBar)
	if err != nil {
		return nil, err
	}

	// tampilkan barcode ke PDF
	pdf.Image(tmpFile, 20, pdf.GetY(), 60, 15, false, "", 0, "")
	pdf.Ln(20)

	pdf.Line(5, pdf.GetY(), 100, pdf.GetY())
	pdf.Ln(5)

	// =========================
	// DETAILS
	// =========================
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(20, 5, "Service:")
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(0, 5, shipment.ServiceType)
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 8)
	pdf.Cell(20, 5, "Weight:")
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(0, 5, fmt.Sprintf("%.2f Kg", shipment.Weight))
	pdf.Ln(8)

	// =========================
	// SENDER
	// =========================
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(0, 6, " SENDER (Pengirim)", "", 0, "L", true, 0, "")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0, 5,
		fmt.Sprintf("%s\n%s",
			shipment.SenderName,
			shipment.SenderAddress,
		),
		"",
		"L",
		false,
	)
	pdf.Ln(3)

	// =========================
	// RECEIVER
	// =========================
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(0, 6, " RECEIVER (Penerima)", "", 0, "L", true, 0, "")
	pdf.Ln(6)

	pdf.SetFont("Arial", "B", 9)
	pdf.MultiCell(0, 5, shipment.ReceiverName, "", "L", false)

	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0, 5,
		fmt.Sprintf("Phone: %s", shipment.ReceiverPhone),
		"",
		"L",
		false,
	)

	pdf.MultiCell(0, 5, shipment.ReceiverAddress, "", "L", false)
	pdf.Ln(5)

	// =========================
	// FOOTER
	// =========================
	pdf.Line(5, pdf.GetY(), 100, pdf.GetY())
	pdf.Ln(3)

	pdf.SetFont("Arial", "I", 7)
	pdf.CellFormat(
		0,
		5,
		fmt.Sprintf(
			"Created at: %s",
			shipment.CreatedAt.Format("2006-01-02 15:04:05"),
		),
		"",
		0,
		"C",
		false,
		0,
		"",
	)

	var buf bytes.Buffer

	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	// hapus file temporary
	_ = os.Remove(tmpFile)

	return buf.Bytes(), nil
}
