package gsin

type (
	// Sin error
	Sin interface {
		error
		Message() string
		SetMessage(string)
		Context() map[string]interface{}
		SecretContext() string
		SetContext(map[string]interface{})
		Cause() error     // Get the nearest cause
		SetCause(error)   // Set the nearest cause
		Causes() []error  // Get all causes, the last one is the root cause
		RootCause() error // Get the root cause
	}

	// baseSin base error
	baseSin struct {
		message       string
		context       map[string]interface{}
		secretContext string
		causes        []error
		cause         error
		rootCause     error
	}
)

// InitError new Error
func InitError(err error, message string, cause error, context map[string]interface{}) error {
	sin := err.(Sin)
	sin.SetMessage(message)
	sin.SetContext(context)
	sin.SetCause(cause)
	return err
}

func (err baseSin) Error() string {
	return err.message
}

func (err *baseSin) Message() string {
	return err.message
}

func (err *baseSin) SetMessage(message string) {
	err.message = message
}

func (err *baseSin) Context() map[string]interface{} {
	return err.context
}

func (err *baseSin) SetContext(context map[string]interface{}) {
	err.context = context
}

func (err *baseSin) SecretContext() string {
	return err.secretContext
}

// SetSecretContext set (unencrypted) context to secret-context
func (err *baseSin) SetSecretContext(context string) {
	// TODO encrypt context
	err.secretContext = context
}

func (err *baseSin) Cause() error {
	return err.cause
}

func (err *baseSin) SetCause(cause error) {
	if cause == nil {
		return
	}
	err.causes = newCauses(cause)
	causeSize := len(err.causes)
	if causeSize > 0 {
		err.cause = err.causes[0]
		err.rootCause = err.causes[causeSize-1]
	}
}

func (err *baseSin) Causes() []error {
	return err.causes
}

func (err *baseSin) RootCause() error {
	return err.rootCause
}

// newCauses create error stack trace
func newCauses(err error) []error {
	causes := []error{err}
	if sin, ok := err.(Sin); ok {
		causes = append(causes, sin.Causes()...)
	}
	return causes
}
