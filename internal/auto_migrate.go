//go:build migrate

package internal

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nmfzone/privy-cake-store/config"
	"github.com/nmfzone/privy-cake-store/internal/logger"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func init() {
	config.InitConfig()
	logger.InitLogger()

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", "mysql://"+config.Dbdsn())
		if err == nil {
			break
		}

		log.Warn().Msgf("Migrate: trying to connect to db, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatal().Msgf("Migrate: can't connect to db: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Msgf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info().Msgf("Migrate: no change")
		return
	}

	log.Info().Msgf("Migrate: up success")
}
