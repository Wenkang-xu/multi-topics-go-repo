package avro_test_b

type B struct {
	Schema_id int64
}

func NewAvro_example_topic_value() B {

	r := B{}
	r.Schema_id = 54

	return r
}
