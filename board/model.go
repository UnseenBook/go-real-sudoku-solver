package board

var filledList = [9]int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Collection holds all the Boards
type Collection []Board

// Board holds the Board itself and other info for the algorithm
type Board struct {
	grid  [9][9]field
	score int
}

type field struct {
	value         int8
	possibilities list
}

type list []int8

func (l list) remove(toRemove int8) {
	for i, value := range l {
		if l.in(value) {
			l[i] = l[len(l)-1]
			l = l[:len(l)-1]
			return
		}
	}
}

func (l list) add(r list) list {
	var newList list
	copy(newList, l)
	for _, value := range r {
		if !newList.in(value) {
			newList = append(newList, value)
		}
	}

	return newList
}

func (l list) inverse() list {
	newList := make(list, 9-len(l), 9)
	for _, value := range filledList {
		if !l.in(value) {
			newList = append(newList, value)
		}
	}

	return newList
}

func (l list) in(i int8) bool {
	for _, value := range l {
		if value == i {
			return true
		}
	}

	return false
}

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Less(i, j int) bool {
	return c[i].score < c[j].score
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
