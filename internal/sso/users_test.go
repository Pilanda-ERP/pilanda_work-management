package sso

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"windshift/internal/database"
)

func newUserStoreTestDB(t *testing.T) (*UserStore, database.Database) {
	t.Helper()
	db, err := database.NewSQLiteDBWithPoolSizes(filepath.Join(t.TempDir(), "windshift.db"), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Initialize(); err != nil {
		_ = db.Close()
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return NewUserStore(db), db
}

func insertTestSSOProvider(t *testing.T, db database.Database) *SSOProvider {
	t.Helper()
	var id int
	err := db.QueryRow(`
		INSERT INTO sso_providers (slug, name, provider_type, enabled, auto_provision_users, require_verified_email)
		VALUES ('azure-ad', 'Azure AD', 'oidc', true, true, true)
		RETURNING id
	`).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	return &SSOProvider{ID: id, Slug: "azure-ad", AutoProvisionUsers: true, RequireVerifiedEmail: true}
}

// insertPasswordlessUser creates a user the way the admin UI does when no
// password is set: password_hash stays NULL.
func insertPasswordlessUser(t *testing.T, db database.Database, email, username string) int {
	t.Helper()
	var id int
	err := db.QueryRow(`
		INSERT INTO users (email, username, first_name, last_name)
		VALUES (?, ?, 'Test', 'User')
		RETURNING id
	`, email, username).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func countUsers(t *testing.T, db database.Database) int {
	t.Helper()
	var n int
	if err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestFindUserByEmail_NullPasswordHash(t *testing.T) {
	store, db := newUserStoreTestDB(t)
	id := insertPasswordlessUser(t, db, "Jane.Doe@example.com", "jane.doe")

	user, err := store.FindUserByEmail("jane.doe@example.com")
	if err != nil {
		t.Fatalf("FindUserByEmail() error = %v", err)
	}
	if user.ID != id {
		t.Errorf("FindUserByEmail() id = %d, want %d", user.ID, id)
	}
	if user.PasswordHash != "" {
		t.Errorf("FindUserByEmail() PasswordHash = %q, want empty", user.PasswordHash)
	}
}

func TestFindOrCreateUser_LinksExistingPasswordlessUser(t *testing.T) {
	store, db := newUserStoreTestDB(t)
	provider := insertTestSSOProvider(t, db)
	id := insertPasswordlessUser(t, db, "Jane.Doe@example.com", "jane.doe")
	before := countUsers(t, db)

	claims := &OIDCClaims{
		Subject:               "subject-1",
		Email:                 "jane.doe@example.com",
		EmailVerified:         true,
		EmailVerifiedProvided: true,
	}

	result, err := store.FindOrCreateUser(provider, claims)
	if err != nil {
		t.Fatalf("FindOrCreateUser() error = %v", err)
	}
	if result.User.ID != id {
		t.Errorf("FindOrCreateUser() user id = %d, want existing %d", result.User.ID, id)
	}
	if result.IsNewUser {
		t.Error("FindOrCreateUser() IsNewUser = true, want false")
	}
	if got := countUsers(t, db); got != before {
		t.Errorf("user count = %d, want %d", got, before)
	}

	account, err := store.FindExternalAccount(provider.ID, "subject-1")
	if err != nil {
		t.Fatalf("FindExternalAccount() error = %v", err)
	}
	if account.UserID != id {
		t.Errorf("linked user id = %d, want %d", account.UserID, id)
	}

	// Second login goes through the linked external account.
	result, err = store.FindOrCreateUser(provider, claims)
	if err != nil {
		t.Fatalf("FindOrCreateUser() second login error = %v", err)
	}
	if result.User.ID != id {
		t.Errorf("FindOrCreateUser() second login user id = %d, want %d", result.User.ID, id)
	}
}

func TestFindOrCreateUser_LookupErrorIsNotProvisioning(t *testing.T) {
	store, db := newUserStoreTestDB(t)
	provider := insertTestSSOProvider(t, db)
	insertPasswordlessUser(t, db, "jane.doe@example.com", "jane.doe")
	// A row that exists but cannot be scanned must not be treated as missing.
	if _, err := db.ExecWrite(`UPDATE users SET is_active = NULL WHERE email = ?`, "jane.doe@example.com"); err != nil {
		t.Fatal(err)
	}
	before := countUsers(t, db)

	_, err := store.FindOrCreateUser(provider, &OIDCClaims{
		Subject:               "subject-1",
		Email:                 "jane.doe@example.com",
		EmailVerified:         true,
		EmailVerifiedProvided: true,
	})
	if err == nil {
		t.Fatal("FindOrCreateUser() error = nil, want lookup error")
	}
	if errors.Is(err, ErrUserNotFound) {
		t.Errorf("FindOrCreateUser() error = %v, must not be ErrUserNotFound", err)
	}
	if !strings.Contains(err.Error(), "failed to look up user by email") {
		t.Errorf("FindOrCreateUser() error = %v, want lookup error instead of a create attempt", err)
	}
	if got := countUsers(t, db); got != before {
		t.Errorf("user count = %d, want %d", got, before)
	}
}
