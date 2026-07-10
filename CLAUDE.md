# CLAUDE.md — pilanda_work-management (Windshift — Aufgaben/Taskmanagement)

> **Aufgaben/Taskmanagement der Pilanda-Unit** auf Basis von Windshift (Go + Svelte 5,
> self-hosted, AGPL-3.0; Upstream `codeberg.org/realigned/windshift-core`). **Owner: Dominik.**
> Eigenständiger Stack neben der Bench — **kein Frappe**, kein Bench-Mount.
> In der Shell: „ToDo" unter „Allgemein" + Dashboard-Karte (kein Nav-Modul, Entscheid 07.07.2026).

## Maßgebliche Doku (SSOT — verweisen, nicht duplizieren)

- **App-Fachplan (Wahrheit/Status):** [`Entwicklungsplan.md`](Entwicklungsplan.md)
- **Master-Plan des Stacks:** `../pilanda/ENTWICKLUNGSPLAN.md` (§1 Inventar, §5 Eingliederung)
- Die Root-`README.md` und die übrigen Root-Dokumente (BUILD/CONTRIBUTING/…) sind
  **Upstream-Text** von Windshift — nicht unsere Projektdoku.

## Grundregeln

- **Owner Dominik:** Windshift-Quellcode nur lesen; Änderungen/Upstream-Sync nur mit
  Dominik abgestimmt (Fork-Divergenz verteuert Updates). Unsere Dateien hier: nur
  `Entwicklungsplan.md` + `CLAUDE.md`.
- **AGPL-3.0:** bei Modifikation + Netzwerk-Bereitstellung an Externe gilt
  Quellcode-Offenlegungspflicht — vor Kunden-/Portal-Nutzung prüfen.
- **Wahrheit ist Pflicht:** Status = verifizierter Stand, siehe Fachplan.

## Git / Workflow (stack-weit)

- **Pushes NUR auf `develop`.** `main` zieht Dominik. Remote: `Pilanda-ERP/pilanda_work-management`.
- **Nur geänderte/neue Dateien stagen** — nie `git add -A`.
