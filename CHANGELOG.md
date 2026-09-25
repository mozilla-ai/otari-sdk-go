# Changelog

## [0.3.1](https://github.com/mozilla-ai/otari-sdk-go/compare/v0.3.0...v0.3.1) (2026-09-25)


### Bug Fixes

* **release:** tag releases as plain vX.Y.Z so go get can resolve them ([#34](https://github.com/mozilla-ai/otari-sdk-go/issues/34)) ([3ba7154](https://github.com/mozilla-ai/otari-sdk-go/commit/3ba7154eb128464547c3a42f27a9ff1554c637a8))

## [0.3.0](https://github.com/mozilla-ai/otari-sdk-go/compare/otari-sdk-go-v0.2.0...otari-sdk-go-v0.3.0) (2026-09-25)


### ⚠ BREAKING CHANGES

* a base URL must be the gateway origin. A base URL ending in /v1 was previously trimmed and is now sent as-is, which will not resolve.

### Features

* regenerate client core and address the gateway at /api/v1 ([#28](https://github.com/mozilla-ai/otari-sdk-go/issues/28)) ([98d7aeb](https://github.com/mozilla-ai/otari-sdk-go/commit/98d7aeb04f9325adbaf7e548c6990d6b464311f6))


### Bug Fixes

* **ci:** make the endpoint-coverage check offline and deterministic ([#26](https://github.com/mozilla-ai/otari-sdk-go/issues/26)) ([864abd7](https://github.com/mozilla-ai/otari-sdk-go/commit/864abd7a5d9823624ac5ae2787950b04bdedbb33))
* **control-plane:** map generated errors to typed SDK errors ([#20](https://github.com/mozilla-ai/otari-sdk-go/issues/20)) ([a9de4a6](https://github.com/mozilla-ai/otari-sdk-go/commit/a9de4a67da60101b9b89db8fce57d96a2b4de82d))
* read the renamed Otari-Attempt-ID response header ([#33](https://github.com/mozilla-ai/otari-sdk-go/issues/33)) ([d811546](https://github.com/mozilla-ai/otari-sdk-go/commit/d811546ff0e21389f19756d9bbbfe9472eb9c806)), closes [#32](https://github.com/mozilla-ai/otari-sdk-go/issues/32)

## [0.2.0](https://github.com/mozilla-ai/otari-sdk-go/compare/otari-sdk-go-v0.1.1...otari-sdk-go-v0.2.0) (2026-06-16)


### Features

* add image generation and audio (speech/transcription) methods ([#15](https://github.com/mozilla-ai/otari-sdk-go/issues/15)) ([8d2660d](https://github.com/mozilla-ai/otari-sdk-go/commit/8d2660db1d5713e2245cab790bd116bf4a4d6de4))

## [0.1.1](https://github.com/mozilla-ai/otari-sdk-go/compare/otari-sdk-go-v0.1.0...otari-sdk-go-v0.1.1) (2026-06-12)


### Features

* independent release automation + surface gateway spec version ([#11](https://github.com/mozilla-ai/otari-sdk-go/issues/11)) ([0b9b85e](https://github.com/mozilla-ai/otari-sdk-go/commit/0b9b85e6698af37ee2500b17de1913f30c6d1fae))
* wrap /v1/messages/count_tokens (regenerate core + ergonomic method) ([#9](https://github.com/mozilla-ai/otari-sdk-go/issues/9)) ([68fa34a](https://github.com/mozilla-ai/otari-sdk-go/commit/68fa34a821d19824dcb4bdc93eb9f5069546745b))


### Bug Fixes

* regenerate SDK client core so message.reasoning is a string ([#13](https://github.com/mozilla-ai/otari-sdk-go/issues/13)) ([34491f7](https://github.com/mozilla-ai/otari-sdk-go/commit/34491f7e3b1f2728ecf53ac282c84a4db30d6c1c))
