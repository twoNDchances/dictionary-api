# Dictionary API - 📖 An API Backend simulator

Dictionary API is an API Backend used for learning about Kubernetes in the system deployment section with StatefulSet.

---

## Installation

```bash
git clone https://github.com/twoNDchances/dictionary-api.git
```

---

## Usage

- Docker (build)

```bash
cd dictionary-api
docker compose up -d
```

- Docker (pull & MySQL server already setup)

```bash
docker run -d --name=dictionary-api -p 8080:8080 secondchances/dictionary-api
```

- Manualy (MySQL server already setup)

```bash
cd dictionary-api
$(cat export.txt)
go run .
```
