// Some simple graphs using gonum/plot

package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

// Create a histogram
func histogram(vals []float64, bins int) {

	// Convert floats to plotter array
	vals1 := createPlotterVals(vals)

	// Create histogram
	hist, err := plotter.NewHist(vals1, bins)
	if err != nil {
		panic(err)
	}

	// Create plot
	p := plot.New()
	p.Title.Text = "Histogram"
	p.Add(hist)

	// Save to SVG file
	err = p.Save(1000, 500, "hist.svg")
	if err != nil {
		panic(err)
	}
}

// Create a line chart
func lineChart(vals []float64) {

	// Convert values to XYer
	xys := createPlotterXYs(vals)

	// Create line graph
	g, err := plotter.NewLine(xys)
	if err != nil {
		panic(err)
	}

	// Create plot
	p := plot.New()
	p.Title.Text = "Line Chart"
	p.Add(g)

	// Save to SVG file
	err = p.Save(1000, 500, "line.svg")
	if err != nil {
		panic(err)
	}
}

// Convert list of floats to plotter.Values object
func createPlotterVals(traces []float64) plotter.Values {
	var plotData plotter.Values
	for _, x := range traces {
		plotData = append(plotData, x)
	}
	return plotData
}

// Convert list of floats to plotter.XYs object
func createPlotterXYs(traces []float64) plotter.XYs {
	var plotData plotter.XYs
	for i, y := range traces {
		plotData = append(plotData, plotter.XY{X: float64(i), Y: float64(y)})
	}
	return plotData
}
