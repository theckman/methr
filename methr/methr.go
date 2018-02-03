package methr

// NonPointer will have no pointer methods.
type NonPointer struct {
	Value int
}

// Set is set
func (np NonPointer) Set(i int) {
	np.Value = i
}

// Get is get
func (np NonPointer) Get() int {
	return np.Value
}

// Pointer will have only pointer methods.
type Pointer struct {
	Value int
}

// Set is set
func (p *Pointer) Set(i int) {
	p.Value = i
}

// Get is get
func (p *Pointer) Get() int {
	return p.Value
}

// Both will have both pointer and non-pointer methods.
type Both struct {
	Value int
}

// Set is set
func (b *Both) Set(i int) {
	b.Value = i
}

// Get is get
func (b Both) Get() int {
	return b.Value
}
