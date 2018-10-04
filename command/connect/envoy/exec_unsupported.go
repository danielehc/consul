// +build !linux,!darwin

package envoy

func execEnvoy(binary string, passThroughArgs []string, bootstrapJson []byte) error {
	return errUnsupportedOS
}
