package httpserver

import (
	"context"
	"net/http"
	"time"
)

// Значения по умолчнию для сервера
const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultAddr            = ":8000"
	defaultShutdownTimeout = 3 * time.Second
)

// Структура сервера, где также реализуется gracfull shutdown
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// Конструктор принимающий http.Handler и options
// где options означает что пользователь задаёт свои настройки для сервера
// Например: меняет read timeout на 10 секунд когда как по умолчанию настроено 5 секунд
// Config pattern
func New(handler http.Handler, opts ...Option) *Server {
	
	// Содаём http.Server где указываем значения по умолчанию
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		Addr:         defaultAddr,
	}

	// Создаём свою структуру где имплементируется gracful shutdown
	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}

	// Тут применяютися пользовательские настройки на сервер
	for _, opt := range opts {
		opt(s)
	}

	// Запускаем наш сервер, операция не блокнируящая
	s.start()

	return s
}


func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}