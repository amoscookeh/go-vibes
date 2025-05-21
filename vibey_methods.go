package vibes

// VibeyMethod represents the vibey alternatives to standard HTTP methods
type VibeyMethod string

const (
	// VIBE is the vibey alternative to GET
	VIBE VibeyMethod = "VIBE"

	// MANIFEST is the vibey alternative to POST
	MANIFEST VibeyMethod = "MANIFEST"

	// ALIGN is the vibey alternative to PUT
	ALIGN VibeyMethod = "ALIGN"

	// RELEASE is the vibey alternative to DELETE
	RELEASE VibeyMethod = "RELEASE"
)

// ToHTTPMethod converts a VibeyMethod to the corresponding standard HTTP method
func (v VibeyMethod) ToHTTPMethod() string {
	switch v {
	case VIBE:
		return "GET"
	case MANIFEST:
		return "POST"
	case ALIGN:
		return "PUT"
	case RELEASE:
		return "DELETE"
	default:
		return string(v)
	}
}

// FromHTTPMethod converts a standard HTTP method to the corresponding VibeyMethod
func FromHTTPMethod(method string) VibeyMethod {
	switch method {
	case "GET":
		return VIBE
	case "POST":
		return MANIFEST
	case "PUT":
		return ALIGN
	case "DELETE":
		return RELEASE
	default:
		return VibeyMethod(method)
	}
}
