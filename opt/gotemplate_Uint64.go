package opt

import (
	"fmt"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Uint64 struct {
	V       uint64
	Defined bool
}

// Creates an optional type with a given value.
func OUint64(v uint64) Uint64 {
	return Uint64{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Uint64) Get(deflt uint64) uint64 {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Uint64) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Uint64(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Uint64) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Uint64{}
	} else {
		v.V = l.Uint64()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v *Uint64) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalJSON implements a standard json marshaler interface.
func (v *Uint64) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Uint64) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Uint64) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
