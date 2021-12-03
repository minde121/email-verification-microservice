package domain

import (
errs "github.com/pkg/errors"
"testing"
)

func TestParserSad(t *testing.T) {
tables := []struct {
fullemail string
}{
{"nesakysiu"},
{"m..g@stud.vgtu.lt"},
{"miau miau miau@gmail.com"},
{"jonas@.com"},
{"Whosabigboy?@ImaBigboy"},
{"petras@gmail.com"},
}
parseTest := NewAddressParser()
for _, table := range tables {
_, err := parseTest.Parse(table.fullemail)
if err == nil {
t.Log("Valid_address_syntax")
}
if err != nil {
if ErrSyntaxError == errs.Cause(err) {
t.Log("syntax_error")
} else {
t.Error("unexpected_error")
}
}
}
}
