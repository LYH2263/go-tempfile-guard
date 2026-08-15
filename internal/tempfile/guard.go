package tempfile

import "os"

type Guard struct {
	Tmp   string
	Final string
}

func New(tmp, final string) *Guard {
	return &Guard{Tmp: tmp, Final: final}
}

func (g *Guard) Save(fn func(tmp string) error) error {
	tmp := g.Tmp
	defer func() {
		_ = os.Remove(tmp)
	}()
	if err := fn(tmp); err != nil {
		return err
	}
	return os.Rename(tmp, g.Final)
}

