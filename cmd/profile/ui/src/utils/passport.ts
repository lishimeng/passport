import { Local } from '/@/utils/storage';
// const passportLogoutUrl = "https://passport.thingplecloud.com/logout"
// const passportLoginUrl = "https://passport.thingplecloud.com/login"

function genUrl(url: string, path: string) {
  return url + "?path=" + path
}

function getLogoutUrl(path: string) {
  let passportLogoutUrl = Local.get('themeConfig') ? Local.get('themeConfig').passportLogoutUrl : ""
  return genUrl(passportLogoutUrl, path)
}

function getLoginUrl(path: string) {
  let passportLoginUrl = Local.get('themeConfig') ? Local.get('themeConfig').passportLoginUrl : ""
  return genUrl(passportLoginUrl, path)
}

export function logout() {
  let href = urlFilter(window.location.href)
  let fullPath = encodeURIComponent(href)
  let url = getLogoutUrl(fullPath)
  window.location.replace(url)
}

export function login() {
  let href = urlFilter(window.location.href)
  let fullPath = encodeURIComponent(href)
  let url = getLoginUrl(fullPath)
  window.location.replace(url)
}

var urlFilter = function (url: string) {
  const key = "token"
  var e = eval('/' + key + '=[^&]*&?/g')
  return url.replace(e, '')
}