package shm

import "github.com/linuxdeepin/go-x11-client"

type BadSegError struct {
	x.ResourceErrorBase
}

func (e BadSegError) Error() string {
	return "shm.BadSegError" + e.Msg()
}

func readBadSegError(r *x.Reader) x.Error {
	return BadSegError{x.ReadResourceErrorBase(r)}
}

type CompletionEvent struct {
}

func readCompletionEvent(r *x.Reader, v *CompletionEvent) error {
	return nil
}

// #WREQ
func encodeQueryVersion() (b x.RequestBody) {
	return
}

type QueryVersionReply struct {
	SharedPixmaps bool
	MajorVersion  uint16
	MinorVersion  uint16
	Uid           uint16
	Gid           uint16
	PixmapFormat  uint8
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SharedPixmaps = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Uid = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Gid = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.PixmapFormat = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(15)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeAttach(shmSeg Seg, shmId uint32, readOnly bool) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(shmSeg)).
		Write4b(shmId).
		Write1b(x.BoolToUint8(readOnly)).
		WritePad(3).
		End()
	return
}

// #WREQ
func encodeDetach(shmSeg Seg) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(shmSeg)).
		End()
	return
}

// #WREQ
func encodePutImage(drawable x.Drawable, gc x.GContext, totalWidth,
	totalHeight, srcX, srcY, srcWidth, srcHeight uint16, dstX, dstY int16,
	depth, format uint8, sendEvent bool, shmSeg Seg, offset uint32) (b x.RequestBody) {

	b.AddBlock(9).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc)).
		Write2b(totalWidth).
		Write2b(totalHeight).
		Write2b(srcX).
		Write2b(srcY).
		Write2b(srcWidth).
		Write2b(srcHeight).
		Write2b(uint16(dstX)).
		Write2b(uint16(dstY)).
		Write1b(depth).
		Write1b(format).
		Write1b(x.BoolToUint8(sendEvent)).
		WritePad(1).
		Write4b(uint32(shmSeg)).
		Write4b(offset).
		End()
	return
}

// #WREQ
func encodeGetImage(drawable x.Drawable, X, y int16, width, height uint16,
	planeMask uint32, format uint8, shmSeg Seg, offset uint32) (b x.RequestBody) {
	b.AddBlock(7).
		Write4b(uint32(drawable)).
		Write2b(uint16(X)).
		Write2b(uint16(y)).
		Write2b(width).
		Write2b(height).
		Write4b(planeMask).
		Write1b(format).
		WritePad(3).
		Write4b(uint32(shmSeg)).
		Write4b(offset).
		End()
	return
}

type GetImageReply struct {
	Depth  uint8
	Visual x.VisualID
	Size   uint32
}

func readGetImageReply(r *x.Reader, v *GetImageReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Visual = x.VisualID(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Size = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeCreatePixmap(pid x.Pixmap, drawable x.Drawable,
	width, height uint16, depth uint8, shmSeg Seg, offset uint32) (b x.RequestBody) {
	b.AddBlock(6).
		Write4b(uint32(pid)).
		Write4b(uint32(drawable)).
		Write2b(width).
		Write2b(height).
		Write1b(depth).
		WritePad(3).
		Write4b(uint32(shmSeg)).
		Write4b(offset).
		End()
	return
}