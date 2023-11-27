package avro_test_c

type C struct {
	Schema_id int64
}

func NewAvro_example_topic_value() C {

	r := C{}
	r.Schema_id = 54

	return r
}
