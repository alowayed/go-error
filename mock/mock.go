package mock

type (
	Context struct{}
)

func (c *Context) JSON(statusCode int, obj interface{}) {}
