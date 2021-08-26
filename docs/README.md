# IRS
**Purpose** | **[Configuration](CONFIGURATION.md)** | **[Running](RUNNING.md)** | **[Client](../pkg/client/README.md)**

---

## Purpose

Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and
CLI for creating, parsing, validating, and transforming IRS electronic Filing Information Returns
Electronically (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and
CSV/TEXT file formats.

| Input      | Output     |
|------------|------------|
| JSON       | JSON       |
| ASCII FIRE | ASCII FIRE |
|            | PDF Form   |
|            | SQL        |

## Supporting Documentation

- [Become an Authorized e-file Provider](https://www.irs.gov/e-file-providers/become-an-authorized-e-file-provider)
- [Modernized e-File (MeF) User Guides and Publications](https://www.irs.gov/e-file-providers/modernized-e-file-mef-user-guides-and-publications)
- [Publication 1220](https://www.irs.gov/pub/irs-pdf/p1220.pdf)

---
**[Next - Configuration](CONFIGURATION.md)**
