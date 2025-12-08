// Package base provides common types and utilities for LiveTemplate components.
//
// All components embed the Base struct which provides ID handling and common functionality.
// Components use functional options for configuration and implement action handlers
// that are automatically registered with the LiveTemplate framework.
package base

// Base provides common functionality for all components.
// Components should embed this struct to gain ID handling and common utilities.
//
// Example:
//
//	type Dropdown struct {
//	    base.Base
//	    Options []Option
//	    Selected *Option
//	}
type Base struct {
	// ID is the unique identifier for this component instance.
	// Used in action names like "toggle_{ID}", "select_{ID}".
	id string

	// namespace is the component type, e.g., "dropdown", "tabs".
	// Used for action naming and template resolution.
	namespace string

	// Styled indicates whether to use Tailwind CSS classes (true) or semantic HTML only (false).
	Styled bool
}

// NewBase creates a new Base with the given ID and namespace.
//
// Example:
//
//	func New(id string, opts ...Option) *Dropdown {
//	    d := &Dropdown{
//	        Base: base.NewBase(id, "dropdown"),
//	    }
//	    for _, opt := range opts {
//	        opt(d)
//	    }
//	    return d
//	}
func NewBase(id, namespace string) Base {
	return Base{
		id:        id,
		namespace: namespace,
		Styled:    true, // Default to styled (Tailwind CSS)
	}
}

// ID returns the component's unique identifier.
func (b *Base) ID() string {
	return b.id
}

// Namespace returns the component's type namespace.
func (b *Base) Namespace() string {
	return b.namespace
}

// ActionName generates a namespaced action name for this component.
// For example: ActionName("toggle") returns "toggle_myid" if ID is "myid".
func (b *Base) ActionName(action string) string {
	return action + "_" + b.id
}

// SetStyled sets whether to use Tailwind CSS classes.
func (b *Base) SetStyled(styled bool) {
	b.Styled = styled
}

// IsStyled returns true if the component should use Tailwind CSS classes.
func (b *Base) IsStyled() bool {
	return b.Styled
}
