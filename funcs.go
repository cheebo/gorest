package rest

func ErrorInternal(msg string) error {
	if len(msg) == 0 {
		msg = "Internal error"
	}
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    500,
			ErrMessage: msg,
		},
	}
}

func ErrorBadRequest(msg string) error {
	if len(msg) == 0 {
		msg = "Bad request"
	}
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    400,
			ErrMessage: msg,
		},
	}
}

func ErrorUnauthorized() error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    401,
			ErrMessage: "Unauthorized",
		},
	}
}

func AccessForbidden() error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    403,
			ErrMessage: "Forbidden",
		},
	}
}

func ErrorNotFound(msg string) error {
	if len(msg) == 0 {
		msg = "Not found"
	}
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    404,
			ErrMessage: msg,
		},
	}
}

func ErrorConflict(msg string) error {
	if len(msg) == 0 {
		msg = "Conflict"
	}
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    409,
			ErrMessage: msg,
		},
	}
}

func ErrorLocked(msg string) error {
	if len(msg) == 0 {
		msg = "Locked"
	}
	return ErrResp{
		Meta: ErrMeta{
			ErrCode:    423,
			ErrMessage: msg,
		},
	}
}
