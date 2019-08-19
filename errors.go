package goqr

import "errors"

const (
	QR_SUCCESS = iota
	QR_ERROR_NO_QRCODE
	QR_ERROR_INVALID_GRID_SIZE
	QR_ERROR_INVALID_VERSION
	QR_ERROR_FORMAT_ECC
	QR_ERROR_DATA_ECC
	QR_ERROR_UNKNOWN_DATA_TYPE
	QR_ERROR_DATA_OVERFLOW
	QR_ERROR_DATA_UNDERFLOW
)

var (
	ERR_NO_QR_CODE        = errors.New("no QR code in image")
	ERR_INVALID_GRID_SIZE = errors.New("invalid grid size")
	ERR_INVALID_VERSION   = errors.New("invalid version")
	ERR_FORMAT_ECC        = errors.New("ecc format error")
	ERR_DATA_ECC          = errors.New("ecc data error")
	ERR_UNKNOWN_DATA_TYPE = errors.New("unknown data type")
	ERR_DATA_OVERFLOW     = errors.New("data overflow")
	ERR_DATA_UNDERFLOW    = errors.New("data underflow")
)
var errMap map[int]error

func init() {
	errMap = make(map[int]error)
	errMap [QR_ERROR_NO_QRCODE] = ERR_NO_QR_CODE
	errMap[QR_ERROR_INVALID_GRID_SIZE] = ERR_INVALID_GRID_SIZE
	errMap[QR_ERROR_INVALID_VERSION] = ERR_INVALID_VERSION
	errMap[QR_ERROR_FORMAT_ECC] = ERR_FORMAT_ECC
	errMap[QR_ERROR_DATA_ECC] = ERR_DATA_ECC
	errMap[QR_ERROR_UNKNOWN_DATA_TYPE] = ERR_UNKNOWN_DATA_TYPE
	errMap[QR_ERROR_DATA_OVERFLOW] = ERR_DATA_OVERFLOW
	errMap[QR_ERROR_DATA_UNDERFLOW] = ERR_DATA_UNDERFLOW
}

func Err(code int) error {
	return errMap[code]
}
