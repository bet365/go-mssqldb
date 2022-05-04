package ntlm

import "github.com/microsoft/go-mssqldb/integratedauth"

// AuthProvider handles NTLM SSPI Windows Authentication
var AuthProvider integratedauth.Provider = integratedauth.ProviderFunc(getAuth)
