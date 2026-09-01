followmyjourney-profile-service

Minimal Go microservice providing Profile REST API used by followmyjourney-front-end.

Structure:
- cmd/server: application entrypoint
- internal/api: HTTP handlers and routing
- internal/service: business logic
- internal/store: storage interfaces/implementations
- internal/model: domain models
- api: OpenAPI spec
- deploy: Cloud Run manifest

Use `make build` / `make run` for local work; `make docker-build` and `make deploy-cloudrun` for CI/CD.
