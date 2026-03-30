2026-02-17T21:51:54.2720596Z Current runner version: '2.331.0'
2026-02-17T21:51:54.2744064Z ##[group]Runner Image Provisioner
2026-02-17T21:51:54.2745027Z Hosted Compute Agent
2026-02-17T21:51:54.2745626Z Version: 20260123.484
2026-02-17T21:51:54.2746237Z Commit: 6bd6555ca37d84114959e1c76d2c01448ff61c5d
2026-02-17T21:51:54.2747056Z Build Date: 2026-01-23T19:41:17Z
2026-02-17T21:51:54.2747682Z Worker ID: {fd081f8e-72e4-41e0-8e2e-aa5cd45d7a5e}
2026-02-17T21:51:54.2748359Z Azure Region: westcentralus
2026-02-17T21:51:54.2749018Z ##[endgroup]
2026-02-17T21:51:54.2751041Z ##[group]Operating System
2026-02-17T21:51:54.2751752Z Ubuntu
2026-02-17T21:51:54.2752266Z 24.04.3
2026-02-17T21:51:54.2752739Z LTS
2026-02-17T21:51:54.2753276Z ##[endgroup]
2026-02-17T21:51:54.2753831Z ##[group]Runner Image
2026-02-17T21:51:54.2754368Z Image: ubuntu-24.04
2026-02-17T21:51:54.2754981Z Version: 20260209.23.1
2026-02-17T21:51:54.2756222Z Included Software: https://github.com/actions/runner-images/blob/ubuntu24/20260209.23/images/ubuntu/Ubuntu2404-Readme.md
2026-02-17T21:51:54.2757764Z Image Release: https://github.com/actions/runner-images/releases/tag/ubuntu24%2F20260209.23
2026-02-17T21:51:54.2758709Z ##[endgroup]
2026-02-17T21:51:54.2760496Z ##[group]GITHUB_TOKEN Permissions
2026-02-17T21:51:54.2762444Z Contents: read
2026-02-17T21:51:54.2762994Z Metadata: read
2026-02-17T21:51:54.2763630Z Packages: read
2026-02-17T21:51:54.2764161Z ##[endgroup]
2026-02-17T21:51:54.2766320Z Secret source: None
2026-02-17T21:51:54.2767161Z Prepare workflow directory
2026-02-17T21:51:54.3154040Z Prepare all required actions
2026-02-17T21:51:54.3192342Z Getting action download info
2026-02-17T21:51:54.6992347Z Download action repository 'github/dependabot-action@main' (SHA:2c339a9695a8285e56f36eef47ae73bf772509d7)
2026-02-17T21:51:55.7004346Z Complete job name: Dependabot
2026-02-17T21:51:55.7780827Z ##[group]Run mkdir -p  ./dependabot-job-1247168785-1771365107
2026-02-17T21:51:55.7781783Z [36;1mmkdir -p  ./dependabot-job-1247168785-1771365107[0m
2026-02-17T21:51:55.7834821Z shell: /usr/bin/bash -e {0}
2026-02-17T21:51:55.7836110Z ##[endgroup]
2026-02-17T21:51:55.8140428Z ##[group]Run github/dependabot-action@main
2026-02-17T21:51:55.8141138Z env:
2026-02-17T21:51:55.8141604Z   DEPENDABOT_DISABLE_CLEANUP: 1
2026-02-17T21:51:55.8142175Z   DEPENDABOT_ENABLE_CONNECTIVITY_CHECK: 0
2026-02-17T21:51:55.8142995Z   GITHUB_TOKEN: ***
2026-02-17T21:51:55.8143839Z   GITHUB_DEPENDABOT_JOB_TOKEN: ***
2026-02-17T21:51:55.8144798Z   GITHUB_DEPENDABOT_CRED_TOKEN: ***
2026-02-17T21:51:55.8145379Z   GITHUB_REGISTRIES_PROXY: ***
2026-02-17T21:51:55.8145896Z ##[endgroup]
2026-02-17T21:51:56.0215258Z 🤖 ~ starting update ~
2026-02-17T21:51:56.0234174Z Fetching job details
2026-02-17T21:51:56.6889976Z ##[group]Pulling updater images
2026-02-17T21:51:56.6973212Z Pulling image ghcr.io/dependabot/dependabot-updater-npm:929cd12fa29d79b59f38970ba313c752d2608ad7 (attempt 1)...
2026-02-17T21:51:56.8646735Z Successfully sent metric (dependabot.action.ghcr_image_pull) to remote API endpoint
2026-02-17T21:52:09.8530171Z Pulled image ghcr.io/dependabot/dependabot-updater-npm:929cd12fa29d79b59f38970ba313c752d2608ad7
2026-02-17T21:52:09.8551275Z Pulling image ghcr.io/dependabot/proxy:v2.0.20260129233510@sha256:aee1af4a514c0c5e573f3b33a51f9f2b9c58234cb011ea4d44b9e05aec92436c (attempt 1)...
2026-02-17T21:52:10.0365719Z Successfully sent metric (dependabot.action.ghcr_image_pull) to remote API endpoint
2026-02-17T21:52:11.0062898Z Pulled image ghcr.io/dependabot/proxy:v2.0.20260129233510@sha256:aee1af4a514c0c5e573f3b33a51f9f2b9c58234cb011ea4d44b9e05aec92436c
2026-02-17T21:52:11.0065113Z ##[endgroup]
2026-02-17T21:52:12.9423651Z Starting update process
2026-02-17T21:52:12.9424699Z Created proxy container: c680794852e7a080d31caa68d68923020a5b29f2cc96a0c801c6c1756f4b555e
2026-02-17T21:52:13.2127354Z Created container: 61f7dd09efc3381e6bc889952f6f44ab2fe2276088b6f0fae505cc9326ae09d1
2026-02-17T21:52:13.2521754Z   proxy | 2026/02/17 21:52:13 proxy starting, commit: dd760113d8edd581443a6aaafe80d4e4025251ab
2026-02-17T21:52:13.2524216Z   proxy | 2026/02/17 21:52:13 Listening (:1080)
2026-02-17T21:52:13.3274601Z Started container 61f7dd09efc3381e6bc889952f6f44ab2fe2276088b6f0fae505cc9326ae09d1
2026-02-17T21:52:13.3776670Z updater | Updating certificates in /etc/ssl/certs...
2026-02-17T21:52:14.2504522Z updater | rehash: warning: skipping ca-certificates.crt,it does not contain exactly one certificate or CRL
2026-02-17T21:52:14.2579458Z updater | 1 added, 0 removed; done.
2026-02-17T21:52:14.2584606Z updater | Running hooks in /etc/ca-certificates/update.d...
2026-02-17T21:52:14.2605885Z updater | done.
2026-02-17T21:52:14.5576581Z updater | fetch_files command is no longer used directly
2026-02-17T21:52:16.5107125Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Starting job processing
2026-02-17T21:52:16.5125559Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Job definition: {"job":{"command":"security","allowed-updates":[{"dependency-type":"direct","update-type":"all"}],"commit-message-options":{"prefix":null,"prefix-development":null,"include-scope":null},"credentials-metadata":[{"type":"git_source","host":"github.com"}],"debug":null,"dependencies":["pbkdf2"],"dependency-groups":[],"dependency-group-to-refresh":null,"existing-pull-requests":[{"pr-number":114,"dependencies":[{"dependency-name":"ini","dependency-version":"1.3.8"}]},{"pr-number":115,"dependencies":[{"dependency-name":"y18n","dependency-version":"3.2.2"}]},{"pr-number":118,"dependencies":[{"dependency-name":"hosted-git-info","dependency-version":"2.8.9"}]},{"pr-number":125,"dependencies":[{"dependency-name":"normalize-url","dependency-version":"4.5.1"}]},{"pr-number":132,"dependencies":[{"dependency-name":"tar","dependency-version":"4.4.15"}]},{"pr-number":133,"dependencies":[{"dependency-name":"path-parse","dependency-version":"1.0.7"}]},{"pr-number":136,"dependencies":[{"dependency-name":"tar","dependency-version":"4.4.19"}]},{"pr-number":150,"dependencies":[{"dependency-name":"pathval","dependency-version":"1.1.1"}]},{"pr-number":151,"dependencies":[{"dependency-name":"ajv","dependency-version":"6.12.6"}]},{"pr-number":166,"dependencies":[{"dependency-name":"copy-props","dependency-version":"2.0.5"}]},{"pr-number":177,"dependencies":[{"dependency-name":"decode-uri-component","dependency-version":"0.2.2"}]},{"pr-number":179,"dependencies":[{"dependency-name":"qs","dependency-version":"6.5.3"}]},{"pr-number":180,"dependencies":[{"dependency-name":"express","dependency-version":"4.18.2"}]},{"pr-number":181,"dependencies":[{"dependency-name":"cookiejar","dependency-version":"2.1.4"}]},{"pr-number":182,"dependencies":[{"dependency-name":"http-cache-semantics","dependency-version":"4.1.1"}]},{"pr-number":195,"dependencies":[{"dependency-name":"simple-get","dependency-version":"2.8.2"}]},{"pr-number":196,"dependencies":[{"dependency-name":"cross-fetch","dependency-version":"2.2.6"}]},{"pr-number":206,"dependencies":[{"dependency-name":"get-func-name","dependency-version":"2.0.2"}]},{"pr-number":207,"dependencies":[{"dependency-name":"browserify-sign","dependency-version":"4.2.2"}]},{"pr-number":211,"dependencies":[{"dependency-name":"es5-ext","dependency-version":"0.10.63"}]},{"pr-number":213,"dependencies":[{"dependency-name":"express","dependency-version":"4.19.2"}]},{"pr-number":219,"dependencies":[{"dependency-name":"express","dependency-version":"4.20.0","directory":"/"}]},{"pr-number":224,"dependencies":[{"dependency-name":"secp256k1","dependency-version":"3.8.1","directory":"/"}]},{"pr-number":230,"dependencies":[{"dependency-name":"base-x","dependency-version":"3.0.11","directory":"/"}]},{"pr-number":233,"dependencies":[{"dependency-name":"pbkdf2","dependency-version":"3.1.3","directory":"/"}]},{"pr-number":238,"dependencies":[{"dependency-name":"cipher-base","dependency-version":"1.0.6","directory":"/"}]}],"existing-group-pull-requests":[],"experiments":{"record-ecosystem-versions":true,"record-update-job-unknown-error":true,"proxy-cached":true,"enable-record-ecosystem-meta":true,"enable-corepack-for-npm-and-yarn":true,"enable-private-registry-for-corepack":true,"enable-shared-helpers-command-timeout":true,"avoid-duplicate-updates-package-json":true,"allow-refresh-for-existing-pr-dependencies":true,"allow-refresh-group-with-all-dependencies":true,"enable-enhanced-error-details-for-updater":true,"gradle-lockfile-updater":true,"enable-exclude-paths-subdirectory-manifest-files":true,"group-membership-enforcement":true,"deprecate-close-command":true,"deprecate-reopen-command":true,"deprecate-merge-command":true,"deprecate-cancel-merge-command":true,"deprecate-squash-merge-command":true,"disable-close-command":true,"disable-reopen-command":true,"disable-merge-command":true,"disable-cancel-merge-command":true,"disable-squash-merge-command":true},"ignore-conditions":[],"lockfile-only":false,"max-updater-run-time":2700,"package-manager":"npm_and_yarn","requirements-update-strategy":null,"reject-external-code":false,"security-advisories":[{"dependency-name":"pbkdf2","patched-versions":[],"unaffected-versions":[],"affected-versions":[">= 1.0.0 <= 3.1.2"]},{"dependency-name":"pbkdf2","patched-versions":[],"unaffected-versions":[],"affected-versions":[">= 3.0.10 <= 3.1.2"]}],"security-updates-only":true,"source":{"provider":"github","repo":"Uniswap/v2-core","branch":null,"api-endpoint":"https://api.github.com/","hostname":"github.com","directories":["/."]},"updating-a-pull-request":false,"update-subdependencies":false,"vendor-dependencies":false,"enable-beta-ecosystems":false,"repo-private":false,"multi-ecosystem-update":false,"exclude-paths":null}}
2026-02-17T21:52:16.5345995Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1095 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/git.store' {}
2026-02-17T21:52:16.5397819Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1095 completed with status: pid 1095 exit 0
2026-02-17T21:52:16.5407610Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:16.5415472Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1104 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:16.5454670Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1104 completed with status: pid 1104 exit 0
2026-02-17T21:52:16.5455545Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:16.5461775Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1112 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:16.5504768Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1112 completed with status: pid 1112 exit 0
2026-02-17T21:52:16.5505689Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:16.5511880Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1119 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:16.5556532Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1119 completed with status: pid 1119 exit 0
2026-02-17T21:52:16.5557420Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:16.5563606Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1126 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:16.5608324Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1126 completed with status: pid 1126 exit 0
2026-02-17T21:52:16.5609210Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:16.5615253Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1133 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:16.5660179Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1133 completed with status: pid 1133 exit 0
2026-02-17T21:52:16.5661108Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:16.5666353Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1141 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:16.5708486Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1141 completed with status: pid 1141 exit 0
2026-02-17T21:52:16.5709404Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:16.5898468Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1148 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:16.5900506Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1148 completed with status: pid 1148 exit 0
2026-02-17T21:52:16.5901625Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.02 seconds
2026-02-17T21:52:16.5906629Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1155 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:16.5944318Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1155 completed with status: pid 1155 exit 0
2026-02-17T21:52:16.5944935Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:16.5952158Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1162 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:16.5997113Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1162 completed with status: pid 1162 exit 0
2026-02-17T21:52:16.5998071Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:16.6004162Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1169 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:16.6048748Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Process PID: 1169 completed with status: pid 1169 exit 0
2026-02-17T21:52:16.6049994Z 2026/02/17 21:52:16 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:16.6060253Z updater | 2026/02/17 21:52:16 INFO <job_1247168785> Started process PID: 1176 with command: {} git clone --no-tags --depth 1 --recurse-submodules --shallow-submodules https://github.com/Uniswap/v2-core /home/dependabot/dependabot-updater/repo {}
2026-02-17T21:52:16.8123821Z   proxy | 2026/02/17 21:52:16 [002] GET https://github.com:443/Uniswap/v2-core/info/refs?service=git-upload-pack
2026-02-17T21:52:16.8125519Z   proxy | 2026/02/17 21:52:16 [002] * authenticating git server request (host: github.com)
2026-02-17T21:52:17.0242546Z   proxy | 2026/02/17 21:52:17 [002] 200 https://github.com:443/Uniswap/v2-core/info/refs?service=git-upload-pack
2026-02-17T21:52:17.0644519Z   proxy | 2026/02/17 21:52:17 [004] POST https://github.com:443/Uniswap/v2-core/git-upload-pack
2026-02-17T21:52:17.0645540Z 2026/02/17 21:52:17 [004] * authenticating git server request (host: github.com)
2026-02-17T21:52:17.1577244Z   proxy | 2026/02/17 21:52:17 [004] 200 https://github.com:443/Uniswap/v2-core/git-upload-pack
2026-02-17T21:52:17.1934867Z   proxy | 2026/02/17 21:52:17 [006] POST https://github.com:443/Uniswap/v2-core/git-upload-pack
2026-02-17T21:52:17.1935904Z 2026/02/17 21:52:17 [006] * authenticating git server request (host: github.com)
2026-02-17T21:52:17.2930632Z   proxy | 2026/02/17 21:52:17 [006] 200 https://github.com:443/Uniswap/v2-core/git-upload-pack
2026-02-17T21:52:17.4728629Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1176 completed with status: pid 1176 exit 0
2026-02-17T21:52:17.4751640Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.87 seconds
2026-02-17T21:52:17.4753400Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1218 with command: {} git -C /home/dependabot/dependabot-updater/repo ls-files --stage {}
2026-02-17T21:52:17.4799379Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1218 completed with status: pid 1218 exit 0
2026-02-17T21:52:17.4800776Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.4853022Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1234 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/git.store' {}
2026-02-17T21:52:17.4904078Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1234 completed with status: pid 1234 exit 0
2026-02-17T21:52:17.4921715Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.4941963Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1242 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.4986287Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1242 completed with status: pid 1242 exit 0
2026-02-17T21:52:17.4987428Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5029323Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1249 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.5068912Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1249 completed with status: pid 1249 exit 0
2026-02-17T21:52:17.5070365Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5076617Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1256 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.5120266Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1256 completed with status: pid 1256 exit 0
2026-02-17T21:52:17.5121172Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5126294Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1263 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.5170370Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1263 completed with status: pid 1263 exit 0
2026-02-17T21:52:17.5171403Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.5176224Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1270 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.5228029Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1270 completed with status: pid 1270 exit 0
2026-02-17T21:52:17.5229932Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5246795Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1277 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.5277131Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1277 completed with status: pid 1277 exit 0
2026-02-17T21:52:17.5278135Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.5292614Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1285 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.5334133Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1285 completed with status: pid 1285 exit 0
2026-02-17T21:52:17.5336747Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5345051Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1293 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.5390793Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1293 completed with status: pid 1293 exit 0
2026-02-17T21:52:17.5394421Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5399988Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1300 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.5447561Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1300 completed with status: pid 1300 exit 0
2026-02-17T21:52:17.5450836Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5457729Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1307 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.5509206Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1307 completed with status: pid 1307 exit 0
2026-02-17T21:52:17.5510491Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.5520076Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1314 with command: {} git lfs pull --include .yarn,./yarn/cache {}
2026-02-17T21:52:17.6274978Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1314 completed with status: pid 1314 exit 0
2026-02-17T21:52:17.6275977Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.08 seconds
2026-02-17T21:52:17.6324294Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1355 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/git.store' {}
2026-02-17T21:52:17.6365540Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1355 completed with status: pid 1355 exit 0
2026-02-17T21:52:17.6366695Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6373694Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1364 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.6413015Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1364 completed with status: pid 1364 exit 0
2026-02-17T21:52:17.6414038Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6420142Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1371 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.6463213Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1371 completed with status: pid 1371 exit 0
2026-02-17T21:52:17.6464218Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6470610Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1378 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.6510096Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1378 completed with status: pid 1378 exit 0
2026-02-17T21:52:17.6511133Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6516990Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1385 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.6553063Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1385 completed with status: pid 1385 exit 0
2026-02-17T21:52:17.6554385Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6558903Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1393 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.6594845Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1393 completed with status: pid 1393 exit 0
2026-02-17T21:52:17.6595852Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6601301Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1400 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.6647891Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1400 completed with status: pid 1400 exit 0
2026-02-17T21:52:17.6649175Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.6654047Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1408 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.6699709Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1408 completed with status: pid 1408 exit 0
2026-02-17T21:52:17.6700750Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6708222Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1416 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.6742411Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1416 completed with status: pid 1416 exit 0
2026-02-17T21:52:17.6743320Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6749951Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1425 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.6791171Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1425 completed with status: pid 1425 exit 0
2026-02-17T21:52:17.6792547Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6798161Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1434 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.6835257Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1434 completed with status: pid 1434 exit 0
2026-02-17T21:52:17.6836263Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.6844403Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1441 with command: {} git rev-parse HEAD {}
2026-02-17T21:52:17.6888093Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1441 completed with status: pid 1441 exit 0
2026-02-17T21:52:17.6891474Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.6936850Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1458 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/git.store' {}
2026-02-17T21:52:17.6981377Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1458 completed with status: pid 1458 exit 0
2026-02-17T21:52:17.6982438Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.6987892Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1466 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.7030062Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1466 completed with status: pid 1466 exit 0
2026-02-17T21:52:17.7031342Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7035261Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1473 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.7082265Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1473 completed with status: pid 1473 exit 0
2026-02-17T21:52:17.7083239Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7088691Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1480 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.7133283Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1480 completed with status: pid 1480 exit 0
2026-02-17T21:52:17.7134258Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7157438Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1488 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.7192939Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1488 completed with status: pid 1488 exit 0
2026-02-17T21:52:17.7193952Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7199075Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1495 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.7241944Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1495 completed with status: pid 1495 exit 0
2026-02-17T21:52:17.7243031Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7250470Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1502 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.7292440Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1502 completed with status: pid 1502 exit 0
2026-02-17T21:52:17.7293084Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7298995Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1509 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.7338810Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1509 completed with status: pid 1509 exit 0
2026-02-17T21:52:17.7339829Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7360883Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1516 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.7399339Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1516 completed with status: pid 1516 exit 0
2026-02-17T21:52:17.7400888Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7415760Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1523 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.7456299Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1523 completed with status: pid 1523 exit 0
2026-02-17T21:52:17.7458218Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7472519Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1531 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.7516067Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1531 completed with status: pid 1531 exit 0
2026-02-17T21:52:17.7517681Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7534684Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1540 with command: {} git rev-parse HEAD {}
2026-02-17T21:52:17.7585940Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1540 completed with status: pid 1540 exit 0
2026-02-17T21:52:17.7593883Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7704799Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1555 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/git.store' {}
2026-02-17T21:52:17.7752232Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1555 completed with status: pid 1555 exit 0
2026-02-17T21:52:17.7754015Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7759839Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1564 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.7803633Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1564 completed with status: pid 1564 exit 0
2026-02-17T21:52:17.7806163Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.7810002Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1572 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.7851958Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1572 completed with status: pid 1572 exit 0
2026-02-17T21:52:17.7853545Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7858656Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1579 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.7902868Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1579 completed with status: pid 1579 exit 0
2026-02-17T21:52:17.7903874Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7910338Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1586 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.7952765Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1586 completed with status: pid 1586 exit 0
2026-02-17T21:52:17.7953758Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.7959115Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1593 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.8001520Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1593 completed with status: pid 1593 exit 0
2026-02-17T21:52:17.8002598Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8008175Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1600 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.8053341Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1600 completed with status: pid 1600 exit 0
2026-02-17T21:52:17.8054264Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8060174Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1607 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.8104786Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1607 completed with status: pid 1607 exit 0
2026-02-17T21:52:17.8105768Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8111018Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1614 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.8155511Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1614 completed with status: pid 1614 exit 0
2026-02-17T21:52:17.8156569Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8171342Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1622 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.8197819Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1622 completed with status: pid 1622 exit 0
2026-02-17T21:52:17.8198942Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8205332Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1629 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.8250874Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1629 completed with status: pid 1629 exit 0
2026-02-17T21:52:17.8251996Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8298313Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1643 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/git.store' {}
2026-02-17T21:52:17.8347426Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1643 completed with status: pid 1643 exit 0
2026-02-17T21:52:17.8348359Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8356148Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1651 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.8392859Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1651 completed with status: pid 1651 exit 0
2026-02-17T21:52:17.8395954Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8405602Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1658 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.8444799Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1658 completed with status: pid 1658 exit 0
2026-02-17T21:52:17.8445394Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8454558Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1665 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.8489028Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1665 completed with status: pid 1665 exit 0
2026-02-17T21:52:17.8490500Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8495566Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1672 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.8538542Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1672 completed with status: pid 1672 exit 0
2026-02-17T21:52:17.8540025Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8545703Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1679 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.8590197Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1679 completed with status: pid 1679 exit 0
2026-02-17T21:52:17.8590931Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8596343Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1686 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:17.8638224Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1686 completed with status: pid 1686 exit 0
2026-02-17T21:52:17.8639256Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8644098Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1693 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:17.8693113Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1693 completed with status: pid 1693 exit 0
2026-02-17T21:52:17.8694152Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8699168Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1702 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:17.8742936Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1702 completed with status: pid 1702 exit 0
2026-02-17T21:52:17.8743865Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8748721Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1710 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:17.8793696Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1710 completed with status: pid 1710 exit 0
2026-02-17T21:52:17.8794850Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:17.8805655Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1717 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:17.8835650Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1717 completed with status: pid 1717 exit 0
2026-02-17T21:52:17.8837468Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:17.8844865Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1725 with command: {} git lfs pull --include .yarn,./yarn/cache {}
2026-02-17T21:52:17.9198453Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Process PID: 1725 completed with status: pid 1725 exit 0
2026-02-17T21:52:17.9200390Z 2026/02/17 21:52:17 INFO <job_1247168785> Total execution time: 0.04 seconds
2026-02-17T21:52:17.9254943Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:17.9256912Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:17.9307475Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:17.9316256Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Fetching version for package manager: yarn
2026-02-17T21:52:17.9324512Z updater | 2026/02/17 21:52:17 INFO <job_1247168785> Started process PID: 1758 with command: {} corepack yarn -v {}
2026-02-17T21:52:18.2263001Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Process PID: 1758 completed with status: pid 1758 exit 0
2026-02-17T21:52:18.2263913Z 2026/02/17 21:52:18 INFO <job_1247168785> Total execution time: 0.29 seconds
2026-02-17T21:52:18.2264619Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Installed version of yarn: 4.9.2
2026-02-17T21:52:18.2265374Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:18.2275071Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:18.2276005Z 2026/02/17 21:52:18 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:18.2276817Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:18.2277579Z 2026/02/17 21:52:18 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:18.2280599Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:18.2281726Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:18.2282584Z 2026/02/17 21:52:18 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:18.2283400Z 2026/02/17 21:52:18 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:18.2287050Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:18.2287893Z 2026/02/17 21:52:18 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:18.2291208Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:18.2292153Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:18.2293020Z 2026/02/17 21:52:18 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:18.2293848Z 2026/02/17 21:52:18 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:18.2296441Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:18.2297171Z 2026/02/17 21:52:18 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:18.2308148Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:18.2310108Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:18.2310918Z 2026/02/17 21:52:18 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:18.2312028Z 2026/02/17 21:52:18 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:18.2313741Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:18.2318700Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Installing "yarn@1"
2026-02-17T21:52:18.2343565Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Started process PID: 1771 with command: {} corepack prepare yarn@1 --activate {}
2026-02-17T21:52:18.4804798Z   proxy | 2026/02/17 21:52:18 [009] GET https://registry.npmjs.org:443/yarn
2026-02-17T21:52:18.5226380Z   proxy | 2026/02/17 21:52:18 [010] GET https://repo.yarnpkg.com:443/tags
2026-02-17T21:52:18.5811903Z   proxy | 2026/02/17 21:52:18 [009] 200 https://registry.npmjs.org:443/yarn
2026-02-17T21:52:18.6402173Z   proxy | 2026/02/17 21:52:18 [010] 200 https://repo.yarnpkg.com:443/tags
2026-02-17T21:52:18.6834191Z   proxy | 2026/02/17 21:52:18 [012] GET https://registry.yarnpkg.com:443/yarn/-/yarn-1.22.22.tgz
2026-02-17T21:52:18.7835795Z   proxy | 2026/02/17 21:52:18 [012] 200 https://registry.yarnpkg.com:443/yarn/-/yarn-1.22.22.tgz
2026-02-17T21:52:18.8833313Z   proxy | 2026/02/17 21:52:18 [014] GET https://registry.npmjs.org:443/yarn/1.22.22
2026-02-17T21:52:18.9242300Z   proxy | 2026/02/17 21:52:18 [014] 200 https://registry.npmjs.org:443/yarn/1.22.22
2026-02-17T21:52:18.9384129Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Process PID: 1771 completed with status: pid 1771 exit 0
2026-02-17T21:52:18.9386281Z 2026/02/17 21:52:18 INFO <job_1247168785> Total execution time: 0.71 seconds
2026-02-17T21:52:18.9387372Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> yarn@1 successfully installed.
2026-02-17T21:52:18.9388252Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Activating currently installed version of yarn: 1
2026-02-17T21:52:18.9388970Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Fetching version for package manager: yarn
2026-02-17T21:52:18.9396246Z updater | 2026/02/17 21:52:18 INFO <job_1247168785> Started process PID: 1784 with command: {} corepack yarn -v {}
2026-02-17T21:52:19.1738944Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Process PID: 1784 completed with status: pid 1784 exit 0
2026-02-17T21:52:19.1740027Z 2026/02/17 21:52:19 INFO <job_1247168785> Total execution time: 0.24 seconds
2026-02-17T21:52:19.1740764Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version of yarn: 1.22.22
2026-02-17T21:52:19.1744376Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.1745151Z 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.1750000Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.1751825Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:19.1753961Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.1755160Z 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.1756131Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.1757358Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.1762338Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.1763386Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:19.1767985Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.1769080Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.1782805Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.1784060Z 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.1788159Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.1791715Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:19.1792889Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.1794108Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.1795474Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.1796434Z 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.1797508Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.1799313Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:19.1800658Z 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.1801916Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.1803057Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.1804218Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.1806784Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.1809302Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:19.1810711Z 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.1811585Z 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.1812492Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.1813534Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.1815129Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.1816159Z 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 4.9.2
2026-02-17T21:52:19.1816908Z 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.1817491Z 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.2985920Z   proxy | 2026/02/17 21:52:19 [016] POST /update_jobs/1247168785/record_ecosystem_versions
2026-02-17T21:52:19.4889227Z   proxy | 2026/02/17 21:52:19 [016] 204 /update_jobs/1247168785/record_ecosystem_versions
2026-02-17T21:52:19.4928451Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Base commit SHA: ee547b17853e71ed4e0101ccfd52e70d5acded58
2026-02-17T21:52:19.4929887Z 2026/02/17 21:52:19 INFO <job_1247168785> Finished job processing
2026-02-17T21:52:19.4939049Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Starting job processing
2026-02-17T21:52:19.4961332Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Detected package manager: yarn
2026-02-17T21:52:19.4962390Z 2026/02/17 21:52:19 INFO <job_1247168785> Resolving package manager for: yarn
2026-02-17T21:52:19.4964194Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Guessed version info "yarn" : "1"
2026-02-17T21:52:19.4965899Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Fetching version for package manager: yarn
2026-02-17T21:52:19.4972848Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Started process PID: 1796 with command: {} corepack yarn -v {}
2026-02-17T21:52:19.6573451Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Process PID: 1796 completed with status: pid 1796 exit 0
2026-02-17T21:52:19.6574768Z 2026/02/17 21:52:19 INFO <job_1247168785> Total execution time: 0.16 seconds
2026-02-17T21:52:19.6575785Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Installed version of yarn: 1.22.22
2026-02-17T21:52:19.6577097Z 2026/02/17 21:52:19 INFO <job_1247168785> Installed version for yarn: 1.22.22
2026-02-17T21:52:19.6578313Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for yarn
2026-02-17T21:52:19.6579404Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> No version requirement found for yarn
2026-02-17T21:52:19.6581953Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Running node command: node -v
2026-02-17T21:52:19.6586477Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Started process PID: 1808 with command: {} node -v {}
2026-02-17T21:52:19.6626806Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Process PID: 1808 completed with status: pid 1808 exit 0
2026-02-17T21:52:19.6628582Z 2026/02/17 21:52:19 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:19.6629824Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Command executed successfully: node -v
2026-02-17T21:52:19.6630668Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Processing engine constraints for node
2026-02-17T21:52:19.6633641Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Parsed constraints for node: >=10
2026-02-17T21:52:19.6670339Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Started process PID: 1810 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:19.8593204Z updater | 2026/02/17 21:52:19 INFO <job_1247168785> Process PID: 1810 completed with status: pid 1810 exit 0
2026-02-17T21:52:19.8594284Z 2026/02/17 21:52:19 INFO <job_1247168785> Total execution time: 0.19 seconds
2026-02-17T21:52:20.0838237Z   proxy | 2026/02/17 21:52:20 [018] POST /update_jobs/1247168785/update_dependency_list
2026-02-17T21:52:20.2663486Z   proxy | 2026/02/17 21:52:20 [018] 204 /update_jobs/1247168785/update_dependency_list
2026-02-17T21:52:20.3664695Z   proxy | 2026/02/17 21:52:20 [020] POST /update_jobs/1247168785/increment_metric
2026-02-17T21:52:20.4494409Z   proxy | 2026/02/17 21:52:20 [020] 204 /update_jobs/1247168785/increment_metric
2026-02-17T21:52:20.4530075Z updater | 2026/02/17 21:52:20 INFO <job_1247168785> Starting security update job for Uniswap/v2-core
2026-02-17T21:52:20.4571547Z updater | 2026/02/17 21:52:20 INFO <job_1247168785> Checking if pbkdf2 3.0.17 needs updating
2026-02-17T21:52:20.5544941Z   proxy | 2026/02/17 21:52:20 [022] GET https://registry.npmjs.org/pbkdf2
2026-02-17T21:52:20.5962241Z   proxy | 2026/02/17 21:52:20 [022] 200 https://registry.npmjs.org/pbkdf2
2026-02-17T21:52:20.7317500Z   proxy | 2026/02/17 21:52:20 [024] HEAD https://registry.npmjs.org/pbkdf2/-/pbkdf2-3.1.5.tgz
2026-02-17T21:52:20.7679487Z   proxy | 2026/02/17 21:52:20 [024] 200 https://registry.npmjs.org/pbkdf2/-/pbkdf2-3.1.5.tgz
2026-02-17T21:52:20.8038713Z updater | 2026/02/17 21:52:20 INFO <job_1247168785> Latest version is 3.1.5
2026-02-17T21:52:20.8995921Z   proxy | 2026/02/17 21:52:20 [026] GET https://registry.npmjs.org/@uniswap%2Fv2-core
2026-02-17T21:52:21.0780088Z   proxy | 2026/02/17 21:52:21 [026] 200 https://registry.npmjs.org/@uniswap%2Fv2-core
2026-02-17T21:52:21.0864118Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> VulnerabilityAuditor: starting audit
2026-02-17T21:52:21.0887400Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> VulnerabilityAuditor: missing lockfile
2026-02-17T21:52:21.1835512Z   proxy | 2026/02/17 21:52:21 [028] HEAD https://registry.npmjs.org/pbkdf2/-/pbkdf2-3.1.3.tgz
2026-02-17T21:52:21.2184685Z   proxy | 2026/02/17 21:52:21 [028] 200 https://registry.npmjs.org/pbkdf2/-/pbkdf2-3.1.3.tgz
2026-02-17T21:52:21.2255131Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Started process PID: 1818 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:21.8432637Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Process PID: 1818 completed with status: pid 1818 exit 0
2026-02-17T21:52:21.8433835Z 2026/02/17 21:52:21 INFO <job_1247168785> Total execution time: 0.62 seconds
2026-02-17T21:52:21.8451217Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Started process PID: 1826 with command: {} git reset HEAD --hard {}
2026-02-17T21:52:21.8554142Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Process PID: 1826 completed with status: pid 1826 exit 0
2026-02-17T21:52:21.8555340Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:21.8560552Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Started process PID: 1834 with command: {} git clean -fx {}
2026-02-17T21:52:21.8606750Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Process PID: 1834 completed with status: pid 1834 exit 0
2026-02-17T21:52:21.8608012Z 2026/02/17 21:52:21 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:21.8635554Z updater | 2026/02/17 21:52:21 INFO <job_1247168785> Started process PID: 1842 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:22.0322825Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1842 completed with status: pid 1842 exit 0
2026-02-17T21:52:22.0323882Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.17 seconds
2026-02-17T21:52:22.1340939Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1858 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/repo/git.store' {}
2026-02-17T21:52:22.1392403Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1858 completed with status: pid 1858 exit 0
2026-02-17T21:52:22.1396561Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1400181Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1866 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:22.1443783Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1866 completed with status: pid 1866 exit 0
2026-02-17T21:52:22.1445197Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1454112Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1873 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:22.1504327Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1873 completed with status: pid 1873 exit 0
2026-02-17T21:52:22.1505912Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1510244Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1880 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:22.1548688Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1880 completed with status: pid 1880 exit 0
2026-02-17T21:52:22.1551604Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:22.1555860Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1887 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:22.1601876Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1887 completed with status: pid 1887 exit 0
2026-02-17T21:52:22.1603273Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1609281Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1895 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:22.1646439Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1895 completed with status: pid 1895 exit 0
2026-02-17T21:52:22.1647404Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:22.1652966Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1902 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:22.1698964Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1902 completed with status: pid 1902 exit 0
2026-02-17T21:52:22.1700136Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1705597Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1909 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:22.1751978Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1909 completed with status: pid 1909 exit 0
2026-02-17T21:52:22.1753148Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1757290Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1916 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:22.1803367Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1916 completed with status: pid 1916 exit 0
2026-02-17T21:52:22.1804613Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1809391Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1924 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:22.1854589Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1924 completed with status: pid 1924 exit 0
2026-02-17T21:52:22.1855853Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:22.1863047Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1932 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:22.1903789Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Process PID: 1932 completed with status: pid 1932 exit 0
2026-02-17T21:52:22.1905028Z 2026/02/17 21:52:22 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:22.1911397Z updater | 2026/02/17 21:52:22 INFO <job_1247168785> Started process PID: 1939 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:22.9148936Z   proxy | 2026/02/17 21:52:22 [030] GET https://registry.yarnpkg.com:443/pbkdf2
2026-02-17T21:52:22.9539068Z   proxy | 2026/02/17 21:52:22 [030] 200 https://registry.yarnpkg.com:443/pbkdf2
2026-02-17T21:52:22.9728972Z   proxy | 2026/02/17 21:52:22 [036] GET https://registry.yarnpkg.com:443/ripemd160
2026-02-17T21:52:22.9731610Z   proxy | 2026/02/17 21:52:22 [037] GET https://registry.yarnpkg.com:443/create-hmac
2026-02-17T21:52:22.9732833Z 2026/02/17 21:52:22 [038] GET https://registry.yarnpkg.com:443/safe-buffer
2026-02-17T21:52:22.9737249Z   proxy | 2026/02/17 21:52:22 [039] GET https://registry.yarnpkg.com:443/to-buffer
2026-02-17T21:52:22.9747292Z   proxy | 2026/02/17 21:52:22 [040] GET https://registry.yarnpkg.com:443/sha.js
2026-02-17T21:52:23.0127742Z   proxy | 2026/02/17 21:52:23 [036] 200 https://registry.yarnpkg.com:443/ripemd160
2026-02-17T21:52:23.0211647Z   proxy | 2026/02/17 21:52:23 [042] GET https://registry.yarnpkg.com:443/hash-base
2026-02-17T21:52:23.0518796Z   proxy | 2026/02/17 21:52:23 [037] 200 https://registry.yarnpkg.com:443/create-hmac
2026-02-17T21:52:23.0666686Z   proxy | 2026/02/17 21:52:23 [038] 200 https://registry.yarnpkg.com:443/safe-buffer
2026-02-17T21:52:23.0796820Z   proxy | 2026/02/17 21:52:23 [039] 200 https://registry.yarnpkg.com:443/to-buffer
2026-02-17T21:52:23.0819423Z   proxy | 2026/02/17 21:52:23 [042] 200 https://registry.yarnpkg.com:443/hash-base
2026-02-17T21:52:23.0837045Z   proxy | 2026/02/17 21:52:23 [040] 200 https://registry.yarnpkg.com:443/sha.js
2026-02-17T21:52:23.0948666Z   proxy | 2026/02/17 21:52:23 [046] GET https://registry.yarnpkg.com:443/isarray
2026-02-17T21:52:23.0990488Z   proxy | 2026/02/17 21:52:23 [047] GET https://registry.yarnpkg.com:443/typed-array-buffer
2026-02-17T21:52:23.1005633Z   proxy | 2026/02/17 21:52:23 [048] GET https://registry.yarnpkg.com:443/readable-stream
2026-02-17T21:52:23.1347765Z   proxy | 2026/02/17 21:52:23 [047] 200 https://registry.yarnpkg.com:443/typed-array-buffer
2026-02-17T21:52:23.1420049Z   proxy | 2026/02/17 21:52:23 [046] 200 https://registry.yarnpkg.com:443/isarray
2026-02-17T21:52:23.1467112Z   proxy | 2026/02/17 21:52:23 [052] GET https://registry.yarnpkg.com:443/call-bound
2026-02-17T21:52:23.1477644Z   proxy | 2026/02/17 21:52:23 [053] GET https://registry.yarnpkg.com:443/es-errors
2026-02-17T21:52:23.1489772Z   proxy | 2026/02/17 21:52:23 [054] GET https://registry.yarnpkg.com:443/is-typed-array
2026-02-17T21:52:23.1739203Z   proxy | 2026/02/17 21:52:23 [048] 200 https://registry.yarnpkg.com:443/readable-stream
2026-02-17T21:52:23.1811091Z   proxy | 2026/02/17 21:52:23 [052] 200 https://registry.yarnpkg.com:443/call-bound
2026-02-17T21:52:23.1890919Z   proxy | 2026/02/17 21:52:23 [057] GET https://registry.yarnpkg.com:443/call-bind-apply-helpers
2026-02-17T21:52:23.1902753Z   proxy | 2026/02/17 21:52:23 [058] GET https://registry.yarnpkg.com:443/get-intrinsic
2026-02-17T21:52:23.2098894Z   proxy | 2026/02/17 21:52:23 [053] 200 https://registry.yarnpkg.com:443/es-errors
2026-02-17T21:52:23.2193546Z   proxy | 2026/02/17 21:52:23 [054] 200 https://registry.yarnpkg.com:443/is-typed-array
2026-02-17T21:52:23.2249192Z   proxy | 2026/02/17 21:52:23 [060] GET https://registry.yarnpkg.com:443/which-typed-array
2026-02-17T21:52:23.2372793Z   proxy | 2026/02/17 21:52:23 [057] 200 https://registry.yarnpkg.com:443/call-bind-apply-helpers
2026-02-17T21:52:23.2428014Z   proxy | 2026/02/17 21:52:23 [062] GET https://registry.yarnpkg.com:443/function-bind
2026-02-17T21:52:23.2507609Z   proxy | 2026/02/17 21:52:23 [058] 200 https://registry.yarnpkg.com:443/get-intrinsic
2026-02-17T21:52:23.2651600Z   proxy | 2026/02/17 21:52:23 [069] GET https://registry.yarnpkg.com:443/gopd
2026-02-17T21:52:23.2657545Z   proxy | 2026/02/17 21:52:23 [070] GET https://registry.yarnpkg.com:443/has-symbols
2026-02-17T21:52:23.2671286Z   proxy | 2026/02/17 21:52:23 [071] GET https://registry.yarnpkg.com:443/get-proto
2026-02-17T21:52:23.2679423Z   proxy | 2026/02/17 21:52:23 [072] GET https://registry.yarnpkg.com:443/hasown
2026-02-17T21:52:23.2681028Z 2026/02/17 21:52:23 [060] 200 https://registry.yarnpkg.com:443/which-typed-array
2026-02-17T21:52:23.2695110Z   proxy | 2026/02/17 21:52:23 [073] GET https://registry.yarnpkg.com:443/math-intrinsics
2026-02-17T21:52:23.2707054Z   proxy | 2026/02/17 21:52:23 [074] GET https://registry.yarnpkg.com:443/es-object-atoms
2026-02-17T21:52:23.2765333Z   proxy | 2026/02/17 21:52:23 [076] GET https://registry.yarnpkg.com:443/es-define-property
2026-02-17T21:52:23.2783703Z   proxy | 2026/02/17 21:52:23 [062] 200 https://registry.yarnpkg.com:443/function-bind
2026-02-17T21:52:23.2833748Z   proxy | 2026/02/17 21:52:23 [078] GET https://registry.yarnpkg.com:443/available-typed-arrays
2026-02-17T21:52:23.2993864Z   proxy | 2026/02/17 21:52:23 [069] 200 https://registry.yarnpkg.com:443/gopd
2026-02-17T21:52:23.3077941Z   proxy | 2026/02/17 21:52:23 [080] GET https://registry.yarnpkg.com:443/call-bind
2026-02-17T21:52:23.3107713Z   proxy | 2026/02/17 21:52:23 [071] 200 https://registry.yarnpkg.com:443/get-proto
2026-02-17T21:52:23.3157458Z   proxy | 2026/02/17 21:52:23 [070] 200 https://registry.yarnpkg.com:443/has-symbols
2026-02-17T21:52:23.3173807Z   proxy | 2026/02/17 21:52:23 [083] GET https://registry.yarnpkg.com:443/for-each
2026-02-17T21:52:23.3210128Z   proxy | 2026/02/17 21:52:23 [084] GET https://registry.yarnpkg.com:443/has-tostringtag
2026-02-17T21:52:23.3243350Z   proxy | 2026/02/17 21:52:23 [072] 200 https://registry.yarnpkg.com:443/hasown
2026-02-17T21:52:23.3305820Z   proxy | 2026/02/17 21:52:23 [086] GET https://registry.yarnpkg.com:443/dunder-proto
2026-02-17T21:52:23.3417699Z   proxy | 2026/02/17 21:52:23 [073] 200 https://registry.yarnpkg.com:443/math-intrinsics
2026-02-17T21:52:23.3584518Z   proxy | 2026/02/17 21:52:23 [074] 200 https://registry.yarnpkg.com:443/es-object-atoms
2026-02-17T21:52:23.3612787Z   proxy | 2026/02/17 21:52:23 [078] 200 https://registry.yarnpkg.com:443/available-typed-arrays
2026-02-17T21:52:23.3621436Z   proxy | 2026/02/17 21:52:23 [080] 200 https://registry.yarnpkg.com:443/call-bind
2026-02-17T21:52:23.3625564Z   proxy | 2026/02/17 21:52:23 [083] 200 https://registry.yarnpkg.com:443/for-each
2026-02-17T21:52:23.3675067Z   proxy | 2026/02/17 21:52:23 [076] 200 https://registry.yarnpkg.com:443/es-define-property
2026-02-17T21:52:23.3677498Z   proxy | 2026/02/17 21:52:23 [084] 200 https://registry.yarnpkg.com:443/has-tostringtag
2026-02-17T21:52:23.3731325Z   proxy | 2026/02/17 21:52:23 [090] GET https://registry.yarnpkg.com:443/possible-typed-array-names
2026-02-17T21:52:23.3771043Z   proxy | 2026/02/17 21:52:23 [091] GET https://registry.yarnpkg.com:443/is-callable
2026-02-17T21:52:23.3789292Z   proxy | 2026/02/17 21:52:23 [086] 200 https://registry.yarnpkg.com:443/dunder-proto
2026-02-17T21:52:23.3791056Z   proxy | 2026/02/17 21:52:23 [092] GET https://registry.yarnpkg.com:443/set-function-length
2026-02-17T21:52:23.4137863Z   proxy | 2026/02/17 21:52:23 [091] 200 https://registry.yarnpkg.com:443/is-callable
2026-02-17T21:52:23.4149968Z   proxy | 2026/02/17 21:52:23 [090] 200 https://registry.yarnpkg.com:443/possible-typed-array-names
2026-02-17T21:52:23.4167053Z   proxy | 2026/02/17 21:52:23 [092] 200 https://registry.yarnpkg.com:443/set-function-length
2026-02-17T21:52:23.4254678Z   proxy | 2026/02/17 21:52:23 [095] GET https://registry.yarnpkg.com:443/define-data-property
2026-02-17T21:52:23.4263134Z   proxy | 2026/02/17 21:52:23 [096] GET https://registry.yarnpkg.com:443/has-property-descriptors
2026-02-17T21:52:23.4657658Z   proxy | 2026/02/17 21:52:23 [095] 200 https://registry.yarnpkg.com:443/define-data-property
2026-02-17T21:52:23.4706853Z   proxy | 2026/02/17 21:52:23 [096] 200 https://registry.yarnpkg.com:443/has-property-descriptors
2026-02-17T21:52:23.5828001Z updater | 2026/02/17 21:52:23 INFO <job_1247168785> Process PID: 1939 completed with status: pid 1939 exit 0
2026-02-17T21:52:23.5828940Z 2026/02/17 21:52:23 INFO <job_1247168785> Total execution time: 1.39 seconds
2026-02-17T21:52:23.5854496Z updater | 2026/02/17 21:52:23 INFO <job_1247168785> Started process PID: 1952 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:23.7556691Z updater | 2026/02/17 21:52:23 INFO <job_1247168785> Process PID: 1952 completed with status: pid 1952 exit 0
2026-02-17T21:52:23.7557485Z 2026/02/17 21:52:23 INFO <job_1247168785> Total execution time: 0.17 seconds
2026-02-17T21:52:23.9322178Z updater | 2026/02/17 21:52:23 INFO <job_1247168785> Started process PID: 1960 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:24.5506461Z updater | 2026/02/17 21:52:24 INFO <job_1247168785> Process PID: 1960 completed with status: pid 1960 exit 0
2026-02-17T21:52:24.5507260Z 2026/02/17 21:52:24 INFO <job_1247168785> Total execution time: 0.62 seconds
2026-02-17T21:52:24.5534153Z updater | 2026/02/17 21:52:24 INFO <job_1247168785> Started process PID: 1968 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:25.1714558Z updater | 2026/02/17 21:52:25 INFO <job_1247168785> Process PID: 1968 completed with status: pid 1968 exit 0
2026-02-17T21:52:25.1715605Z 2026/02/17 21:52:25 INFO <job_1247168785> Total execution time: 0.62 seconds
2026-02-17T21:52:25.2687107Z   proxy | 2026/02/17 21:52:25 [098] GET https://registry.npmjs.org/@uniswap%2Fv2-core
2026-02-17T21:52:25.2689412Z 2026/02/17 21:52:25 [098] 200 https://registry.npmjs.org/@uniswap%2Fv2-core (cached)
2026-02-17T21:52:25.2733659Z updater | 2026/02/17 21:52:25 INFO <job_1247168785> Requirements to unlock own
2026-02-17T21:52:25.3655623Z   proxy | 2026/02/17 21:52:25 [100] GET https://registry.npmjs.org/@uniswap%2Fv2-core
2026-02-17T21:52:25.3656624Z   proxy | 2026/02/17 21:52:25 [100] 200 https://registry.npmjs.org/@uniswap%2Fv2-core (cached)
2026-02-17T21:52:25.3718870Z updater | 2026/02/17 21:52:25 INFO <job_1247168785> Requirements update strategy widen_ranges
2026-02-17T21:52:25.3744476Z updater | 2026/02/17 21:52:25 INFO <job_1247168785> Started process PID: 1976 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:26.0026721Z updater | 2026/02/17 21:52:26 INFO <job_1247168785> Process PID: 1976 completed with status: pid 1976 exit 0
2026-02-17T21:52:26.0027761Z 2026/02/17 21:52:26 INFO <job_1247168785> Total execution time: 0.63 seconds
2026-02-17T21:52:26.0053201Z updater | 2026/02/17 21:52:26 INFO <job_1247168785> Started process PID: 1984 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:26.6540773Z updater | 2026/02/17 21:52:26 INFO <job_1247168785> Process PID: 1984 completed with status: pid 1984 exit 0
2026-02-17T21:52:26.6541793Z 2026/02/17 21:52:26 INFO <job_1247168785> Total execution time: 0.65 seconds
2026-02-17T21:52:26.6568052Z updater | 2026/02/17 21:52:26 INFO <job_1247168785> Started process PID: 1992 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:27.2709772Z updater | 2026/02/17 21:52:27 INFO <job_1247168785> Process PID: 1992 completed with status: pid 1992 exit 0
2026-02-17T21:52:27.2710755Z 2026/02/17 21:52:27 INFO <job_1247168785> Total execution time: 0.61 seconds
2026-02-17T21:52:27.2737490Z updater | 2026/02/17 21:52:27 INFO <job_1247168785> Started process PID: 2000 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:27.8829104Z updater | 2026/02/17 21:52:27 INFO <job_1247168785> Process PID: 2000 completed with status: pid 2000 exit 0
2026-02-17T21:52:27.8830216Z 2026/02/17 21:52:27 INFO <job_1247168785> Total execution time: 0.61 seconds
2026-02-17T21:52:27.8864005Z updater | 2026/02/17 21:52:27 INFO <job_1247168785> Started process PID: 2008 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:28.5112417Z updater | 2026/02/17 21:52:28 INFO <job_1247168785> Process PID: 2008 completed with status: pid 2008 exit 0
2026-02-17T21:52:28.5113407Z 2026/02/17 21:52:28 INFO <job_1247168785> Total execution time: 0.63 seconds
2026-02-17T21:52:28.5141034Z updater | 2026/02/17 21:52:28 INFO <job_1247168785> Started process PID: 2016 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:29.1298403Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2016 completed with status: pid 2016 exit 0
2026-02-17T21:52:29.1300153Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.62 seconds
2026-02-17T21:52:29.1314657Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Updating pbkdf2 from 3.0.17 to 3.1.5
2026-02-17T21:52:29.1333678Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2024 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:29.2934652Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2024 completed with status: pid 2024 exit 0
2026-02-17T21:52:29.2935638Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.16 seconds
2026-02-17T21:52:29.3908992Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2032 with command: {} git reset HEAD --hard {}
2026-02-17T21:52:29.3991374Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2032 completed with status: pid 2032 exit 0
2026-02-17T21:52:29.3992340Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:29.4000834Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2039 with command: {} git clean -fx {}
2026-02-17T21:52:29.4043566Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2039 completed with status: pid 2039 exit 0
2026-02-17T21:52:29.4044572Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:29.4107473Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2054 with command: {} git config --global credential.helper '!/home/dependabot/common/lib/dependabot/../../bin/git-credential-store-immutable --file /home/dependabot/dependabot-updater/repo/git.store' {}
2026-02-17T21:52:29.4150941Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2054 completed with status: pid 2054 exit 0
2026-02-17T21:52:29.4152878Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4159811Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2063 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:29.4203798Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2063 completed with status: pid 2063 exit 0
2026-02-17T21:52:29.4204411Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:29.4209937Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2071 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:29.4247825Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2071 completed with status: pid 2071 exit 0
2026-02-17T21:52:29.4248829Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4253885Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2079 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:29.4301047Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2079 completed with status: pid 2079 exit 0
2026-02-17T21:52:29.4302003Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:29.4307575Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2087 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:29.4353336Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2087 completed with status: pid 2087 exit 0
2026-02-17T21:52:29.4353960Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:29.4359390Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2095 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:29.4397661Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2095 completed with status: pid 2095 exit 0
2026-02-17T21:52:29.4398635Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4404732Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2103 with command: {} git config --global --replace-all url.https://github.com/.insteadOf ssh://git@github.com/ {}
2026-02-17T21:52:29.4450103Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2103 completed with status: pid 2103 exit 0
2026-02-17T21:52:29.4451120Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:29.4456326Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2111 with command: {} git config --global --add url.https://github.com/.insteadOf ssh://git@github.com: {}
2026-02-17T21:52:29.4497995Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2111 completed with status: pid 2111 exit 0
2026-02-17T21:52:29.4499014Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4505110Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2118 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com: {}
2026-02-17T21:52:29.4547228Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2118 completed with status: pid 2118 exit 0
2026-02-17T21:52:29.4547953Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4554091Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2127 with command: {} git config --global --add url.https://github.com/.insteadOf git@github.com/ {}
2026-02-17T21:52:29.4592298Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2127 completed with status: pid 2127 exit 0
2026-02-17T21:52:29.4593042Z 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4601797Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2134 with command: {} git config --global --add url.https://github.com/.insteadOf git://github.com/ {}
2026-02-17T21:52:29.4636230Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Process PID: 2134 completed with status: pid 2134 exit 0
2026-02-17T21:52:29.4637460Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Total execution time: 0.0 seconds
2026-02-17T21:52:29.4647481Z updater | 2026/02/17 21:52:29 INFO <job_1247168785> Started process PID: 2141 with command: node /opt/npm_and_yarn/run.js
2026-02-17T21:52:30.1746233Z   proxy | 2026/02/17 21:52:30 [102] GET https://registry.yarnpkg.com:443/pbkdf2
2026-02-17T21:52:30.1748677Z 2026/02/17 21:52:30 [102] 200 https://registry.yarnpkg.com:443/pbkdf2 (cached)
2026-02-17T21:52:30.1917727Z   proxy | 2026/02/17 21:52:30 [108] GET https://registry.yarnpkg.com:443/create-hmac
2026-02-17T21:52:30.1939644Z   proxy | 2026/02/17 21:52:30 [108] 200 https://registry.yarnpkg.com:443/create-hmac (cached)
2026-02-17T21:52:30.1946438Z   proxy | 2026/02/17 21:52:30 [109] GET https://registry.yarnpkg.com:443/safe-buffer
2026-02-17T21:52:30.1947799Z 2026/02/17 21:52:30 [109] 200 https://registry.yarnpkg.com:443/safe-buffer (cached)
2026-02-17T21:52:30.1958376Z   proxy | 2026/02/17 21:52:30 [110] GET https://registry.yarnpkg.com:443/ripemd160
2026-02-17T21:52:30.1959860Z 2026/02/17 21:52:30 [110] 200 https://registry.yarnpkg.com:443/ripemd160 (cached)
2026-02-17T21:52:30.1967926Z   proxy | 2026/02/17 21:52:30 [111] GET https://registry.yarnpkg.com:443/sha.js
2026-02-17T21:52:30.1969212Z 2026/02/17 21:52:30 [111] 200 https://registry.yarnpkg.com:443/sha.js (cached)
2026-02-17T21:52:30.1980776Z   proxy | 2026/02/17 21:52:30 [112] GET https://registry.yarnpkg.com:443/to-buffer
2026-02-17T21:52:30.1982124Z 2026/02/17 21:52:30 [112] 200 https://registry.yarnpkg.com:443/to-buffer (cached)
2026-02-17T21:52:30.2117817Z   proxy | 2026/02/17 21:52:30 [116] GET https://registry.yarnpkg.com:443/hash-base
2026-02-17T21:52:30.2123977Z   proxy | 2026/02/17 21:52:30 [116] 200 https://registry.yarnpkg.com:443/hash-base (cached)
2026-02-17T21:52:30.2193265Z   proxy | 2026/02/17 21:52:30 [117] GET https://registry.yarnpkg.com:443/typed-array-buffer
2026-02-17T21:52:30.2194297Z 2026/02/17 21:52:30 [117] 200 https://registry.yarnpkg.com:443/typed-array-buffer (cached)
2026-02-17T21:52:30.2201419Z   proxy | 2026/02/17 21:52:30 [118] GET https://registry.yarnpkg.com:443/isarray
2026-02-17T21:52:30.2204172Z 2026/02/17 21:52:30 [118] 200 https://registry.yarnpkg.com:443/isarray (cached)
2026-02-17T21:52:30.2301549Z   proxy | 2026/02/17 21:52:30 [123] GET https://registry.yarnpkg.com:443/readable-stream
2026-02-17T21:52:30.2303080Z 2026/02/17 21:52:30 [123] 200 https://registry.yarnpkg.com:443/readable-stream (cached)
2026-02-17T21:52:30.2349263Z   proxy | 2026/02/17 21:52:30 [124] GET https://registry.yarnpkg.com:443/es-errors
2026-02-17T21:52:30.2351242Z 2026/02/17 21:52:30 [124] 200 https://registry.yarnpkg.com:443/es-errors (cached)
2026-02-17T21:52:30.2415454Z   proxy | 2026/02/17 21:52:30 [125] GET https://registry.yarnpkg.com:443/is-typed-array
2026-02-17T21:52:30.2416971Z   proxy | 2026/02/17 21:52:30 [125] 200 https://registry.yarnpkg.com:443/is-typed-array (cached)
2026-02-17T21:52:30.2424014Z   proxy | 2026/02/17 21:52:30 [126] GET https://registry.yarnpkg.com:443/call-bound
2026-02-17T21:52:30.2425345Z 2026/02/17 21:52:30 [126] 200 https://registry.yarnpkg.com:443/call-bound (cached)
2026-02-17T21:52:30.2515585Z   proxy | 2026/02/17 21:52:30 [130] GET https://registry.yarnpkg.com:443/which-typed-array
2026-02-17T21:52:30.2522955Z   proxy | 2026/02/17 21:52:30 [130] 200 https://registry.yarnpkg.com:443/which-typed-array (cached)
2026-02-17T21:52:30.2548180Z   proxy | 2026/02/17 21:52:30 [131] GET https://registry.yarnpkg.com:443/call-bind-apply-helpers
2026-02-17T21:52:30.2554325Z   proxy | 2026/02/17 21:52:30 [131] 200 https://registry.yarnpkg.com:443/call-bind-apply-helpers (cached)
2026-02-17T21:52:30.2565763Z   proxy | 2026/02/17 21:52:30 [132] GET https://registry.yarnpkg.com:443/get-intrinsic
2026-02-17T21:52:30.2580015Z   proxy | 2026/02/17 21:52:30 [132] 200 https://registry.yarnpkg.com:443/get-intrinsic (cached)
2026-02-17T21:52:30.2763924Z   proxy | 2026/02/17 21:52:30 [141] GET https://registry.yarnpkg.com:443/call-bind
2026-02-17T21:52:30.2770581Z   proxy | 2026/02/17 21:52:30 [141] 200 https://registry.yarnpkg.com:443/call-bind (cached)
2026-02-17T21:52:30.2771797Z 2026/02/17 21:52:30 [142] GET https://registry.yarnpkg.com:443/available-typed-arrays
2026-02-17T21:52:30.2775012Z 2026/02/17 21:52:30 [142] 200 https://registry.yarnpkg.com:443/available-typed-arrays (cached)
2026-02-17T21:52:30.2781148Z   proxy | 2026/02/17 21:52:30 [143] GET https://registry.yarnpkg.com:443/for-each
2026-02-17T21:52:30.2782354Z 2026/02/17 21:52:30 [143] 200 https://registry.yarnpkg.com:443/for-each (cached)
2026-02-17T21:52:30.2856606Z   proxy | 2026/02/17 21:52:30 [144] GET https://registry.yarnpkg.com:443/get-proto
2026-02-17T21:52:30.2858982Z   proxy | 2026/02/17 21:52:30 [144] 200 https://registry.yarnpkg.com:443/get-proto (cached)
2026-02-17T21:52:30.2869102Z   proxy | 2026/02/17 21:52:30 [145] GET https://registry.yarnpkg.com:443/function-bind
2026-02-17T21:52:30.2873135Z   proxy | 2026/02/17 21:52:30 [145] 200 https://registry.yarnpkg.com:443/function-bind (cached)
2026-02-17T21:52:30.2880562Z   proxy | 2026/02/17 21:52:30 [146] GET https://registry.yarnpkg.com:443/gopd
2026-02-17T21:52:30.2886833Z   proxy | 2026/02/17 21:52:30 [146] 200 https://registry.yarnpkg.com:443/gopd (cached)
2026-02-17T21:52:30.2899174Z   proxy | 2026/02/17 21:52:30 [148] GET https://registry.yarnpkg.com:443/hasown
2026-02-17T21:52:30.2900392Z   proxy | 2026/02/17 21:52:30 [148] 200 https://registry.yarnpkg.com:443/hasown (cached)
2026-02-17T21:52:30.2901852Z   proxy | 2026/02/17 21:52:30 [149] GET https://registry.yarnpkg.com:443/has-tostringtag
2026-02-17T21:52:30.2903714Z   proxy | 2026/02/17 21:52:30 [149] 200 https://registry.yarnpkg.com:443/has-tostringtag (cached)
2026-02-17T21:52:30.3005114Z   proxy | 2026/02/17 21:52:30 [153] GET https://registry.yarnpkg.com:443/es-object-atoms
2026-02-17T21:52:30.3009090Z   proxy | 2026/02/17 21:52:30 [153] 200 https://registry.yarnpkg.com:443/es-object-atoms (cached)
2026-02-17T21:52:30.3017412Z   proxy | 2026/02/17 21:52:30 [154] GET https://registry.yarnpkg.com:443/math-intrinsics
2026-02-17T21:52:30.3018551Z   proxy | 2026/02/17 21:52:30 [154] 200 https://registry.yarnpkg.com:443/math-intrinsics (cached)
2026-02-17T21:52:30.3026716Z   proxy | 2026/02/17 21:52:30 [155] GET https://registry.yarnpkg.com:443/has-symbols
2026-02-17T21:52:30.3030206Z   proxy | 2026/02/17 21:52:30 [155] 200 https://registry.yarnpkg.com:443/has-symbols (cached)
2026-02-17T21:52:30.3114261Z   proxy | 2026/02/17 21:52:30 [160] GET https://registry.yarnpkg.com:443/es-define-property
2026-02-17T21:52:30.3131829Z   proxy | 2026/02/17 21:52:30 [161] GET https://registry.yarnpkg.com:443/set-function-length
2026-02-17T21:52:30.3141144Z   proxy | 2026/02/17 21:52:30 [161] 200 https://registry.yarnpkg.com:443/set-function-length (cached)
2026-02-17T21:52:30.3161009Z   proxy | 2026/02/17 21:52:30 [160] 200 https://registry.yarnpkg.com:443/es-define-property (cached)
2026-02-17T21:52:30.3163792Z   proxy | 2026/02/17 21:52:30 [162] GET https://registry.yarnpkg.com:443/is-callable
2026-02-17T21:52:30.3164751Z   proxy | 2026/02/17 21:52:30 [162] 200 https://registry.yarnpkg.com:443/is-callable (cached)
2026-02-17T21:52:30.3165739Z   proxy | 2026/02/17 21:52:30 [163] GET https://registry.yarnpkg.com:443/dunder-proto
2026-02-17T21:52:30.3168284Z   proxy | 2026/02/17 21:52:30 [164] GET https://registry.yarnpkg.com:443/possible-typed-array-names
2026-02-17T21:52:30.3170112Z   proxy | 2026/02/17 21:52:30 [163] 200 https://registry.yarnpkg.com:443/dunder-proto (cached)
2026-02-17T21:52:30.3172014Z   proxy | 2026/02/17 21:52:30 [164] 200 https://registry.yarnpkg.com:443/possible-typed-array-names (cached)
2026-02-17T21:52:30.3286101Z   proxy | 2026/02/17 21:52:30 [167] GET https://registry.yarnpkg.com:443/define-data-property
2026-02-17T21:52:30.3287743Z 2026/02/17 21:52:30 [167] 200 https://registry.yarnpkg.com:443/define-data-property (cached)
2026-02-17T21:52:30.3296792Z   proxy | 2026/02/17 21:52:30 [168] GET https://registry.yarnpkg.com:443/has-property-descriptors
2026-02-17T21:52:30.3303909Z   proxy | 2026/02/17 21:52:30 [168] 200 https://registry.yarnpkg.com:443/has-property-descriptors (cached)
2026-02-17T21:52:30.4593189Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Process PID: 2141 completed with status: pid 2141 exit 0
2026-02-17T21:52:30.4593959Z 2026/02/17 21:52:30 INFO <job_1247168785> Total execution time: 0.99 seconds
2026-02-17T21:52:30.4621946Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Started process PID: 2154 with command: {} git status --untracked-files all --porcelain v1 . {}
2026-02-17T21:52:30.4675344Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Process PID: 2154 completed with status: pid 2154 exit 0
2026-02-17T21:52:30.4676215Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:30.4687262Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Started process PID: 2162 with command: {} git status --untracked-files all --porcelain v1 .yarn/cache {}
2026-02-17T21:52:30.4733323Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Process PID: 2162 completed with status: pid 2162 exit 0
2026-02-17T21:52:30.4734374Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:30.4741966Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Started process PID: 2169 with command: {} git status --untracked-files all --porcelain v1 .yarn/install-state.gz {}
2026-02-17T21:52:30.4796896Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Process PID: 2169 completed with status: pid 2169 exit 0
2026-02-17T21:52:30.4797904Z 2026/02/17 21:52:30 INFO <job_1247168785> Total execution time: 0.01 seconds
2026-02-17T21:52:30.4810259Z updater | 2026/02/17 21:52:30 INFO <job_1247168785> Submitting pbkdf2 pull request for creation
2026-02-17T21:52:30.5700143Z   proxy | 2026/02/17 21:52:30 [170] GET https://api.github.com:443/repos/Uniswap/v2-core/commits?per_page=100
2026-02-17T21:52:30.5701314Z 2026/02/17 21:52:30 [170] * authenticating github api request with token for api.github.com
2026-02-17T21:52:31.0241468Z   proxy | 2026/02/17 21:52:31 [170] 200 https://api.github.com:443/repos/Uniswap/v2-core/commits?per_page=100
2026-02-17T21:52:31.3966843Z   proxy | 2026/02/17 21:52:31 [172] GET https://registry.npmjs.org/pbkdf2/latest
2026-02-17T21:52:31.4382661Z   proxy | 2026/02/17 21:52:31 [172] 200 https://registry.npmjs.org/pbkdf2/latest
2026-02-17T21:52:31.4800482Z   proxy | 2026/02/17 21:52:31 [174] GET https://api.github.com:443/repos/browserify/pbkdf2/releases?per_page=100
2026-02-17T21:52:31.4801411Z   proxy | 2026/02/17 21:52:31 [174] * authenticating github api request with token for api.github.com
2026-02-17T21:52:31.6675599Z   proxy | 2026/02/17 21:52:31 [174] 200 https://api.github.com:443/repos/browserify/pbkdf2/releases?per_page=100
2026-02-17T21:52:31.6985247Z   proxy | 2026/02/17 21:52:31 [176] GET https://api.github.com:443/repos/browserify/pbkdf2/contents/
2026-02-17T21:52:31.6986414Z 2026/02/17 21:52:31 [176] * authenticating github api request with token for api.github.com
2026-02-17T21:52:31.8224736Z   proxy | 2026/02/17 21:52:31 [176] 200 https://api.github.com:443/repos/browserify/pbkdf2/contents/
2026-02-17T21:52:31.8579060Z   proxy | 2026/02/17 21:52:31 [178] GET https://api.github.com:443/repos/browserify/pbkdf2/contents/CHANGELOG.md?ref=master
2026-02-17T21:52:31.8580111Z   proxy | 2026/02/17 21:52:31 [178] * authenticating github api request with token for api.github.com
2026-02-17T21:52:32.0227095Z   proxy | 2026/02/17 21:52:32 [178] 200 https://api.github.com:443/repos/browserify/pbkdf2/contents/CHANGELOG.md?ref=master
2026-02-17T21:52:32.1224670Z   proxy | 2026/02/17 21:52:32 [180] GET https://github.com/browserify/pbkdf2.git/info/refs?service=git-upload-pack
2026-02-17T21:52:32.1225610Z 2026/02/17 21:52:32 [180] * authenticating git server request (host: github.com)
2026-02-17T21:52:32.1960121Z   proxy | 2026/02/17 21:52:32 [180] 200 https://github.com/browserify/pbkdf2.git/info/refs?service=git-upload-pack
2026-02-17T21:52:32.2459702Z   proxy | 2026/02/17 21:52:32 [182] GET https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.0.17
2026-02-17T21:52:32.2460487Z 2026/02/17 21:52:32 [182] * authenticating github api request with token for api.github.com
2026-02-17T21:52:32.4341219Z   proxy | 2026/02/17 21:52:32 [182] 200 https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.0.17
2026-02-17T21:52:32.4547420Z   proxy | 2026/02/17 21:52:32 [184] GET https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.1.5
2026-02-17T21:52:32.4548422Z 2026/02/17 21:52:32 [184] * authenticating github api request with token for api.github.com
2026-02-17T21:52:32.6515278Z   proxy | 2026/02/17 21:52:32 [184] 200 https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.1.5
2026-02-17T21:52:32.7240867Z   proxy | 2026/02/17 21:52:32 [186] GET https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.0.17
2026-02-17T21:52:32.7242681Z   proxy | 2026/02/17 21:52:32 [186] 200 https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.0.17 (cached)
2026-02-17T21:52:32.7479252Z   proxy | 2026/02/17 21:52:32 [188] GET https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.1.5
2026-02-17T21:52:32.7480921Z 2026/02/17 21:52:32 [188] 200 https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.1.5 (cached)
2026-02-17T21:52:32.7792743Z   proxy | 2026/02/17 21:52:32 [190] GET https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.0.17
2026-02-17T21:52:32.7794021Z 2026/02/17 21:52:32 [190] 200 https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.0.17 (cached)
2026-02-17T21:52:32.8061210Z   proxy | 2026/02/17 21:52:32 [192] GET https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.1.5
2026-02-17T21:52:32.8062941Z   proxy | 2026/02/17 21:52:32 [192] 200 https://api.github.com:443/repos/browserify/pbkdf2/commits?sha=v3.1.5 (cached)
2026-02-17T21:52:32.9197359Z   proxy | 2026/02/17 21:52:32 [194] GET https://registry.npmjs.org/pbkdf2
2026-02-17T21:52:32.9198309Z 2026/02/17 21:52:32 [194] 200 https://registry.npmjs.org/pbkdf2 (cached)
2026-02-17T21:52:33.1345306Z   proxy | 2026/02/17 21:52:33 [196] POST /update_jobs/1247168785/create_pull_request
2026-02-17T21:52:33.8107173Z   proxy | 2026/02/17 21:52:33 [196] 204 /update_jobs/1247168785/create_pull_request
2026-02-17T21:52:33.9105785Z   proxy | 2026/02/17 21:52:33 [198] POST /update_jobs/1247168785/record_ecosystem_meta
2026-02-17T21:52:33.9962196Z   proxy | 2026/02/17 21:52:33 [198] 204 /update_jobs/1247168785/record_ecosystem_meta
2026-02-17T21:52:34.0975333Z   proxy | 2026/02/17 21:52:34 [200] PATCH /update_jobs/1247168785/mark_as_processed
2026-02-17T21:52:34.2458625Z   proxy | 2026/02/17 21:52:34 [200] 204 /update_jobs/1247168785/mark_as_processed
2026-02-17T21:52:34.2514002Z updater | 2026/02/17 21:52:34 INFO <job_1247168785> Finished job processing
2026-02-17T21:52:34.2547667Z updater | 2026/02/17 21:52:34 INFO Results:
2026-02-17T21:52:34.2548172Z +-------------------------------------------+
2026-02-17T21:52:34.2548660Z |    Changes to Dependabot Pull Requests    |
2026-02-17T21:52:34.2549119Z +---------+---------------------------------+
2026-02-17T21:52:34.2549834Z | created | pbkdf2 ( from 3.0.17 to 3.1.5 ) |
2026-02-17T21:52:34.2550303Z +---------+---------------------------------+
2026-02-17T21:52:34.4857769Z Cleaned up container 61f7dd09efc3381e6bc889952f6f44ab2fe2276088b6f0fae505cc9326ae09d1
2026-02-17T21:52:34.4941652Z   proxy | 2026/02/17 21:52:34 41/100 calls cached (41%)
2026-02-17T21:52:34.4942369Z 2026/02/17 21:52:34 Posting metrics to remote API endpoint
2026-02-17T21:52:34.7088356Z 🤖 ~ finished ~
2026-02-17T21:52:34.7199893Z Post job cleanup.
2026-02-17T21:52:34.8838465Z Cleaning up orphan processes
