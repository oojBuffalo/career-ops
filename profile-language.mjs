// profile-language.mjs — output-language resolution for headless engines.
//
// This fork is English-only: `language.output` in config/profile.yml is
// accepted for upstream-profile compatibility but ALWAYS resolves to English.
// The parse/instruction API shape is kept identical to upstream so engine call
// sites (`outputLanguageInstruction(parseOutputLanguage(profile))`) merge
// cleanly, while the resolution itself is pinned.

const OUTPUT_LANGUAGE = 'en';

/** Resolve the output language from profile YAML. Always English in this fork. */
export function parseOutputLanguage() {
  return OUTPUT_LANGUAGE;
}

/** Build the canonical output-language rule injected into every model prompt. */
export function outputLanguageInstruction() {
  return [
    `Write all human-facing output in English, including the full A–G`,
    `evaluation and the machine-readable summary's free-text fields, regardless`,
    `of the language of these instructions or the job description. Keep`,
    `market-specific terms when relevant, but explain them in English`,
    `when needed. This system is English-only: the job description's language`,
    `never changes the output language.`,
  ].join(' ');
}
