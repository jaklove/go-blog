package convert

import "strconv"

type StrTo string

func (s StrTo)String()string  {
	return  string(s)
}

func (s StrTo)Int()(int,error)  {
	atoi, err := strconv.Atoi(s.String())
	return atoi,err
}

func (s StrTo)MustInt()int  {
	i, _ := s.Int()
	return i
}

func (s StrTo)Uint32()(uint32,error)  {
	atoi, err := strconv.Atoi(s.String())
	return uint32(atoi),err
}

func (s StrTo)MustInt32()uint32  {
	u, _ := s.Uint32()
	return  u
}

