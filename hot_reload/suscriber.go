package hotreload

// Suscriber
type Suscriber interface {
	Update([]byte) error
}
