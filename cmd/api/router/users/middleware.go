// Code generated by hertz generator.

package users

import (
	"HuaTug.com/cmd/api/router/authfunc"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteuserMw() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}

func _getuserinfoMw() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}

func _loginuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserMw() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}

func _checkUserExistsByIdMv() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}

func _verifycodeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _sendcodeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryuserMw() []app.HandlerFunc {
	// your code...
	return authfunc.Auth()
}
