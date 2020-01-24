package main

import (
  "bufio"
  "io"
  "errors"
)

const PACKET_SIGN = byte(0xAA)

func SendPacket(w io.Writer, data []byte) (err error) {
  l := len(data)
  if(!(0 < l && l < 0x10000)) {
    return errors.New("[Error] Expect len(data) in (0, 0x10000)")
  }
  lh, ll := byte((l >> 8) & 0xFF), byte((l) & 0xFF)
  writer := bufio.NewWriter(w)
  _, err = writer.Write([] byte {PACKET_SIGN, lh, ll, (PACKET_SIGN + lh + ll)})
  if err == nil {
    _, err = writer.Write(data)
  }
  if err == nil {
    err = writer.Flush()
  }
  return err
}

func ReadPacket(r io.Reader) (data []byte, err error) {
  reader := bufio.NewReader(r)
  b, err := reader.ReadByte()
  if err != nil {
    return []byte{}, err
  }
  if b != PACKET_SIGN {
    return []byte{}, errors.New("Sign mismatch")
  }
  checksum := b

  l := 0
  for i := 2; i > 0; i-- {
    b, err = reader.ReadByte()
    if err != nil {
      return []byte{}, err
    }
    l = (l << 8) + int(b)
    checksum += b
  }

  b, err = reader.ReadByte()
  if err != nil {
    return []byte{}, err
  }
  if checksum != b {
    return []byte{}, errors.New("Checksum mismatch")
  }

  buffer := make([]byte, l)
  for i, n := 0, 0; i < l; i += n {
    n, err = reader.Read(buffer[i:])
    if err != nil {
      return []byte{}, err
    }
  }
  return buffer, nil
}
