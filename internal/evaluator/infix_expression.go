package evaluator

import "github.com/Ryoga-exe/monkey/internal/object"

type InfixPair struct {
	Left  object.ObjectType
	Right object.ObjectType
}

type InfixFunction func(left, right object.Object) object.Object

var infixOperator = map[InfixPair]map[string]InfixFunction{
	{Left: object.INTEGER_OBJ, Right: object.INTEGER_OBJ}: {
		"+": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return &object.Integer{Value: leftVal + rightVal}
		},
		"-": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return &object.Integer{Value: leftVal - rightVal}
		},
		"*": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return &object.Integer{Value: leftVal * rightVal}
		},
		"/": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return &object.Integer{Value: leftVal / rightVal}
		},
		"<": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return nativeBoolToBooleanObject(leftVal < rightVal)
		},
		">": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return nativeBoolToBooleanObject(leftVal > rightVal)
		},
		"==": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return nativeBoolToBooleanObject(leftVal == rightVal)
		},
		"!=": func(left, right object.Object) object.Object {
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			return nativeBoolToBooleanObject(leftVal != rightVal)
		},
	},
	{Left: object.STRING_OBJ, Right: object.STRING_OBJ}: {
		"+": func(left, right object.Object) object.Object {
			leftVal := left.(*object.String).Value
			rightVal := right.(*object.String).Value
			return &object.String{Value: leftVal + rightVal}
		},
		"==": func(left, right object.Object) object.Object {
			leftVal := left.(*object.String).Value
			rightVal := right.(*object.String).Value
			return nativeBoolToBooleanObject(leftVal == rightVal)
		},
		"!=": func(left, right object.Object) object.Object {
			leftVal := left.(*object.String).Value
			rightVal := right.(*object.String).Value
			return nativeBoolToBooleanObject(leftVal != rightVal)
		},
	},
}
