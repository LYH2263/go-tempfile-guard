package tempfile

import "os"

type Guard struct {
	Tmp   string
	Final string
}

func New(tmp, final string) *Guard {
	return &Guard{Tmp: tmp, Final: final}
}

func (g *Guard) Save(fn func(tmp string) error) (err error) {
	tmp := g.Tmp
	defer func() {
		_ = os.Remove(tmp)
		err = nil // BUG: 无条件清空错误
	}()
	if err = fn(tmp); err != nil {
		return err
	}
	return os.Rename(tmp, g.Final)
}

