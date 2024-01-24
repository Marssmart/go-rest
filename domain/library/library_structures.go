package library

import "strings"

type Library interface {
	Borrow(name string) bool
	Return(name string) bool
	List() []string
	ListAvailable() []string
	Exists(name string) bool
	Add(b Book)
}

type library struct {
	availability map[string]int
}

type Book struct {
	Name  string
	Count int
}

var DefaultLibrary = CreateDefaultLibrary()

func CreateDefaultLibrary() Library {
	lib := &library{make(map[string]int)}
	lib.availability["hp"] = 7
	lib.availability["lotr"] = 3
	return lib
}

func (l *library) Borrow(name string) bool {
	nameLower := strings.ToLower(name)
	bookAvailability, ok := l.availability[nameLower]
	if ok && bookAvailability > 0 {
		l.availability[nameLower] = bookAvailability - 1
		return true
	}
	return false
}

func (l *library) Return(name string) bool {
	nameLower := strings.ToLower(name)
	bookAvailability, ok := l.availability[nameLower]
	if ok {
		l.availability[nameLower] = bookAvailability + 1
		return true
	}
	return false
}

func (l *library) List() []string {
	keys := make([]string, 0, len(l.availability))
	for k := range l.availability {
		keys = append(keys, k)
	}

	return keys
}

func (l *library) ListAvailable() []string {
	keys := make([]string, 0, len(l.availability))
	for k := range l.availability {
		if l.availability[k] > 0 {
			keys = append(keys, k)
		}
	}

	return keys
}

func (l *library) Exists(name string) bool {
	_, ok := l.availability[name]
	return ok
}

func (l *library) Add(b Book) {
	l.availability[b.Name] = b.Count
}

func (b *Book) Valid() bool {
	return len(b.Name) > 0 && b.Count > 0
}
