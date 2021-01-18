package convert

import "strconv"

type Strto string
func (s Strto)String()string {
	return string(s)
}

func (s Strto)Int()(int,error)  {
	i, e := strconv.Atoi(s.String())
	return  i,e
}

func (s Strto)MustInt()int  {
	i, _ := s.Int()
	return i
}

func (s Strto)UInt32()(uint32,error)  {
	i, e := strconv.Atoi(s.String())
	return uint32(i),e
}

func (s Strto)MustUInt32()uint32  {
	v, _ := s.UInt32()
	return v
}
