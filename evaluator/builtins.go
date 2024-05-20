package evaluator

import "ra/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1.",
					len(args))
			}

			argument, ok := args[0].(*object.String)
			if !ok {
				return newError("argument to `len` not supported. got=%s, want=%s.",
					args[0].Type(), object.STRING_OBJ)
			}

			return &object.Integer{Value: int64(len(argument.Value))}
		},
	},
}
