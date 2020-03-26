package pmsX003

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)


func TestPMS7003(t *testing.T) {
	rd := strings.NewReader(string([]byte{
		//test garbage,
		0x01, 0x02,
		// hdr
		0x42,0x4d,
		// size
		0x00,0x1c,
		// TSI
		0x00,0x11,
		0x00,0x16,
		0x00,0x19,
		// Atm
		0x00,0x11,
		0x00,0x16,
		0x00,0x19,
		// counts
		0x0c,0x8d,
		0x03,0x6b,
		0x00,0x7b,
		0x00,0x08,
		0x00,0x04,
		0x00,0x04,
		// reserved
		0x97,0x00,
		// checksum
		0x03,0x54,
	}))
	frame, err := DecodeFrame(rd)
	require.NoError(t,err)
	assert.Equal(t,17,int(frame.PM1_0_tsi),"PM1.0_TSI")
	assert.Equal(t,22,int(frame.PM2_5_tsi),"PM2.5_TSI")
	assert.Equal(t,25,int(frame.PM10_tsi),"PM10_TSI")
	assert.Equal(t,17,int(frame.PM1_0_atm),"PM1.0 atmospheric")
	assert.Equal(t,22,int(frame.PM2_5_atm),"PM2.5 atmospheric")
	assert.Equal(t,25,int(frame.PM10_atm),"PM10 atmospheric")
	assert.Equal(t,3213,int(frame.PCount_0_3),"particle count 0.3")
	assert.Equal(t,875,int(frame.PCount_0_5),"particle count 0.5")
	assert.Equal(t,123,int(frame.PCount_1_0),"particle count 1.0")
	assert.Equal(t,8,int(frame.PCount_2_5),"particle count 2.5")
	assert.Equal(t,4,int(frame.PCount_5_0),"particle count 5.0")
	assert.Equal(t,4,int(frame.PCount_10),"particle count 10")
	assert.Equal(t,int(frame.ChecksumRead),int(frame.ChecksumCalculated),"checksum")
	assert.Equal(t,852,int(frame.ChecksumRead),"checksum")
}
