<!--
PR title must follow Conventional Commits (checked by "Lint PR Title"):
  <type>(<optional scope>): <lowercase subject without trailing period>
  types:  feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert
  scopes: frontend, backend, api, db, auth, ui, deps, docker
Target branch: develop (main is pulled by the owner).
-->

### Problem
<!-- What is broken or missing, and for whom? Link the issue (e.g. Windshiftapp/core#123 or WI-123). -->

### Change
<!-- What this PR does, per area. Delete lines that do not apply. -->
- **Backend:**
- **Database:**
- **API v2:**
- **Frontend:**

Not included / follow-ups:

### Tests
<!-- Commands you ran and their result; say what you did NOT run. -->
- [ ] `go test -count=1 ./...` (affected packages) / `go vet` / `gofmt -l` clean
- [ ] Frontend: `npm run check`, `npx svelte-check --threshold error`, `npm run build`
- [ ] Manual check (local instance, desktop + mobile if UI changed)
- Not run:

### Checklist
- [ ] Schema change under `internal/database/schema/*.sql` is paired with a new migration in `migrations.go` (SQLite **and** PostgreSQL) — `scripts/check-schema-migration-pairing.sh --staged`
- [ ] v2 routes changed: `internal/restapi/v2/contract-metadata.json` and `api/openapi-v2.json` regenerated (`go run ./scripts/openapi-v2-metadata`, `go run ./scripts/openapi-v2-generate`)
- [ ] New UI strings added to `en` and the fully covered locales `de`, `fr`, `ru`, `zh-CN`
- [ ] New `*_test.go` files have a `.gitignore` exception (test sources are ignored by default)
- [ ] `scripts/check-layering.sh` and `scripts/check-handler-db-access.sh` pass
- [ ] No secrets, credentials or customer data in code, tests or screenshots
- [ ] CLA signed (comment `I have read the CLA Document and I hereby sign the CLA` once)
- [ ] Upstream relevance: could this go to Windshiftapp/core as well? (fork divergence makes updates expensive)
