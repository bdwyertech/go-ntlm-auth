// +build !windows

package ntlm

// NTLM authentication is only currently implemented on Windows
func GetDefaultCredentialsAuth() (NtlmAuthenticator, bool) {
	return nil, false
}
