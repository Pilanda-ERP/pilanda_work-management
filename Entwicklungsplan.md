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
- **Dev-Login = der ZENTRALE Testnutzer `t.tester@lcs.local` via SSO** (Marco 11.07.2026,
  Klarstellung): KEIN eigener Windshift-User im Env — der eine Stack-Testnutzer (SSOT
  `pilanda/_devenv/.env` + Dex + Frappe-Seed, BENCH §1) meldet sich per „Login with LCS SSO"
  an; Windshift legt die lokale Zeile beim ersten Login automatisch an (Auto-Provisioning,
  belegt 10.07.). Der Env-Admin `admin@lcs.local` ist nur Bootstrap/Verwaltung, kein Tages-Login.
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
  `http://localhost:8088`. **Ersteinrichtung headless erledigt (10.07.2026)** —
  Erst-Admin `admin@lcs.local` per `POST /api/setup/complete` (kein Browser-Wizard
  mehr); reproduzierbar via `_devenv/setup/windshift-sso-setup.ps1` (idempotent,
  Zugangsdaten aus Env, Dev-Defaults in `.env.example`). Nachweis: `GET
  /api/setup/status` → `{"setup_completed":true,"admin_user_created":true,...}`,
  Admin-Login `POST /api/auth/login` → `/api/auth/me` liefert `is_system_admin:true`.
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
- [x] **SSO Dev-Teil KOMPLETT — „ein Login für alles" E2E verifiziert (10.07.2026):**
  - [x] Dex-Client `windshift` (`_devenv/dex/config.yaml`, Secret Dev-Wert,
    RedirectURI `http://localhost:8088/api/sso/callback/windshift`) — Discovery 200.
  - [x] Windshift-OIDC-Provider `LCS SSO` in der windshift-DB (slug `windshift`,
    Issuer `http://host.docker.internal:5556/dex`, `auto_provision_users`,
    Client-Secret app-verschlüsselt). Angelegt per Admin-API (HTTPS-Dummy-Issuer,
    da `ValidateExternalURL` HTTPS erzwingt) + SQL-Downgrade des Issuers auf die
    Dex-HTTP-URL; reproduzierbar im Setup-Helfer. Nachweis: `GET /api/sso/status`
    → `{"enabled":true,"provider_name":"LCS SSO",...}` (Login-Seite zeigt den
    SSO-Knopf).
  - [x] **E2E-Login verifiziert (10.07.2026, nach expliziter Marco-Freigabe des
    Dev-Schalters):** `ALLOW_LOCAL_CONNECTIONS=true` am `windshift`-Service aktiv
    (NUR Dev — Windshifts SSRF-Schutz blockte sonst den Dex-Call auf die private
    Host-Gateway-IP; in Prod mit Entra/HTTPS bleibt der Schalter AUS). Kompletter
    OIDC-Code-Flow durchgespielt: `/api/sso/login/windshift` → Dex-Login
    `t.tester@lcs.local` → 303 mit Code → Callback 302 + Session-Cookie →
    `GET /api/auth/me` = `{"email":"t.tester@lcs.local", …}` — Nutzer wurde per
    **Auto-Provisioning** beim ersten SSO-Login angelegt (id 2, email_verified).
    Hinweis Cookie-Verhalten: Windshift setzt state/pkce/session-Cookies mit
    `Secure` — Browser akzeptieren das auf `localhost` (trustworthy origin),
    Skript-Clients müssen die Cookies explizit mitführen.
  - [ ] Prod: LCS/Entra mit dem stack-weiten SSO-Auftrag (IT); keine Doppel-User-
    verwaltung. Umstieg = Issuer/Client-IDs tauschen (gleicher Code-Flow).
- [x] **Absprungpunkte gebaut** (Entscheid 10.07.): stabiler Deep-Link `<base>/item/<ITEM_KEY>`
  (base aus `site_config` `windshift_url`, Default `http://localhost:8088`; Link-Logik SSOT in
  `pilanda_engineering.api`). Alle drei Flächen erledigt: 5 Team-Cockpits, PM-Inspektor, PLS-AP-Ebene.
  - [x] **Windshift-Spalte der 5 Team-Cockpits verdrahtet (10.07.2026):** echter Deep-Link bei
    gepflegtem `Task.custom_eng_windshift_task`, sonst „nicht verknüpft"; alle 5 Bundles gebaut,
    HTTP 200, Render-Zweig nachgewiesen (Feld + Lese-API + Frontends).
  - [x] **PM-Gantt-Absprung gebaut (10.07.2026):** dezenter Deep-Link im Vorgangs-Inspektor
    (`pilanda_pm`), nur bei gepflegtem `Task.custom_eng_windshift_task`; das Gantt-Backend hängt
    `windshift_url` je Vorgang an (Link-Logik SSOT `pilanda_engineering.api.get_windshift_links`).
    Nachweis: Bundle HTTP 200 + Link-Zweig, `/app/pm-gantt` 200, E2E `TASK-2026-00525` →
    `http://localhost:8088/item/ETECH-1` in Projekt `PROJ-0011`, PM-Tests grün (Details pilanda_pm-Plan).
  - [x] **PLS-Absprung gebaut (10.07.2026, Welle 2):** die AP-/Task-Ebene ist jetzt als **Projekt-Cockpit-Drill**
    da — Klick auf eine Portfolio-Ampel-Zeile (`/app/pls-portfolio`) → Desk-Page `pls-projekt-cockpit`
    (eigenes Bundle, SSOT-Baustein **PpTechCockpit@2** `mode="project"`) mit allen Technik-APs des Projekts
    über alle Teams. **Die Windshift-Spalte dort IST der PLS-AP-Absprung.** Der frühere Realitäts-Einwand
    (Cockpit team-, nicht projektgefiltert → irreführend) ist mit dem PROJEKT-Modus der Klammer-API aufgelöst
    (jetzt projektgefiltert, Deep-Link am Task = 1:1-Ziel, Link-Logik weiter SSOT `pilanda_engineering.api`).
    Nachweis: Build grün, `/app/pls-projekt-cockpit` 200, E2E PROJ-0011 → 15 APs, ETECH-1 →
    `http://localhost:8088/item/ETECH-1`. Details im pilanda_pls-Plan.
  - **Hinweis Erst-Setup:** Die Ziel-Tasks in Windshift entstehen erst, wenn Marco den
    Windshift-Setup-Wizard durchläuft und die ersten Tasks anlegt — bis dahin rendern die Links
    korrekt, laufen aber ggf. in 404.
- [ ] **Prod-Betrieb später**: wo die eine geteilte Instanz fürs Haus läuft
  (Server/Backup) — mit IT, wenn es soweit ist

## Grenzen / ehrliche Hinweise
- Dieses Repo ist der **Quellcode-Stand** von Windshift in unserer Org; das laufende
  System ist Dominiks Instanz — Betriebs-/Konfigurationsstand ist hier **nicht** abgebildet.
- Kein Frappe, kein Bench-Mount, keine DocTypes — die Theme-/Fixture-Regeln der
  Frappe-Apps gelten hier nicht automatisch; Design-SSOT-Angleich nur per Entscheid (s. Offen).
- Die Feature-Beschreibung in der Root-`README.md` ist Upstream-Werbung, nicht unser
  verifizierter Nutzungs-Scope; was wir wirklich nutzen, definiert die ERP-Verzahnung (s. Offen).
