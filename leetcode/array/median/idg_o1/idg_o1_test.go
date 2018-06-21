package idg_o1

import "testing"

/*
["RandomizedSet","insert","remove","insert","remove","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom","getRandom"]
[[],[0],[0],[-1],[0],[],[],[],[],[],[],[],[],[],[]]
*/
func Test(t *testing.T) {
	s := Constructor()
	s.Insert(0)
	s.Remove(0)
	s.Insert(-1)
	s.Remove(0)
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
	s.GetRandom()
}
