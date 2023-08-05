import {Local} from "/@/utils/storage";

export function getOpenUrl(referrer: any, path: any,token:any) {
    var openUrl = ''
    if (path) {
        if (path.indexOf("?") < 0) {
            openUrl = path + "?token=" + token
        } else {
            openUrl = path + "&token=" + token
        }
    } else {
        openUrl = referrer + "#/" + "?token=" + token
    }
    Local.set("openUrl",openUrl)
    return openUrl;
}