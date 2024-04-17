package usecases

// Database Configuration
const (
	DbName = "database.sqlite"
	TbName = "users_cipher"
)

// SPADE configuration
const (
	// NumUsers for now let's just assume that each user is
	// going to encrypt a single file, so NumUsers is going
	// to be the number of hypnogram files
	NumUsers = 10
	// MaxVecSize should be defined to a fixed number, where
	// the system is going to work with it by running spade's
	// setup phase. Actually the length of mpk and msk is set
	// to this value
	MaxVecSize = 1000
	// PaddingItem basically it must be anything larger than
	// the maximum value of biggest element of the vector of
	// hypnograms, in this case for us is 10, so any x > 10 works
	// why it is 22? just for fun :)
	PaddingItem = 22
)
