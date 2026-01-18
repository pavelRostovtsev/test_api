package test_any_client

type AnyStruct struct {
	AnyField1 string
	AnyField2 string
	AnyField3 string
}

func (c *Client) GetPreparedAnyData() AnyStruct {
	return AnyStruct{
		AnyField1: "value1",
		AnyField2: "value2",
		AnyField3: c.input.AnyData,
	}
}
