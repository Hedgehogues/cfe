package core_

import (
	"github.com/pkg/errors"
	"strings"
)


var InternalError = errors.New("problem with internal logic")
var BadHypertextError = errors.New("hypertext was changed")
var NotFoundObjectError = errors.New("parser is not found first anchor of requested element")
var NotFoundCtxError = errors.New("parser is not found first anchor of context element")


// anchor is pairs substring for detect of object
// Example: <x>a</x>
// In this case we must detect object "a" by ctx anchor "<x>" and finish anchor "</x>"
type Anchor struct {
	start  string
	finish string
}

// NewAnchor creates new anchor
func NewAnchor(s, f string) *Anchor {
	return &Anchor{
		start:  s,
		finish: f,
	}
}

// anchor is pairs substring for detect of object
// Example: <x>Hello world!<z><x>a</x></z>
// In this case we must detect object "a" by ctx anchor "<x>" and finish anchor "</x>" with context <z>. Etc., first,
// we detect <z> and after that, we detect <x>
type CtxAnchor struct {
	ctx    string
	anchor *Anchor
}

// NewCtxAnchor creates new context anchor. If you need to detect object since position of text, you need to use context
// Extractor
func NewCtxAnchor(ctxSubstr string, anchor *Anchor) *CtxAnchor {
	return &CtxAnchor{
		ctx:    ctxSubstr,
		anchor: anchor,
	}
}

// Extract is container for data. This one intend to save data about extracted object of text and position them in the
// text
type Extract struct {
	Object string // extracted Object
	SPos   int    // ctx position
	FPos   int    // finish position
}

func initCounts(count *uint32) uint32 {
	c := uint32(0)
	if count != nil {
		c = *count
	}
	return c
}

// ExtractObject extracts object by anchor
func ExtractObject(text string, anchor *Anchor) (*Extract, error) {
	i0 := strings.Index(text, anchor.start)
	if i0 == -1 {
		return nil, NotFoundObjectError
	}
	i0 = i0 + len(anchor.start)

	i1 := strings.Index(text[i0:], anchor.finish)
	if i1 == -1 {
		return nil, BadHypertextError
	}

	newPos := i0 + i1
	object := text[i0:newPos]

	e := Extract{
		FPos:   newPos,
		Object: object,
		SPos:   i0,
	}

	return &e, nil
}

func ExtractCtxObject(text string, ctxAnchor *CtxAnchor) (*Extract, error) {
	a := NewAnchor(ctxAnchor.ctx, "")
	e, err := ExtractObject(text, a)
	if err == NotFoundObjectError {
		return nil, NotFoundCtxError
	}
	if err != nil {
		return nil, err
	}
	subPage := text[e.FPos:]

	descr, err := ExtractObject(subPage, ctxAnchor.anchor)
	if err != nil {
		return nil, err
	}
	descr.FPos += e.FPos-1
	descr.SPos += e.FPos-1

	return descr, nil
}

// ExtractCtxObjects detects objects by ctx and finish anchor. This function used to ExtractObject
//
// Parameters. count is amount of objects. This is optional parameter. If you set them, it is optimized memory
// allocation.
func ExtractObjects(text string, anchor *Anchor, count *uint32) ([]*Extract, error) {
	c := initCounts(count)
	names := make([]*Extract, 0, c)
	subText := text // source text needs for check corrects extraction objects from text
	offset := 0     // this offset set shift by text for different ids
	for {
		// notice. this is stateful function. Each next text takes with offset
		extract, err := ExtractObject(subText, anchor)
		if err == NotFoundObjectError {
			break
		}
		if err != nil {
			return nil, err
		}
		subText = subText[extract.FPos:]
		newOffset := offset + extract.FPos
		extract.SPos += offset
		extract.FPos += offset
		offset = newOffset
		if text[extract.SPos:extract.FPos] != extract.Object {
			// This condition checks correct extraction of Object (this is assertion)
			return nil, InternalError
		}
		names = append(names, extract)
	}
	return names, nil
}

// ExtractCtxObjects detects context and after that detects objects in this context. This method is used to context
// anchor. Context anchor consists object for detect. This function used to ExtractObject. Notice: there are some
// differences from ExtractObjects
//
// Parameters. count is amount of objects. This is optional parameter. If you set them, it is optimized memory
// allocation.
func ExtractCtxObjects(text string, ctxAnchor *CtxAnchor, count *uint32) ([]*Extract, error) {
	c := initCounts(count)
	names := make([]*Extract, 0, c)
	subText := text // source text needs for check corrects extraction objects from text
	offset := 0     // this offset set shift by text for different ids
	ctxA := NewAnchor(ctxAnchor.ctx, "")
	for {
		ctxE, err := ExtractObject(subText, ctxA)
		if err == NotFoundObjectError {
			break
		}
		if err != nil {
			return nil, err
		}
		subText = subText[ctxE.FPos:]
		offset := offset + ctxE.FPos

		e, err := ExtractObject(subText, ctxAnchor.anchor)
		if err != nil {
			return nil, err
		}
		subText = subText[e.FPos:]
		newOffset := offset + e.FPos
		e.SPos += offset
		e.FPos += offset
		offset = newOffset
		if text[e.SPos:e.FPos] != e.Object {
			// This condition checks correct extraction of Object (this is assertion)
			// This case is unreachable
			return nil, InternalError
		}
		names = append(names, e)
	}
	return names, nil
}

// TODO: name "CFE" (Context, For-loop, Extract). Declarative style:
//  Ctx(For(Ctx(Extract())))
