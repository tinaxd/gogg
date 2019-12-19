package gogg

// #cgo pkg-config: vorbis
// #include <vorbis/codec.h>
// #include <ogg/ogg.h>
// #include <stdlib.h>
import "C"
import "unsafe"

const (
	OvFalse = -1
	OvEOF   = -2
	OvHole  = -3

	OvEread      = -128
	OvEfault     = -129
	OvEimpl      = -130
	OvEinval     = -131
	OvEnotvorbis = -132
	OvEbadheader = -133
	OvEversion   = -134
	OvEnotaudio  = -135
	OvEbadpacket = -136
	OvEbadlink   = -137
	OvEnoseek    = -138
)

type VorbisInfo struct {
	VorbisInfo C.vorbis_info
}

type VorbisDspState struct {
	VorbisDspState C.vorbis_dsp_state
}

type VorbisBlock struct {
	VorbisBlock C.vorbis_block
}

type VorbisComment struct {
	VorbisComment C.vorbis_comment
}

func VorbisInfoInit(vi *VorbisInfo) {
	C.vorbis_info_init(&vi.VorbisInfo)
}

func VorbisInfoClear(vi *VorbisInfo) {
	C.vorbis_info_clear(&vi.VorbisInfo)
}

func VorbisInfoBlockSize(vi *VorbisInfo, zo int) int {
	return int(C.vorbis_info_blocksize(&vi.VorbisInfo, C.int(zo)))
}

func VorbisCommentInit(vc *VorbisComment) {
	C.vorbis_comment_init(&vc.VorbisComment)
}

func VorbisCommentAdd(vc *VorbisComment, comment string) {
	cstr := C.CString(comment)
	defer C.free(unsafe.Pointer(cstr))
	C.vorbis_comment_add(&vc.VorbisComment, cstr)
}

func VorbisCommentAddTag(vc *VorbisComment, tag string, contents string) {
	ctag := C.CString(tag)
	ccontents := C.CString(contents)
	defer func() {
		C.free(unsafe.Pointer(ctag))
		C.free(unsafe.Pointer(ccontents))
	}()
	C.vorbis_comment_add_tag(&vc.VorbisComment, ctag, ccontents)
}

func VorbisCommentQuery(vc *VorbisComment, tag string, count int) string {
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))
	cret := C.vorbis_comment_query(&vc.VorbisComment, ctag, C.int(count))
	return C.GoString(cret)
}

func VorbisCommentQueryCount(vc *VorbisComment, tag string) int {
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))
	return int(C.vorbis_comment_query_count(&vc.VorbisComment, ctag))
}

func VorbisCommentClear(vc *VorbisComment) {
	C.vorbis_comment_clear(&vc.VorbisComment)
}

func VorbisBlockInit(v *VorbisDspState, vb *VorbisBlock) int {
	return int(C.vorbis_block_init(&v.VorbisDspState, &vb.VorbisBlock))
}

func VorbisBlockClear(vb *VorbisBlock) int {
	return int(C.vorbis_block_clear(&vb.VorbisBlock))
}

func VorbisDspClear(v *VorbisDspState) {
	C.vorbis_dsp_clear(&v.VorbisDspState)
}

func VorbisGranuleTime(v *VorbisDspState, granulepos int64) float64 {
	return float64(C.vorbis_granule_time(&v.VorbisDspState, C.long(granulepos)))
}

func VorbisVersionString() string {
	return C.GoString(C.vorbis_version_string())
}

func VorbisAnalysisInit(v *VorbisDspState, vi *VorbisInfo) int {
	return int(C.vorbis_analysis_init(&v.VorbisDspState, &vi.VorbisInfo))
}

// TODO: Not implemented!
func VorbisCommentheaderOut(vc *VorbisComment) {
}
