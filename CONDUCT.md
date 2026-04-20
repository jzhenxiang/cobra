## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org/). release.

### Backward be a lower priority than a high severity.slack.com/archives/CD3LP1199) as the primary means is to foster open communication with all users and contributors.

### Breaking Changes
Breaking changes are generally/pflag`, `spf13/viper` etc...
  - Some version updates may be acceptable for picking up bug fixes, but maintainers must exercise caution when reviewing.

There may, at times, need to be exceptions where breaking changes are allowed in release branches. These are at the discretion of the project's maintainers, and must be carefully considered before merging.

### CI Testing
Maintainers will ensure the Cobra test suite utilizes the current supported versions of Golang.

### Disclaimer
Changes to this document and the contents therein are at the discretion of the maintainers.

---

> **Personal fork note:** This is my personal fork for learning and experimentation. Changes here do not reflect the policies of the upstream spf13/cobra project.
>
> **TODO:** The CVE section above appears to have garbled/incomplete text — worth fixing in a PR to upstream once I understand the intended policy better.
>
> **TODO:** The Backward Compatibility section also appears truncated/garbled — the full text should describe which two major releases are maintained and link to the slack channel properly.
>
> **Study note (2024-01-15):** Reading through the Breaking Changes section — key takeaway is that release branches are treated as stable and breaking changes there require explicit maintainer sign-off. Good pattern to follow in my own projects.
>
> **Study note (2024-02-03):** The two-major-releases compatibility window is a useful policy. For my own CLIs I should decide early whether I want to commit to something similar or keep it informal while the project is small.
>
> **Study note (2024-03-10):** Noticed the "Examples" paragraph under Breaking Changes is missing list formatting — items like "Removing or renaming exported constant, variable" and the pflag/viper version update note should probably be a proper bullet list. Flagging here as a reminder to check the upstream source and potentially submit a formatting fix.
