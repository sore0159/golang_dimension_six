package models

type Status struct {
	Durability byte
	Sprint     bool
	CoverMove  bool
	InvCover   bool
	SpotLvl    byte
	AvailRange [2]bool
}

func (f *Frame) Status() Status {
	st := Status{
		Durability: byte(f.Basics + len(f.Systems)),
		Sprint:     true,
	}
	if f.Rockets > 0 {
		st.AvailRange[0] = true
	}
	var hasBlue bool
	for _, sy := range f.Systems {
		switch sy {
		case MOBILITY:
			st.CoverMove = true
		case DEFENSE:
			if hasBlue {
				st.InvCover = true
			} else {
				hasBlue = true
			}
		case SUPPORT:
			st.SpotLvl += 1
		//case MELEE:
		case DIRECT_FIRE:
			st.Sprint = false
			st.AvailRange[0] = true
		case ARTILLERY:
			st.Sprint = false
			st.AvailRange[1] = true
		}
	}
	if st.Sprint {
		st.CoverMove = true
	}
	return st
}
