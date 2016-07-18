Do this to generate your change history

    git log --pretty=format:'  * [%h](https://github.com/pact-foundation/pact-go/commit/%h) - %s (%an, %ad)' vX.Y.Z..HEAD | grep -v wip

### 0.0.2 (1 July 2016)

  * [f6e39b2](https://github.com/pact-foundation/pact-go/commit/f6e39b2) - Merge pull request #4 from pact-foundation/feature/publish-from-tag (Matt Fellows, Fri Jul 1 19:21:01 2016 +1000)
  * [deb83ef](https://github.com/pact-foundation/pact-go/commit/deb83ef) - docs(verification): add documentation for automatic consumer verification via pact broker (Matt Fellows, Fri Jul 1 19:13:42 2016 +1000)
  * [4e1539b](https://github.com/pact-foundation/pact-go/commit/4e1539b) - feature(example): add logout to Go Kit consumer (Matt Fellows, Wed Jun 15 08:58:45 2016 +1000)
  * [0cc51cd](https://github.com/pact-foundation/pact-go/commit/0cc51cd) - feat(example): dded login failure to Go Kit example (Matt Fellows, Wed Jun 15 08:50:53 2016 +1000)
  * [70c4469](https://github.com/pact-foundation/pact-go/commit/70c4469) - feat(example): simple bootstrap form for Go Kit UI (Matt Fellows, Tue Jun 14 21:30:49 2016 +1000)
  * [d5f5c51](https://github.com/pact-foundation/pact-go/commit/d5f5c51) - docs(example): add how to run go-kit consumer (Matt Fellows, Tue Jun 14 10:46:23 2016 +1000)
  * [5e33690](https://github.com/pact-foundation/pact-go/commit/5e33690) - docs(readme): update heading level for provider states (Matt Fellows, Tue Jun 14 10:44:26 2016 +1000)
  * [a118875](https://github.com/pact-foundation/pact-go/commit/a118875) - docs(example): update go-kit readme (Matt Fellows, Tue Jun 14 07:39:42 2016 +1000)
  * [ff5e832](https://github.com/pact-foundation/pact-go/commit/ff5e832) - feat(publish): automatically write pact file on shutdown/teardown (Matt Fellows, Tue Jun 14 01:41:17 2016 +1000)
  * [eba0324](https://github.com/pact-foundation/pact-go/commit/eba0324) - docs(example): go-kit test consumer and pact tests (Matt Fellows, Tue Jun 14 01:40:49 2016 +1000)
  * [3cc44ab](https://github.com/pact-foundation/pact-go/commit/3cc44ab) - docs(readme): added provider states documentation (Matt Fellows, Mon Jun 13 23:00:51 2016 +1000)
  * [eb9e6a4](https://github.com/pact-foundation/pact-go/commit/eb9e6a4) - docs(readme): added links to examples in readme (Matt Fellows, Sun Jun 12 10:35:13 2016 +1000)
  * [0bbe0be](https://github.com/pact-foundation/pact-go/commit/0bbe0be) - docs(example): move gokit readme into go-kit folder (Matt Fellows, Sun Jun 12 10:29:47 2016 +1000)
  * [0001703](https://github.com/pact-foundation/pact-go/commit/0001703) - feat(example): create example Go Kit microservice Pact testing (Matt Fellows, Sun Jun 12 10:26:21 2016 +1000)
  * [33f5394](https://github.com/pact-foundation/pact-go/commit/33f5394) - refactor(states): add types for provider state testing (Matt Fellows, Sun Jun 12 10:25:14 2016 +1000)
  * [ea2c807](https://github.com/pact-foundation/pact-go/commit/ea2c807) - feat(writepact): allow client to control when pact file is written (Matt Fellows, Sat Jun 11 23:18:26 2016 +1000)
  * [879a3b4](https://github.com/pact-foundation/pact-go/commit/879a3b4) - fix(daemon): fix daemon unable to find mock + verification service when called from $PATH (Matt Fellows, Sat Jun 11 13:01:25 2016 +1000)
  * [0ea229b](https://github.com/pact-foundation/pact-go/commit/0ea229b) - docs(godoc): update main preamble (Matt Fellows, Sat Jun 11 12:44:35 2016 +1000)
  * [7101b78](https://github.com/pact-foundation/pact-go/commit/7101b78) - docs(godoc): add package godoc comments and examples (Matt Fellows, Sat Jun 11 11:37:19 2016 +1000)
  * [eae4100](https://github.com/pact-foundation/pact-go/commit/eae4100) - test(cli): improve test coverage for daemon command (Matt Fellows, Sat Jun 11 09:36:13 2016 +1000)
  * [ef33b71](https://github.com/pact-foundation/pact-go/commit/ef33b71) - refactor(daemon): cleanup pointer usage in daemon code (Matt Fellows, Sat Jun 11 08:12:35 2016 +1000)
  * [6a9ea3d](https://github.com/pact-foundation/pact-go/commit/6a9ea3d) - docs(readme): tidy up bulleted list formatting (Matt Fellows, Fri Jun 10 20:51:14 2016 +1000)
  * [7f64f9c](https://github.com/pact-foundation/pact-go/commit/7f64f9c) - feat(publish): publish pacts to a broker (Matt Fellows, Fri Jun 10 20:44:23 2016 +1000)

### 0.0.1 (7 June 2016)

  * [bb5faac](https://github.com/pact-foundation/pact-go/commit/bb5faac) - docs(readme): cleanup bullet points in readme (Matt Fellows, Tue Jun 7 06:54:20 2016 +1000)
  * [3255d7f](https://github.com/pact-foundation/pact-go/commit/3255d7f) - feat(verifier): complete implementation of pact verification feature (Matt Fellows, Mon Jun 6 22:41:44 2016 +1000)
  * [a4c9d35](https://github.com/pact-foundation/pact-go/commit/a4c9d35) - test(verifier): increase test coverage for verification failure scenarios (Matt Fellows, Mon Jun 6 22:02:32 2016 +1000)
  * [d0b875a](https://github.com/pact-foundation/pact-go/commit/d0b875a) - test(integration): remove print statement in E2E test (Matt Fellows, Sun Jun 5 15:03:13 2016 +1000)
  * [671a63c](https://github.com/pact-foundation/pact-go/commit/671a63c) - chore(build): exclude vendor directory from golint (Matt Fellows, Sun Jun 5 14:53:00 2016 +1000)
  * [0cd4f08](https://github.com/pact-foundation/pact-go/commit/0cd4f08) - chore(vendor): vendor dependent packages (Matt Fellows, Sun Jun 5 14:48:31 2016 +1000)
  * [5727682](https://github.com/pact-foundation/pact-go/commit/5727682) - docs(readme): remove code climate as its useless (Matt Fellows, Sun Jun 5 14:37:44 2016 +1000)
  * [610a54d](https://github.com/pact-foundation/pact-go/commit/610a54d) - docs(readme): format log example as Go code (Matt Fellows, Sun Jun 5 14:31:49 2016 +1000)
  * [db83055](https://github.com/pact-foundation/pact-go/commit/db83055) - feat(logging): standardise log and output messages (Matt Fellows, Sun Jun 5 14:25:01 2016 +1000)
  * [90b94e2](https://github.com/pact-foundation/pact-go/commit/90b94e2) - fix(readme): fix broken link in readme (Matt Fellows, Sun Jun 5 11:07:30 2016 +1000)
  * [5cf2b95](https://github.com/pact-foundation/pact-go/commit/5cf2b95) - Merge pull request #2 from pact-foundation/feature/body-matching (Matt Fellows, Sun Jun 5 11:05:38 2016 +1000)
  * [6b6aada](https://github.com/pact-foundation/pact-go/commit/6b6aada) - docs(readme): improved examples (Matt Fellows, Sun Jun 5 11:02:54 2016 +1000)
  * [20354df](https://github.com/pact-foundation/pact-go/commit/20354df) - chore(build): fix codeclimate configuration (Matt Fellows, Sun Jun 5 08:02:13 2016 +1000)
  * [78a0932](https://github.com/pact-foundation/pact-go/commit/78a0932) - chore(readme): updated build/coveralls badges (Matt Fellows, Sat Jun 4 21:19:52 2016 +1000)
  * [4077448](https://github.com/pact-foundation/pact-go/commit/4077448) - docs(contributing): added contribution guidelines (Matt Fellows, Sat Jun 4 20:47:08 2016 +1000)
  * [97be321](https://github.com/pact-foundation/pact-go/commit/97be321) - chore(build): build for Pact Foundation in Wercker (Matt Fellows, Sat Jun 4 19:17:00 2016 +1000)
  * [1ce233c](https://github.com/pact-foundation/pact-go/commit/1ce233c) - chore(build): moved to pact-foundation repository (Matt Fellows, Sat Jun 4 18:59:53 2016 +1000)
  * [ee3f1a8](https://github.com/pact-foundation/pact-go/commit/ee3f1a8) - style(readme): using Go syntax highlighter (Matt Fellows, Sat Jun 4 18:56:28 2016 +1000)
  * [7c3c70b](https://github.com/pact-foundation/pact-go/commit/7c3c70b) - feat(consumerdsl): working v1 consumer DSL interface (Matt Fellows, Sat Jun 4 18:52:38 2016 +1000)
  * [3b859cf](https://github.com/pact-foundation/pact-go/commit/3b859cf) - test(daemon): RPC Client error tests (Matt Fellows, Sat Jun 4 09:38:58 2016 +1000)
  * [7d8a0fd](https://github.com/pact-foundation/pact-go/commit/7d8a0fd) - test(service): remove need for a running pact mock service in tests (Matt Fellows, Fri Jun 3 22:05:12 2016 +1000)
  * [98b7c57](https://github.com/pact-foundation/pact-go/commit/98b7c57) - test(service): improved test coverage for services (Matt Fellows, Fri Jun 3 21:08:25 2016 +1000)
  * [c3f11e1](https://github.com/pact-foundation/pact-go/commit/c3f11e1) - test(daemon): coverage for shutting down the daemon via RPC (Matt Fellows, Fri Jun 3 08:19:35 2016 +1000)
  * [83c75c2](https://github.com/pact-foundation/pact-go/commit/83c75c2) - test(daemon): RPC client testing (Matt Fellows, Thu Jun 2 21:17:28 2016 +1000)
  * [3c71bba](https://github.com/pact-foundation/pact-go/commit/3c71bba) - test(service): initial service test coverage (Matt Fellows, Wed Jun 1 21:09:27 2016 +1000)
  * [27d50d4](https://github.com/pact-foundation/pact-go/commit/27d50d4) - test(daemon): RPC style tests to ensure client can work (Matt Fellows, Wed Jun 1 08:47:20 2016 +1000)
  * [82692af](https://github.com/pact-foundation/pact-go/commit/82692af) - style(daemon): formatting (Matt Fellows, Tue May 31 21:36:17 2016 +1000)
  * [5ccfb64](https://github.com/pact-foundation/pact-go/commit/5ccfb64) - test(daemon): working tests with RPC integration-style tests (Matt Fellows, Tue May 31 18:36:22 2016 +1000)
  * [3f764a6](https://github.com/pact-foundation/pact-go/commit/3f764a6) - test(daemon): working daemon tests (Matt Fellows, Tue May 31 10:30:53 2016 +1000)
  * [24e7c9b](https://github.com/pact-foundation/pact-go/commit/24e7c9b) - test(daemon): improved daemon test coverage (Matt Fellows, Tue May 31 07:30:17 2016 +1000)
  * [345cd9f](https://github.com/pact-foundation/pact-go/commit/345cd9f) - test(daemon): daemon testing (Matt Fellows, Mon May 30 22:41:15 2016 +1000)
  * [8c13832](https://github.com/pact-foundation/pact-go/commit/8c13832) - test(daemon): improved daemon test coverage for listing servers (Matt Fellows, Wed May 25 22:10:42 2016 +1000)
  * [5e74ac6](https://github.com/pact-foundation/pact-go/commit/5e74ac6) - chore(build): added coveralls publish step (Matt Fellows, Sun May 22 18:31:15 2016 +1000)
  * [5288cdf](https://github.com/pact-foundation/pact-go/commit/5288cdf) - e3aeb6a chore(build): fix tests (Matt Fellows, Sun May 22 18:27:16 2016 +1000)
  * [7149465](https://github.com/pact-foundation/pact-go/commit/7149465) - chore(build): linting and windows deps (Matt Fellows, Sun May 22 18:18:37 2016 +1000)
  * [ace86a2](https://github.com/pact-foundation/pact-go/commit/ace86a2) - chore(build): wercker build and required scripting (Matt Fellows, Sat May 21 20:36:22 2016 +1000)
  * [2b5b06e](https://github.com/pact-foundation/pact-go/commit/2b5b06e) - feat(consumerdsl): initial shell for Pact Go (Matt Fellows, Sat May 21 13:27:21 2016 +1000)