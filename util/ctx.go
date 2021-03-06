package echoapp_util

import (
	echoapp "github.com/gw123/echo-app"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ctxLogger struct {
	logger *logrus.Entry
	fields logrus.Fields
}

type Record struct {
	TraceId   int64  `json:"trace_id"`
	CreatedAt int64  `json:"uint_64"`
	Point     string `json:"point"`
	Place     string `json:"place"`
}

const ctxRequestIdKey = echo.HeaderXRequestID

var (
	defaultLogger *logrus.Logger
	isDebug       = false
)

const (
	//ctxkeys
	ctxUserKey      = "&userKey{}"
	ctxUserIdKey    = "&userIdKey{}"
	ctxComKey       = "&comKey{}"
	ctxUserRolesKey = "&userRolesKey{}"
	ctxLoggerKey    = "&loggerKey{}"
)

func SetCtxUserId(ctx echo.Context, userId int64) {
	AddField(ctx, "user_id", strconv.FormatInt(userId, 10))
	ctx.Set(ctxUserIdKey, userId)
}

func GetCtxtUserId(ctx echo.Context) (int64, error) {
	userId, ok := ctx.Get(ctxUserIdKey).(int64)
	if !ok {
		return 0, errors.New("get ctxUserId flied")
	}
	return userId, nil
}

func SetCtxUser(ctx echo.Context, user *echoapp.User) {
	ctx.Set(ctxUserKey, user)
}

func GetCtxtUser(ctx echo.Context) (*echoapp.User, error) {
	user, ok := ctx.Get(ctxUserKey).(*echoapp.User)
	if !ok {
		return nil, errors.New("get ctxUser flied")
	}
	return user, nil
}

func SetCtxCompany(ctx echo.Context, company *echoapp.Company) {
	ctx.Set(ctxComKey, company)
}

func GetCtxtCompany(ctx echo.Context) (*echoapp.Company, error) {
	company, ok := ctx.Get(ctxComKey).(*echoapp.Company)
	if !ok {
		return nil, errors.New("get ctxCompany flied")
	}
	return company, nil
}

func SetCtxUserRoles(ctx echo.Context, company []echoapp.Role) {
	ctx.Set(ctxComKey, company)
}

func GetCtxtUserRoles(ctx echo.Context) ([]echoapp.Role, error) {
	roles, ok := ctx.Get(ctxComKey).([]echoapp.Role)
	if !ok {
		return nil, errors.New("get userRoles flied")
	}
	return roles, nil
}

func SetCtxJwsPayload(ctx echo.Context, payload string) {
	ctx.Set(ctxComKey, payload)
}

func GetCtxtJwsPayload(ctx echo.Context) (string, error) {
	payload, ok := ctx.Get(ctxComKey).(string)
	if !ok {
		return "", errors.New("get ctxPayload flied")
	}
	return payload, nil
}

func SetDebug(flag bool) {
	isDebug = flag
}

// 为了方便创建一个默认的Logger
func DefaultLogger() *logrus.Logger {
	if defaultLogger == nil {
		defaultLogger = logrus.New()
	}
	defaultLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		FieldMap:         nil,
		CallerPrettyfier: nil,
	})
	return defaultLogger
}

// 为了方便创建一个默认的Logger
func DefaultJsonLogger() *logrus.Logger {
	if defaultLogger == nil {
		defaultLogger = logrus.New()
	}
	defaultLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      isDebug,
	})
	return defaultLogger
}

func NewDefaultEntry() *logrus.Entry {
	return logrus.NewEntry(DefaultLogger())
}

// 添加logrus.Entry到context, 这个操作添加的logrus.Entry在后面AddFields和Extract都会使用到
func ToContext(ctx echo.Context, entry *logrus.Entry) {
	l := &ctxLogger{
		logger: entry,
		fields: logrus.Fields{},
	}
	ctx.Set(ctxLoggerKey, l)
}

//添加日志字段到日志中间件(ctx_logrus)，添加的字段会在后面调用 info，debug，error 时候输出
func AddFields(ctx echo.Context, fields logrus.Fields) {
	l, ok := ctx.Get(ctxLoggerKey).(*ctxLogger)
	if !ok || l == nil {
		return
	}
	for k, v := range fields {
		l.fields[k] = v
	}
}

//添加日志字段到日志中间件(ctx_logrus)，添加的字段会在后面调用 info，debug，error 时候输出
func AddField(ctx echo.Context, key, val string) {
	l, ok := ctx.Get(ctxLoggerKey).(*ctxLogger)
	if !ok || l == nil {
		return
	}
	l.fields[key] = val
}

// 添加一个追踪规矩id 用来聚合同一次请求, 注意要用返回的contxt 替换传入的ctx
func AddRequestId(ctx echo.Context, requestId string) {
	ctx.Set(ctxRequestIdKey, requestId)
}

//导出requestId
func ExtractRequestId(ctx echo.Context) string {
	l, ok := ctx.Get(ctxRequestIdKey).(string)
	if !ok {
		return ""
	}
	return l
}

//导出ctx_logrus日志库
func ExtractEntry(ctx echo.Context) *logrus.Entry {
	l, ok := ctx.Get(ctxLoggerKey).(*ctxLogger)
	if !ok || l == nil {
		return logrus.NewEntry(logrus.New())
	}

	fields := logrus.Fields{}
	for k, v := range l.fields {
		fields[k] = v
	}

	requestId := ExtractRequestId(ctx)
	if requestId != "" {
		fields[ctxRequestIdKey] = requestId
	}
	return l.logger.WithFields(fields)
}
