package evaluator

import (
	"fmt"
	"ra/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1.",
					len(args))
			}

			switch argument := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(argument.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(argument.Elements))}
			default:
				return newError("argument to `len` not supported. got=%s, want={%s or %s}.",
					args[0].Type(), object.STRING_OBJ, object.ARRAY_OBJ)
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1.",
					len(args))
			}

			argument, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `first` not supported. got=%s, want=%s.",
					args[0].Type(), object.ARRAY_OBJ)
			}

			if len(argument.Elements) == 0 {
				return NULL
			}

			return argument.Elements[0]
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1.",
					len(args))
			}

			argument, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `last` not supported. got=%s, want=%s.",
					args[0].Type(), object.ARRAY_OBJ)
			}

			if len(argument.Elements) == 0 {
				return NULL
			}

			return argument.Elements[len(argument.Elements)-1]
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1.",
					len(args))
			}

			argument, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `rest` not supported. got=%s, want=%s.",
					args[0].Type(), object.ARRAY_OBJ)
			}

			length := len(argument.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, argument.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2.",
					len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `push` not supported. got=%s, want=%s.",
					args[0].Type(), object.ARRAY_OBJ)
			}

			length := len(array.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, array.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
