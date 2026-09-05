# zynko-oracle · `go-two-fer-sharewith-conformance-oracle`

**A deterministic, re-checkable conformance oracle for `two-fer` (go).**

## Proven
Measured on the canonical Exercism corpus — **3 input/output pairs, 3 distinct outputs** — produced by *running* the reference in a sealed sandbox, not asserted.

## Scope (declared)
The corpus is the canonical Exercism test data for `two-fer`. Inputs outside that set are **not covered**; this oracle decides agreement on the published corpus only and makes no claim of general correctness.

## Provenance
Reference: the Exercism reference solution for `two-fer` (go; MIT, Exercism), body unchanged. Proven by the exercism testsuite (pin=19cf0be7577f6464), re-executed by harvest in a sealed sandbox (unshare -rn) before this bundle was generated.

## License
Apache-2.0 for the scaffolding; the reference body retains its upstream MIT (Exercism) license.
