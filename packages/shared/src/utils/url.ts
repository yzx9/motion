export function joinPath(base: string, path: string): string {
  if (base.endsWith("/") && path.startsWith("/")) {
    return base + path.slice(1)
  } else if (!base.endsWith("/") && !path.startsWith("/")) {
    return base + "/" + path
  } else {
    return base + path
  }
}

const EXTERNAL_URL_RE = /^https?:/i

export function startsWithProtocol(url: string): boolean {
  return EXTERNAL_URL_RE.test(url)
}
