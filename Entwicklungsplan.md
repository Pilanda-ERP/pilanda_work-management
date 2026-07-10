# Entwicklungsplan — pilanda_work-management (Aufgaben / Taskmanagement der Unit — Windshift)
> Master: pilanda/ENTWICKLUNGSPLAN.md · Theme-Mitbau: pilanda_theme/CONTRIBUTING.md
> Stand: 10.07.2026 · Regel: NUR echte Zustände abhaken — Wahrheit ist Pflicht.

Rolle: **Taskmanagement/„Aufgaben" der Pilanda-Unit** auf Basis von **Windshift**
(Open-Source Work-Management: Kanban-Boards, Workflows, Custom Fields, Zeiterfassung,
Test-/Asset-Management; Go + Svelte 5, self-hosted, SQLite/PostgreSQL). **Owner: Dominik.**
Eigenständiger Stack **neben** der Bench (wie die salesbot-Scout-Runtime) — kein Frappe.

## Bindende Entscheide
- **Kein Nav-Modul** (Marco 07.07.2026): Aufruf als **„ToDo" unter „Allgemein" + Dashboard-Karte**
  in der Pilanda-Shell.
- **Betriebsmodell (Marco 10.07.2026, präzisiert):** WIR betreiben eine **eigene
  Windshift-Instanz, gebaut aus DIESEM Repo (`develop`), in Docker** — eigene Datenbank
  im Docker-Volume, per API an Frappe.
- **DB = PostgreSQL von Anfang an (Marco 10.07.2026):** Dev=Prod, **kein späterer
  DB-Umzug**. MariaDB wird von Windshift nicht unterstützt; die Mitnutzung der
  Frappe-DB wurde bewusst verworfen (Datenhoheit, Kopplung nur per API). Eigener
  Compose-Service `windshift-db` (postgres:17, Volume `windshift-db-data`), DSN
  `postgresql://windshift:windshift@windshift-db:5432/windshift?sslmode=disable`.
  **Dominiks Server-Installation ist sein Test-/Feature-Labor mit Usern und für unseren
  Betrieb NICHT relevant** — wir verlinken nie dorthin. Code-Fluss: Dominiks Welt →
  dieses Repo (`develop`) → unsere Docker-Instanz. **Update-Takt bewusst:** git pull →
  Docker-Rebuild → Kurztest (kein Auto-Folgen).
- **Task-URL-Schema leiten WIR selbst ab** (aus Repo-Code/laufender Instanz) — keine
  Abhängigkeit von Dominiks Installation.
- **Upstream-Kette:** `codeberg.org/realigned/windshift-core` (AGPL-3.0) → Dominiks Welt →
  dieses Repo. Die Root-`README.md` ist Upstream-Text. Upstream-Sync = Dominik.
- **Lizenz AGPL-3.0 beachten:** bei eigenen Modifikationen + Netzwerk-Bereitstellung besteht
  Quellcode-Offenlegungspflicht gegenüber den Nutzern — für interne Nutzung unkritisch,
  bei Kunden-/Extern-Zugriff (Customer Portal!) vorher prüfen.
- Git stack-weit: nur `develop` pushen, `main` zieht Dominik (Master-SSOT).
- Änderungen am Windshift-Code nur mit Dominik abgestimmt — jede Fork-Divergenz
  verteuert künftige Upstream-Updates.
- **Optik = Upstream-Look** (Marco 10.07.2026): KEIN `--pp-*`-Angleich — Update-Fähigkeit
  geht vor Einheitsoptik; bewusste, dokumentierte Ausnahme vom Theme-Zielbild.
- **Scope-Start = NUR Technik** (Marco 10.07.2026): Taskmanagement zunächst für die
  Engineering-Bereiche (`pilanda_engineering`, `pilanda_projectengineering`,
  `pilanda_technicaldocu`, `pilanda_electricalengineering`, `pilanda_researchdevelop` —
  die drei Team-Repos seit 10.07. eingegliedert). Alle anderen Bereiche später.
- **Kopplung = NUR Verlinkung** (Marco 10.07.2026): Absprungpunkte aus **PLS-Gantt** und
  **PM-Gantt** (Gantt-Balken → Arbeitspaket PLS → Task in der Technik) als Links —
  keine Datenkopplung/kein Sync.

## Erledigt
- [x] Windshift-Instanz läuft (Dominik) und ist seit 07.07.2026 in die Pilanda-Shell
  eingebunden: „ToDo" unter „Allgemein" + Dashboard-Karte (kein Nav-Modul)
- [x] §5-Eingliederung ins Repo-Inventar (10.07.2026): `develop` angelegt + Default,
  Repo-Beschreibung, `pilanda-dev.code-workspace` + cSpell, Master §1 (19 Repos),
  ARCHITEKTUR §2, Architektur-Diagramm, BENCH-Matrix (als Nicht-Bench-Stack), dieser Fachplan
- [x] **Eigene Docker-Instanz in der Dev-Env (10.07.2026)** — Compose-Services
  `windshift` (build aus diesem Repo/Dockerfile, `pilanda-windshift:local`) +
  `windshift-db` (postgres:17) in `pilanda/_devenv/docker-compose.dev.yml`,
  Compose-Projekt `pilanda`, Port **8088**, Volumes `pilanda_windshift-data`
  (Attachments/Plugins/Prompts) + `pilanda_windshift-db-data` (PostgreSQL).
  **Nachweise:** HTTP `GET http://localhost:8088` → **200** (`<title>Windshift -
  Work Management</title>`); `GET /api/setup/status` → 200
  `{"setup_completed":false,...}`; PostgreSQL belegt (235 Tabellen in DB
  `windshift`, Log `postgres timestamp backfill complete`); CORS/CSRF-Origin
  `http://localhost:8088`. **Ersteinrichtung = manueller Setup-Wizard durch Marco**
  im Browser (Erst-Admin + Module) — kein Env/CLI-Bootstrap; danach Auth aktiv.
- [x] **Task-URL-Schema abgeleitet (aus `frontend/src/lib/router.js`, Code-Fund;
  praktisch klickbar erst nach Setup/Login):**
  - Stabiler, menschenlesbarer Deep-Link (**für PLS-/PM-Gantt-Absprünge empfohlen**):
    `http://localhost:8088/workspace/<WORKSPACE_KEY>/item/<ITEM_NUMBER>`
    bzw. Kurzform `http://localhost:8088/item/<WORKSPACE_KEY>-<ITEM_NUMBER>`
    (Item-Key-Format `PROJ-123` = `workspace_key`-`workspace_item_number`,
    `utils/itemKey.js`).
  - Numerische Form (so erzeugt die eingebaute „Copy Link"-Aktion, `ItemDetail.svelte`):
    `http://localhost:8088/workspaces/<workspaceId>/items/<itemId>`.
  - Windshift nennt Tasks „Work Items" (Route-View `item-detail`).

## Offen — wird wirklich gebaut
- [ ] **SSO**: Windshift kann OIDC — Dev: an unseren Dex hängen (wie Frappe);
  Prod: LCS/Entra mit dem stack-weiten SSO-Auftrag (IT); keine Doppel-Userverwaltung
- [ ] **Absprungpunkte bauen** (Entscheid 10.07.): Windshift-Spalte der 5 Team-Cockpits
  + Link-Punkte im PLS-/PM-Gantt auf unsere Instanz verdrahten, sobald das URL-Schema
  abgeleitet ist (kein Dominik-Blocker mehr)
- [ ] **Prod-Betrieb später**: wo die eine geteilte Instanz fürs Haus läuft
  (Server/Backup) — mit IT, wenn es soweit ist

## Grenzen / ehrliche Hinweise
- Dieses Repo ist der **Quellcode-Stand** von Windshift in unserer Org; das laufende
  System ist Dominiks Instanz — Betriebs-/Konfigurationsstand ist hier **nicht** abgebildet.
- Kein Frappe, kein Bench-Mount, keine DocTypes — die Theme-/Fixture-Regeln der
  Frappe-Apps gelten hier nicht automatisch; Design-SSOT-Angleich nur per Entscheid (s. Offen).
- Die Feature-Beschreibung in der Root-`README.md` ist Upstream-Werbung, nicht unser
  verifizierter Nutzungs-Scope; was wir wirklich nutzen, definiert die ERP-Verzahnung (s. Offen).
