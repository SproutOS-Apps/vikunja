# Deploying vikunja on SproutOS

This repository is an exact deployment mirror of its upstream. Keep this file as the only mirror-only source change.

## Build target

- Repository root: `.`
- Dockerfile: `Dockerfile`
- Production branch: `main`

## Agent checklist

1. Read the upstream README and compose files before changing deployment settings.
2. Identify the Postgres, Valkey, search, and object-storage services the application actually uses; provision only those services.
3. Put every returned connection URI in SproutOS project environment variables. Never commit credentials.
4. Build the listed Dockerfile locally before deploying. If it expects release artifacts, add a build step to the deployment workflow rather than changing this mirror.
5. Add the SproutOS GitHub Actions workflow with `contents: read` and `id-token: write`, then wait for its terminal deployment result.