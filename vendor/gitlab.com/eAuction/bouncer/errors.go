package bouncer

type badrequest struct {
	msg string
}

func (e *badrequest) Error() string {
	return e.msg
}

func (*badrequest) BadRequest() bool {
	return true
}

type notallowed struct {
	msg string
}

func (e *notallowed) Error() string {
	return e.msg
}

func (*notallowed) NotAllowed() bool {
	return true
}
