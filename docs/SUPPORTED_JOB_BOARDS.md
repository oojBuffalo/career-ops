# Supported Job Boards

Career-Ops scans job sources through provider modules in `providers/`. Each
non-helper `*.mjs` file maps to one supported source. Files prefixed with `_`
are shared helpers and are not loaded as providers.

| Board | Type (API / RSS / parser) | Notes |
| --- | --- | --- |
| 4 Day Week | API | Reads the public `https://4dayweek.io/api/jobs` JSON feed (4-day-week / reduced-hours roles). Configure with `provider: 4dayweek`; paginates `?page=N` up to `max_pages` (default 3), drops expired postings, then scanner filters apply. |
| Amazon / AWS | API | Auto-detects `amazon.jobs` careers URLs and queries the public amazon.jobs search API. The board is one global endpoint, so narrow it with an `amazon:` config block (`loc_query`, `base_query`, `category`, …) whose keys pass through as query params. Configure with `provider: amazon`. |
| Ashby | API | Auto-detects `https://jobs.ashbyhq.com/<slug>` boards and uses Ashby's public posting API. |
| Avature | Parser | Auto-detects `<tenant>.avature.net` career sites and parses the public server-rendered job list (`/careers/SearchJobs?jobOffset=N`, 6 results/page). A branded custom domain that proxies Avature needs `provider: avature` + `api:` pointing at the Avature origin. Paginates up to `max_pages` (default 50). |
| BambooHR | API | Auto-detects `<tenant>.bamboohr.com` careers pages, reads `/careers/list`, and follows public detail endpoints for job URLs. |
| Breezy HR | API | Auto-detects `<tenant>.breezy.hr` boards and reads the public JSON position feed. |
| Comeet / Spark Hire Recruit | API | Uses Comeet's public careers API. Provide the full API URL with `api:` or `careers_url`; it cannot derive the endpoint from a branded careers page. |
| Cornerstone OnDemand | API | Reads hosted CSOD career sites (`<tenant>.csod.com/ux/ats/careersite/...`). The search API wants a bearer token, but the career-site home page embeds an anonymous JWT (no login), so each fetch bootstraps that token then pages the public search API. Point `careers_url` at the csod.com URL (or set `provider: csod`). |
| CryptocurrencyJobs | RSS | Curated Web3/crypto job board. Reads the public `https://cryptocurrencyjobs.co/index.xml` RSS feed (listings are 100% remote); the feed URL is hardcoded in the provider, so no `careers_url` or `api` field is needed. Configure with `provider: cryptocurrencyjobs`. |
| EchoJobs | API | Reads the board-wide `https://echojobs.io/api/jobs` JSON feed (tech jobs aggregated from company ATS boards). Configure with `provider: echojobs`; paginates `?page=N` up to `max_pages` (default 3), then scanner filters apply. Job URLs point at the original ATS posting. |
| Flowxtra | API | Reads the public, no-auth, cross-tenant `https://app.flowxtra.com/api/central/jobs` JSON feed (board-wide — every company hosted on Flowxtra, not one tenant). Configure with `provider: flowxtra`; paginates `?page=N` up to `max_pages` (default 3), filtered to `status=Live`, then scanner filters apply. Job URLs use the API's ready-made `urlJobApplay` apply link on flowxtra.com. |
| Gem | API | Auto-detects `https://jobs.gem.com/<boardId>` boards and batches the listing plus per-job detail queries into a single POST to Gem's public GraphQL endpoint. |
| Greenhouse | API | Handles explicit `api:` URLs and auto-detects public Greenhouse board URLs for the boards API. |
| Hacker News (Who is hiring?) | API | Finds the current monthly "Ask HN: Who is hiring?" thread via the Algolia HN search API and parses top-level comments as postings (the pipe-delimited "Company \| Role \| Location \| URL" convention is extracted defensively; free-form comments keep the first line as title). Configure with `provider: hackernews` in a `job_boards:` entry. |
| HigherEdJobs | RSS | Reads the public `https://www.higheredjobs.com/rss/categoryFeed.cfm?catID={catID}` feed and parses it in-process. Configure with `provider: higheredjobs` and optional `cat_id` (default 68 = Higher Education). Not auto-detected — requires explicit `provider:` config. |
| Himalayas | API | Reads the board-wide `https://himalayas.app/jobs/api?limit=50` JSON remote-jobs feed. Configure with `provider: himalayas` in a `job_boards:` entry. |
| IBM Careers | API | Posts to IBM's public careers search API and supports optional IBM facet filters in the portal entry. |
| JibeApply | API | Auto-detects `https://<slug>.jibeapply.com/jobs` careers URLs (rewriting `/jobs` to the public `/api/jobs` endpoint); paginates `?page=N` up to `max_pages` (default 50), warning if a tenant's postings exceed the cap. Also supports branded/iCIMS-hosted sites at their own `/jobs` path via an explicit `provider: jibeapply` + `api:` URL. |
| Jobicy | API | Reads the board-wide `https://jobicy.com/api/v2/remote-jobs?count=50` JSON feed (remote-jobs aggregator). Configure with `provider: jobicy` in a `job_boards:` entry. |
| Jobspresso | RSS | Reads the public WordPress `https://jobspresso.co/?feed=job_feed` XML job feed and parses it in-process. Configure with `provider: jobspresso` in a `job_boards:` entry. |
| LaraJobs | RSS | Reads the board-wide `https://larajobs.com/feed` RSS feed (Laravel / PHP jobs) and parses it in-process. Configure with `provider: larajobs`; company and location come from the feed's `job:` namespace. |
| Lever | API | Auto-detects `https://jobs.(eu.)?lever.co/<slug>` boards and uses Lever's public postings endpoint. |
| Local parser | Parser | Runs an in-repo parser command from `portals.yml`. Use this for stable SSR or HTML pages that need a custom extractor. |
| NoDesk | RSS | Reads the public `https://nodesk.co/remote-jobs/index.xml` feed and parses it in-process. Configure with `provider: nodesk`. |
| Personio | RSS | Auto-detects `<slug>.jobs.personio.de` or `.com` hosts and parses the public XML jobs feed. |
| Phenom People | API | Reads Phenom "CareerConnect" career sites (e.g. `careers.allianz.com`) via the public no-auth `POST {origin}/widgets` JSON endpoint on the branded host. Point `careers_url`/`api` at the Phenom host (or set `provider: phenom`). |
| Pinpoint | API | Auto-detects `<slug>.pinpointhq.com` boards and reads the public zero-auth `/postings.json` per-tenant feed. |
| Radancy | Parser | Reads Radancy (TalentBrew) career sites (e.g. `careers.munichre.com`) via the server-rendered `/{lang}/search-jobs?p={N}` results page (1-based pagination). Select with `provider: radancy`. |
| Recruitee | API | Auto-detects `<slug>.recruitee.com` boards and uses the public per-tenant offers API. |
| RemoteOK | API | Reads the board-wide `https://remoteok.com/api` JSON feed; scanner filters decide which rows are relevant. |
| Remotive | API | Reads the board-wide `https://remotive.com/api/remote-jobs` JSON feed, then applies local scanner filters. |
| Rippling | API | Auto-detects `https://ats.rippling.com/<slug>/jobs` careers pages and reads the public zero-auth board API (`api.rippling.com/platform/api/ats/v1/board/<slug>/jobs`). |
| SAP SuccessFactors | Parser | Reads SF Recruiting Marketing (RMK) career sites — branded boards like `jobs.sap.com`, `jobs.zf.com`, `jobs.schaeffler.com` — via the public no-auth `/tile-search-results/?startrow=N` HTML fragment. Branded hosts carry no "successfactors" string, so select with `provider: successfactors` + `api:` the board origin. |
| SmartRecruiters | API | Auto-detects SmartRecruiters careers URLs or uses `provider: smartrecruiters` for branded custom domains. |
| a16z speedrun talent network | API | Board-wide aggregator for the a16z speedrun + wider a16z portfolio (~200 startups). Auto-detects `speedrun-talent-network.com` URLs and reads the public zero-auth `/api/v1/jobs` JSON feed (OpenAPI at `/api/v1/openapi.json`); paginates `?page=N` (0-indexed) up to `max_pages` (default 3). Optional `q:` (or `keywords:`) runs the feed's server-side full-text search with synonym expansion. Unannounced roles arrive pre-masked as "Stealth". Operated by the a16z speedrun program's talent team (provenance: the site links a16z.com's privacy/terms/disclosures; owner contact jmazer@a16z.com). Configure with `provider: a16z-speedrun-talent`. |
| Teamtailor | RSS | Auto-detects `<slug>.teamtailor.com` career sites and reads the public zero-auth `/jobs.rss` per-tenant feed. For a branded careers domain, set `provider: teamtailor` and it reads `/jobs.rss` off that host. Job links may point at a branded custom domain; location comes from the `tt:` city/country tags, falling back to `Remote` when a posting carries no `tt:city`/`tt:country` but its `remoteStatus` is remote (`fully` or `temporary`). |
| The Muse | API | Reads the public `https://www.themuse.com/api/public/jobs?page={n}` JSON feed, fetching all pages sequentially before normalizing. Configure with `provider: themuse` in a `job_boards:` entry. |
| We Work Remotely | RSS | Reads the public `https://weworkremotely.com/remote-jobs.rss` feed and parses it in-process. |
| Workable | Parser | Auto-detects `https://apply.workable.com/<slug>` and parses Workable's public markdown jobs feed. |
| Workday | API | Auto-detects `<tenant>.<instance>.myworkdayjobs.com[/<locale>]/<site>` careers URLs and posts to the public CXS jobs endpoint; paginates via offset up to `max_pages` (default 100), warning if a tenant's postings exceed the cap. |
| Working Nomads | API | Reads the board-wide `https://www.workingnomads.com/api/exposed_jobs/` JSON feed, then applies scanner filters. |

When adding a new provider, add a new non-helper module under `providers/` and
update this table in the same PR.
