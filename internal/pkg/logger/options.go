package logger

type options struct {
	level      string
	filename   string
	maxSize    int // MB
	maxBackups int
	maxAge     int
	localtime  bool
	compress   bool
}

type Option struct {
	apply func(option *options)
}

func defaultOptions() options {
	return options{
		level:      "info",
		filename:   "./log.log",
		maxSize:    100,
		maxBackups: 5,
		maxAge:     30,
		localtime:  false,
		compress:   true,
	}
}

func WithLevel(level string) Option {
	return Option{
		apply: func(option *options) {
			option.level = level
		},
	}
}

func WithFilename(filename string) Option {
	return Option{
		apply: func(option *options) {
			option.filename = filename
		},
	}
}

func WithMaxSize(size int) Option {
	return Option{
		apply: func(option *options) {
			option.maxSize = size
		},
	}
}

func WithMaxBackups(num int) Option {
	return Option{
		apply: func(option *options) {
			option.maxBackups = num
		},
	}
}

func WithAge(age int) Option {
	return Option{
		apply: func(option *options) {
			option.maxAge = age
		},
	}
}

func WithLocaltime(flag bool) Option {
	return Option{
		apply: func(option *options) {
			option.localtime = flag
		},
	}
}

func WithCompress(flag bool) Option {
	return Option{
		apply: func(option *options) {
			option.compress = flag
		},
	}
}
