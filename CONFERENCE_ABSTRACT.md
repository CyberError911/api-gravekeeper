## Conference Submission Abstract

**Title:** The Forgotten Backend: Using Log Analysis to Audit the Undocumented Half of Your APIs

**Word Count:** ~280 words

---

Your API inventory is incomplete. Traditional security tools—SAST scanners, DAST platforms, runtime monitors—excel at finding bugs in code they can see. But they're blind to a critical gap: **the discrepancy between what your code defines and what's actually running**.

This blind spot creates two security disasters:

1. **Zombie APIs**: Endpoints sitting in production code, never accessed, never patched. Dead wood accumulating in your attack surface.
2. **Shadow APIs**: Undocumented services that users depend on but security teams don't monitor. Invisible to your DevSecOps pipeline.

Neither SAST nor DAST catches these. They assume defined code is live, and live traffic aligns with source. It doesn't.

**API-Gravekeeper** solves this by implementing a novel Code-vs-Log analysis engine. Instead of scanning code in isolation, it compares source definitions against actual request logs, surfacing both misalignments in seconds.

The tool introduces two innovations that make this practical:

**Route Normalization** tackles the dynamic path problem. When a user accesses `/users/42` and your code defines `/users/:id`, traditional diff tools see a mismatch. Our normalizer learns to convert numeric IDs and UUIDs into standard placeholders, achieving true endpoint matching across thousands of requests.

**Git Blame Integration** bridges security and engineering accountability. Every zombie route shows who committed the dead code and when—no more arguing about who's responsible for cleanup.

Built in Go with zero external dependencies beyond git, API-Gravekeeper runs locally on your infrastructure. It's framework-agnostic, supporting Python Flask today with planned support for FastAPI, Django, and Node.js frameworks.

Attendees will leave with a production-ready open-source tool, a methodology for true DevSecOps API inventory auditing, and a new perspective on API security that extends beyond code analysis.

---

**Speaker:** CyberError911 (GitHub)  
**Repository:** https://github.com/CyberError911/api-gravekeeper
