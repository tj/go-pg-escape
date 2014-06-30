package escape

import "github.com/bmizerany/assert"
import "testing"

func TestEscape(t *testing.T) {
	{
		s := Escape("SELECT %I FROM %I WHERE %I=%L", "some stuff", "some table", "some column", "some value")
		exp := `SELECT "some stuff" FROM "some table" WHERE "some column"='some value'`
		assert.Equal(t, exp, s)
	}

	{
		s := Escape("COPY %s", "something")
		assert.Equal(t, "COPY something", s)
	}

	{
		s := Escape("COPY something")
		assert.Equal(t, "COPY something", s)
	}
}

func TestLiteral(t *testing.T) {
	{
		s := Literal(`hey`)
		assert.Equal(t, `'hey'`, s)
	}

	{
		s := Literal(`Hello World`)
		assert.Equal(t, `'Hello World'`, s)
	}

	{
		s := Literal(`O'Reilly`)
		assert.Equal(t, `'O''Reilly'`, s)
	}

	{
		s := Literal(`\\whoop\\`)
		assert.Equal(t, `E'\\\\whoop\\\\'`, s)
	}
}

func TestIdent(t *testing.T) {
	{
		s := Ident(`foo`)
		assert.Equal(t, `foo`, s)
	}

	{
		s := Ident(`_foo`)
		assert.Equal(t, `_foo`, s)
	}

	{
		s := Ident(`foo_$_bar`)
		assert.Equal(t, `foo_$_bar`, s)
	}

	{
		s := Ident(`test.some.stuff`)
		assert.Equal(t, `"test.some.stuff"`, s)
	}

	{
		s := Ident(`test."some".stuff`)
		assert.Equal(t, `"test.""some"".stuff"`, s)
	}

	{
		s := Ident(`join`)
		assert.Equal(t, `"join"`, s)
	}

	{
		s := Ident(`desc`)
		assert.Equal(t, `"desc"`, s)
	}
}
