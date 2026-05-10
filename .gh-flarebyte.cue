package ghflarebyte

project: {
  org:  "flarebyte"
  repo: "snake-knot-picker"
}

sync: {
  mode: "push"
}

repository: {
  description:   "Validation library for Go with a compact schema language built from CLI-style arguments."
  defaultBranch: "main"
  homepage:      "https://github.com/flarebyte/snake-knot-picker"
  visibility:    "public"
  template:      false
  topics: [
    "go",
    "validation",
    "schema",
    "cli",
  ]
  labels: [
    {
      name:        "bug"
      color:       "d73a4a"
      description: "Something isn't working"
    },
    {
      name:        "documentation"
      color:       "0075ca"
      description: "Improvements or additions to documentation"
    },
    {
      name:        "duplicate"
      color:       "cfd3d7"
      description: "This issue or pull request already exists"
    },
    {
      name:        "enhancement"
      color:       "a2eeef"
      description: "New feature or request"
    },
    {
      name:        "good first issue"
      color:       "7057ff"
      description: "Good for newcomers"
    },
    {
      name:        "help wanted"
      color:       "008672"
      description: "Extra attention is needed"
    },
    {
      name:        "invalid"
      color:       "e4e669"
      description: "This doesn't seem right"
    },
    {
      name:        "question"
      color:       "d876e3"
      description: "Further information is requested"
    },
    {
      name:        "wontfix"
      color:       "ffffff"
      description: "This will not be worked on"
    },
  ]
  features: {
    issues:                       true
    wiki:                         false
    projects:                     false
    discussions:                  false
    autoMerge:                    true
    mergeCommit:                  false
    rebaseMerge:                  false
    squashMerge:                  true
    squashMergeCommitMessage:     "pr-title"
    deleteBranchOnMerge:          true
    allowForking:                 false
    allowUpdateBranch:            false
    advancedSecurity:             true
    secretScanning:               true
    secretScanningPushProtection: true
  }
}

build: {
  language:             "go"
  mode:                 "library"
  packages: [
    "./...",
  ]
  runTests:             true
  outputDir:            "build"
  checksumFile:         "build/checksums.txt"
  artifactTargetSuffix: true
  targets: [
    "darwin-arm64",
    "linux-amd64",
  ]
}

release: {
  versionSource:    "main.project.yaml"
  tagPrefix:        "v"
  notesMode:        "generate-notes"
  includeArtifacts: false
  artifactDir:      "build"
  includeChecksums: true
}
