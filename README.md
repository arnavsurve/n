# n

A simple and extensible CLI for personal note-taking.

## Features

- **Daily notes**: Automatically templated daily notes
- **Quick inbox**: Fast capture with timestamps
- **Project notes**: Timestamped project notes with slugs
- **Search**: Ripgrep + fzf integration for fast note discovery

## Installation

```bash
make install
```

This will build and install the binary to `~/.local/bin/n`.

## Usage

### Daily note (default)
```bash
n          # Opens today's daily note
n daily    # Same as above
```

### Quick inbox capture
```bash
n inbox follow up on email thread
```
Appends a timestamped entry to `~/notes/inbox/inbox.md` and opens it in your editor.

### New project note
```bash
n new graph indexing    # Creates ~/notes/projects/20251008-1547-graph-indexing.md
n new                   # Creates ~/notes/projects/20251008-1547-untitled.md
```

### Search notes
```bash
n search
```
Uses ripgrep and fzf to search all notes. Requires `rg`, `fzf`, and `bat` to be installed.

## Configuration

The CLI uses the following defaults:
- **Notes directory**: `~/notes/`
- **Editor**: `$EDITOR` environment variable (defaults to `nvim`)

Directory structure:
```
~/notes/
├── daily/      # Daily notes (YYYY-MM-DD.md)
├── inbox/      # Quick captures (inbox.md)
├── projects/   # Project notes (YYYYMMDD-HHMM-slug.md)
└── scratch/    # Scratch space
```

## Replacing zsh functions

After installing, you can remove the old zsh functions from `~/.zshrc`:
- `n()` - lines 178-195
- `ni()` - lines 198-204
- `nn()` - lines 207-225
- `nr()` - lines 228-234

## Dependencies

- Go 1.24+
- ripgrep (`rg`) - for search
- fzf - for search
- bat - for search preview

## Development

```bash
make build    # Build binary
make install  # Build and install to ~/.local/bin
make clean    # Remove binary and clean build cache
```
