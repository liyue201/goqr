package goqr

const (
	QR_PIXEL_WHITE  = 0
	QR_PIXEL_BLACK  = 1
	QR_PIXEL_REGION = 2

	QR_MAX_REGIONS = 254

	QR_MAX_CAPSTONES = 32
	QR_MAX_GRIDS     = 8

	QR_PERSPECTIVE_PARAMS = 8
)

type qr_pixel_t = uint8

type Point struct {
	x int
	y int
}

type QrRegion struct {
	seed     Point
	count    int
	capstone int
}

type QrCapstone struct {
	ring    int
	stone   int
	corners [4]Point
	center  Point
	c       [QR_PERSPECTIVE_PARAMS] float64
	qrGrid  int
}

type QrGrid struct {
	/* Capstone indices */
	caps [3] int

	/* Alignment pattern region and corner */
	alignRegion int
	align       Point

	/* Timing pattern endpoints */
	tpep  [3] Point
	hscan int
	vscan int

	/* Grid size and perspective transform */
	gridSize int
	c        [QR_PERSPECTIVE_PARAMS] float64
}

/************************************************************************
 * QR-code Version information database
 */

const (
	QR_MAX_VERSION   = 40
	QR_MAX_ALIGNMENT = 7
)

type QrRsParams struct {
	bs int /* Small block size */
	dw int /* Small data words */
	ns int /* Number of small blocks */
}

type QrVersionInfo struct {
	dataBytes int
	apat      [QR_MAX_ALIGNMENT] int
	ecc       [4]  QrRsParams
}

type polygonScoreData struct {
	ref     Point
	scores  [4] int
	corners []Point
}

type neighbour struct {
	index    int
	distance float64
}
