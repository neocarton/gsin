package gsin

type (
	// SecurityError user is not allowed to access data or features
	SecurityError struct {
		baseSin
	}

	// ParameterInvalid inputs are invalid
	ParameterInvalid struct {
		baseSin
	}

	// DataNotFound data not found
	DataNotFound struct {
		baseSin
	}

	// DataConflict intput conflict with existing data
	DataConflict struct {
		baseSin
	}

	// Error system error
	Error struct {
		baseSin
	}
)
