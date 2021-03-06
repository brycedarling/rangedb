package rangedberror

import (
	"fmt"
	"strings"
)

// UnexpectedSequenceNumber is an error containing expected and next sequence numbers.
type UnexpectedSequenceNumber struct {
	Expected           uint64
	NextSequenceNumber uint64
}

// NewUnexpectedSequenceNumberFromString constructs an UnexpectedSequenceNumber error.
func NewUnexpectedSequenceNumberFromString(input string) *UnexpectedSequenceNumber {
	pieces := strings.Split(input, "unexpected sequence number:")
	if len(pieces) < 2 {
		return &UnexpectedSequenceNumber{}
	}

	var expected, next uint64
	_, err := fmt.Sscanf(pieces[1], "%d, next: %d", &expected, &next)
	if err != nil {
		return &UnexpectedSequenceNumber{}
	}

	return &UnexpectedSequenceNumber{
		Expected:           expected,
		NextSequenceNumber: next,
	}
}

func (e UnexpectedSequenceNumber) Error() string {
	return fmt.Sprintf("unexpected sequence number: %d, next: %d",
		e.Expected,
		e.NextSequenceNumber,
	)
}
