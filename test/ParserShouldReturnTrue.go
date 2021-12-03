package domain

import (
"testing"
)

func TestParserHappy(t *testing.T) {
tables := []struct {
fullemail string
username string
domain string
}{
{"povilas.lajus@stud.vgtu.lt", "mindaugas.gaidys", "stud.vgtu.lt"},
{"nesakysiu", "syntax_error", "syntax_error"},
{"m..g@stud.vgtu.lt", "syntax_error", "stud.vgtu.lt"},
{"hi@example.org", "hi", "example.org"},
{" @gmail.com", "syntax_error", "gmail.com"},
{"povilas@", "povilas", "syntax_error"},
}
parseTest := NewAddressParser()
for _, table := range tables {
temp, err := parseTest.Parse(table.fullemail)
if err == nil {
t.Log("username: ", temp.LocalPart, "; domain: ", temp.Domain)
}
if err != nil {
t.Error("error")
}
if temp.LocalPart != table.username {
t.Error("error")
}
if temp.Domain != table.domain {
t.Error("error")
}
}
}
