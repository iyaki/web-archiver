# web-archiver

A command-line utility to save web pages from a [sitemap](https://www.sitemaps.org/) to the [Web Archive](https://archive.org/) (archive.org).

## Features

- Parse standard XML sitemaps
- Save pages to Web Archive with full captures (screenshots, outlinks)
- Filter pages by last modification date
- Concurrent processing for faster archiving

## Installation

### From Source

Requires Go 1.22 or later:

```bash
go install github.com/iyaki/web-archiver/v2@latest
```

### From Releases

Download pre-built binaries from the [Releases page](https://github.com/iyaki/web-archiver/releases).

## Configuration

This program requires Web Archive API credentials as environment variables:

- `WAYBACK_S3_ACCESS_KEY` - Your Web Archive S3 access key
- `WAYBACK_S3_SECRET_KEY` - Your Web Archive S3 secret key

Get your API keys from [Web Archive S3-Like API](https://archive.org/account/s3.php).

## Usage

```bash
web-archiver <sitemap_url> [<date>]
```

### Arguments

- `sitemap_url` (required): URL to the sitemap XML file
- `date` (optional): Filter date in ISO format (YYYY-MM-DD). Only URLs with `lastMod` newer than this date will be saved.

### Examples

Save all URLs from a sitemap:

```bash
web-archiver https://example.com/sitemap.xml
```

Save only URLs modified since a specific date:

```bash
web-archiver https://example.com/sitemap.xml 2024-05-01
```
