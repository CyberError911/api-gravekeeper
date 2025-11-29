## Conference Submission Abstract

**Title:** The Forgotten Backend: Using Log Analysis to Audit the Undocumented 50% of Your APIs and Assign Accountability

**Word Count:** ~300 words 

---

Your API inventory is incomplete. Traditional security tools—SAST scanners, DAST platforms, OpenAPI specs—all fail at the same fundamental problem: **they can't spot the discrepancy between what your code defines and what's actually running.**

This creates two security disasters:

1. **Zombie APIs**: Endpoints in production code, never accessed, never patched. Attack surface bloat.
2. **Shadow APIs**: Undocumented services users depend on. Invisible to your security model.

Neither SAST nor DAST catches this gap. They assume defined code equals live code. It doesn't.

**API-Gravekeeper** solves this with a unified Code-vs-Log analysis engine: compare source definitions against real request logs, surface misalignments in seconds. But here's the innovation that changes the game:

**Git Blame as the final pipeline step.** Every zombie route retrieved with author, email, commit date, and line number. This single feature transforms a security audit into an Engineering Accountability Report—cutting remediation time from days of manual investigation to minutes of instant task assignment.

The tool introduces two technical innovations:

**Route Normalization** tackles dynamic paths intelligently. When users access /users/42 but code defines /users/:id, our heuristic-based normalizer learns to convert numeric IDs and UUIDs into placeholders, achieving true endpoint matching across millions of requests—fast, lightweight, no ML required.

**Non-Intrusive Architecture** requires only source code and existing Nginx/Apache logs. No agents, API gateways, or OpenAPI specs needed. A single Go binary, zero external dependencies beyond git. Fits naturally into DevOps pipelines.

Built in Go and open-sourced under MIT, API-Gravekeeper operates locally on your infrastructure. Framework support starts with Flask, with planned extensions to FastAPI, Django, Express.js, and Go frameworks.

Attendees will leave with a production-ready tool, a methodology for true DevSecOps API inventory auditing, and one core insight: **the developer's name is the most powerful accountability mechanism** for turning security findings into engineering action.

---

**Speaker:** CyberError911 (GitHub)  
**Repository:** https://github.com/CyberError911/api-gravekeeper
