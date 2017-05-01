package rest

func ErrorInternal(msg string) error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode: 500,
			ErrMessage: msg,
		},
	}
}

func ErrorBadRequest(msg string) error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode: 400,
			ErrMessage: msg,
		},
	}
}

func ErrorUnauthorized() error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode: 401,
			ErrMessage: "Unauthorized",
		},
	}
}

func ErrorNotFound(msg string) error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode: 404,
			ErrMessage: msg,
		},
	}
}


func ErrorLocked() error {
	return ErrResp{
		Meta: ErrMeta{
			ErrCode: 423,
			ErrMessage: "Locked",
		},
	}
}
