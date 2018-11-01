package routererror

type ResourceNotFoundError struct{}

func (ResourceNotFoundError) Error() string {
	return "resource not found"
}
