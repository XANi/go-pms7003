package pmsX003

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PMS1003, PMS5003, PMS7003
// MSB, LSB
type Pms7003Frame  struct {
	// * 0x42,0x4d - start
	// * frame length - normally 32 bytes
	Length uint16 // 0x001C = 28
	// µg/m3
	PM1_0_tsi uint16
	PM2_5_tsi uint16
	PM10_tsi  uint16
	PM1_0_atm uint16
	PM2_5_atm uint16
	PM10_atm  uint16
	// count/100cm3
	PCount_0_3 uint16
	PCount_0_5 uint16
	PCount_1_0 uint16
	PCount_2_5 uint16
	PCount_5_0 uint16
	PCount_10  uint16
	rsrvd      uint16
	ChecksumRead   uint16 //  cksum=byte01+..+byte30
	ChecksumCalculated uint16
}

// find first frame and decode it. Gives up after seeking 64 bytes
func DecodeFrame(r io.Reader) (Pms7003Frame,error) {
	bytes := 0
	b := make([]byte,1)
	w := make([]byte,2)
	var frame Pms7003Frame
 	for {
		if bytes >= 64 {
			return Pms7003Frame{}, fmt.Errorf("could not find start packet after %d",bytes)
		}
		ctr, err := r.Read(b)
		if err != nil || ctr != 1 {
			return Pms7003Frame{}, fmt.Errorf("reader error:[1/%d] %s", ctr, err)
		}
		bytes += ctr
		if b[0] == 0x42 {
			ctr, err := r.Read(b)

			if err != nil || ctr != 1 {
				return Pms7003Frame{}, fmt.Errorf("reader error:[1/%d] %s", ctr, err)
			}
			if b[0] == 0x4d {
				break
			}
			bytes += ctr
		}
	}
	var csum uint16
	ctr, err := r.Read(w)
	if err != nil || ctr != 2 {
		return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)
	}
	if w[0] != 0x00 || w[1] != 0x1c {
		return Pms7003Frame{}, fmt.Errorf("expected 0x001c:[%+v] %s", w, err)
	}
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PM1_0_tsi = binary.BigEndian.Uint16(w) ; csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PM2_5_tsi = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PM10_tsi = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)

	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PM1_0_atm = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PM2_5_atm = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PM10_atm = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)

	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PCount_0_3 = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PCount_0_5 = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PCount_1_0 = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PCount_2_5 = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PCount_5_0 = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.PCount_10 = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.rsrvd = binary.BigEndian.Uint16(w); csum += binary.BigEndian.Uint16(w)
	ctr, err = r.Read(w); if ctr != 2 { return Pms7003Frame{}, fmt.Errorf("reader error:[%+v] %s", w, err)}
	frame.ChecksumRead = binary.BigEndian.Uint16(w)
	frame.ChecksumCalculated = csum

	return frame,nil
}
