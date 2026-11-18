# IRS
**Purpose** | **[Configuration](CONFIGURATION.md)** | **[Running](RUNNING.md)**

---

## Purpose

> **This repository is archived.** The IRS FIRE system retires on 19 November 2026 at 3 p.m. Eastern. IRIS does not accept Publication 1220 files. This project remains a Publication 1220 reader and writer for historical files.

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
