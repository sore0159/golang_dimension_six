package models

type System int

const (
	MOBILITY    System = iota + 1
	DEFENSE     System = iota + 1
	SUPPORT     System = iota + 1
	MELEE       System = iota + 1
	DIRECT_FIRE System = iota + 1
	ARTILLERY   System = iota + 1
)
const (
	COLOR_WHITE = iota
	COLOR_GREEN
	COLOR_BLUE
	COLOR_YELLOW
	COLOR_RED
)

func (m System) Color() int {
	switch m {
	case MOBILITY:
		return COLOR_GREEN
	case DEFENSE:
		return COLOR_BLUE
	case SUPPORT:
		return COLOR_YELLOW
	default:
		//case MELEE, DIRECT_FIRE, ARTILLERY:
		return COLOR_RED
	}
}
