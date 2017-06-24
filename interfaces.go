package strumt

// Prompter defines a generic common prompt
type Prompter interface {
	PromptString() string
	NextOnError(error) string
}

// LinePrompter defines a one line prompter
// that will ask only for one user input.
// To mark prompter as ending prompter
// NextOnSucces must returns an empty string
type LinePrompter interface {
	Prompter
	NextOnSuccess(string) string
	Parse(string) error
}

// MultilinePrompter defines a mutiline prompter
// that will let the possibility to the user to
// provide several input, result is provided as
// an input slice.
// To mark prompter as ending prompter
// NextOnSucces must returns an empty string
type MultilinePrompter interface {
	Prompter
	NextOnSuccess([]string) string
	Parse([]string) error
}

// PromptRenderer can be implemented to customize
// the way prompt is rendered, prompt string given by
// PromptString is given as parameter
type PromptRenderer interface {
	PrintPrompt(string)
}

// ErrorRenderer can be implemented to customize
// the way an error return by Parse is rendered
type ErrorRenderer interface {
	PrintError(err error)
}
