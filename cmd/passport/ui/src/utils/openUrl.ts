import {Local} from "/@/utils/storage";

export function getOpenUrl(referrer:any, path:any) {
    var openUrl = ''
    if (path) {
        if (path.indexOf("?") < 0) {
            openUrl = referrer + "#" + path + "?token=" + Local.get("token")
        } else {
            openUrl = referrer + "#" + path + "&token=" + Local.get("token")
        }
    } else {
        openUrl = referrer + "#/" + "?token=" + Local.get("token")
    }
    // console.log("openUrl",openUrl)
    return openUrl;
}