package php

import (
	_ "embed"
)

func mapScalaType(protoType string) string {
	switch protoType {
	case "uint64", "sint64", "int64", "fixed64", "sfixed64", "string", "int32", "sint32", "uint32", "fixed32", "sfixed32":
		return "int"
	case "float", "double":
		return "double"
	case "bool":
		return "boolean"
	case "bytes":
		return "string"
	}

	return ""
}
