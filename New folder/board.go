package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	err := publish()
	if err != nil {
		panic(err)
	}
}

func publish() error {
	pdf := gofpdf.New("", "in", "", "")
	pdf.AddPage()
	pdf.SetMargins(0.75, 0.5, 0.75)
	pdf.LinearGradient(0.75, 0.5, 7, 0.5, 0, 0, 255, 255, 255, 0, 0, 0, 1, 0)
	pdf.SetDrawColor(0, 255, 255)
	pdf.SetLineWidth(0.06)
	var x float64
	for x = 0; x < 7; x += .5 {
		pdf.Rect(x+0.75, 0.5, .5, .5, "D")
	}
	pdf.SetFillColor(210, 180, 140)
	pdf.Rect(2.75, 2.5, 3, 4, "F")
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetLineWidth(0.08)
	addSpaces(pdf,
		newSpace(3, 2, "purple"),
		newSpace(2, 2, "red"),
		newSpace(1, 2, "orange"),
		newSpace(1, 3, "blue"),
		newSpace(1, 4, "green"),
		newSpace(1, 5, "yellow"),
		newSpace(1, 6, "purple"),
		newSpace(2, 6, "red"),
		newSpace(3, 6, "orange"),
		newSpace(4, 6, "blue"),
		newSpace(5, 6, "green"),
		newSpace(5, 5, "yellow"),
		newSpace(5, 4, "red"),
		newSpace(5, 3, "purple"),
		newSpace(5, 2, "blue"),
		newSpace(5, 1, "orange"),
		newSpace(4, 1, "yellow"),
		newSpace(3, 1, "green"),
		newSpace(2, 1, "red"),
		newSpace(1, 1, "purple"),
		newSpace(0, 1, "blue"),
		newSpace(0, 2, "orange"),
		newSpace(0, 3, "green"),
		newSpace(0, 4, "yellow"),
		newSpace(0, 5, "purple"),
		newSpace(0, 6, "red"),
		newSpace(0, 7, "orange"),
		newSpace(1, 7, "blue"),
		newSpace(2, 7, "green"),
		newSpace(3, 7, "yellow"),
		newSpace(4, 7, "red"),
		newSpace(5, 7, "purple"),
		newSpace(6, 7, "blue"),
		newSpace(6, 6, "red"),
		newSpace(6, 5, "purple"),
		newSpace(6, 4, "blue"),
		newSpace(6, 3, "orange"),
		newSpace(6, 2, "green"),
		newSpace(6, 1, "yellow"),
	)
	pdf.SetDrawColor(0, 0, 0)
	poly(pdf, "D",
		point(1, 7),
		point(6, 7),
		point(6, 1),
		point(0, 1),
		point(0, 8),
		point(7, 8),
		point(7, 1),
		point(6, 1),
		point(6, 7),
		point(1, 7),
		point(1, 2),
		point(5, 2),
		point(5, 6),
		point(2, 6),
		point(2, 3),
		point(4, 3),
		point(4, 2),
		point(1, 2),
	)
	pdf.SetFont("Courier", "B", 14)
	pdf.SetTextColor(0, 0, 0)
	for x := 1.0; x <= 6.0; x++ {
		pdf.Text(7.0, x+1.0, "START")
	}
	pdf.Text(4, 5, "GOAL")
	pdf.SetFillColor(0, 0, 0)
	for x := 0.0; x < 6.0; x++ {
		smiley(pdf, x, 9)
	}
	return pdf.OutputFileAndClose("game.pdf")
}

func poly(pdf *gofpdf.Fpdf, style string, points ...gofpdf.PointType) {
	pdf.Polygon(points, style)
}

func point(x, y float64) gofpdf.PointType {
	return gofpdf.PointType{X: x + 0.75, Y: y + 0.5}
}

func addSpaces(pdf *gofpdf.Fpdf, spaces ...Space) {
	for _, space := range spaces {
		pdf.SetFillColor(space.Color.R, space.Color.G, space.Color.B)
		pdf.Rect(space.X+0.75, space.Y+0.5, 1, 1, "FD")
	}
}

var colors = map[string]Color{
	"blue":   color(0, 0, 255),
	"red":    color(255, 0, 0),
	"green":  color(0, 255, 0),
	"yellow": color(255, 255, 0),
	"purple": color(255, 0, 255),
	"orange": color(255, 165, 0),
}

func newSpace(x, y float64, color string) Space {
	return Space{
		X:     x,
		Y:     y,
		Color: colors[color],
	}
}

type Space struct {
	X     float64
	Y     float64
	Color Color
}

func color(r, g, b int) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

type Color struct {
	R int
	G int
	B int
}

func smiley(pdf *gofpdf.Fpdf, x, y float64) {
	pdf.Circle(0.75+x+0.5, 1+y, 0.4, "D")
	pdf.Circle(1.1+x, 0.8+y, 0.05, "DF")
	pdf.Circle(1.4+x, 0.8+y, 0.05, "DF")
	pdf.Arc(0.75+x+0.5, 1.1+y, 0.2, 0.2, 180, 90, 270, "FD")
}
