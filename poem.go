package fate

import (
	"fmt"
	"github.com/xormsharp/builder"
	"github.com/xormsharp/xorm"
)

//Character 字符
type Poem struct {
	Wid                      int      `xorm:"pk wid"`
	Poet                     string   `xorm:"default() notnull poet"`                               //诗人
	Decade            		 string   `xorm:"default() notnull decade"` 							//朝代
	Poem                     string   `xorm:"default() notnull poem"`                               //诗名
	Content                  string   `xorm:"default() notnull content"`                            //内容
}

func getPoems(engine *xorm.Engine, fn func(engine *xorm.Engine) *xorm.Session) ([]*Poem, error) {
	s := fn(engine)
	var c []*Poem
	e := s.Find(&c)
	if e == nil {
		return c, nil
	}
	return nil, fmt.Errorf("character list get error:%w", e)
}

func poem(name string) func(engine *xorm.Engine) *xorm.Session {
	return func(engine *xorm.Engine) *xorm.Session {
		return engine.Where(builder.Like{"content", name})
	}
}
