package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func drawGraph(events []Event, title string, line string, path string) error {

	fmt.Println("Plotting " + title)
	xy := make(plotter.XYs, len(events))

	for i := 0; i < len(events); i++ {
		xy[i].X = float64(i)
		//xy[i].X = events[i].Date
		xy[i].Y = float64(events[i].Cases)
	}

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Cases"

	err = plotutil.AddLines(p, line, xy)
	if err != nil {
		return err
	}

	if err := p.Save(20*vg.Centimeter, 10*vg.Centimeter, path); err != nil {
		return err
	}
	fmt.Println("Saved " + title)
	return nil
}
