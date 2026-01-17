# GetStream Chat + PocketBase Demo Backend

Minimal backend demo for a freelance marketplace, built for a technical assessment.

The project demonstrates:

* authentication and data storage with PocketBase,
* real-time chat integration using GetStream Chat,
* clean backend architecture with Go,
* soft delete and access control via PocketBase rules.

---

## Architecture overview

* **Go API** — main application, request handling, domain logic
* **PocketBase (embedded)** — authentication, projects, proposals, milestones
* **GetStream Chat** — real-time messaging (channels and messages)

PocketBase is used as an embedded database and auth provider.
Chat data (messages, channels) is **not stored** in PocketBase and is handled exclusively by GetStream.

---

## Features implemented

* Email/password authentication (PocketBase)
* Projects (job posts)
* Proposals linked to projects (sent / accepted / rejected)
* Milestones per project
* Real-time chat per project via GetStream
* Soft delete for all domain entities
* Access rules enforced at the PocketBase level

---

## Database schema (PocketBase)

The database schema is defined using PocketBase collections.

A JSON export of the collections schema is provided here:

```
pb/schema/collections.json
```

### Importing the schema

To import the schema into a fresh PocketBase instance:

1. Start PocketBase
2. Open the Admin UI
3. Go to **Settings → Import / Export**
4. Import `pb/schema/collections.json`

---

## Configuration

Create a local `.env` file from the example:

```bash
cp .env.example .env
```

Required variables:

* application port
* PocketBase admin credentials
* GetStream Chat credentials

---

## Running locally (without Docker)

```bash
make run
```

The API will be available at:

```
http://localhost:<PORT>
```

---

## Running with Docker

```bash
make docker-run
```

This will:

* build the application image,
* start the API service,
* persist PocketBase data using a Docker volume.

---

## Notes on scope

This is a **minimal demo** created for a technical assessment.

Out of scope:

* payments and escrow logic,
* notifications,
* search and filtering,
* full video/audio call implementation (GetStream Video API not wired).

---

## Security considerations

* No hard deletes — soft delete is enforced via schema and rules
* Access rules restrict data visibility to project participants
* Secrets are provided via environment variables only
* PocketBase admin credentials are not committed

---

## Tech stack

* Go
* PocketBase (embedded)
* GetStream Chat
* Docker / Docker Compose

## License

MIT
