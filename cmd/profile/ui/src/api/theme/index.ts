import { get } from '/@/utils/request';
const baseURL = "/api"
export const getThemeConfigApi = (p: object) => get( baseURL + "/theme" , p);