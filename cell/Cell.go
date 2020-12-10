package cell

func SetEmptyCells(boardSize int) *Cell {

	board := make([]string, boardSize*boardSize)
	for cell := range board {
		board[cell] = "-"
	}

	return &Cell{
		cells: board,
	}

}

func (c *Cell) Cells() *[]string {
	return &c.cells
}

// MarkCellIfEmpty marks the respective cell if it is empty
func (c *Cell) MarkCellIfEmpty(location int, playerMark string) error {

	if (c.cells)[location] != "-" {
		return &cellAlreadyOccupied{cellOccupied: "This cell is occupied, enter a different cell number"}
	}
	(c.cells)[location] = playerMark
	return nil

}

func (err *cellAlreadyOccupied) Error() string {
	return err.cellOccupied
}

// CellAlreadyOccupied throws error if cell is already occupied
type cellAlreadyOccupied struct {
	cellOccupied string
}

// Cell will contain each cell on board
type Cell struct {
	cells []string
}
