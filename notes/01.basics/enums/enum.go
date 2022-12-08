package enums

type Month int

const (
	January Month = iota + 1
	February
	March
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) String() string {
	switch m {
	case January:
		return "january"
	case February:
		return "february"
	case March:
		return "march"
	case May:
		return "may"
	case June:
		return "june"
	case July:
		return "july"
	case August:
		return "august"
	case September:
		return "september"
	case October:
		return "october"
	case November:
		return "november"
	case December:
		return "december"
	}
	return "unknown"
}

func MyBirthMonth() Month {
	return December
}
