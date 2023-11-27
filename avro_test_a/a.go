package avro_test_a

type A struct {
	Schema_id int64
}

func NewAvro_example_topic_value() A {

	r := A{}
	r.Schema_id = 54

	return r
}
