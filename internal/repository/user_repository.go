package repository

import (
	"context"
	"database/sql"
	"github.com/DDnagasawa/FocusFlow/internal/model"
)

// UserRepository interface for user data operations
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int) error
}

// MySQLUserRepository implements UserRepository for MySQL
type MySQLUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository creates a new MySQLUserRepository
func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *MySQLUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO Users (username, email, password_hash, notification_enabled, theme, focus_mode_settings, adhd_diagnosis_date, adhd_severity, medications)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, user.Username, user.Email, user.PasswordHash, user.NotificationEnabled, user.Theme, user.FocusModeSettings, user.ADHDDiagnosisDate, user.ADHDSeverity, user.Medications)
	return err
}

// GetUser retrieves a user from the database by ID
func (r *MySQLUserRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT * FROM Users WHERE id = ?", id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt,
		&user.LastLogin, &user.NotificationEnabled, &user.Theme, &user.FocusModeSettings,
		&user.ADHDDiagnosisDate, &user.ADHDSeverity, &user.Medications,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user's information in the database
func (r *MySQLUserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, `
        UPDATE Users SET username=?, email=?, notification_enabled=?, theme=?,
        focus_mode_settings=?, adhd_diagnosis_date=?, adhd_severity=?, medications=?
        WHERE id=?
    `, user.Username, user.Email, user.NotificationEnabled, user.Theme,
		user.FocusModeSettings, user.ADHDDiagnosisDate, user.ADHDSeverity, user.Medications,
		user.ID)
	return err
}

// DeleteUser deletes a user from the database
func (r *MySQLUserRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM Users WHERE id = ?", id)
	return err
}
