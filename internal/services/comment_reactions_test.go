package services

import (
	"errors"
	"path/filepath"
	"testing"

	"windshift/internal/database"
	"windshift/internal/models"
	"windshift/internal/repository"
)

func newCommentReactionTestService(t *testing.T) (*CommentService, database.Database) {
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
	return NewCommentService(db), db
}

func insertReactionFixture(t *testing.T, db database.Database, query string, args ...any) int {
	t.Helper()
	var id int
	if err := db.QueryRow(query+" RETURNING id", args...).Scan(&id); err != nil {
		t.Fatal(err)
	}
	return id
}

// reactionFixture creates a workspace item with one comment by alice and
// returns the comment ID with the user IDs of alice and bob.
func reactionFixture(t *testing.T, db database.Database) (commentID, alice, bob int) {
	t.Helper()
	alice = insertReactionFixture(t, db, `INSERT INTO users (email, username, first_name, last_name) VALUES ('alice@example.com', 'alice', 'Alice', 'Doe')`)
	bob = insertReactionFixture(t, db, `INSERT INTO users (email, username, first_name, last_name) VALUES ('bob@example.com', 'bob', '', '')`)
	workspace := insertReactionFixture(t, db, `INSERT INTO workspaces (name, key) VALUES ('Reactions', 'REACT')`)
	item := insertReactionFixture(t, db, `INSERT INTO items (workspace_id, workspace_item_number, title, frac_index) VALUES (?, 1, 'Item', 'a0')`, workspace)
	commentID = insertReactionFixture(t, db, `INSERT INTO comments (item_id, author_id, content) VALUES (?, ?, 'Please review')`, item, alice)
	return commentID, alice, bob
}

func TestCommentReactionsGroupInKeyOrderAndAreIdempotent(t *testing.T) {
	service, db := newCommentReactionTestService(t)
	commentID, alice, bob := reactionFixture(t, db)

	for _, step := range []struct {
		user     int
		reaction string
	}{{bob, "eyes"}, {alice, "thumbs_up"}, {bob, "thumbs_up"}, {bob, "thumbs_up"}} {
		if _, err := service.AddReaction(commentID, step.user, step.reaction); err != nil {
			t.Fatalf("AddReaction(%d, %q) error = %v", step.user, step.reaction, err)
		}
	}

	got, err := service.ReactionsForComment(commentID)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Reaction != "thumbs_up" || got[1].Reaction != "eyes" {
		t.Fatalf("reactions = %+v, want thumbs_up then eyes", got)
	}
	if got[0].Count != 2 || got[0].UserIDs[0] != alice || got[0].UserNames[0] != "Alice Doe" || got[0].UserNames[1] != "bob" {
		t.Fatalf("thumbs_up = %+v, want Alice Doe then bob once each", got[0])
	}

	if err := service.RemoveReaction(commentID, bob, "thumbs_up"); err != nil {
		t.Fatal(err)
	}
	if err := service.RemoveReaction(commentID, bob, "thumbs_up"); err != nil {
		t.Fatalf("removing a missing reaction should be a no-op, got %v", err)
	}
	comments := []models.Comment{{ID: commentID, Source: "human"}, {ID: -1, Source: "approval"}}
	if err := service.AttachReactions(comments); err != nil {
		t.Fatal(err)
	}
	if r := comments[0].Reactions; len(r) != 2 || r[0].Count != 1 || r[0].UserIDs[0] != alice {
		t.Fatalf("attached reactions = %+v, want alice's thumbs_up and bob's eyes", r)
	}
	if comments[1].Reactions != nil {
		t.Fatalf("approval row got reactions: %+v", comments[1].Reactions)
	}
}

func TestCommentReactionsRejectUnknownKeysAndComments(t *testing.T) {
	service, db := newCommentReactionTestService(t)
	commentID, alice, _ := reactionFixture(t, db)

	if _, err := service.AddReaction(commentID, alice, "rocket"); !errors.Is(err, ErrInvalidCommentReaction) {
		t.Fatalf("AddReaction(rocket) error = %v, want ErrInvalidCommentReaction", err)
	}
	if _, err := service.AddReaction(commentID+100, alice, "check"); !errors.Is(err, repository.ErrNotFound) {
		t.Fatalf("AddReaction(missing comment) error = %v, want ErrNotFound", err)
	}
}

func TestCommentReactionsCascadeWithTheComment(t *testing.T) {
	service, db := newCommentReactionTestService(t)
	commentID, alice, _ := reactionFixture(t, db)
	if _, err := service.AddReaction(commentID, alice, "check"); err != nil {
		t.Fatal(err)
	}
	if err := service.Delete(commentID); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := db.QueryRow(`SELECT COUNT(*) FROM comment_reactions`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("comment_reactions rows after comment delete = %d, want 0", n)
	}
}
