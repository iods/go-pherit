package globals

// NOTE: Globals, in the traditional sense, are bad. But there are some
// situations where we want a value that is constant per invocation but
// invocation may be in various places, such as a test or via CLI, and placing
// such a value in any given package leads to import issues. So think of this
// package as "freeing us from circular dependency hell". These values should
// only ever be set at startup, but simply available to reference from anywhere.

var (
	server = false // Set to true if the Go Pher-it project will be acting as a server
)