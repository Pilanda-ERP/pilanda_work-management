# Entwicklungsplan — pilanda_work-management (Aufgaben / Taskmanagement der Unit — Windshift)
> Master: pilanda/ENTWICKLUNGSPLAN.md · Theme-Mitbau: pilanda_theme/CONTRIBUTING.md
> Stand: 10.07.2026 · Regel: NUR echte Zustände abhaken — Wahrheit ist Pflicht.

Rolle: **Taskmanagement/„Aufgaben" der Pilanda-Unit** auf Basis von **Windshift**
(Open-Source Work-Management: Kanban-Boards, Workflows, Custom Fields, Zeiterfassung,
Test-/Asset-Management; Go + Svelte 5, self-hosted, SQLite/PostgreSQL). **Owner: Dominik.**
Eigenständiger Stack **neben** der Bench (wie die salesbot-Scout-Runtime) — kein Frappe.

## Bindende Entscheide
- **Kein Nav-Modul** (Marco 07.07.2026): Aufruf als **„ToDo" unter „Allgemein" + Dashboard-Karte**
  in der Pilanda-Shell (aktuell `localhost:8088`; Einbindung pilanda `f2ffc8e`/`97607f5`).
- **Upstream = `codeberg.org/realigned/windshift-core`** (AGPL-3.0). Die Root-`README.md` ist
  Upstream-Text (nennt GitHub als Push-Mirror des Upstreams — das betrifft das Upstream-Repo,
  nicht dieses). Upstream-Sync-Strategie = Dominik.
- **Lizenz AGPL-3.0 beachten:** bei eigenen Modifikationen + Netzwerk-Bereitstellung besteht
  Quellcode-Offenlegungspflicht gegenüber den Nutzern — für interne Nutzung unkritisch,
  bei Kunden-/Extern-Zugriff (Customer Portal!) vorher prüfen.
- Git stack-weit: nur `develop` pushen, `main` zieht Dominik (Master-SSOT).
- Änderungen am Windshift-Code nur mit Dominik abgestimmt — jede Fork-Divergenz
  verteuert künftige Upstream-Updates.
- **Optik = Upstream-Look** (Marco 10.07.2026): KEIN `--pp-*`-Angleich — Update-Fähigkeit
  geht vor Einheitsoptik; bewusste, dokumentierte Ausnahme vom Theme-Zielbild.
- **Scope-Start = NUR Technik** (Marco 10.07.2026): Taskmanagement zunächst für die
  Engineering-Bereiche (engineering, projectengineering; künftig tecdoc, e-engineering,
  researchdevelop). Alle anderen Bereiche später.
- **Kopplung = NUR Verlinkung** (Marco 10.07.2026): Absprungpunkte aus **PLS-Gantt** und
  **PM-Gantt** (Gantt-Balken → Arbeitspaket PLS → Task in der Technik) als Links —
  keine Datenkopplung/kein Sync.

## Erledigt
- [x] Windshift-Instanz läuft (Dominik) und ist seit 07.07.2026 in die Pilanda-Shell
  eingebunden: „ToDo" unter „Allgemein" + Dashboard-Karte (kein Nav-Modul)
- [x] §5-Eingliederung ins Repo-Inventar (10.07.2026): `develop` angelegt + Default,
  Repo-Beschreibung, `pilanda-dev.code-workspace` + cSpell, Master §1 (19 Repos),
  ARCHITEKTUR §2, Architektur-Diagramm, BENCH-Matrix (als Nicht-Bench-Stack), dieser Fachplan

## Offen — wird wirklich gebaut
- [ ] **Betriebsmodell klären** [Dominik]: wo läuft die produktive Instanz
  (aktuell Dominiks `:8088`), Deployment/Backup, SQLite vs. PostgreSQL
- [ ] **SSO** [Dominik/IT]: Windshift kann OIDC — Anbindung an LCS/Entra ID, sobald der
  stack-weite SSO-Auftrag (IT) steht; keine Doppel-Userverwaltung
- [ ] **Absprungpunkte bauen** (Entscheid 10.07. s. oben): Link-Punkte im PLS-Gantt und
  PM-Gantt (Balken → Arbeitspaket → Technik-Task). Voraussetzung: Ziel-URL-Schema mit
  Dominik klären (Board-/Task-Adressierung in Windshift), dann UI-Verdrahtung in
  `pilanda_pls`/`pilanda_pm`

## Grenzen / ehrliche Hinweise
- Dieses Repo ist der **Quellcode-Stand** von Windshift in unserer Org; das laufende
  System ist Dominiks Instanz — Betriebs-/Konfigurationsstand ist hier **nicht** abgebildet.
- Kein Frappe, kein Bench-Mount, keine DocTypes — die Theme-/Fixture-Regeln der
  Frappe-Apps gelten hier nicht automatisch; Design-SSOT-Angleich nur per Entscheid (s. Offen).
- Die Feature-Beschreibung in der Root-`README.md` ist Upstream-Werbung, nicht unser
  verifizierter Nutzungs-Scope; was wir wirklich nutzen, definiert die ERP-Verzahnung (s. Offen).
