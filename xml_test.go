package gojson

import (
	"testing"
)

func TestXmlDecodeNoError(t *testing.T) {
	data := XmlDecodeNoError([]byte(`<xml>    <AppId><![CDATA[wxe2c]]></AppId>    <Encrypt><![CDATA[yqBRLWH/QXHA==]]></Encrypt></xml>`))
	t.Log(data)
	t.Logf("%T", data)
}

func TestXmlEncodeNoError(t *testing.T) {
	data := XmlEncodeNoError(XmlDecodeNoError([]byte(`<xml>    <AppId><![CDATA[wxe2c]]></AppId>    <Encrypt><![CDATA[yqBRLWH/QXHA==]]></Encrypt></xml>`)))
	t.Log(data)
	t.Logf("%T", data)
}
