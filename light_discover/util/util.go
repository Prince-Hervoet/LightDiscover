package util

import (
	"encoding/binary"
	"errors"
)

const (
	MAGIC_NUMBER          = uint16(27126)
	CONTENT_TYPE_HEART    = 1 << 0
	CONTENT_TYPE_REGISTER = 1 << 1
	CONTENT_TYPE_DISCOVER = 1 << 2
	CONTENT_TYPE_UPDATE   = 1 << 3
	CONTENT_HEADER_LENGTH = 4
	SOCKET_KEY_NAME       = "serverSocket"
	PONG                  = "pong"
)

type EfProtocol struct {
	MagicNum     uint16
	ContentType  uint8
	HeaderLength uint16
	DataLength   uint32
	Data         []byte
}

func HandleMagicNumber(magicNumberBytes []byte, protocol *EfProtocol) error {
	if magicNumberBytes == nil || len(magicNumberBytes) < 2 {
		return errors.New("magicNumber error")
	}
	magicNum := binary.BigEndian.Uint16(magicNumberBytes) // 将byte切片转换为BigEndian格式的uint32整数
	magicNumx := ^magicNum
	if (magicNumx & MAGIC_NUMBER) == 0 {
		protocol.MagicNum = magicNum
		return nil
	}
	return errors.New("magic error")
}

func HandleContentType(contentType []byte, protocol *EfProtocol) error {
	if contentType == nil || len(contentType) < 1 {
		return errors.New("contentType error")
	}
	contentTypeNum := uint8(binary.BigEndian.Uint16(contentType))
	protocol.ContentType = contentTypeNum
	return nil
}

func HandleDataHeader(lenInfo []byte, protocol *EfProtocol) (uint16, uint32, error) {
	if lenInfo == nil || len(lenInfo) != 6 {
		return 0, 0, errors.New("header error")
	}
	headerLen := binary.BigEndian.Uint16(lenInfo[0:2])
	dataLen := binary.BigEndian.Uint32(lenInfo[2:6])
	protocol.HeaderLength = headerLen
	protocol.DataLength = dataLen
	return headerLen, dataLen, nil
}
