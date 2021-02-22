#  (2021-02-22)



## [0.17.2](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.17.1...0.17.2) (2021-02-22)


### Bug Fixes

* **i18n:** fix error where locale file not being loaded ([#374](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/374)) ([ce40d2a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ce40d2ada03f9a99a01a32fc7be5738ae2c72def))



## [0.17.1](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.17.0...0.17.1) (2021-02-22)


### Bug Fixes

* **login:** fix nil-pointer error ([#373](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/373)) ([8820492](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8820492056c6000af7130d5fd64da9f78b23719e))



# [0.17.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.16.0...0.17.0) (2021-02-19)


### Bug Fixes

* i18n errors ([#353](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/353)) ([654cfb7](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/654cfb778655dbbaf1a8c2ddfead5a14814dd400))
* invalid YAML ([8f4fff8](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8f4fff80f0fc8cde516ae4c72ac4f31c655aa75c))
* service account i18n ([#344](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/344)) ([a7d631e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a7d631ea22558457fc7b4392a4a952eb8f9a557d))
* use yq only if version >= 4 ([#367](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/367)) ([79f2afa](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/79f2afab8c6cc256da8cd2e21c87a5980cbe13b6))


### Features

* **kafka topic:** add topic commands ([#309](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/309)) ([9e399f5](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/9e399f588eb144b7695261df992f388ad5ca17a2))
* **whoami:** add whoami command ([#356](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/356)) ([421b165](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/421b165b1092f8e3d6096a5eb337433ff543b87c)), closes [#339](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/339)



# [0.16.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.15.1...0.16.0) (2021-02-10)


### Bug Fixes

* add ability to force delete ([#329](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/329)) ([fb53f6b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/fb53f6b6825e7358136b114a3be5b1021add7bb2))
* refresh token if no access token is provided ([#326](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/326)) ([c931bd2](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c931bd2bfc4fa6ab5bea44cf4c06d8ef11c44edc))
* **kafka delete:** confirm name only to delete ([#321](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/321)) ([a51a7db](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a51a7db426d6c3c0e28b0b89a833c11e928cae67))


### Features

* **kafka create:** use a positional argument for Kafka create ([#330](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/330)) ([3a0e30f](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3a0e30f96e56784cc2745c5a12c554ad982e5972))



## [0.15.1](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.15.0...0.15.1) (2021-02-04)


### Bug Fixes

* **kafka delete:** add async=true to ensure Kafka can be deleted ([#314](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/314)) ([87eddb6](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/87eddb6f4e3f55d454c573d898c312753b1a99e7))
* **kafka topic:** change topic command to singular form ([#308](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/308)) ([3d36326](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3d363266f83760bbe6e6b14850aa7a6d56283069))



# [0.15.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.14.1...0.15.0) (2021-01-28)


### Bug Fixes

* handle "MGD-SERV-API-36" error code ([#305](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/305)) ([8ca3f1a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8ca3f1a08e83f4186b323a27463cb156ed758bbd))


### Features

* **status:** add root-level status command ([#301](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/301)) ([ce30137](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ce30137448fbaafe7b21b358fcf457b11861fea5))



## [0.14.1](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.14.0...0.14.1) (2021-01-28)


### Bug Fixes

* print only single topics ([#300](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/300)) ([be76612](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/be76612d027594f3a80f6013a3cbeaaf5a4332e7))



# [0.14.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.13.2...0.14.0) (2021-01-26)


### Bug Fixes

* **cluster info:** rename command info to status ([#289](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/289)) ([25a7eb2](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/25a7eb2b44249dfe45ebcfbf51a7b7b8f34369ac)), closes [#282](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/282)
* remove unused function ([#275](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/275)) ([5896729](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/58967297f3f9fc4f9ab2009387f00bb1b244c74c))
* **connection:** only refresh tokens when needed ([#274](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/274)) ([1c1056e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1c1056e1f7e55fefd77e774db1e8c04e0c2b3705))
* BootstrapServerHost nil pointer ([#269](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/269)) ([e8eda42](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/e8eda42b7d313f4bcae0b76909cb3c95b515e43a))
* refactor cluster connect to use new format of the CRD's ([#247](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/247)) ([8e59246](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8e59246fdb9fa78c0c3aa0cc5409c0efc084da47))
* **docs:** remove the docs command ([#267](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/267)) ([ed00c08](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ed00c0840859082dba5a7de1929b3d214b700ca9))


### Features

* **login page:** use Patternfly empty state template ([#292](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/292)) ([cc10856](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/cc10856494a69161eff4a180904c28d7364799f4))
* standardise colors for printing to console ([#291](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/291)) ([2c7f7f0](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/2c7f7f02ceca25f748a1f42fed70395aa48a7cc9))



## [0.13.2](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.13.1...0.13.2) (2021-01-21)


### Bug Fixes

* pointer error when bootstrap host is empty ([#266](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/266)) ([7281b8b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/7281b8b3088af61787f73177648fec40fe938008))



## [0.13.1](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.13.0...0.13.1) (2021-01-21)


### Bug Fixes

* **status:** fix pointer error ([#262](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/262)) ([300ffda](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/300ffdaa4abd9f84b2120b6ca903ab238cc4ed08))



# [0.13.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.12.0...0.13.0) (2021-01-21)


### Bug Fixes

* negate flag value check ([#254](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/254)) ([db524a6](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/db524a6870fdee38e50b8be51159dcd3f7745f7b))


### Features

* **serviceaccount:** add interactive mode for the reset credentials command ([#248](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/248)) ([16699e5](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/16699e574fed5d5e4bbdce6e33ce6d8646004a42))



# [0.12.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.11.0...0.12.0) (2021-01-20)


### Bug Fixes

* remove kafka credentials format ([#245](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/245)) ([ffbb807](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ffbb8070ea1d9178a8dd1c8c6f7c6b402e0734a3))



# [0.11.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.10.0...0.11.0) (2021-01-19)


### Bug Fixes

* standardize table output format flag ([#233](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/233)) ([a4b5e65](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a4b5e65ee876b0c65275d300399da128e054ddf0))
* usused option value ([#231](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/231)) ([18c3f89](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/18c3f895f421dc0eb17d5402901db4b8a1cac48a))
* **serviceaccount:** remove ability to force delete service accounts ([#230](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/230)) ([79061f0](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/79061f0f18e6a14a5878d962bfd9eb1f8e523704))


### Features

* **kafka:** require name confirmation ([#227](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/227)) ([f661229](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f661229f011c3c395ae7b030ff4d8aac38752be6))
* **kafka > create:** add interactive create mode ([#236](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/236)) ([ee46c6c](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ee46c6c03c8db658a93f2f1ba253f66193b38a2e))
* **kafka>list:** append port to end of bootstrap URL ([#234](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/234)) ([277244a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/277244a1256d2d6e4586a7cd995cc54079c601d4))
* **status:** print Bootstrap URL ([#235](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/235)) ([6ecaf6a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/6ecaf6a3ac3150b5e65efcb428128544f9275965))



# [0.10.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.9.3...0.10.0) (2021-01-14)


### Bug Fixes

* **topics:** missing connection option ([#223](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/223)) ([c5d4c3f](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c5d4c3fd2dc8a625676d55577317b83ead2516f7))


### Features

* add service account CRUD commands ([#216](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/216)) ([3e2f9bc](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3e2f9bc57b1afc0b47d6164779d4512ccbe35276))



## [0.9.3](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.9.2...0.9.3) (2021-01-11)


### Bug Fixes

* pointer error when bootstrap host is empty ([#214](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/214)) ([c05ccb5](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c05ccb50ad80173ba6ffd558175e86612a7ec6e4))


### Features

* **login:** add ability to provide custom openid scope ([#210](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/210)) ([79e7c0e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/79e7c0e8fdaae59f08c9606819ed60a568805547))



## [0.9.2](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.9.1...0.9.2) (2021-01-05)


### Bug Fixes

* ensure context is cancelled when finished ([#198](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/198)) ([5398d57](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5398d572ff7b23d6ad010f14074f0f8f56586745))



## [0.9.1](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.9.0...0.9.1) (2021-01-05)



# [0.9.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.8.0...0.9.0) (2020-12-15)


### Bug Fixes

* append :443 to BootstrapServerHost ([#176](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/176)) ([1cc38a5](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1cc38a5cb3f0e1a23252895977d705f0f9a8143b))
* do not use a pointer for a slice ([37ddd3a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/37ddd3ae7cdc8220a880333f362f0ae75938a316))


### Features

* add insecure data plane ([#127](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/127)) ([8ca363b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8ca363bd986567a844f4bdaeea7f505c1215de59))



# [0.8.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.7.1...0.8.0) (2020-12-14)


### Features

* print sso url in login ([#167](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/167)) ([699da53](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/699da53ae851a2c46f204a5bf40bf6ea235c9ece))



## [0.7.1](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.7.0...0.7.1) (2020-12-14)


### Bug Fixes

* display API error reason ([#164](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/164)) ([202d056](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/202d056e15e0edee05f8ced18eea81a177b7b2bc))



# [0.7.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.6.0...0.7.0) (2020-12-11)


### Bug Fixes

* Initial version of SASL/Plain support for topic creation ([#161](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/161)) ([6df805a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/6df805a2a1eb78cfa0154148abd68252c5233e29))
* list command with pagination ([#156](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/156)) ([5ab441d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5ab441dcb27afda32cbd14661387a2d4d957b532))
* remove credentials file ([a3468ed](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a3468ed631c4d5e8a28a1dc8aeb7b79a5c22e111))
* return error ([#159](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/159)) ([8056bc8](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8056bc82dc57064a5f5023a2639b33018933cde7))



# [0.6.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.5.0...0.6.0) (2020-12-10)


### Bug Fixes

* bump version to 0.6.0 ([7d964f4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/7d964f40a0b97b8e81728782c41b922c207af8dc))
* navigation for cli documentation ([#150](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/150)) ([6f05da0](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/6f05da0afb008e9905ce2ab20c924cd3da696d4d))
* pandoc trying to remove twice ([#152](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/152)) ([ca134f3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ca134f384250c865ff4da312f23571145acd7844))
* remove trailing % from stdout/stderr messages ([#147](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/147)) ([d41cf3e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/d41cf3e25af463553e629f068f9fbb9a18c74783))



# [0.5.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.4.0...0.5.0) (2020-12-10)


### Bug Fixes

* change default client ID and remove token login ([#146](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/146)) ([bd1beba](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/bd1beba05dd4c4e46ddc481d320d840f2e49b5f3))



# [0.4.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.3.0...0.4.0) (2020-12-09)


### Bug Fixes

* adding kuberentes secret as output ([#138](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/138)) ([4cda41c](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/4cda41c8b9b9f16ef4c44192fdf9476089779d36))
* CR name in credentials ([b2928a3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/b2928a302b9fd926ddb1a0ed5a5ca947975ef6a6))
* rename kafka cluster to kafka instance ([#144](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/144)) ([e67cd0d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/e67cd0d135aae18b988c38616ac81fe7b98cbcf9))


### Features

* auto-use kafka cluster after creation ([#142](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/142)) ([1c2a4c0](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1c2a4c098e8b0a5ca61cb19c1fb0091aef3512e2))
* refactor connect to use top level group ([#139](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/139)) ([1328345](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1328345dddac5e6d3116c417f1e38b8a083fbc48))



# [0.3.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.2.0...0.3.0) (2020-12-08)


### Bug Fixes

* add -n flag for create ([#119](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/119)) ([c4aeddb](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c4aeddb7c8053c834514e883cd07ce0b24f648e9))
* add missing builders file ([67c5375](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/67c5375409e3361ece0bfd9b54fb66394b5a3713))
* change apiversion for connect command ([b57c68a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/b57c68a6a74b0f642f57ad30b656cb3cd55afb71))
* Cleanup of the documentation topics ([9a1b45b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/9a1b45bb362d7e27bc771dbcf72310dc45e85d7f))
* make auth url hard-coded ([#102](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/102)) ([e762865](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/e7628650ba84dd3747c866fadfc8a5535e2da94c))
* Make CR using namespaced scope ([#116](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/116)) ([9aa1406](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/9aa1406f9a2848c693a8b1a565b30718ad7fdc31))
* make create work ([#133](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/133)) ([12ff528](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/12ff52882ceb758f85747aed6b9ad3e3b9079698))
* parse API URL to get host and scheme ([#106](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/106)) ([1e5c2f4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1e5c2f4d2e515049ae7cc4526ada4a4f4358e3c6))
* remove trailing slash from url ([#103](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/103)) ([1aa3ad3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1aa3ad39e97f0693d3a9409f0ed2a6655713580b))
* Rename cr version ([#113](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/113)) ([a69c829](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a69c8298b354c945c8ef061b8b1ea831d23a99b3))
* unused flag for linting ([88a3957](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/88a3957e68456301e542ed9315d39990c723a03f))
* update branch ([b0cb9fe](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/b0cb9fe41ee90ee08de9d4251e18c5726b36e8f2))


### Features

* allow using the currently selected Kafka cluster in the describe command ([#114](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/114)) ([325c6ea](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/325c6ea04ab237ef358985d76f17326fbd4b8bd2))
* expanded help for credentials command ([#120](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/120)) ([c3f50fd](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c3f50fd2d7bc0a227fcae3b9d60d43c9da30e54a))
* rhoas kafka connect command ([#85](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/85)) ([8afe30d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8afe30d6c330b2c842d69494d80290f190c05bde))
* show message on login success ([31d0f35](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/31d0f3540987f210ba990348e417fb06d48f36a0))
* token-based login ([#132](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/132)) ([a08e51c](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a08e51c75b3083440e9ff350db12d4497f309c90))
* update OPENAPI spec for Service Account ([#121](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/121)) ([c9805af](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c9805afa42e44d81e8f49fc1f2963d02a0145579))
* wip: validate kafka name ([#131](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/131)) ([5b46929](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5b46929143f776325b561696a65f2c0da20a4eb5))
* **cmd:** add YAML output format ([f050dab](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f050dabdc574f33e0bc71b9ce0c871ff7706150b))



# [0.2.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/0.1.0...0.2.0) (2020-11-20)



# [0.1.0](https://github.com/bf2fc6cc711aee1a0c2a/cli/compare/8b8f02000dde5c7d55be5b8922f27dc99af07f68...0.1.0) (2020-11-18)


### Bug Fixes

* add authz ([b402415](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/b4024154a404ed28bd6c59a59cd3721d10deff71))
* add basic documentation ([#67](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/67)) ([2efe4c3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/2efe4c333d7514bffbb40f33a5cda41b9a7c410e))
* Add dummy test targetr ([fc7afbb](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/fc7afbbef8edf16d64dc4591acffff77709a6d6b))
* add formatting check to PR's ([bea223a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/bea223a1d3b0234c694a0eb44c8f540bf347f210))
* add handy kafka docker compose to the mock ([3d41f6f](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3d41f6f241adf9595a2036857b615cfba18ef92e))
* add initial version of goreleaser ([3651762](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3651762d38612179f05f41dbde4cccfe02bbbeda))
* add minor fixes ([5e8b172](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5e8b172091ad65eb4fd328d3cc1e85f325624228))
* add missing elements to guide ([1cf2161](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1cf2161556973a16a3d8a26d84c5df6202480beb))
* add package ([977d3bd](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/977d3bddba327f9c93e812634ab3c1a257ffd3e7))
* add release process docs ([ad42abf](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ad42abfd864722ec1cd3adaf4d674a52973df39b))
* apply fedback by [@wtrocki](https://github.com/wtrocki) ([d6310c4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/d6310c499846989b962e4708a670016aae179a16))
* build for mac and linux ([7e214f1](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/7e214f18174f817813f812b0d157789afb50abe9))
* build pipeline  ([3db6ae3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3db6ae348abdbbe1d56815758e544b948e3b5a58))
* cleanup commands documents for usability ([#69](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/69)) ([8436101](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/843610112e454b6afc756c01cd4ba5ff81cc5676))
* disable documentation creator ([7572c4d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/7572c4dac04143bbaec1986ad807f371a178da31))
* disable invalid printing for login/logout ([6f1950d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/6f1950dfdb071ccf3c23cdbe633537beb23ad7d6))
* Do not require gopath on build ([835ae3e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/835ae3e2ede44a5d248e469bff56db3484a1f8bc))
* documentation generator ([8d71656](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8d716568c41ab19838609106159226776eec1596))
* formatting of the status command ([010349b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/010349b811f21e45eba59f3918fc47a16b208329))
* general improvements to make file ([f8e3143](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f8e3143a121687b56b081939a41c47d9514da476))
* Guide for running this docs ([9fe0b2c](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/9fe0b2cba4dd5ff9f73a4c85c00f4f50e63f91d4))
* minor changes for the demo ([a8ffbd8](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a8ffbd889f057496c20857030e8333a7660f9f94))
* minor fixes ([f2d23e3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f2d23e332e3a44aee8662cd40c28874a1e57ae46))
* reduce golang versions ([8155373](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8155373ca9f7edac5112c00a832587c4e8479e04))
* remove function used to test bot ([66c9a2c](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/66c9a2c468af75eb66fd16c614ef8d6d995054a6))
* remove operator from the repository ([19ee8f4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/19ee8f4fe853c357f83143d0ede20e28f09cc71e))
* Remove token mock ([#66](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/66)) ([17e4d39](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/17e4d396ea4830aecebd5554d2b4d54cab8e6585))
* **kafka:** change default region to "us-west-1" ([71f2148](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/71f2148ba9a2bd1ad96b1638c1c546b5695db5c7))
* **login:** make token required for now until a proper login flow is figured out ([be28892](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/be28892c4bc0fd1fdbac61b921db29b0c5073b17)), closes [/github.com/bf2fc6cc711aee1a0c2a/cli/issues/48#issuecomment-726061600](https://github.com//github.com/bf2fc6cc711aee1a0c2a/cli/issues/48/issues/issuecomment-726061600)
* make credentials file more secure ([110b1e8](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/110b1e87982028b8cb0af92a7d1aae828d53e87d))
* Remove architecture for cli ([e481511](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/e481511baf6eb50e0d47872f18179808886ff8d8))
* remove vendor folder. It should not be used with packages ([09e0049](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/09e0049b415af48f0f35cef6854e06fada1c2350))
* rename yml file ([3c5858e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3c5858ed3473ca9624529e5544e52a710f702c8c))
* reorganization of the structure ([b2c2d20](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/b2c2d203d6aff7ffcdd1f3fc947051557534ed35))
* reorganize script for api updates ([c5aca5c](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c5aca5ca428ca3191b157934990c6345bd1c2f20))
* resolve confusion around authorization command ([ed120c2](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ed120c27b5f529f11d506ba6a29177b778416125))
* resolve formatting problems ([08481cf](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/08481cff990d00be54e94462fd4ad8a9a20a55b6))
* revert changes for formatting ([3d7ba66](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3d7ba666ae4edf2c536100d6b0470872adb8936e))
* update api ([f4a4891](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f4a489104b2577e938bad479310a73c72b4027f5))
* Update gomod version ([81297d2](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/81297d29a0c0ad795945ca6f337f8908b147a736))
* **cmd:** typo in command name ([f8bbc79](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f8bbc79de9a5a569f132ad58bcfec0f9606ebc71))
* **kafka:** create command returns 202 and always require async=true ([d6867e0](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/d6867e08e6c5a9fae23485c8c186f5a486a7b80a))
* **kafka:** delete status code results is 204 and not 200; ([74eead6](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/74eead6c1127bf21639e8b813fce82ee5b181313))
* **kafka:** stop command execution when user is not loggen in ([acdf6f6](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/acdf6f626dc077d7415278d87a815b6ae9e7317f)), closes [#17](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/17)
* **login:** check token expiration before sending request to control plane ([8b262f3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8b262f31176b627b8cf9e3d1d74d94d109e20fab)), closes [#22](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/22)
* **login:** make staging the default environment and do not require "url" ([7dd6aa8](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/7dd6aa898620c32b31e5049b678c7355d1c6b91c)), closes [#18](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/18)
* add demo setup ([fb1a9a3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/fb1a9a370bc19a467d77bc3a2c690f50357e6e21))
* add docusaurus for the demo ([5f52e23](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5f52e230d63452f31164aa9f074968e26e90147f))
* add error handling ([6f879ee](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/6f879ee85905e21b03c9506f52c3e7e8a653921d))
* add extra commands ([b71f313](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/b71f313fce6f94346c3f4972620cd54eb072ad68))
* Add logout ([6415187](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/6415187e5e51f6e370ec1bf83c60428d04133cd5))
* add missing files to client ([82f4c7b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/82f4c7b86c30a2c6bd10c20bf1d5d795098b2441))
* add new info to readme ([780519b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/780519b528de98a0fcfa2c266d2ac5e942980f9b))
* add spec for operator to read config ([f71d367](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f71d3671129d66e301af721bd0e52253d7e08d54))
* Add support for credentials ([e366398](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/e3663989a713d3c3b317107f63979871a3ed6b91))
* additional commands and formatting ([418c303](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/418c30377633a59b43cd7632d56e9064686e731e))
* Base for the unit and integration tests ([4bff4bc](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/4bff4bc5b740edc5ab08869221be4835de227b3b))
* build issue with wrong arg ([307266b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/307266bf2a0e739ae9ca8ec919828c9ec947be90))
* change namespace ([465c94b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/465c94b493fcaabe3bd0812e8def38a9b5cb8d95))
* CMD backbone ([8761a81](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8761a81f5ce456e272ef56727993960d274ade1c))
* command completion ([94f4429](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/94f44294141e646ea3be46795bbe6358ae4bc2d3))
* format for the cli ([62cde70](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/62cde7075f4d09f71a7045ff1aa2cb3cc2d1b2a8))
* formatting ([41293f2](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/41293f232d5ec745aafb89e63b91acd641149946))
* functional operator ([c2840c5](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c2840c5b8c47bd72735f1f377593d61b74e2a520))
* improve architecture ([43e7df3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/43e7df33936bff87b072a52207d3685aed86e774))
* Improve commands ([2c7c877](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/2c7c877b5b7b5665e145363f5510ca9ec16f4949))
* improve deletion script ([2fe22ae](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/2fe22ae48044eca2418d598dd42e8a178fa2f1af))
* Improve formatting ([1cbac9e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1cbac9e9f54e1afc50a0cb1fa4d4ae5b0dac3303))
* Initial architecture ([93cf653](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/93cf6535fe4c20b9c90f78afde42ee294f40c45a))
* list command ([ec22648](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ec22648b1847f2190e2c3e7543d1ee1dd2a82150))
* makefile install problem ([5b31a7b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5b31a7bacc81b46912515407e0a0bcd4f22c5e20))
* minor fixes based on the approved spec ([f424cdf](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/f424cdfc379c99d29521f9b9ed18899eac4ea5e7))
* minor improvements ([53c28f7](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/53c28f7aeca50e268b705ff66087a332a64f5ba3))
* mock ([e6333e4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/e6333e40bdc08189fd619952468209e1de7d2319))
* mock index page ([9891cda](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/9891cdaf16accdb98931a503126898d79c2212ec))
* move package to root ([16cdf80](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/16cdf80ca1b8d42c423656380c12272b899405a3))
* multi_az to boolean ([7af3c10](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/7af3c10789d76d5d5e4079535f0ab56c01f16e26))
* name issue ([c849039](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c849039835c96a58b4c6c2d02f12e9dcd51bbd21))
* openapi make file ([1521cb9](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/1521cb9626233e4b52c6388a11088ba0cd894b08))
* provide script for the provisioning of the clusters ([0dc469d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/0dc469deed471d8b98890fb667b55c237dc49e39))
* remove duplicate ([078029b](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/078029baf896d022ffea4ed707359ecfd69aa6f7))
* rename cli ([a1db3e4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a1db3e440a534b2cef107a65d525b1c8966d0824))
* rename cli ([0760cf7](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/0760cf7acd771b5130a9979d42857f114e5a42ae))
* rename client ([a2fa470](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a2fa47085d58686d6fef2f87c7c7b8bfcd3bd979))
* rename command ([3b9f14e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/3b9f14e89ad74480360c6ff85340fe899a4c6316))
* rename folder ([28f5c00](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/28f5c00045e29a73e88299fab2bf8e76fde9d26b))
* rename operator ([2dd1d16](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/2dd1d1640f4efd1e844b6ea4b675169f20b17054))
* support for all commands ([88226c9](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/88226c99c855dab841a26a1cc778a12028ca3f9d))
* support for help in browser ([7932671](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/793267185de8b9296fc5bf58018721766153762f))
* support for the create with some missing environment abstraction ([5dc2d59](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/5dc2d5904c5e34ed066bc9bdd2e6da019d79ccb3))
* support loging flow ([dfe6593](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/dfe6593744ed6a2cb31e7d22a2ac49365d96506a))
* switch to github package name ([dc5d676](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/dc5d6769ed8a5aa11eefc32dd5a70cc474bc9c0b))
* typo ([8f571f4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8f571f49ec3cc42c70d94f835d7c91c33ddb7e25))
* Use golang setup action ([20b2dcc](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/20b2dcc5abfc3e0a6a6071a46b78c04294887d31))
* Use make when building command ([9ff776d](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/9ff776dc2b559f339647f4b118c9c39f19cabbad))
* use packge name ([d8788d9](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/d8788d94028d06c6af25763fcb244fb4b84371d9))
* website backbone ([0cf02f3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/0cf02f3e18c42533c868d46e54c2951b23502697))


### Features

* **cmd:** Display message if there are no clusters ([#45](https://github.com/bf2fc6cc711aee1a0c2a/cli/issues/45)) ([34a30b7](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/34a30b7576f1d761febbab5f93061a68f9fc62fd))
* **login:** login using the --token flow ([78ed20a](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/78ed20ab1ec6314272a8555fbb8a3b703ce81f25))
* add config ([ff49a5f](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/ff49a5f0336d43725a922021657d6abf580c071e))
* add status command ([090341e](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/090341e9650d996e506a50d329e65f8ec22d2dae))
* **kafka:** add mocked version of topics command ([c9e499f](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/c9e499f9fe0a49fca3a2cac59e3b706bade59573))
* mock server used for the demo purposes ([88f4638](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/88f463863d8565f41ecd1eddd63fcd3110beb9a4))
* open browser according to OS ([a642046](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a6420461b486284fe82a1007dd1ecf3a827b5d39))
* OpenAPI generated client ([a6a96e4](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a6a96e44656bd974eecba033c539800dc580e843))
* Openshift CR's ([8b8f020](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/8b8f02000dde5c7d55be5b8922f27dc99af07f68))
* Operator using SDK ([a0286c5](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/a0286c54a54a7ee25ce84979faba945a5a25b6c3))
* positional argument to reference Kafka ([be88ec3](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/be88ec330af6ab7b173c5377d338dc33356b1774))
* print kafka instances to table ([15ea6b9](https://github.com/bf2fc6cc711aee1a0c2a/cli/commit/15ea6b943023a18093e32219ebb2f20233d42928))



