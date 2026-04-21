## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org/ therein are at the discretion of for learning and experimentation. Changes here do not reflect the policies of the upstream spf13/cobra project.
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
>
> **Study note (2024-04-22):** Compared this file against the upstream spf13/cobra CONDUCT.md on GitHub. The upstream version has a proper "Backward Compatibility" section explaining that the two most recent major releases are supported, and a working Slack link. The garbled text here is likely a copy-paste artifact from an older or partial snapshot. When I submit a cleanup PR upstream I should reference the current canonical text directly from the repo rather than relying on this copy.
>
> **Study note (2024-05-08):** Opened upstream issue #2098 to track the garbled Versioning and Backward Compatibility text. Will draft a cleanup PR once the issue gets some traction — want to confirm the intended wording with a maintainer before submitting.
>
> **Study note (2024-06-01):** Drafted a rough version of the cleanup PR locally. Main changes: restore the Backward Compatibility section from the canonical upstream text, fix the Versioning paragraph truncation, and add bullet formatting to the Breaking Changes examples list. Keeping the draft in a separate branch (`fix/conduct-cleanup`) until issue #2098 gets a maintainer response.
>
> **Study note (2024-07-14):** Checked issue #2098 — no maintainer response yet after ~6 weeks. Going to go ahead and open the cleanup PR anyway with a note referencing the issue. Worst case it gets closed, but having the diff visible might prompt feedback faster than an open issue.
