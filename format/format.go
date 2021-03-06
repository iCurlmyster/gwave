package format

import (
	"encoding/binary"
	"io"
)

// currently not supporting Non-PCM and Extensible fmt types

// Header is a simple layout of the format for a Wave file
type Header struct {
	// "RIFF" or "RIFX"
	ByteType [4]byte
	// File Size
	Size int32
	// "WAVE"
	HeaderType [4]byte
	// "fmt "
	FmtMarker [4]byte
	// Format size
	FmtSize int32
	// 1 is PCM type
	FmtType int16
	// number of channels
	NumChannels int16
	// Number of samples per second (Hz) 44100 is CD quality
	SampleRate int32
	// Sample Rate * BitsPerSample * NumChannels / 8. Constant Bit Rate (CBR)
	BytesPerSecond int32
	// NumChannels * BitsPerSample / 8. Number of bytes for one sample including all channels
	BlockAlign int16
	// 16 would be two 8 bytes to sample from. usually 8, 16, 24, or 32 (currently not handling 24)
	BitsPerSample int16
	// "data"
	DataMarker [4]byte
	// Size of data
	DataSize int32
}

const (
	// PcmType is to represent the PCM type
	PcmType = 1
	// CdSampleRate is the sample rate for a wave file that a CD would have
	CdSampleRate = 44100
	// DatSampleRate sample rate for DAT
	DatSampleRate = 48000
)

var (
	// Wav represents WAVE type
	Wav = [4]byte{'W', 'A', 'V', 'E'}
	// FmtMarker represents the fmt marker in the header
	FmtMarker = [4]byte{'f', 'm', 't', ' '}
	// Riff represents the Byte order type for little endian for the wave file
	Riff = [4]byte{'R', 'I', 'F', 'F'}
	// Rifx represents the Byte order type for big endian for the wave file
	Rifx = [4]byte{'R', 'I', 'F', 'X'}
	// DataMarker represents the DATA marker in the header.
	DataMarker = [4]byte{'d', 'a', 't', 'a'}
	// RiffByteOrder value to compare when converting bytes
	RiffByteOrder = binary.LittleEndian
	// RifxByteOrder value to compare when converting bytes
	RifxByteOrder = binary.BigEndian
)

// WaveWriter defines how to interact with a Wave File object for writing
type WaveWriter interface {
	io.Writer
	io.WriterAt
	FileHeader() Header
	AllocateDataSize(size int32)
}

// FileByteOrder returns the byte order the file is
func (h Header) FileByteOrder() binary.ByteOrder {
	if h.ByteType[3] == 'F' {
		return RiffByteOrder
	}
	return RifxByteOrder
}

// GetByteCount returns the count of bytes in a sample
func (h Header) GetByteCount() int {
	switch h.BitsPerSample {
	case 8:
		{
			return 1
		}
	case 16:
		{
			return 2
		}
	default:
		{
			return 4
		}
	}
}
