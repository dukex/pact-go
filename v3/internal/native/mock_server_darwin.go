package native

/*
#cgo LDFLAGS: ${SRCDIR}/../../../libs/libpact_mock_server_ffi.dylib

// Library headers
typedef int bool;
#define true 1
#define false 0

void init(char* log);
int create_mock_server(char* pact, char* addr, bool tls);
int mock_server_matched(int port);
char* mock_server_mismatches(int port);
bool cleanup_mock_server(int port);
int write_pact_file(int port, char* dir);
void free_string(char* s);
char* get_tls_ca_certificate();
int create_plugin_mock_server(int port, char* config);
int add_plugin_interaction(int port, char* config);
char* plugin_mock_server_mismatches(int port);
int plugin_mock_server_matched(int port);
bool cleanup_plugin_mock_server(int port);

*/
import "C"
