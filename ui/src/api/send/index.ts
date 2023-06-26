import {get, post} from "/@/utils/request";

const baseUrl = "/api"
export const sendCodeApi = (p: object) => get(baseUrl + "/sign_in/sendCode", p);
