package cell

// MarkCellIfEmpty marks the respective cell if it is empty
func MarkCellIfEmpty(board *[]string, location int, playerMark string) error {

	if (*board)[location] != "-" {
		return &cellAlreadyOccupied{cellOccupied: "This cell is occupied, enter a different cell number"}
	}
	(*board)[location] = playerMark
	return nil

}

func (err *cellAlreadyOccupied) Error() string {
	return err.cellOccupied
}

// CellAlreadyOccupied throws error if cell is already occupied
type cellAlreadyOccupied struct {
	cellOccupied string
}
