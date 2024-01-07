package main

import (
	"awesomeProject/main/balistics"
	"awesomeProject/main/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type tableData struct {
	velocities      [3]float64
	positiveDegrees [3]float64
	negativeDegrees [3]float64
	positiveMils    [3]float64
	negativeMils    [3]float64
}

func (d *tableData) Length() (int, int) {
	return 5, 3
}

func (d *tableData) GetValue(row, col int) string {
	switch row {
	case 0:
		if col == 0 {
			return fmt.Sprintf("Close (%.0f m/s)", d.velocities[col])
		}
		if col == 1 {
			return fmt.Sprintf("Medium (%.0f m/s)", d.velocities[col])
		}
		if col == 2 {
			return fmt.Sprintf("Far (%.0f m/s)", d.velocities[col])
		}
		return ""
	case 1:
		return fmt.Sprintf("%.2f °", d.positiveDegrees[col])
	case 2:
		return fmt.Sprintf("%.2f °", d.negativeDegrees[col])
	case 3:
		return fmt.Sprintf("%.2f mils", d.positiveMils[col])
	case 4:
		return fmt.Sprintf("%.2f mils", d.negativeMils[col])
	default:
		return ""
	}
}

func (d *tableData) SetValue(row, col int, value string) {
	// This table is read-only, so this method can be left empty.
}

func main() {
	data := &tableData{
		velocities: [3]float64{70, 140, 200},
	}

	a := app.New()
	w := a.NewWindow("Ballistics Calculator Mk6 Mortar")

	w.Resize(fyne.NewSize(470, 400))

	targetDistanceLabel := widget.NewLabel("Target Distance")
	targetDistanceEntry := widget.NewEntry()

	heightDifferenceLabel := widget.NewLabel("Height Difference")
	heightDifferenceEntry := widget.NewEntry()

	table := widget.NewTable(
		func() (int, int) { return data.Length() },
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			label.Resize(fyne.NewSize(150, 100))
			return label
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data.GetValue(i.Row, i.Col))
		},
	)

	for i := 0; i < 3; i++ {
		table.SetColumnWidth(i, 150)
	}

	process := func() {
		targetDistance, err1 := strconv.ParseFloat(targetDistanceEntry.Text, 64)
		heightDifference, err2 := strconv.ParseFloat(heightDifferenceEntry.Text, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input")
			return
		}

		for i := 0; i < 3; i++ {
			positive := balistics.AngleOfReachPositive(data.velocities[i], targetDistance, heightDifference)
			negative := balistics.AngleOfReachNegative(data.velocities[i], targetDistance, heightDifference)
			data.positiveDegrees[i] = utils.RadiansToDegrees(positive)
			data.negativeDegrees[i] = utils.RadiansToDegrees(negative)
			data.positiveMils[i] = utils.DegreesToMils(data.positiveDegrees[i])
			data.negativeMils[i] = utils.DegreesToMils(data.negativeDegrees[i])
			fmt.Println(data.positiveDegrees[i], data.negativeDegrees[i], data.positiveMils[i], data.negativeMils[i])
		}

		table.Refresh()
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: targetDistanceLabel.Text, Widget: targetDistanceEntry},
			{Text: heightDifferenceLabel.Text, Widget: heightDifferenceEntry},
		},
		OnSubmit: process,
	}

	w.SetContent(container.NewAdaptiveGrid(1,
		form,
		table,
	))

	w.ShowAndRun()
}
