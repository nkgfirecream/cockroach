// Code generated by "stringer -type=serverErrFieldType"; DO NOT EDIT

package pgwire

import "fmt"

const (
	_serverErrFieldType_name_0 = "serverErrFieldSQLState"
	_serverErrFieldType_name_1 = "serverErrFieldSrcFile"
	_serverErrFieldType_name_2 = "serverErrFieldSrcLineserverErrFieldMsgPrimary"
	_serverErrFieldType_name_3 = "serverErrFieldSrcFunctionserverErrFieldSeverity"
)

var (
	_serverErrFieldType_index_0 = [...]uint8{0, 22}
	_serverErrFieldType_index_1 = [...]uint8{0, 21}
	_serverErrFieldType_index_2 = [...]uint8{0, 21, 45}
	_serverErrFieldType_index_3 = [...]uint8{0, 25, 47}
)

func (i serverErrFieldType) String() string {
	switch {
	case i == 67:
		return _serverErrFieldType_name_0
	case i == 70:
		return _serverErrFieldType_name_1
	case 76 <= i && i <= 77:
		i -= 76
		return _serverErrFieldType_name_2[_serverErrFieldType_index_2[i]:_serverErrFieldType_index_2[i+1]]
	case 82 <= i && i <= 83:
		i -= 82
		return _serverErrFieldType_name_3[_serverErrFieldType_index_3[i]:_serverErrFieldType_index_3[i+1]]
	default:
		return fmt.Sprintf("serverErrFieldType(%d)", i)
	}
}
