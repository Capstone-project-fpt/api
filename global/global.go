package global

import (
	"github.com/api/pkg/logger"
	"github.com/api/pkg/setting"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hibiken/asynq"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	Config       setting.Config
	I18nBundle   *i18n.Bundle
	Localizer    *i18n.Localizer
	Logger       *logger.LoggerZap
	Db           *gorm.DB
	RDb          *redis.Client
	AsyncqClient *asynq.Client
	AwsSession   *session.Session
	S3Client     *s3.S3
	Validator    *validator.Validate
)
