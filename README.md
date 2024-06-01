# web-archiver

Utility to save on the [Web Archive](https://archive.org/) pages from a [sitemap](https://www.sitemaps.org/)

## Usage

```shell
web-archiver <sitemap URI> [<date>]
```

This programs needs 2 environment variables:

- WAYBACK_S3_ACCESS_KEY
- WAYBACK_S3_SECRET_KEY

Keys are obtained from [Web Archive S3-Like API](https://archive.org/account/s3.php)

The program will save all the entries present in the sitemap with a
`lastMod` property newer than the provided date.

If no date is provided all entries will be saved

### Examples

```shell
# Save only URLs with `lastMod` newer than 2024-05-01
web-archiver https://example.com/sitemap.xml 2024-05-01
```

Only URLs modified since 2024-05-01 will be saved

```shell
# Save all the URLs present in the sitemap
web-archiver https://example.com/sitemap.xml
```
