// Regression test for the /api/run prerequisite check (fork PR #3 review).
// The route gates each kind on a real core file (needsScript) and its prompts
// tell the agent which modes/*.md to follow. After the fork renamed
// modes/oferta.md → modes/offer.md, a stale reference here returned HTTP 400
// for every web evaluation. The route is TypeScript, so this test validates
// the source directly: every modes/*.md and needsScript file the route names
// must exist in the repository root.
//
// Run:  node --test tests/lib/run-route-prereqs.test.mjs

import { test } from "node:test";
import assert from "node:assert/strict";
import { readFileSync, existsSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";

const WEB_DIR = join(dirname(fileURLToPath(import.meta.url)), "..", "..");
const REPO_ROOT = join(WEB_DIR, "..");
const ROUTE = join(WEB_DIR, "src", "app", "api", "run", "route.ts");

const source = readFileSync(ROUTE, "utf8");

test("run route: every referenced modes/*.md exists in the repo root", () => {
  const refs = [...new Set(source.match(/modes\/[a-z0-9._-]+\.md/gi) ?? [])];
  assert.ok(refs.length > 0, "expected the route to reference at least one modes/*.md");
  for (const ref of refs) {
    // User-layer files (modes/_profile.md, modes/_custom.md) are created from
    // their .template.md counterparts during onboarding, so a fresh checkout
    // ships only the template — accept either.
    const template = ref.replace(/\.md$/, ".template.md");
    assert.ok(
      existsSync(join(REPO_ROOT, ref)) || existsSync(join(REPO_ROOT, template)),
      `route.ts references ${ref}, which does not exist — stale rename?`,
    );
  }
});

test("run route: every needsScript prerequisite exists in the repo root", () => {
  const block = source.match(/needsScript[^=]*=\s*\{([^}]*)\}/);
  assert.ok(block, "expected a needsScript map in route.ts");
  const files = [...block[1].matchAll(/:\s*"([^"]+)"/g)].map((m) => m[1]);
  assert.ok(files.length > 0, "expected needsScript to gate at least one kind");
  for (const f of files) {
    assert.ok(existsSync(join(REPO_ROOT, f)), `needsScript gates on ${f}, which does not exist — every run of that kind would 400`);
  }
});
