package chat

type Config struct {
	BaseUrl         string `yaml:"baseUrl" json:"baseUrl" bson:"baseUrl" validate:"required"`
	AppKey          string `yaml:"appKey" json:"appKey" bson:"appKey" validate:"required"`
	Port            int    `yaml:"port" json:"port" bson:"port" validate:"required"`
	IntervalSeconds int    `yaml:"intervalSeconds" json:"intervalSeconds" bson:"intervalSeconds" validate:"required"`
	MaxLength       int    `yaml:"maxLength" json:"maxLength" bson:"maxLength" validate:"required"`
	Cors            bool   `yaml:"cors" json:"cors" bson:"cors" validate:""`
}
