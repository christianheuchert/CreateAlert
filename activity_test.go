package ParseXpertMessage

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	xpertTestMsg := "{\"DataBusReports\":[{\"Attributes\":[],\"CustomDataBytes\":\"RUxQAAACAAEANwPnwTgDIACQeqir4QAyAFgAAYAD//4KcZA0f4j/5QAGAAAD58E4rEylXysM/9kABgAAA+fBOK5MpWcrCf/RACwAAAPnwTiuTKVnKw//0AAsAAAD58E4rkylZysK/9AALAAAA+fBOKxMpV8rDf/PACwAAAPnwTiuTKVnKwj/zwAsAAAD58E4uPhTsO9G/8kAAQAAA+fBOLj4U7qYhv/CAAEAAAPnwTi8m2hsYe3/wAAGAAAD58E4uPhT8w1i/70AAQAAA+fBOLj4U7DvR/+6AJUAAAPnwTjUuS/ttjb/ugALAAAD58E41Lkv7bY7/7oACwAAA+fBODy9xbq11P+5AAEAAAPnwTi4+FO6mIj/uAAoAAAD58E4vJtobGH0/7YALAAAA+fBOLybaGxh9v+2ACwAAAPnwThi+FO6mIT/tgBkAAAD58E4vJtobGH3/7YALAAAA+fBOLybaGxh9f+2ACwAAAPnwTi8m2hsYfL/tQAsAAAD58E4uPhTupiH/7UAZAAAA+fBOLj4U/MNZP+0ADgAAAPnwThq+FPzDWb/swA4AAAD58E4cr3FBher/7MACwAAA+fBODy9xQYXrP+xADwAAAPnwThqvcUGF67/sQA8AAAD58E41Lkv7bZB/7EAnQAAA+fBOGL4U/MNYP+xAKEAAAPnwTi4+FPzDWP/sQChAAAD58E4uPhTSbjj/7EABgAAA+fBODy9xbq11v+vADQAAAPnwTjUuS/ttj7/rwCdAAAD58E41Lkv7bZA/68AnQAAA+fBOGq9xbq11P+uADQAAAPnwTgQDGtvEMr/rgAEAAAD58E41Lkv7bZC/60AnQAAA+fBOGAysRnElv+rAAQAAAPnwThivcW6tdb/qgCEAAAD58E4zpQ1FRof/6oALAAAA+fBOASiIuLlXf+qAAsAAAPnwTg8vcW6tdX/qQCEAAAD58E41pQ1FRof/6kALAAAA+fBODy9xSLzKv+pAAYAAAPnwTjKlDUVGh//qQAsAAAD58E40pQ1FRof/6kALAAAA+fBOMCUNRUaHv+pAAYAAAPnwTis20h26F7/qAALAAAD58E4\",\"DataTimestamp\":\"1970-01-01T00:00:00\",\"DataUniqueID\":\"00:90:7A:A8:AB:E1\",\"Events\":[],\"LastEvent\":{},\"LastEventIndex\":-1,\"ReceivedTimestamp\":\"1970-01-01T00:00:00\",\"DateCreated\":\"1970-01-01T00:00:00\",\"DateUpdated\":\"1970-01-01T00:00:00\"}],\"DataCorrelations\":[],\"DeviceReports\":[{\"Attributes\":[],\"AUTOID_EPC\":\"MAC_00:90:7A:A8:AB:E1\",\"DataTimestamp\":\"1970-01-01T00:00:00\",\"DeviceActionBit\":true,\"DeviceLogID\":12157461,\"DeviceModel\":\"E50\",\"DeviceSerialNumber\":\"00:90:7A:A8:AB:E1\",\"DeviceType\":\"RFID_ACTIVE_802_11\",\"DeviceUniqueID\":\"00:90:7A:A8:AB:E1\",\"DeviceUniqueID_DisplayName\":\"MAC\",\"Events\":[],\"Item\":{\"ItemId\":400001774,\"DateCreated\":\"1970-01-01T00:00:00\",\"DateUpdated\":\"1970-01-01T00:00:00\"},\"LastEvent\":{\"StartDateTime\":\"2023-03-23T14:11:52.5839418-04:00\",\"SystemName\":\"RULEPROCESSOR\"},\"LastEventIndex\":-1,\"LOCMessageContent\":\"0A:71:90:34:7F:88:-27;AC:4C:A5:5F:2B:0C:-39;AE:4C:A5:67:2B:09:-47;AE:4C:A5:67:2B:0F:-48;AE:4C:A5:67:2B:0A:-48;AC:4C:A5:5F:2B:0D:-49;AE:4C:A5:67:2B:08:-49;B8:F8:53:B0:EF:46:-55;B8:F8:53:BA:98:86:-62;BC:9B:68:6C:61:ED:-64;B8:F8:53:F3:0D:62:-67;B8:F8:53:B0:EF:47:-70;D4:B9:2F:ED:B6:36:-70;D4:B9:2F:ED:B6:3B:-70;3C:BD:C5:BA:B5:D4:-71;B8:F8:53:BA:98:88:-72;BC:9B:68:6C:61:F4:-74;BC:9B:68:6C:61:F6:-74;62:F8:53:BA:98:84:-74;BC:9B:68:6C:61:F7:-74;BC:9B:68:6C:61:F5:-74;BC:9B:68:6C:61:F2:-75;B8:F8:53:BA:98:87:-75;B8:F8:53:F3:0D:64:-76;6A:F8:53:F3:0D:66:-77;72:BD:C5:06:17:AB:-77;3C:BD:C5:06:17:AC:-79;6A:BD:C5:06:17:AE:-79;D4:B9:2F:ED:B6:41:-79;62:F8:53:F3:0D:60:-79;B8:F8:53:F3:0D:63:-79;B8:F8:53:49:B8:E3:-79;3C:BD:C5:BA:B5:D6:-81;D4:B9:2F:ED:B6:3E:-81;D4:B9:2F:ED:B6:40:-81;6A:BD:C5:BA:B5:D4:-82;10:0C:6B:6F:10:CA:-82;D4:B9:2F:ED:B6:42:-83;60:32:B1:19:C4:96:-85;62:BD:C5:BA:B5:D6:-86;CE:94:35:15:1A:1F:-86;04:A2:22:E2:E5:5D:-86;3C:BD:C5:BA:B5:D5:-87;D6:94:35:15:1A:1F:-87;3C:BD:C5:22:F3:2A:-87;CA:94:35:15:1A:1F:-87;D2:94:35:15:1A:1F:-87;C0:94:35:15:1A:1E:-87;AC:DB:48:76:E8:5E:-88\",\"MessageId\":\"55\",\"MessageType\":\"RTLS_RFID_ACTIVE_802_11\",\"ReceivedTimestamp\":\"2023-03-23T18:11:52.6464061Z\",\"RFID\":{\"EPC\":\"AC:4C:A5:5F:2B:0C\",\"IsValid\":true,\"ReadCount\":2,\"Readers\":[{\"AntennaDisplayName\":\"AC:4C:A5:5F:2B:0C\",\"AntennaID\":6224,\"FirstDetectedTimestamp\":\"2023-03-23T18:11:52.6932953Z\",\"HopTableValue\":1,\"IsDefaultLocation\":true,\"LastDetectedTimestamp\":\"2023-03-23T18:11:52.6932953Z\",\"MaxTransmitPower\":80,\"MinReceiveSensitivity\":20,\"ReadCount\":1,\"ReaderDisplayName\":\"Auditorium\",\"ReaderID\":6224,\"RSSI\":39,\"RSSI_dbm\":-39,\"RTLSModelID\":4150,\"RTLSModelMapID\":4365,\"X\":1611.0,\"Y\":553.0}],\"Timestamp\":\"2023-03-23T18:11:52.6932953Z\"},\"RTLSAddress\":{},\"RTLSContact\":{},\"RTLSGeo\":{},\"RTLSGPS\":{},\"RTLSModel2D\":{\"IsValid\":true,\"PosDisplayName\":\"TestHospital\",\"PosMapID\":4150,\"PosModelID\":4150,\"PosX\":1611.0,\"PosY\":553.0,\"PosZoneIDs\":\"[{\\\"ZoneID\\\":5215,\\\"ZoneName\\\":\\\"Audiotorium\\\",\\\"ZoneType\\\":\\\"Open\\\"}]\",\"Timestamp\":\"2023-03-23T18:11:52.771404Z\"},\"Sensor\":{},\"SequenceNumber\":55,\"Status\":{\"BatteryLevel1\":88.0,\"ChargerConnected\":true,\"Data2\":\"TkE=\",\"DeviceReportReason\":3,\"IsValid\":true,\"Timestamp\":\"2023-03-23T18:11:52.6464061Z\"},\"DateCreated\":\"1970-01-01T00:00:00\",\"DateUpdated\":\"1970-01-01T00:00:00\"}],\"ItemInfo\":{\"ItemId\":400001774,\"DateCreated\":\"1970-01-01T00:00:00\",\"DateUpdated\":\"1970-01-01T00:00:00\"},\"ProximityReports\":[],\"ReceivedTimestamp\":\"2023-03-23T18:11:52.6464061Z\",\"SchemaName\":\"XpertSchema.XpertMessage.XpertMessage\",\"SchemaVersion\":\"1\"}"
	input := &Input{XpertMessage: xpertTestMsg}
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)
	assert.Equal(t, "00:90:7A:A8:AB:E1", output.DeviceMAC)
}
