package goseed

func ThrowsServerErr(err error) {
	panic(HttpError{Code: 500, Msg: err.Error()})
}

func ThrowsBadRequestErr(err error) {
	panic(HttpError{Code: 400, Msg: err.Error()})
}
