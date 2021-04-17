// Package settings - реализует модели данных для хранения настроек сервера и их частичного автозаполнения
package settings

// GetTRulesForGuest - Возвращает заполненный список ролей по всем таблицам будущей базы данных для гостя
func GetTRulesForGuest() SQLTRules {
	return SQLTRules{
		TRule{
			TName:      "public.\"Files\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"users\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     true,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"hashes\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     true,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"confirmations\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"password_resets\"",
			SELECT:     true,
			INSERT:     false,
			UPDATE:     false,
			DELETE:     false,
			REFERENCES: false,
		},
		TRule{
			TName:      "secret.\"totp\"",
			SELECT:     true,
			INSERT:     true,
			UPDATE:     true,
			DELETE:     true,
			REFERENCES: true,
		},
	}
}
