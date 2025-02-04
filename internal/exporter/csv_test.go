package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCsv = `
name, power.draw [W]
NVIDIA GeForce RTX 2080 SUPER, 30.14 W
Some Dummy GPU, 12.34 W
`
)

func TestParseCsvIntoTable(t *testing.T) {
	t.Parallel()

	parsed, err := parseCSVIntoTable(testCsv, []qField{"name", "power.draw"})

	assert.NoError(t, err)
	assert.Len(t, parsed.rows, 2)
	assert.Equal(t, []rField{"name", "power.draw [W]"}, parsed.rFields)

	cell00 := cell{qField: "name", rField: "name", rawValue: "NVIDIA GeForce RTX 2080 SUPER"}
	cell01 := cell{qField: "power.draw", rField: "power.draw [W]", rawValue: "30.14 W"}
	cell10 := cell{qField: "name", rField: "name", rawValue: "Some Dummy GPU"}
	cell11 := cell{qField: "power.draw", rField: "power.draw [W]", rawValue: "12.34 W"}

	row0 := row{
		qFieldToCells: map[qField]cell{"name": cell00, "power.draw": cell01},
		cells:         []cell{cell00, cell01},
	}

	row1 := row{
		qFieldToCells: map[qField]cell{"name": cell10, "power.draw": cell11},
		cells:         []cell{cell10, cell11},
	}

	expected := table{
		rows:    []row{row0, row1},
		rFields: []rField{"name", "power.draw [W]"},
		qFieldToCells: map[qField][]cell{
			"name":       {cell00, cell10},
			"power.draw": {cell01, cell11},
		},
	}

	assert.Equal(t, expected, parsed)
}
