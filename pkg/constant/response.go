package constant

const (
	SUCCESS               = 20000
	VALIDATE_ERROR        = 40000
	NOT_FOUND             = 40400
	INTERNAL_SERVER_ERROR = 50000
)

var CodeMsg = map[int]string{
	SUCCESS:               "success",
	VALIDATE_ERROR:        "validate error",
	NOT_FOUND:             "not found",
	INTERNAL_SERVER_ERROR: "server unknown error",
}
