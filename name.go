package fate

import (
	"fmt"
	"github.com/godcong/chronos"
	"github.com/godcong/yi"
	"strconv"
	"strings"
)

//Name 姓名
type Name struct {
	FirstName   []*Character //名姓
	LastName    []*Character
	born        *chronos.Calendar
	baZi        *BaZi
	baGua       *yi.Yi //周易八卦
	zodiac      *Zodiac
	score       int  // 网络名称评分
	poemNum     int  // 包含的诗词量
	poem        Poem // 诗词
	zodiacPoint int
}

// String ...
func (n Name) String() string {
	var s string
	for _, l := range n.LastName {
		s += l.Ch
	}
	for _, f := range n.FirstName {
		s += f.Ch
	}
	return s
}

func (n Name) getFirstName() string {
	var s string
	for _, f := range n.FirstName {
		s += f.Ch
	}
	return s
}

// Strokes ...
func (n Name) Strokes() string {
	var s []string
	for _, l := range n.LastName {
		s = append(s, strconv.Itoa(l.ScienceStroke))
	}

	for _, f := range n.FirstName {
		s = append(s, strconv.Itoa(f.ScienceStroke))
	}
	return strings.Join(s, ",")
}

// PinYin ...
func (n Name) PinYin() string {
	var s string
	for _, l := range n.LastName {
		s += "[" + strings.Join(l.PinYin, ",") + "]"
	}

	for _, f := range n.FirstName {
		s += "[" + strings.Join(f.PinYin, ",") + "]"
	}
	return s
}

// WuXing ...
func (n Name) WuXing() string {
	var s string
	for _, l := range n.LastName {
		s += l.WuXing
	}
	for _, f := range n.FirstName {
		s += f.WuXing
	}
	return s
}

// XiYongShen ...
func (n Name) XiYongShen() string {
	return n.baZi.XiYongShen()
}

func (n Name) getScore() string {
	return fmt.Sprintf("%d", n.score)
}

func (n Name) getPoet() string {
	return fmt.Sprintf("%s(%s)", n.poem.Poet, n.poem.Decade)
}

func (n Name) getPoem() string {
	return n.poem.Poem
}

func (n Name) getPoemContent() string {
	return n.poem.Content
}

func (n Name) getPoemNum() string {
	return fmt.Sprintf("%d", n.poemNum)
}

func createName(impl *fateImpl, f1 *Character, f2 *Character) *Name {
	lastSize := len(impl.lastChar)
	last := make([]*Character, lastSize, lastSize)
	copy(last, impl.lastChar)
	ff1 := *f1
	ff2 := *f2
	first := []*Character{&ff1, &ff2}

	return &Name{
		FirstName: first,
		LastName:  last,
	}
}

// BaGua ...
func (n *Name) BaGua() *yi.Yi {
	if n.baGua == nil {
		lastSize := len(n.LastName)
		shang := getStroke(n.LastName[0])
		if lastSize > 1 {
			shang += getStroke(n.LastName[1])
		}
		xia := getStroke(n.FirstName[0]) + getStroke(n.FirstName[1])
		n.baGua = yi.NumberQiGua(xia, shang, shang+xia)
	}

	return n.baGua
}

// BaZi ...
func (n Name) BaZi() string {
	return n.baZi.String()
}
