const passportLogoutUrl = "https://passport.thingplecloud.com/logout"
const passportLoginUrl = "https://passport.thingplecloud.com/login"

function genUrl(url: string, path: string) {
  return url + "?path=" + path
}

function getLogoutUrl(path: string) {
  return genUrl(passportLogoutUrl, path)
}

function getLoginUrl(path: string) {
  return genUrl(passportLoginUrl, path)
}

export function logout() {
  let fullPath = encodeURIComponent(window.location.href)
  let url = getLogoutUrl(fullPath)
  window.location.replace(url)
}

export function login() {
  let fullPath = encodeURIComponent(window.location.href)
  let url = getLoginUrl(fullPath)
  window.location.replace(url)
}