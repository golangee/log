// +build !debug

package log

// Debug is a build tag determined at build time, so that the compiler can remove dead code.
const Debug = false
