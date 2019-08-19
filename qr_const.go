package goqr

/* Limits on the maximum size of QR-codes and their content. */
const
(
	QR_MAX_BITMAP  = 3917
	QR_MAX_PAYLOAD = 8896

	/* QR-code ECC types. */
	QR_ECC_LEVEL_M = 0
	QR_ECC_LEVEL_L = 1
	QR_ECC_LEVEL_H = 2
	QR_ECC_LEVEL_Q = 3

	/* QR-code data types. */
	QR_DATA_TYPE_NUMERIC = 1
	QR_DATA_TYPE_ALPHA   = 2
	QR_DATA_TYPE_BYTE    = 4
	QR_DATA_TYPE_KANJI   = 8

	/* Common character encodings */
	QR_ECI_ISO_8859_1  = 1
	QR_ECI_IBM437      = 2
	QR_ECI_ISO_8859_2  = 4
	QR_ECI_ISO_8859_3  = 5
	QR_ECI_ISO_8859_4  = 6
	QR_ECI_ISO_8859_5  = 7
	QR_ECI_ISO_8859_6  = 8
	QR_ECI_ISO_8859_7  = 9
	QR_ECI_ISO_8859_8  = 10
	QR_ECI_ISO_8859_9  = 11
	QR_ECI_WINDOWS_874 = 13
	QR_ECI_ISO_8859_13 = 15
	QR_ECI_ISO_8859_15 = 17
	QR_ECI_SHIFT_JIS   = 20
	QR_ECI_UTF_8       = 26
)
