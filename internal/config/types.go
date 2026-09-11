package config

type MailerProvider string
type DBProvider string
type FSProvider string
type LLMProvider string

const (
	MailGmail  MailerProvider = "gmail"
	MailResend MailerProvider = "resend"
)

const (
	DBMongo DBProvider = "mongo"
)

const (
	FSMinio    FSProvider = "minio"
	FSOptional FSProvider = "optional"
)

const (
	LLMOpenAI   LLMProvider = "openai"
	LLMOptional LLMProvider = "optional"
)
