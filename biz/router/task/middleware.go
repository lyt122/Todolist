// Code generated by hertz generator.

package task

import (
	"Todolist/biz/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.JwtMiddleware.MiddlewareFunc(),
	}
}

func _taskMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletetaskbystatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatetaskMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querytaskbystatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createtaskMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteonetaskMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querytaskbykeywordMw() []app.HandlerFunc {
	// your code...
	return nil
}
