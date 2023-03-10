package core

type EngineName struct {
	name string
}

func NewEngineName(name string) EngineName {
	return EngineName{
		name: name,
	}
}

func (e EngineName) GetEngineName() string {
	return e.name
}
