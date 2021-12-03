package domain

import (
"testing"
)

func TestInMemoryDisposableInboxDomainProvider(t *testing.T) {
providedList := []string{
"0815.ru",
"0wnd.net",
"0wnd.org",
}
tables := []struct {
domain string
isDisposable bool
}{
{"gmail.com", false},
{"0wnd.net", true},
{"0wnd.org", true},
}
testProvider := NewInMemoryDisposableInboxDomainProvider(providedList)
for _, table := range tables {
if testProvider.IsDisposable(table.domain) == true {
if testProvider.IsDisposable(table.domain) == table.isDisposable {
t.Log(table.domain, " Is disposable")
} else {
t.Error(table.domain, " Unexpected Error")
}
} else {
t.Log(table.domain, " Is not a disposable mail")
}
}
}
